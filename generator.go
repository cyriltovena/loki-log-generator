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
	NewLogFmtApp,
	NewMysqlLogFmt,
}

// todo logfmt otel, json, errors, popular logs....
