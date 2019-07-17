package logger

import (
  "os"
  "log"
)

type Logger struct {
  Warn  func(...interface{})
}

var logger = &Logger{ log.New(os.Stderr, "WARN: ", log.Lshortfile).Print }

func Get() *Logger {
  return logger
}


