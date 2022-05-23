package main

import (
	"log"
	"main/database"
	"main/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    database.CreateDBConnection()
    database.AutoMigrateDB()

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/create", routes.CreateUser).Methods("POST")
    router.HandleFunc("/show", routes.ShowUser).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
}
