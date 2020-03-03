package rules

import (
        "context"
        "fmt"
        "net/url"
        "strings"
        "sync"
        "time"
        
        
        html_template "html/template"
        
        yaml "gopkg.in/yaml.v2"
        
        "github.com/go-kit/kit/log"
	      "github.com/go-kit/kit/log/level"
	      "github.com/pkg/errors"
	      "github.com/prometheus/common/model"
        
        
      	"github.com/prometheus/prometheus/pkg/labels"
      	"github.com/prometheus/prometheus/pkg/rulefmt"
      	"github.com/prometheus/prometheus/pkg/timestamp"
      	"github.com/prometheus/prometheus/promql"
	      "github.com/prometheus/prometheus/promql/parser"
	      "github.com/prometheus/prometheus/template"
	      "github.com/prometheus/prometheus/util/strutil"
        
)

const (
	// AlertMetricName is the metric name for synthetic alert timeseries.
	alertMetricName = "ALERTS"
	// AlertForStateMetricName is the metric name for 'for' state of alert.
	alertForStateMetricName = "ALERTS_FOR_STATE"

	// AlertNameLabel is the label name indicating the name of an alert.
	alertNameLabel = "alertname"
	// AlertStateLabel is the label name indicating the state of an alert.
	alertStateLabel = "alertstate"
)
    
    
// AlertState denotes the state of an active alert.
type AlertState int


const (
	// StateInactive is the state of an alert that is neither firing nor pending.
	StateInactive AlertState = iota
	// StatePending is the state of an alert that has been active for less than
	// the configured threshold duration.
	StatePending
	// StateFiring is the state of an alert that has been active for longer than
	// the configured threshold duration.
	StateFiring
)


func (s AlertState) String() string {
	switch s {
	case StateInactive:
		return "inactive"
	case StatePending:
		return "pending"
	case StateFiring:
		return "firing"
	}
	panic(errors.Errorf("unknown alert state: %s", s.String()))
}



type Alert struct {
	State AlertState

	Labels      labels.Labels
	Annotations labels.Labels

	// The value at the last evaluation of the alerting expression.
	Value float64
	// The interval during which the condition of this alert held true.
	// ResolvedAt will be 0 to indicate a still active alert.
	ActiveAt   time.Time
	FiredAt    time.Time
	ResolvedAt time.Time
	LastSentAt time.Time
	ValidUntil time.Time
}


func (a *Alert) needsSending(ts time.Time, resendDelay time.Duration) bool {
	if a.State == StatePending {
		return false
	}

	// if an alert has been resolved since the last send, resend it
	if a.ResolvedAt.After(a.LastSentAt) {
		return true
	}

	return a.LastSentAt.Add(resendDelay).Before(ts)
}


