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
    router.HandleFunc("/signup", routes.).Methods("POST")
    router.HandleFunc("/create", routes.CreateUser).Methods("POST")
    router.HandleFunc("/show", routes.ShowUser).Methods("GET")
    router.HandleFunc("/edit", routes.EditUser).Methods("PUT")
    router.HandleFunc("/delete", routes.DeleteUser).Methods("DELETE")
    router.HandleFunc("/showall", routes.ShowAllUser).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}
