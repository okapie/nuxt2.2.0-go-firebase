package main

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "io/ioutil"
    "database/sql"

    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    /* 'imported and not used' error occurred if without _. */
    _ "github.com/go-sql-driver/mysql"
)

type Card struct {
    Id               string    `json:"id"`
    Name             string    `json:"name"`
    UploadedImage    string    `json:"uploadedImage"`
}

type Cards []Card

func getCardList(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:password@/sample")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    rows, err := db.Query("select id, name, uploadedImage from cards")
    if err != nil {
        panic(err.Error())
    }

    var cards []Card

    for rows.Next() {
        var id string
        var name string
        var uploadedImage string

        rows.Scan(&id ,&name, &uploadedImage)
        cards = append(cards, Card{id, name, uploadedImage })
    }

    cardsBytes, _ := json.Marshal(&cards)

    w.Write(cardsBytes)
    db.Close()
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
