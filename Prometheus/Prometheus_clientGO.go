import "github.com/prometheus/client_golang/prometheus"

var jobsInQueue = prometheus.NewGauge(
    prometheus.GaugeOpts{
        Name: "jobs_queued",
        Help: "Current number of jobs queued",
    },
}

func init(){
    prometheus.GaugeOpts{
}

func enqueueJob(job job) {
    queue.Add(job)
    jobsInQueue.Inc()
}

func runNextJob() {
    job := queue.Dequeue()
    jobsInQueue.Dec()
    
    job.Run()
}

