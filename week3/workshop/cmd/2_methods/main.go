package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	var implementation server1
	http.HandleFunc("/article", authMiddleware(articleHandler(implementation)))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func articleHandler(implementation server1) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			implementation.Get(w, r)
		case http.MethodPost:
			implementation.Create(w, r)
		case http.MethodPut:
			implementation.Update(w, r)
		case http.MethodDelete:
			implementation.Delete(w, r)
		default:
			log.Fatal("unsupported method")
		}
	}
}

type server1 struct {
	a string
}

func (s *server1) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create")
}

func (s *server1) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update")
}

func (s *server1) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
}

func (s *server1) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get")
}

// Authorization: Basic (<user:password> : base64)
func authMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fmt.Println(username, password)

		handler.ServeHTTP(w, r)
	}
}
