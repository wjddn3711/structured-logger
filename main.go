package main

import (
	"context"

	"github.com/ggwhite/go-masker"
	jsoniter "github.com/json-iterator/go"
	"github.com/wjddn3711/structured-logger/logger"
	"github.com/wjddn3711/structured-logger/logger/options"
	"github.com/wjddn3711/structured-logger/logger/types"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type logEntry struct {
	StartTime   string `json:"start_time,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	Elapsed     int64  `json:"elapsed,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	URI         string `json:"uri,omitempty"`
	Referer     string `json:"referer,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty" mask:"mobile"`
}

// ToFields converts logEntry to map[string]interface{}
func (le logEntry) ToFields() map[string]interface{} {
	entryMap := make(map[string]interface{})
	t, _ := masker.Struct(le)
	entryBytes, _ := json.Marshal(t)
	_ = json.Unmarshal(entryBytes, &entryMap)

	return entryMap
}

func main() {
	ctx := context.Background()
	log := logger.NewWrapper(types.ZeroLog)
	ctx = log.WithContext(ctx)

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

	log.Info(options.WithMessage("debug message"), options.WithFields(entry))

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	log := logger.FromContext(ctx, types.ZeroLog)

	entry := logEntry{
		Elapsed:     1000,
		StatusCode:  500,
		URI:         "/",
		Referer:     "https://www.naver.com",
		PhoneNumber: "66666666",
	}

	log.Info(options.WithMessage("error message"), options.WithFields(entry))
}
