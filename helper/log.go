package helper

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func InitializeLog() {
	logDir := Config.LogDir
	_ = os.Mkdir(logDir, os.ModePerm)

	f, err := os.OpenFile(logDir+Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetFlags(0)
	log.SetOutput(f)
}

type bodyLogWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (writer bodyLogWriter) Write(b []byte) (int, error) {
	return writer.body.Write(b)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		myBodyLogWriter := &bodyLogWriter{
			ResponseWriter: w,
			body:           &bytes.Buffer{},
		}

		next.ServeHTTP(myBodyLogWriter, r)

		statusCode := r.Response.StatusCode
		response := NoErrorsFound
		level := "INFO"
		if statusCode >= 400 {
			response = myBodyLogWriter.body.String()
			level = "ERROR"
		}
		
		data, err := json.Marshal(&LogStruct{
			Method:          r.Method,
			Level:           level,
			StatusCode:      strconv.Itoa(statusCode),
			Path:            r.URL.String(),
			UserAgent:       r.UserAgent(),
			RemoteIP:        r.RemoteAddr,
			ResponseTime:    time.Since(time.Now()).String(),
			Message:         http.StatusText(statusCode) + ": " + response,
			Version:         "1",
			CorrelationId:   uuid.New().String(),
			AppName:         Config.AppName,
			ApplicationHost: r.Host,
			LoggerName:      "",
			TimeStamp:       time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s\n", data)

		next.ServeHTTP(myBodyLogWriter, r)
	})
}

func LogEvent(level string, message interface{}) {

	data, err := json.Marshal(struct {
		TimeStamp string      `json:"@timestamp"`
		Level     string      `json:"level"`
		Message   interface{} `json:"message"`
	}{TimeStamp: time.Now().Format(time.RFC3339),
		Message: message,
		Level:   level,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", data)

}
