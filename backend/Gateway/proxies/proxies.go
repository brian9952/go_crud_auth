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
    auth_secret_key = os.Getenv("AUTH_API_SECRET")
    product_secret_key = os.Getenv("PRODUCT_API_SECRET")
    )

func addUrl(u *url.URL, r *http.Request) {
    r.URL.Scheme = u.Scheme
    r.URL.Host = u.Host
    r.URL.Path = u.Path
    r.Host = u.Path
}

func generateToken(url_from string, url_to string) (string, error) {
    var key = []byte(auth_secret_key)

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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    url_str := "http://107.102.183.168:8082/login"
    url, err := url.Parse(url_str)
    if err != nil {
        panic("Error when parsing")
    }

    token, _ := generateToken("http://107.102.183.168:8081/v1/api/auth", url_str)

    proxy := httputil.ReverseProxy{Director: func(r *http.Request){
        addUrl(url, r)
        r.Header.Add("api_token", token)
    }}
    proxy.ServeHTTP(w, r)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    url, err := url.Parse("http://107.102.183.168:8082/register")
    if err != nil {
        panic("Error when parsing")
    }

    proxy := httputil.ReverseProxy{Director: func(r *http.Request){
        addUrl(url, r)
    }}
    proxy.ServeHTTP(w, r)
}
