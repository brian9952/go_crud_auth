package main

import (
	"main/database"
	//"main/models"
	"log"
	"main/routes"
    "main/middleware"
	"net/http"

	"github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

type routers struct {
    mainRouter *mux.Router
    authRouter *mux.Router
    productRouter *mux.Router
}

func (r *routers) createRouter() {
    r.mainRouter = mux.NewRouter()
    r.authRouter = r.mainRouter.PathPrefix("/auth").Subrouter()
    r.productRouter = r.mainRouter.PathPrefix("/product").Subrouter()
}

func (r *routers) createFunctions() {
    // router for authentication functions
    r.authRouter.HandleFunc("/register", 
        middleware.Logging(routes.RegisterUser)).Methods("POST")

    r.authRouter.HandleFunc("/login", 
        middleware.Logging(routes.LoginUser)).Methods("POST")

    // router for products management
    r.productRouter.HandleFunc("/create", 
        middleware.Logging(middleware.IsAuthorized(routes.CreateProduct))).Methods("POST")

    r.productRouter.HandleFunc("/show",
        middleware.Logging(middleware.IsAuthorized(routes.ShowProduct))).Methods("GET")

    r.productRouter.HandleFunc("/update",
        middleware.Logging(middleware.IsAuthorized(routes.UpdateProduct))).Methods("POST")

    r.productRouter.HandleFunc("/delete",
        middleware.Logging(middleware.IsAuthorized(routes.DeleteProduct))).Methods("DELETE")

    r.productRouter.HandleFunc("/showall",
        middleware.Logging(middleware.IsAuthorized(routes.ShowAllProduct))).Methods("GET")

}

func (r *routers) serverStart() {
    log.Default().Println("Server started at http://localhost:8081")
    credentials := handlers.AllowCredentials()
    headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
    origins := handlers.AllowedOrigins([]string{"http://107.102.183.168:8080"})
    err := http.ListenAndServe(":8081", handlers.CORS(headers, credentials, methods, origins)(r.mainRouter))
    if err != nil {
        log.Default().Println("Failed to start server!")
        log.Fatal(err)
    }
}

func main() {
    var r_ptr *routers = new(routers)

    // create database connection and migrate
    database.CreateDBConnection()
    database.AutoMigrateDB()

    // create mux router
    r_ptr.createRouter()

    // create routers
    r_ptr.createFunctions()

    // start server
    r_ptr.serverStart()
}
