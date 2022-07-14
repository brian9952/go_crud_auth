package main

import (
	"log"
	"main/middleware"
	"main/proxies"
	"net/http"

	"github.com/gorilla/mux"
)

type Routers struct {
    mainRouter *mux.Router
    authRouter *mux.Router
    productRouter *mux.Router
}

func (r *Routers) createAuthRouter() {
    r.authRouter = r.mainRouter.PathPrefix("/v1/api/auth").Subrouter()

    // login handler
    r.authRouter.HandleFunc("/login", middleware.Logging(proxies.LoginHandler))

    // register handler
    r.authRouter.HandleFunc("/register", middleware.Logging(proxies.RegisterHandler))
}

func (r *Routers) createProductRouter() {
    r.productRouter = r.mainRouter.PathPrefix("/v1/api/product").Subrouter()

    // create product handler
    r.productRouter.HandleFunc("/create_product", middleware.Logging(proxies.AddProductHandler))

    // show product handler
    r.productRouter.HandleFunc("/show_product", middleware.Logging(proxies.ShowProductHandler))
}

func createMainRouter() *mux.Router {
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
