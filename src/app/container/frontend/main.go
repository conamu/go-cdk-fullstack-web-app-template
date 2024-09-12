package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World from Frontend!")

	s := http.Server{
		Addr: ":80",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	panic(s.ListenAndServe())
}
