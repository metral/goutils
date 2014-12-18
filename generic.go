package goutils

import (
	"fmt"
	"log"
	"runtime"
)

// check for errors and panic, if found
func checkForErrors(err error, stderr string, callerNum int) {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(callerNum)
		msg := ""
		if stderr != "" {
			msg = fmt.Sprintf("[Error] in %s[%s:%d] %v: %s",
				runtime.FuncForPC(pc).Name(), fn, line, err, stderr)
		} else {
			msg = fmt.Sprintf("[Error] in %s[%s:%d] %v",
				runtime.FuncForPC(pc).Name(), fn, line, err)
		}
		log.Fatal(msg)
	}
}
