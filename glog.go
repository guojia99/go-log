/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package glog

import "io"

type LogFunc func() []interface{}

// STdGLogger log output in the command terminal.
//
//goland:noinspection ALL
type STdGLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
	LogPrint(Level, ...interface{})
	LogPrintf(Level, string, ...interface{})
	LogPrintln(Level, ...interface{})

	Panicln(args ...interface{})
	Fatalln(args ...interface{})
	Exitln(args ...interface{})
	Errorln(args ...interface{})
	Warningln(args ...interface{})
	Infoln(args ...interface{})
	Debugln(args ...interface{})
	Traceln(args ...interface{})
}

// BasicGLogger the base operation log function.
//
//goland:noinspection ALL
type BasicGLogger interface {
	Log(level Level, args ...interface{})
	Logf(level Level, format string, args ...interface{})
	LogFn(level Level, fn LogFunc)

	Panic(args ...interface{})
	Fatal(args ...interface{})
	Exit(exitCode int, args ...interface{})
	Panicf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Exitf(exitCode int, format string, args ...interface{})
	PanicFn(fn LogFunc)
	FatalFn(fn LogFunc)
	ExitFn(exitCode int, fn LogFunc)

	Error(args ...interface{})
	Warning(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
	ErrorFn(fn LogFunc)
	WarningFn(fn LogFunc)
	InfoFn(fn LogFunc)
	DebugFn(fn LogFunc)
	TraceFn(fn LogFunc)
}

//goland:noinspection ALL
type JsonGLogger interface {
	JsonPanic(arg interface{})
	JsonFatal(arg interface{})
	JsonExit(exitCode int, arg interface{})
	JsonError(arg interface{})
	JsonWarning(arg interface{})
	JsonInfo(arg interface{})
	JsonDebug(arg interface{})
	JsonTrace(arg interface{})
}

type GLoggerParams interface {
	SetV(int)
	SetLevel(level Level)
	SetFlags(flags Flags)
	SetKeyFlags(key string, value string)
	SetWrite(level Level, Write io.Writer)
	SetWriteByMap(s map[Level]io.Writer)
}

//goland:noinspection ALL
type GLogger interface {
	STdGLogger
	BasicGLogger
	JsonGLogger

	Birth(format string, args ...interface{})
	LastWords(format string, args ...interface{})

	V(int) GLogger
	OpenLog(format string, args ...interface{})
	OffLog(format string, args ...interface{})
	Close() error
}
