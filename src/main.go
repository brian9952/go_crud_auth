package main

import (
    "log"
    "net/http"

    "main/database"
    "main/models"
    "main/routes"

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
    r.authRouter.HandleFunc("/register").Methods("POST")
    r.authRouter.HandleFunc("/login").Methods("POST")

    // router for products management
    r.productRouter.HandleFunc("/create").Methods("POST")
    r.productRouter.HandleFunc("/update").Methods("PUT")
    r.productRouter.HandleFunc("/showproduct").Methods("GET")
    r.productRouter.HandleFunc("/showallproduct").Methods("GET")
    r.productRouter.HandleFunc("/deleteproduct").Methods("DELETE")
}

func main() {
    var r_ptr *routers = new(routers)

    // create mux router
    r_ptr.createRouter()

    // create routers
    r_ptr.createFunctions()

}
