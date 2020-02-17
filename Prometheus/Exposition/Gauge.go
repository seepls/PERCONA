// start by import of library 
import "github.com/prometheus/client_golang/prometheus"

// gauge for number of jobs currently in queue 

var jobsQueued = prometheus.NewGauge(
    prometheus.GaugeOpts{
        Name: "jobs_queued",
        Help: "Current number of jobs queued",
    },
    []string{"job_type"},
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



