package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"time"
)

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

// Monitor 监控
func Monitor() gin.HandlerFunc {
	// 注册监控指标
	prometheus.MustRegister(WebActiveTotal)
	prometheus.MustRegister(WebRequestTotal)
	prometheus.MustRegister(WebRequestDuration)

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		//WebActiveTotal.Set(float64(r.Session.Size()))
		WebRequestTotal.With(prometheus.Labels{"method": c.Request.Method, "endpoint": c.Request.URL.Path}).Inc()
		WebRequestDuration.With(prometheus.Labels{"method": c.Request.Method, "endpoint": c.Request.URL.Path}).Observe(duration.Seconds())
	}

}

func ExamplePusher_Push() {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})
	completionTime.SetToCurrentTime()
	if err := push.New("http://pushgateway:9091", "db_backup").
		Collector(completionTime).
		Grouping("db", "customers").
		Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}

var (
	completionTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last completion of a DB backup, successful or not.",
	})
	successTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_success_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})
	duration = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_duration_seconds",
		Help: "The duration of the last DB backup in seconds.",
	})
	records = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_records_processed",
		Help: "The number of records processed in the last DB backup.",
	})
)

func performBackup() (int, error) {
	// Perform the backup and return the number of backed up records and any
	// applicable error.
	// ...
	return 42, nil
}

func ExamplePusher_Add() {
	// We use a registry here to benefit from the consistency checks that
	// happen during registration.
	registry := prometheus.NewRegistry()
	registry.MustRegister(completionTime, duration, records)
	// Note that successTime is not registered.

	pusher := push.New("http://pushgateway:9091", "db_backup").Gatherer(registry)

	start := time.Now()
	n, err := performBackup()
	records.Set(float64(n))
	// Note that time.Since only uses a monotonic clock in Go1.9+.
	duration.Set(time.Since(start).Seconds())
	completionTime.SetToCurrentTime()
	if err != nil {
		fmt.Println("DB backup failed:", err)
	} else {
		// Add successTime to pusher only in case of success.
		// We could as well register it with the registry.
		// This example, however, demonstrates that you can
		// mix Gatherers and Collectors when handling a Pusher.
		pusher.Collector(successTime)
		successTime.SetToCurrentTime()
	}
	// Add is used here rather than Push to not delete a previously pushed
	// success timestamp in case of a failure of this backup.
	if err := pusher.Add(); err != nil {
		fmt.Println("Could not push to Pushgateway:", err)
	}
}
