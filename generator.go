package main

import (
	"time"

	"github.com/cyriltovena/loki-log-generator/flog"
	"github.com/prometheus/common/model"
)

type LogMessage struct {
	Message string
	Level   model.LabelValue
}

type LogGenerator func(t time.Time) LogMessage

var Generators = []LogGenerator{
	func(t time.Time) LogMessage {
		return LogMessage{
			Message: flog.NewApacheCommonLog(t),
			Level:   INFO,
		}
	},
	func(t time.Time) LogMessage {
		return LogMessage{
			Message: flog.NewApacheCombinedLog(t),
			Level:   INFO,
		}
	},
	func(t time.Time) LogMessage {
		return LogMessage{
			Message: flog.NewApacheErrorLog(t),
			Level:   INFO,
		}
	},
	func(t time.Time) LogMessage {
		return LogMessage{
			Message: flog.NewRFC3164Log(t),
			Level:   INFO,
		}
	},
	func(t time.Time) LogMessage {
		return LogMessage{
			Message: flog.NewRFC5424Log(t),
			Level:   INFO,
		}
	},
	func(t time.Time) LogMessage {
		return LogMessage{
			Message: flog.NewCommonLogFormat(t),
			Level:   INFO,
		}
	},
	func(t time.Time) LogMessage {
		return LogMessage{
			Message: flog.NewJSONLogFormat(t),
			Level:   INFO,
		}
	},
}

// todo logfmt otel, json, errors, popular logs....

// ts=2024-02-28T23:04:12.760535253Z caller=http.go:194 level=debug traceID=279a41b78e22cad1 orgID=1218 msg="POST /ingester.v1.IngesterService/Push (200) 1.134561ms"
const HttpLogFmtFormat = `ts=%s caller=http.go:194 level=%s traceID=279a41b78e22cad1 orgID=1218 msg="POST /ingester.v1.IngesterService/Push (200) 1.134561ms"`
