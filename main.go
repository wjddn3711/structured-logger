package main

import (
	"context"

	"github.com/ggwhite/go-masker"
	jsoniter "github.com/json-iterator/go"
	"github.com/wjddn3711/structured-logger/logger"
	"github.com/wjddn3711/structured-logger/logger/types"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type logEntry struct {
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Elapsed     int64  `json:"elapsed"`
	StatusCode  int    `json:"status_code"`
	URI         string `json:"uri"`
	Referer     string `json:"referer"`
	PhoneNumber string `json:"phone_number" mask:"mobile"`
}

func (le logEntry) ToFields() map[string]interface{} {
	entryMap := make(map[string]interface{})
	t, _ := masker.Struct(le)
	entryBytes, _ := json.Marshal(t)
	_ = json.Unmarshal(entryBytes, &entryMap)

	return entryMap
}

func main() {
	ctx := context.Background()
	log := logger.NewLoggerWrapper(types.ZeroLog, ctx)

	entry := logEntry{
		StartTime:   "2021-01-01T00:00:00Z",
		EndTime:     "2021-01-01T00:00:01Z",
		Elapsed:     1000,
		StatusCode:  200,
		URI:         "/",
		Referer:     "https://www.google.com",
		PhoneNumber: "01012345678",
	}

	log.RegisterCommonField("rid", "1234")

	log.Debug(entry, logger.WithMessage("debug message"))
}
