package log

import "log"
import "runtime"

var isOn = false
var l *log.Logger

func DebugOn(ll *log.Logger) {
	isOn = true
	l = ll
}

func Err(err error) bool {
	if err != nil {
		if isOn {
			// notice that we're using 1, so it will actually log where
			// the error happened, 0 = this function, we don't want that.
			_, fn, line, _ := runtime.Caller(1)
			l.Printf("[error] %s:%d %v", fn, line, err)
		}
		return true

	}
	return false
}
