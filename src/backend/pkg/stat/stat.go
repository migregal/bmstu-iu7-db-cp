package stat

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Counter struct {
	prometheus.Counter
}

func NewCounter(namespace, name, help string) Counter {
	return Counter{promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})}
}

type HistogramVec struct {
	*prometheus.HistogramVec
}

func NewHistogramVec(namespace, name, help string, labelNames []string) HistogramVec {
	buckets := []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}

	return HistogramVec{
		promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      name,
			Help:      help,
			Buckets:   buckets,
		}, labelNames),
	}
}
