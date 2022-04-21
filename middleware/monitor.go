package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"kubernetes-admin-backend/config"
	"time"
)

var store = sessions.NewCookieStore([]byte("bzka"))

func monitor() gin.HandlerFunc {
	// 注册监控指标
	prometheus.MustRegister(config.WebActiveTotal)
	prometheus.MustRegister(config.WebRequestTotal)
	prometheus.MustRegister(config.WebRequestDuration)

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		//config.WebActiveTotal.Set(float64(r.Session.Size()))
		config.WebRequestTotal.With(prometheus.Labels{"method": c.Request.Method, "endpoint": c.Request.URL.Path}).Inc()
		config.WebRequestDuration.With(prometheus.Labels{"method": c.Request.Method, "endpoint": c.Request.URL.Path}).Observe(duration.Seconds())
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
