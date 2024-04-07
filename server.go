package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("response"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler)
	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("vim-go")
	panic(s.ListenAndServe())
}
