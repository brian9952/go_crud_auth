package main

import (
	"log"
	"main/middleware"
	"main/proxies"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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

    r.productRouter.PathPrefix("/").HandlerFunc(middleware.Logging(proxies.ProductHandler))
    //r.productRouter.HandleFunc("/{url:[+]}", middleware.Logging(proxies.ProductHandler))

    // create product handler
    // r.productRouter.HandleFunc("/create_product", middleware.Logging(proxies.ProductHandler))

    // show product handler
    // r.productRouter.HandleFunc("/show_products", middleware.Logging(proxies.ProductHandler))
}

func (r *Routers) createMainRouter() {
    r.mainRouter = mux.NewRouter()
}

func main() {
    var router *Routers = new(Routers)
    router.createMainRouter()
    router.createAuthRouter()
    router.createProductRouter()

    log.Default().Println("Service started at " + os.Getenv("GATEWAY_URL"))

    credentials := handlers.AllowCredentials()
    headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
    origins := handlers.AllowedOrigins([]string{"http://107.102.183.168:8081"})
    //origins := handlers.AllowedOrigins([]string{"http://107.102.183.168:8081"})
    err := http.ListenAndServe(":8081", handlers.CORS(headers, credentials, methods, origins)(router.mainRouter))

    if err != nil {
        log.Default().Println("Failed to start service")
        log.Fatal(err)
    }
}
