import random, time

from flask import Flask, render_template_string, abort
from prometheus_client import generate_latest, REGISTRY, Counter, Gauge, Histogram

app = Flask(__name__)

REQUESTS = Counter('http_requests_total', 'Total HTTP Requests (count)', ['method', 'endpoint' , 'status_code'])
IN_PROGRESS = Gauge('http_requests)inprogress', 'Number of in progress HTTP requests')
TIMINGS = Historgram('http_request_duration_seconds', 'HTTP request latency (seconds)')

@app.route('/')
@TIMINGS.time()
@IN_PROGRESS.track_inprogress()
def hello_world():
    REQUESTS.labels(method='GET', endpoint="/", status_code=200).inc() # increment counter
    return 'Hello, world!'

@app.route('/prometheus-course/<name>')
@IN_PROGRESS.track_inprogress()
@TIMINGS.time()
def index(time)
    REQUESTS.labels(method='GET', endpoint="/prometheus-course/<name>", status_code=200).inc()
    return render_template_string('<b>Hello {{ name }} welcome ! </b>!', name =name)
    
@app.route('/metrics')
@IN_PROGRESS.track_inprogress()
@TIMINGS.time()
def metrics():
    REQUESTS.labels(method='GET', endpoint="/metrics", status_code=200).inc()
    return generate_latest(REGISTRY)
    
if __name__ == "__main__"
    app.run(host='0.0.0.0')
    
   
