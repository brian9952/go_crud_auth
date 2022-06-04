package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	l "main/libs"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/felixge/httpsnoop"
)

type Claims struct {
    authorized string
    username string
    role string
    exp string
}

type Log struct {
    method string
    uri string
    referer string
    ipaddr string

    code int
    size int64

    duration time.Duration
    userAgent string
}

func printLog(logStruct *Log) {
    var log_str string
    log_str = logStruct.method + " | " 
    log_str = log_str + logStruct.uri + " | "
    log_str = log_str + logStruct.referer + " | "
    log_str = log_str + logStruct.ipaddr + " | "
    log_str = log_str + strconv.Itoa(logStruct.code) + " | " 
    log_str = log_str + strconv.Itoa(int(logStruct.size)) + " | "
    log_str = log_str + logStruct.duration.String() + " | " 
    log_str = log_str + logStruct.userAgent
    log.Default().Println(log_str)
}

// logging middleware
func Logging(handler http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        logStruct := &Log{
            method: r.Method,
            uri: r.URL.String(),
            referer: r.Header.Get("Referer"),
            userAgent: r.Header.Get("User-Agent"),
        }

        logStruct.ipaddr = r.RemoteAddr

        m := httpsnoop.CaptureMetrics(handler, w, r)

        logStruct.code = m.Code
        logStruct.size = int64(m.Written)
        logStruct.duration = m.Duration

        // print logging
        printLog(logStruct)

    })
}

// check if user is authorized
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var err *l.Error 

        w.Header().Set("Content-type", "application/json")

        if r.Header["Token"] == nil {
            err = l.CreateError("token_404", "Token not found")
            json.NewEncoder(w).Encode(err)
            return
        }

        // get secret key & token
        var key = []byte(os.Getenv("CRUD_SECRETKEY"))

        // jwt parsing
        token, jwtErr := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Error in parsing token")
            }
            return key, nil
        })

        // error handling
        if jwtErr != nil {
            err = l.CreateError("expired_token", "Your token has been expired")
            json.NewEncoder(w).Encode(err)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims);
        if ok && token.Valid {
            if claims["role"] == "admin" {
                r.Header.Set("Role", "Admin")
                handler.ServeHTTP(w, r)
                return
            } else if claims["role"] == "user" {
                r.Header.Set("Role", "User")
                handler.ServeHTTP(w, r)
                return
            }
        } 

        // not authorized
        err = l.CreateError("not_authorized", "You are not authorized!")
        json.NewEncoder(w).Encode(err)

    })
}
