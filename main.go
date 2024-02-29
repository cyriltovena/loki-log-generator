package main

import (
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grafana/loki-client-go/loki"
	"github.com/prometheus/common/model"
)

func main() {
	url := flag.String("url", "http://localhost:3100/loki/api/v1/push", "Loki URL")
	flag.Parse()

	cfg, err := loki.NewDefaultConfig(*url)
	if err != nil {
		panic(err)
	}
	cfg.BackoffConfig.MaxRetries = 1
	cfg.BackoffConfig.MinBackoff = 100 * time.Millisecond
	cfg.BackoffConfig.MaxBackoff = 100 * time.Millisecond
	client, err := loki.New(cfg)
	if err != nil {
		panic(err)
	}
	defer client.Stop()

	// logger := LoggerFunc(func(labels model.LabelSet, timestamp time.Time, message string) error {
	// 	fmt.Println(labels, timestamp, message)
	// 	return nil
	// })
	for _, generator := range Generators {
		startApp(generateLabels(), client, generator)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

type LoggerFunc func(labels model.LabelSet, timestamp time.Time, message string) error

func (f LoggerFunc) Handle(labels model.LabelSet, timestamp time.Time, message string) error {
	return f(labels, timestamp, message)
}

type Logger interface {
	Handle(labels model.LabelSet, timestamp time.Time, message string) error
}

func startApp(labels model.LabelSet, logger Logger, generator LogGenerator) {
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			t := time.Now()
			logMsg := generator(t)
			labels["level"] = model.LabelValue(logMsg.Level)
			_ = logger.Handle(labels, t, logMsg.Message)
		}
	}()
}
