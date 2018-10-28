package log

import (
	"fmt"
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
	dbg string = fmt.Sprintf("%s[ DBG ]%s", fmt.Sprintf(START, ORANGE), RESET)
	inf string = fmt.Sprintf("%s[ INF ]%s", fmt.Sprintf(START, LIGHTBLUE), RESET)
	err string = fmt.Sprintf("%s[ ERR ]%s", fmt.Sprintf(START, LIGHTRED), RESET)
	war string = fmt.Sprintf("%s[ WAR ]%s", fmt.Sprintf(START, YELLOW), RESET)
)

func Dbg(args ...interface{}) {
	fmt.Printf("%s %s", dbg, fmt.Sprintf("%s\n", args...))
}
func Inf(args ...interface{}) {
	fmt.Printf("%s %s", inf, fmt.Sprintf("%s\n", args...))
	//fmt.Printf(" %s %s\n", inf, fmt.Sprintf(args...))
}
func Err(args ...interface{}) {
	fmt.Printf("%s %s", err, fmt.Sprintf("%s\n", args...))
	//fmt.Printf(" %s %s\n", err, fmt.Sprintf(args...))
}
func War(args ...interface{}) {
	fmt.Printf("%s %s", war, fmt.Sprintf("%s\n", args...))
	//fmt.Printf(" %s %s\n", war, fmt.Sprintf(args...))
}
