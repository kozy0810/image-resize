package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func SetCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
	})
}

func main() {
	r := mux.NewRouter()
	r.Use(SetCors)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "ping"}`))
	})

	r.HandleFunc("/images", Upload)

	fmt.Println("Listen to :8080")
	http.ListenAndServe(":8080", r)
}
