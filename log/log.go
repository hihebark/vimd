package log

import (
	"fmt"
	"time"
)

const (
	RESET     = "\033[0m"
	START     = "\033[%sm"
	ORANGE    = "33"
	LIGHTRED  = "91"
	LIGHTBLUE = "94"
	YELLOW    = "93"
)

var (
	dbg        string = fmt.Sprintf("%s[DBG]%s", fmt.Sprintf(START, ORANGE), RESET)
	inf        string = fmt.Sprintf("%s[INF]%s", fmt.Sprintf(START, LIGHTBLUE), RESET)
	err        string = fmt.Sprintf("%s[ERR]%s", fmt.Sprintf(START, LIGHTRED), RESET)
	war        string = fmt.Sprintf("%s[WAR]%s", fmt.Sprintf(START, YELLOW), RESET)
	timeFormat string = "15:04:05"
)

func Dbg(str string, args ...interface{}) {
	fmt.Printf("[%s]%s %s\n", time.Now().Format(timeFormat), dbg, fmt.Sprintf(str, args...))
}
func Inf(str string, args ...interface{}) {
	fmt.Printf("[%s]%s %s\n", time.Now().Format(timeFormat), inf, fmt.Sprintf(str, args...))
}
func Err(str string, args ...interface{}) {
	fmt.Printf("[%s]%s %s\n", time.Now().Format(timeFormat), err, fmt.Sprintf(str, args...))
}
func War(str string, args ...interface{}) {
	fmt.Printf("[%s]%s %s\n", time.Now().Format(timeFormat), war, fmt.Sprintf(str, args...))
}
