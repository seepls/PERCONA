import "github.com/prometheus/client_golang/prometheus"

var jobsInQueue = prometheus.NewGauge(
    prometheus.GaugeOpts{
        Name: "jobs_queued",
        Help: "Current number of jobs queued",
    },
}


var jobsInQueue = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
        Name: "jobs_queued",
        Help: "Current number of jobs queued",
    },
    []string{"job_type"},
}
    
func init(){
    prometheus.GaugeOpts{
}

func enqueueJob(job job) {
    queue.Add(job)
    jobsInQueue.Inc()
    jobsQueued.WithLabelValues(job.type()).Inc()
}

func runNextJob() {
    job := queue.Dequeue()
    jobsInQueue.Dec()
    jobsQueued.WithLabelValues(job.type()).Dec()
    job.Run()
}

