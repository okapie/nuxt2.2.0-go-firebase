package main

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "io/ioutil"

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

func postCard(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("[request body row] " + string(body))

    var card Card
    if err := json.Unmarshal(body, &card); err != nil {
        log.Fatal(err)
    }

    fmt.Fprint(w, "Received Post Request.")
}

func main() {
    allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
    allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

    r := mux.NewRouter()
    r.HandleFunc("/api/v1/cards", getCardList).Methods("GET")
    r.HandleFunc("/api/v1/registration", postCard).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
