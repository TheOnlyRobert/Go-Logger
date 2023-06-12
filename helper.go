package logger

import (
	"fmt"
	"runtime"
	"time"
)

// Returns the name of caller function.
func getFunctionInfo() (info string) {
	pc, _, _, ok := runtime.Caller(6)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	funcName := runtime.FuncForPC(pc).Name()

	return funcName
}

func BuildMessage(logCat LogCat, startTime time.Time, apiKey string, v ...interface{}) string {
	var t time.Duration

	if !startTime.IsZero() {
		t = time.Since(startTime)
	}

	ms := float64(t.Nanoseconds()) / float64(time.Millisecond)

	extraInfo := fmt.Sprintf(
		`ApiKey="%s", Milliseconds="%f", Function="%s", Code="%s", Type="%s"`,
		apiKey,
		ms,
		getFunctionInfo(),
		logCat.Code,
		logCat.Type,
	)

	detailInfo := fmt.Sprint(v...)

	return extraInfo + ", " + fmt.Sprintf(`message="%s"`, detailInfo)
}
