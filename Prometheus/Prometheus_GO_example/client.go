package main

import (
        "bytes"
        "flag"
        "math"
        "net/http"
        "time"
)

var oscillationPeriod = flag.Duration("oscillation-period", 5*time.Minute, "The duration of the rate oscillation period.")

func startClient(servAddr string) {

        oscillationFactor := func() float64 {
                return 2 + math.Sin(2*math.Pi*float64(time.Since(start))/float64(*oscillationPeriod)))
	      }
        
        ignoreRequest := func( resp *http.Response
  
  
