package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

func public(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"message": "hello public!"})
}

func private(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"message": "hello private!"})
}

func main() {
    allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
    allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

    r := mux.NewRouter()
    r.HandleFunc("/public", public).Methods("GET")
    r.HandleFunc("/private", private).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
