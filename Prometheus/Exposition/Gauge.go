import "github.com/prometheus/client_golang/prometheus"

var jobsQueued = prometheus.NewGauge(
    prometheus.GaugeOpts{
        Name: "jobs_queued",
        Help: "Current number of jobs queued",
    },
)

func init(){
    prometheus.MustRegister(jobsQueued)
}

func enqueueJob(job Job) {
    queue.Add(job)
    jobsQueued.Inc()
}
        
func runNextJob() {
    job := queue.Dequeue()
    jobsInQueue.Dec()
    
    job.Run()
    
}



