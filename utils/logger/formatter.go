package logger

import "runtime"

const (
	maxCallersStackGet = 5
	skippedCallers     = 3
	unknownCaller      = "unknown"
)

const (
	attributesTag = "attributes"
)

func getCaller() (string, string) {
	pcs := make([]uintptr, maxCallersStackGet)
	n := runtime.Callers(skippedCallers, pcs)
	pcs = pcs[:n]
	file := unknownCaller
	function := unknownCaller

	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		if !more {
			break
		}

		file := frame.File
		function := frame.Function

		return file, function
	}
	return file, function
}
