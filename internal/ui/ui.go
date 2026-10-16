package ui

import (
	"fmt"
	"os"
)

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	dim    = "\033[2m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
)

func enabled() bool {
	info, err := os.Stdout.Stat()
	return err == nil && (info.Mode()&os.ModeCharDevice) != 0
}

func wrap(code, s string) string {
	if !enabled() {
		return s
	}
	return code + s + reset
}

func Cyan(s string) string   { return wrap(cyan, s) }
func Bold(s string) string   { return wrap(bold, s) }
func Dim(s string) string    { return wrap(dim, s) }
func Red(s string) string    { return wrap(red, s) }
func Green(s string) string  { return wrap(green, s) }
func Yellow(s string) string { return wrap(yellow, s) }

func Banner() {
	fmt.Println(Cyan(Bold(`
   __________  ___ _    ____  _  __
  / ____/ __ \/   / |  / / / |/ /
 / /   / /_/ / /| | | / / /|   / 
/ /___/ _, _/ ___ | |/ / / |   | 
\____/_/ |_/_/  |_|___/_/  |_|_| 
`)))
}

func CrawlLine(n, depth, status int, errMsg, url, parent string) string {
	meta := Dim(fmt.Sprintf("#%d depth=%d", n, depth))
	body := fmt.Sprintf("status=%d err=%s %s %s", status, errMsg, url, parent)
	switch {
	case errMsg != "" || status >= 400:
		body = Red(body)
	case status >= 300:
		body = Yellow(body)
	default:
		body = Green(body)
	}
	return meta + " " + body
}
