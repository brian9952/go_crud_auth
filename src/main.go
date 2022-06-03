package main

import (
	"main/database"
	//"main/models"
	"log"
	"main/routes"
	"net/http"

	"github.com/gorilla/mux"
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
    r.authRouter.HandleFunc("/register", routes.RegisterUser).Methods("POST")
    r.authRouter.HandleFunc("/login", routes.LoginUser).Methods("POST")

    // router for products management
    //r.productRouter.HandleFunc("/create").Methods("POST")
    //r.productRouter.HandleFunc("/update").Methods("PUT")
    //r.productRouter.HandleFunc("/showproduct").Methods("GET")
    //r.productRouter.HandleFunc("/showallproduct").Methods("GET")
    //r.productRouter.HandleFunc("/deleteproduct").Methods("DELETE")
}

func (r *routers) serverStart() {
    log.Default().Println("Server started at http://localhost:8080")
    err := http.ListenAndServe(":8080", r.mainRouter)
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
