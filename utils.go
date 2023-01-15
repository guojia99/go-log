package glog

import (
	"runtime"
	"strings"
)

const packageName = "github.com/guojia99/go-log"

func getCodeRunPackage(startSkip int) (codeLine int, codePath string) {
	for skip := startSkip; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok || pc == 0 || line == 0 || len(file) == 0 {
			return
		}
		codeLine, codePath = line, file
		if !strings.Contains(file, packageName) {
			return
		}
	}
}
