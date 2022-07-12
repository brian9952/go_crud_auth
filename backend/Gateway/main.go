package main

import (
	"log"
	"main/middleware"
	"main/proxies"
	"net/http"

	"github.com/gorilla/mux"
)

func createRouter() *mux.Router {
    mainRouter := mux.NewRouter()
    mainRouter = mainRouter.PathPrefix("/v1/api/auth").Subrouter()

    // login handler
    mainRouter.HandleFunc("/login", middleware.Logging(proxies.LoginHandler))

    // register handler
    mainRouter.HandleFunc("/register", middleware.Logging(proxies.RegisterHandler))

    return mainRouter
}

func main() {
    router := createRouter()

    log.Default().Println("Service started at http://107.102.183.168:8081")
    err := http.ListenAndServe(":8081", router)

    if err != nil {
        log.Default().Println("Failed to start service")
        log.Fatal(err)
    }
}
