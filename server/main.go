package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

type Card struct {
    Id               string    `json:"id"`
    Name             string    `json:"name"`
    UploadedImage    string    `json:"uploadedImage"`
}

type Cards []Card

func getCardList(w http.ResponseWriter, r *http.Request) {
    cards := Cards{
        Card{Id: "1", Name: "TEST001", UploadedImage: "TEST001.png"},
        Card{Id: "2", Name: "TEST002", UploadedImage: "TEST002.png"},
    }

    json.NewEncoder(w).Encode(cards)
}

func main() {
    allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
    allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

    r := mux.NewRouter()
    r.HandleFunc("/api/v1/cards", getCardList).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
