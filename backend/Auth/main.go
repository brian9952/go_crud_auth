package main

import (
	"log"
	"main/database"
	"main/middleware"
	"main/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func createRouter() *mux.Router {
    mainRouter := mux.NewRouter()

    // handlers
    mainRouter.HandleFunc("/login",
        middleware.IsAuthorized(middleware.Logging(routes.LoginUser)))

    mainRouter.HandleFunc("/register",
        middleware.IsAuthorized(middleware.Logging(routes.RegisterUser)))

    return mainRouter
}

func startServer(r *mux.Router) {
    log.Default().Println("Server started at http://localhost:8082")
    credentials := handlers.AllowCredentials()
    headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
    origins := handlers.AllowedOrigins([]string{"http://107.102.183.168:8081"})
    //origins := handlers.AllowedOrigins([]string{"*"})
    err := http.ListenAndServe(":8082", handlers.CORS(headers, credentials, methods, origins)(r))

    // start
    if err != nil {
        log.Default().Println("Failed to start server!")
        log.Fatal(err)
    }
}

func main() {
    // database init
    database.CreateDBConnection()

    mainRouter := createRouter()
    startServer(mainRouter)
}
