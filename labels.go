package main

import (
	"math/rand"

	"github.com/prometheus/common/model"
)

var clusters = []string{
	"us-west-1",
	"us-east-1",
	"us-east-2",
	"eu-west-1",
}

var namespaces = []string{
	"prod",
	"dev",
	"staging",
	"infra",
	"monitoring",
}

var services = []string{
	"api",
	"web",
	"db",
	"cache",
	"queue",
	"worker",
	"cart",
	"checkout",
	"payment",
	"shipping",
	"order",
}

const (
	INFO  = model.LabelValue("info")
	ERROR = model.LabelValue("error")
	WARN  = model.LabelValue("warn")
	DEBUG = model.LabelValue("debug")
)

var level = []model.LabelValue{
	DEBUG,
	INFO,
	WARN,
	ERROR,
}

func randLevel() model.LabelValue {
	return level[rand.Intn(len(level))]
}

func generateLabels() model.LabelSet {
	svc := model.LabelValue(services[rand.Intn(len(services))])
	return model.LabelSet{
		"cluster":   model.LabelValue(clusters[rand.Intn(len(clusters))]),
		"namespace": model.LabelValue(namespaces[rand.Intn(len(namespaces))]),
		"service":   svc,
		"pod":       svc + "-" + model.LabelValue(randSeq(5)),
	}
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
