package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	// многоуровневый базовый путь (роут). в случае повторной регистрации будет паника
	http.HandleFunc("/", rootHandler)

	// фиксированный путь
	http.HandleFunc("/home", homeHandler)

	// многоуровневый путь. После него может быть больше потомков
	http.HandleFunc("/article/", articleHandler)

	// фиксированный путь
	//http.HandleFunc("/article/hello", nil)

	// регулярки не работают.
	//http.HandleFunc("article/*/hello/world", nil)

	// редирект
	http.HandleFunc("/redirect", http.RedirectHandler("http://localhost:9000/home", http.StatusPermanentRedirect).ServeHTTP)
	http.HandleFunc("/redirect/2", redirectHandler)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(_ http.ResponseWriter, r *http.Request) {
	fmt.Println("root")
}

func homeHandler(_ http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
}

func articleHandler(_ http.ResponseWriter, r *http.Request) {
	fmt.Println("article")
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:9000/article", http.StatusPermanentRedirect)
}
