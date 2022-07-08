package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    )

func main() {

    r := mux.NewRouter()

    r.HandleFunc("/hello", func(resp http.ResponseWriter, r *http.Request) {
        fmt.Println("HELLO WORLD")
    })

    log.Print("Listening..")
    http.ListenAndServe(":8080", r)

}
