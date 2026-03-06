package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"week3/workshop/internal/pkg/db"
	"week3/workshop/internal/pkg/repository"
	"week3/workshop/internal/pkg/repository/postgresql"

	"github.com/gorilla/mux"
)

const port = ":9000"
const queryParamKey = "key"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := db.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetPool(ctx).Close()

	articlesRepo := postgresql.NewArticles(database)

	implementation := server1{repo: articlesRepo}

	router := createRouter(implementation)
	http.Handle("/", router)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}

func createRouter(implementation server1) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/article", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			implementation.Create(w, r)
		case http.MethodPut:
			implementation.Update(w, r)
		default:
			fmt.Println("error")
		}
	})

	router.HandleFunc(fmt.Sprintf("/article/{%s:[0-9]+}", queryParamKey), func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			implementation.Get(w, r)
		case http.MethodPost:
			implementation.Update(w, r)
		default:
			fmt.Println("error")
		}
	})

	return router
}

type server1 struct {
	repo *postgresql.ArticleRepos
}

type addArticleRequest struct {
	Name   string `json:"name"`
	Rating int64  `json:"rating"`
}

func (s *server1) Create(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var unm addArticleRequest
	if err = json.Unmarshal(body, &unm); err != nil {
		fmt.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	articleRepo := &repository.Article{
		Name:   unm.Name,
		Rating: unm.Rating,
	}

	id, err := s.repo.Add(r.Context(), articleRepo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	articleRepo.ID = id
	articleJson, _ := json.Marshal(articleRepo)
	w.Write(articleJson)
}

func (s *server1) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update")
}

func (s *server1) Get(w http.ResponseWriter, r *http.Request) {
	key, ok := mux.Vars(r)[queryParamKey]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	article, err := s.repo.GetByID(r.Context(), keyInt)

	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	articleJson, _ := json.Marshal(article)
	w.Write(articleJson)
}
