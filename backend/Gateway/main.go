package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func createProxy(host string) (*httputil.ReverseProxy, error) {
    url, err := url.Parse(host)
    if err != nil {
        return nil, err
    }

    proxy := httputil.NewSingleHostReverseProxy(url)

    originalDirector := proxy.Director
    proxy.Director = func(r *http.Request) {
        originalDirector(r)
        modifyResponse(r)
    }

    proxy.ModifyResponse = modifyResponse()
    return proxy, nil
}

func modifyResponse(){

}

func main() {

}
