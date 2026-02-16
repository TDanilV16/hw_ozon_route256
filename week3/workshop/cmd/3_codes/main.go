package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	var implementation server1
	http.HandleFunc("/article", articleHandler(implementation))
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if string(body) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if string(body) == "asd" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

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
