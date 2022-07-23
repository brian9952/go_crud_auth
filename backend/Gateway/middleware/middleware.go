package middleware

import (
    "time"
    "strconv"
    "log"
    "net/http"

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

