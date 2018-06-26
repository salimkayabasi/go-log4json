package go_log4json

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetFlags(0)
}

type logData struct {
	TS    time.Time `json:"ts"`
	Level string    `json:"level"`
	MSG   string    `json:"msg"`
}

type Logger struct{}

func (lgr *Logger) Log(msg, level string, a ...interface{}) {
	l := logData{
		TS:    time.Now(),
		Level: level,
		MSG:   fmt.Sprintf(msg, a...),
	}
	dt, err := json.Marshal(l)
	if err != nil {
		log.Println("error in marshaling log data", l)
		return
	}
	log.Println(string(dt))
}

//Debug just simplifies logging in that level
func (lgr *Logger) Debug(msg string, a ...interface{}) {
	lgr.Log(msg, "debug", a...)
}

//Info just simplifies logging in that level
func (lgr *Logger) Info(msg string, a ...interface{}) {
	lgr.Log(msg, "info", a...)
}

//Warn just simplifies logging in that level
func (lgr *Logger) Warn(msg string, a ...interface{}) {
	lgr.Log(msg, "warn", a...)
}

//Error just simplifies logging in that level
func (lgr *Logger) Error(msg string, a ...interface{}) {
	lgr.Log(msg, "error", a...)
}

//Fatal just simplifies logging in that level and also it kills the process
func (lgr *Logger) Fatal(msg string, a ...interface{}) {
	lgr.Log(msg, "fatal", a...)
	log.Fatal("must die here!")
}

// NewLogger is factory for logger
func NewLogger() *Logger {
	return &Logger{}
}
