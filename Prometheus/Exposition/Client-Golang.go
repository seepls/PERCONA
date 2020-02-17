package main

import (
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

func main(){
    http.Handle("/metrics",promhttp.Handler())
    panic(http.ListenAndServe(":8080",nil))
}

