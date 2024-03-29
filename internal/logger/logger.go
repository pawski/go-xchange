package logger

import (
	"sync"

	"github.com/Sirupsen/logrus"
)

var defaultLogger *logrus.Logger
var once sync.Once

func Get() *logrus.Logger {
	once.Do(func() {
		defaultLogger = logrus.New()
	})

	return defaultLogger
}

func SetDebugLevel() {
	Get().Level = logrus.DebugLevel
}
