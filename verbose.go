package glog

import (
	"io"
	"sync"
)

type Logger struct {
	lock sync.Mutex
	pool *sync.Pool

	levelFiles map[Level]io.Writer
}

func NewLogger() GLogger {
	return &Logger{}
}
