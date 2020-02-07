 import "github.com/prometheus/client_golang/prometheus"

var jobDurationHistogram = prometheus.NewHistogramVec(
    prometheus.HistorgramOpts{
     Name:   "jobs_duration_seconds",
     Help:   "Jobs_duration distribution",
     Buckets:  []float64{1,2,5,10,20,60},
    },
 []string{"job_type"},
)

start := time.Now()
job.Run()
duration := time.Since(start)
jobsDurationHistogram.WithLabelValues(job.Type()).Observe(duration.Seconds())

prometheus.NewSummary()


  
