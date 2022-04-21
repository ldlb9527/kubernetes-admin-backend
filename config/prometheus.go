package config

import "github.com/prometheus/client_golang/prometheus"

var WebRequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "web_request_total",
		Help: "Number of requests in total",
	},
	// 设置两个标签 请求方法和 路径 对请求总次数在两个
	[]string{"method", "endpoint"},
)

var WebRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "web_request_duration_seconds",
		Help:    "web request duration distribution",
		Buckets: []float64{0.1, 0.3, 0.5, 0.7, 0.9, 1.5},
	},
	[]string{"method", "endpoint"},
)

var WebActiveTotal = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "web_active_total",
		Help: "web active total",
	})
