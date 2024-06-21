package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World from Backend!")

	s := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.HandleFunc("/lol", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		w.Write([]byte("Moin!"))
	})

	panic(s.ListenAndServe())
}
