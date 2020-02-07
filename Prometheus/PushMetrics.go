package main
import (
    "flag"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/prometheus/client_golang/prometheus/push"
)

gatewayUrl:="http://localhost:9091/"

throughputGauge := prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "throughput",
        Help: "Throughput in Mbps",
})


throughputGauge.Set(800)

if err := push.Collectors(
        "throughput_job", push.HostnameGroupingKey(),
        gatewayUrl, throughputGauge
); err != nil {
    fmt.Println("Could not push completion   time to Pushgateway:", err)
}
 
