package logger_test

import (
	"bytes"
	"context"
	"sync"
	"testing"

	"github.com/ggwhite/go-masker"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/wjddn3711/structured-logger/logger"
	"github.com/wjddn3711/structured-logger/logger/options"
	"github.com/wjddn3711/structured-logger/logger/types"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func TestLogger(t *testing.T) {
	var (
		logType = types.ZeroLog
	)

	// logger func Test
	t.Run("WithMessage로 로깅 시, 로그 메시지 생성 테스트", func(t *testing.T) {
		// given
		captureWriter := &captureWriter{}
		zLog := logger.NewWrapper(
			logType,
			options.WithLevel(types.Debug),
			options.WithOutput(captureWriter),
			options.WithTimeFormat("2006-01-02 15:04:05"),
		)

		// when
		zLog.Debug(options.WithMessage("debug message"))

		// then
		entries := captureWriter.Map()
		assert.Equal(t, entries[types.MessageField], "debug message", "로그 메시지가 정확히 캡처되어야 합니다.")
	})

	t.Run("WithFields로 로깅 시, 로그 필드 생성 테스트", func(t *testing.T) {
		// given
		captureWriter := &captureWriter{}
		zLog := logger.NewWrapper(
			logType,
			options.WithLevel(types.Debug),
			options.WithOutput(captureWriter),
			options.WithTimeFormat("2006-01-02 15:04:05"),
		)

		// when
		entry := Example{
			StartTime:   "2021-01-01T00:00:00Z",
			EndTime:     "2021-01-01T00:00:01Z",
			Elapsed:     1000,
			StatusCode:  200,
			PhoneNumber: "01012345678",
		}

		// then
		zLog.Debug(options.WithFields(entry))

		// then
		entries := captureWriter.Map()
		assert.Equal(t, entries["start_time"], "2021-01-01T00:00:00Z", "로그 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["end_time"], "2021-01-01T00:00:01Z", "로그 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["elapsed"], float64(1000), "로그 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["status_code"], float64(200), "로그 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["phone_number"], "0101***5678", "로그 필드가 정확히 캡처되어야 합니다.")
	})

	t.Run("등록된 RequestID가 다른 곳에서도 로깅 되는지 테스트",
		func(t *testing.T) {
			// given
			captureWriter := &captureWriter{}
			zLog := logger.NewWrapper(
				logType,
				options.WithLevel(types.Debug),
				options.WithOutput(captureWriter),
				options.WithTimeFormat("2006-01-02 15:04:05"),
			)
			ctx := zLog.WithContext(context.Background())

			// doSomething debug로 로깅
			doSomething := func(ctx context.Context) {
				log := logger.FromContext(ctx, logType)
				log.Debug()
			}

			// when
			zLog.RegisterCommonField("rid", "1234")
			// context에 로거 등록
			ctx = zLog.WithContext(ctx)

			doSomething(ctx)

			// then
			entries := captureWriter.Map()
			assert.Equal(t, entries["rid"], "1234", "공통 필드가 정확히 캡처되어야 합니다.")
		},
	)

	t.Run("등록된 공통 필드가 다른 곳에서도 로깅 되는지 테스트", func(t *testing.T) {
		// given
		captureWriter := &captureWriter{}
		zLog := logger.NewWrapper(
			logType,
			options.WithLevel(types.Debug),
			options.WithOutput(captureWriter),
			options.WithTimeFormat("2006-01-02 15:04:05"),
		)
		ctx := zLog.WithContext(context.Background())

		// doSomething debug로 로깅
		doSomething := func(ctx context.Context) {
			log := logger.FromContext(ctx, logType)
			log.Debug()
		}

		// when
		entry := Example{
			StartTime:   "2021-01-01T00:00:00Z",
			EndTime:     "2021-01-01T00:00:01Z",
			Elapsed:     1000,
			StatusCode:  200,
			PhoneNumber: "01012345678",
		}
		zLog.RegisterCommonFields(entry)
		ctx = zLog.WithContext(ctx)

		doSomething(ctx)

		// then
		entries := captureWriter.Map()
		assert.Equal(t, entries["start_time"], "2021-01-01T00:00:00Z", "공통 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["end_time"], "2021-01-01T00:00:01Z", "공통 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["elapsed"], float64(1000), "공통 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["status_code"], float64(200), "공통 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["phone_number"], "0101***5678", "공통 필드가 정확히 캡처되어야 합니다.")
	})

	t.Run("엔트리가 등록 된 상태에서 새로운 엔트리로 로깅 되는지 테스트", func(t *testing.T) {
		// given
		captureWriter := &captureWriter{}
		zLog := logger.NewWrapper(
			logType,
			options.WithLevel(types.Debug),
			options.WithOutput(captureWriter),
			options.WithTimeFormat("2006-01-02 15:04:05"),
		)
		// doSomething debug로 로깅
		doSomething := func(ctx context.Context) {
			log := logger.FromContext(ctx, logType)
			entry := Example{
				StartTime:   "2022-01-01T00:00:00Z",
				EndTime:     "2022-01-01T00:00:01Z",
				Elapsed:     2000,
				StatusCode:  400,
				PhoneNumber: "01087654321",
			}
			log.Debug(options.WithFields(entry))
		}

		// when
		entry := Example{
			StartTime:   "2021-01-01T00:00:00Z",
			EndTime:     "2021-01-01T00:00:01Z",
			Elapsed:     1000,
			StatusCode:  200,
			PhoneNumber: "01012345678",
		}
		zLog.Debug(options.WithFields(entry))
		ctx := zLog.WithContext(context.Background())
		doSomething(ctx)

		// then
		entries := captureWriter.Map()
		assert.Equal(t, entries["start_time"], "2022-01-01T00:00:00Z", "로그 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["end_time"], "2022-01-01T00:00:01Z", "로그 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["elapsed"], float64(2000), "로그 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["status_code"], float64(400), "로그 필드가 정확히 캡처되어야 합니다.")
		assert.Equal(t, entries["phone_number"], "0108***4321", "로그 필드가 정확히 캡처되어야 합니다.")
	})
}

type Example struct {
	StartTime   string `json:"start_time,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	Elapsed     int64  `json:"elapsed,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty" mask:"mobile"`
}

func (e Example) ToFields() map[string]interface{} {
	entryMap := make(map[string]interface{})
	// masker.Struct(e)로 구조체 필드에 마스킹 처리
	t, _ := masker.Struct(e)
	entryBytes, _ := json.Marshal(t)
	_ = json.Unmarshal(entryBytes, &entryMap)

	return entryMap
}

type captureWriter struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (w *captureWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buf.Write(p)
}

func (w *captureWriter) String() string {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buf.String()
}

func (w *captureWriter) Map() map[string]interface{} {
	w.mu.Lock()
	defer w.mu.Unlock()
	str := w.buf.String()
	var m map[string]interface{}
	json.Unmarshal([]byte(str), &m)
	return m
}
