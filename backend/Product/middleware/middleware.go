package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"main/libs"
	"net/http"
	"strconv"
	"time"

    "github.com/golang-jwt/jwt/v4"
	"github.com/felixge/httpsnoop"
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

// TODO: check auth api
func IsAuthorizedAPI(handler http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var err *libs.Status
        gatewayTokenStr := r.Header.Get("API-Token")

        if gatewayTokenStr == "" {
            err = libs.CreateErrorMessage("Error: Token not Found")
            json.NewEncoder(w).Encode(err)
            return
        }

        // get api secret key
        var apiKey = []byte(libs.Gateway_api_key)

        
        token, jwtErr := jwt.Parse(gatewayTokenStr, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Error in parsing token")
            }
            return apiKey, nil
        })

        // error handling
        if jwtErr != nil {
            err = libs.CreateErrorMessage("Error: API token is invalid")
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
        err = libs.CreateErrorMessage("Error: You are not authorized! (invalid API token)")
        json.NewEncoder(w).Encode(err)
    })
}

// TODO: check auth user
func IsAuthorizedUser(handler http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var err *libs.Status
        gatewayTokenStr := r.Header.Get("User-Token")

        if gatewayTokenStr == "" {
            err = libs.CreateErrorMessage("Error: Token not Found")
            json.NewEncoder(w).Encode(err)
            return
        }

        log.Default().Println("ERROR IN USER AUTHORIZATION")

        // get api secret key
        var apiKey = []byte(libs.Auth_api_key)

        
        token, jwtErr := jwt.Parse(gatewayTokenStr, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Error in parsing token")
            }
            return apiKey, nil
        })

        // error handling
        if jwtErr != nil {
            err = libs.CreateErrorMessage("Error: User token is invalid")
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
        err = libs.CreateErrorMessage("Error: You are not authorized! (invalid user token)")
        json.NewEncoder(w).Encode(err)
    })
}

