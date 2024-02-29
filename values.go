package main

import (
	"math/rand"
	"strings"
	"time"

	gofakeit "github.com/brianvoe/gofakeit/v7"
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

var orgID = []string{randSeq(4), randSeq(4), randSeq(4), randSeq(4), randSeq(4)}

func randLevel() model.LabelValue {
	r := rand.Intn(100)
	if r < 10 {
		return ERROR
	} else if r < 15 {
		return WARN
	} else {
		return level[rand.Intn(len(level)-2)]
	}
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

func randOrgID() string {
	return orgID[rand.Intn(len(orgID))]
}

func randError() string {
	switch rand.Intn(10) {
	case 0:
		return gofakeit.ErrorDatabase().Error()
	case 1:
		return gofakeit.ErrorGRPC().Error()
	case 2:
		return gofakeit.ErrorObject().Error()
	case 3:
		return gofakeit.ErrorRuntime().Error()
	case 4:
		return gofakeit.ErrorHTTP().Error()
	default:
		return gofakeit.Error().Error()
	}
}

var filesNames = []string{gofakeit.ProductName(), gofakeit.ProductName(), gofakeit.ProductName(), gofakeit.Word(), gofakeit.Word()}

func randFileName() string {
	return strings.ReplaceAll(strings.ToLower(filesNames[rand.Intn(len(filesNames))]), " ", "_")
}

func randDuration() string {
	return (time.Duration(gofakeit.Number(1, 30000)) * time.Millisecond).String()
}
