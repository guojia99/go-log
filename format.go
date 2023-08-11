/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package glog

import "strings"

type Flags int

type Level uint

const (
	PanicLevel Level = iota
	FatalLevel
	ExitLevel
	OpenLevel
	OffLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
	AllLevel
)

var levelString = map[Level]string{
	PanicLevel: "panic",
	FatalLevel: "fatal",
	ExitLevel:  "exit",
	OpenLevel:  "open",
	OffLevel:   "off",
	ErrorLevel: "error",
	WarnLevel:  "warn",
	InfoLevel:  "info",
	DebugLevel: "debug",
	TraceLevel: "trace",
	AllLevel:   "all",
}

func (l Level) String() string {
	if v, ok := levelString[l]; ok {
		return v
	}
	return "unknown"
}

func (l Level) ToUpper() string { return strings.ToUpper(l.String()) }

type Format interface {
	OutPut(message []byte) ([]byte, error)
}
