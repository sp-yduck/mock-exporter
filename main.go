package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metricConfig struct {
	Name  string `yaml:"metricName"`
	Value int64  `yaml:"metricValue"`
	Help  string `yaml:"metricHelp"`
}

func readConfig(filename string) (*string, *int64, *string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var config metricConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config.Name, &config.Value, &config.Help
}

func main() {
	port := flag.Int64("port", 8080, "port to expose metrics on")
	flag.Parse()

	// create new metrics
	metricName, metricValue, metricHelp := readConfig("exporter-config.yml")
	metric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: *metricName,
			Help: *metricHelp,
		},
	)
	// register a new metrics
	prometheus.MustRegister(metric)

	// set vale
	metric.Set(float64(*metricValue))

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
