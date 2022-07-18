package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"main/libs"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/felixge/httpsnoop"
)

var (
    this_url = os.Getenv("PRODUCT_URL")
    )

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

type Data struct {
    Data string `json:"data"`
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

func checkIntegrity(claims jwt.MapClaims) bool {
    url_from := "http://107.102.183.168:8081"
    if claims["authorized"] == true && claims["url_from"] == url_from {
        return true
    }
    return false
}

// TODO: check auth
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var err *libs.Status
        tokenStr := r.Header.Get("API-Token")

        if tokenStr == "" {
            err = libs.CreateErrorMessage("Error: Token not Found")
            json.NewEncoder(w).Encode(err)
            return
        }

        // get secret key
        var key = []byte(libs.Private_Key)
        
        token, jwtErr := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Error in parsing token")
            }
            return key, nil
        })

        // error handling
        if jwtErr != nil {
            err = libs.CreateErrorMessage("Error: Token is invalid")
            json.NewEncoder(w).Encode(err)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if ok && token.Valid {
            if checkIntegrity(claims) {
                r.Header.Set("Authorized", "1")
                handler.ServeHTTP(w, r)
                return
            }
        }

        // not authorized
        err = libs.CreateErrorMessage("Error: You are not authorized!")
        json.NewEncoder(w).Encode(err)
    })
}

