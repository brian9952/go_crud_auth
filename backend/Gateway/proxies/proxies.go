package proxies

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
    auth_api_key = os.Getenv("AUTH_API_SECRET")
    product_api_key = os.Getenv("PRODUCT_API_SECRET")
    url_this = os.Getenv("GATEWAY_URL")
    )

func addUrl(u *url.URL, r *http.Request) {
    r.URL.Scheme = u.Scheme
    r.URL.Host = u.Host
    r.URL.Path = u.Path
    r.Host = u.Host
}

func httpToWs(url *string) {
  
}

func generateToken(key []byte, url_from string, url_to string) (string, error) {

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["authorized"] = true
    claims["url_from"] = url_from
    claims["url_to"] = url_to
    claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

    tokenString, err := token.SignedString(key)
    if err != nil {
        log.Default().Println("Something went wrong")
        log.Default().Println(err.Error())
        return "", err
    }

    return tokenString, nil
}

//func LoginHandler(w http.ResponseWriter, r *http.Request) {
//    url_str := os.Getenv("AUTH_URL") + "/login"
//    url, err := url.Parse(url_str)
//    if err != nil {
//        panic("Error when parsing")
//    }
//
//    jwt_key := []byte(auth_api_key)
//    token, _ := generateToken(jwt_key, url_this, url_str)
//
//    proxy := httputil.ReverseProxy{Director: func(r *http.Request){
//        addUrl(url, r)
//        r.Header.Add("API-Token", token)
//    }}
//    proxy.ServeHTTP(w, r)
//}
//
//func RegisterHandler(w http.ResponseWriter, r *http.Request) {
//    url_str := os.Getenv("AUTH_URL") + "/register"
//    url, err := url.Parse(url_str)
//    if err != nil {
//        panic("Error when parsing")
//    }
//
//    jwt_key := []byte(auth_api_key)
//    token, _ := generateToken(jwt_key, url_this, url_str)
//
//    proxy := httputil.ReverseProxy{Director: func(r *http.Request){
//        addUrl(url, r)
//        r.Header.Add("API-Token", token)
//    }}
//    proxy.ServeHTTP(w, r)
//}

func sliceURL(url string, temp_str string) string {
    var newStr string
    var idx int
    tempStr := temp_str
    //tempStr := "product/"
    // search index
    for i := 0; i < len(url); i++ {
        for j := i + len(tempStr); j < len(url); j++ {
            if url[i:j] == tempStr {
                idx = j
            }
        }
    }

    // get string
    newStr = url[idx:len(url)]
    return newStr
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
    // split url string
    rawUrl := r.URL.Path
    path := sliceURL(rawUrl, "auth/")

    // append string
    url_str := os.Getenv("AUTH_URL") + "/" + path
    url, err := url.Parse(url_str)
    if err != nil {
        panic("Error when parsing")
    }

    jwt_key := []byte(auth_api_key)
    token, _ := generateToken(jwt_key, url_this, url_str)

    proxy := httputil.ReverseProxy{Director: func(r *http.Request){
        addUrl(url, r)
        r.Header.Add("API-Token", token)
    }}
    proxy.ServeHTTP(w, r)
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
    // split url string
    rawUrl := r.URL.Path
    path := sliceURL(rawUrl, "product/")

    // append string
    url_str := os.Getenv("PRODUCT_URL") + "/" + path
    url, err := url.Parse(url_str)
    if err != nil {
        panic("Error when parsing")
    }

    jwt_key := []byte(product_api_key)
    token, _ := generateToken(jwt_key, url_this, url_str)

    proxy := httputil.ReverseProxy{Director: func(r *http.Request){
        addUrl(url, r)
        r.Header.Add("API-Token", token)
    }}
    proxy.ServeHTTP(w, r)
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
    // split url string
    rawUrl := r.URL.Path
    path := sliceURL(rawUrl, "ws/")

    // append string
    url_str := os.Getenv("WEBSOCKET_URL") + path
    url, err := url.Parse(url_str)
    if err != nil {
        log.Default().Panic("Error when parsing")
    }

    proxy := httputil.ReverseProxy{Director: func(r *http.Request) {
        addUrl(url, r)
    }}
    proxy.ServeHTTP(w, r)
}
