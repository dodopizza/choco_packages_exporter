package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func sanitizePromLabelName(str string) string {
	re := regexp.MustCompile(`[\.\-]`)
	result := re.ReplaceAllString(str, "_")
	re = regexp.MustCompile(`^\d`)
	result = re.ReplaceAllString(result, "_$0")
	return result
}

func updateMetric(prometheusCounter *prometheus.Counter) {
	svclog.Info(1, "Preparing metric labels")
	promLabels := make(prometheus.Labels)
	for _, p := range chocoPackages.GetPackages() {
		promLabels[sanitizePromLabelName(p.GetName())] = p.GetVersion()
	}
	svclog.Info(1, "Updating counter object")

	defer func() {
		if err := recover(); err != nil {
			svclog.Error(1, "panic occurred:", err)

			for k, v := range promLabels {
				svclog.Info(1, fmt.Sprintf("key: %s, value: %s", k, v))
			}

			panic("Panic on creating new Prometheus counter")
		}
	}()
	*prometheusCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "winserver_choco_packages",
		Help:        "Chocolatey packages presents on the machine",
		ConstLabels: promLabels,
	})
}

func serveMetrics() {
	promRegistry := prometheus.NewRegistry()
	var promCounter prometheus.Counter

	go func() {
		for {
			updateMetric(&promCounter)
			svclog.Info(1, "Unregistering counter")
			promRegistry.Unregister(promCounter)
			prometheus.DefaultRegisterer.Unregister(promCounter)
			svclog.Info(1, "Registering counter")
			promRegistry.Register(promCounter)
			svclog.Info(1, "Waiting 1 minute for the next updates")
			time.Sleep(10 * time.Second)
		}
	}()

	svclog.Info(1, "Setup metrics output to "+appConfig.getHttpAddress()+"/metrics")
	handler := promhttp.HandlerFor(promRegistry, promhttp.HandlerOpts{})
	http.Handle("/metrics", handler)

	go func() {
		svclog.Info(1, fmt.Sprintf("Starting server on %s", appConfig.getHttpAddress()))
		svclog.Error(1, fmt.Sprintf("Cannot start windows_exporter: %s", http.ListenAndServe(appConfig.getHttpAddress(), nil)))
	}()
}
