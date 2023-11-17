package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	}

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/product", productHandler)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile Page"))
	})
	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to home"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Product Handler"))

	fmt.Fprintf(w, "Product page : %d", idNumb)
}
