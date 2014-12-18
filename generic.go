package goutils

import (
	"fmt"
	"log"
	"runtime"
)

type ErrorParams struct {
	err       error
	stderr    string
	callerNum int
}

// check for errors and panic, if found
func checkForErrors(e ErrorParams) {
	if e.err != nil {
		pc, fn, line, _ := runtime.Caller(e.callerNum)
		msg := ""
		if e.stderr != "" {
			msg = fmt.Sprintf("[Error] in %s[%s:%d] %v: %s",
				runtime.FuncForPC(pc).Name(), fn, line, e.err, e.stderr)
		} else {
			msg = fmt.Sprintf("[Error] in %s[%s:%d] %v",
				runtime.FuncForPC(pc).Name(), fn, line, e.err)
		}
		log.Fatal(msg)
	}
}
