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

func createRouters() *mux.Router {
    router := mux.NewRouter()

    // handlers
    router.HandleFunc("/create_product",
        middleware.Logging(middleware.IsAuthorized(routes.CreateProduct)))

    router.HandleFunc("/show_products",
        middleware.Logging(routes.ShowProduct))

    return router
}

func startServer(r *mux.Router)  {
    log.Default().Println("Server started at http://localhost:8083")
    credentials := handlers.AllowCredentials()
    headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
    origins := handlers.AllowedOrigins([]string{"http://107.102.183.168:8081"})
    //origins := handlers.AllowedOrigins([]string{"*"})
    err := http.ListenAndServe(":8083", handlers.CORS(headers, credentials, methods, origins)(r))

    // start
    if err != nil {
        log.Default().Println("Failed to start server!")
        log.Fatal(err)
    }
}

func main() {
    database.CreateDBConnection()

    mainRouter := createRouters()
    startServer(mainRouter)
}
