/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

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

func (l Logger) Print(i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Printf(s string, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Println(i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) LogPrint(level Level, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) LogPrintf(level Level, s string, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) LogPrintln(level Level, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Panicln(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Fatalln(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Exitln(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Errorln(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Warningln(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Infoln(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Debugln(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Traceln(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Log(level Level, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Logf(level Level, format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) LogFn(level Level, fn LogFunc) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Panic(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Fatal(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Exit(exitCode int, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Panicf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Exitf(exitCode int, format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) PanicFn(fn LogFunc) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) FatalFn(fn LogFunc) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) ExitFn(exitCode int, fn LogFunc) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Error(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Warning(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Info(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Debug(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Trace(args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Errorf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Warningf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Infof(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Debugf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Tracef(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) ErrorFn(fn LogFunc) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) WarningFn(fn LogFunc) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) InfoFn(fn LogFunc) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) DebugFn(fn LogFunc) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) TraceFn(fn LogFunc) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) JsonPanic(arg interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) JsonFatal(arg interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) JsonExit(exitCode int, arg interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) JsonError(arg interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) JsonWarning(arg interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) JsonInfo(arg interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) JsonDebug(arg interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) JsonTrace(arg interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Birth(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) LastWords(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) V(i int) GLogger {
	//TODO implement me
	panic("implement me")
}

func (l Logger) OpenLog(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) OffLog(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Close() error {
	//TODO implement me
	panic("implement me")
}

func NewLogger() GLogger {
	return &Logger{}
}
