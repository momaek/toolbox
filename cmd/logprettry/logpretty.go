package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var (
	// InfoColor ..
	InfoColor = "\033[1;34m%s\033[0m"
	// NoticeColor ..
	NoticeColor = "\033[1;36m%s\033[0m"
	// WarningColor ..
	WarningColor = "\033[38;5;3m%s\033[39;49m"
	// ErrorColor ...
	ErrorColor = "\033[1;31m%s\033[0m"
	// DebugColor ...
	DebugColor = "\033[0;36m%s\033[0m"
	// FileColor ...
	FileColor = "\033[38;5;5m%s\033[39;49m"
	// DateColor ..
	DateColor = "\033[38;5;14m%s\033[39;49m"
	// SlowQueryColor ..
	SlowQueryColor = "\033[38;5;167m%s\033[39;49m"
	// MysqlTimeColor ..s2w
	MysqlTimeColor = "\033[38;5;172m%s\033[39;49m"

	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

// Log request log msg
type Log struct {
	Message      string `json:"message"`
	ReqID        string `json:"reqid"`
	Type         string `json:"type"`
	Level        string `json:"level"`
	Time         string `json:"time"`
	Path         string `json:"path"`
	Method       string `json:"method"`
	Status       int    `json:"status"`
	Latency      string `json:"latency"`
	ClientIP     string `json:"client_ip"`
	Caller       string `json:"caller"`
	Action       string `json:"action"`
	SQL          string `json:"sql"`
	RowsAffected int    `json:"rows_affected"`
	SlowQuery    bool   `json:"slow_query"`
}

func main() {
	count := 0
wait:
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | logpretty")
		return
	}

	if info.Size() == 0 {
		count++
		time.Sleep(100 * time.Millisecond)
		goto wait
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}

		if input == '\n' {
			prettyPrintJSON(string(output))
			output = []rune{}
			continue
		}

		output = append(output, input)
	}
}

func prettyPrintJSON(line string) {
	if !isValidJSON(line) {
		fmt.Println(line)
		return
	}

	var l = Log{}
	err := json.Unmarshal([]byte(line), &l)
	if err != nil {
		fmt.Println(line)
		return
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf(NoticeColor, l.Time))
	buf.WriteString(" [")
	printLog(l, buf)
	fmt.Printf("%s\n", buf.String())
}

func printLog(m Log, buf *bytes.Buffer) {
	reqID := m.ReqID
	buf.WriteString(reqID)
	buf.WriteString("] ")
	buf.WriteString(fmt.Sprintf("["+levelToColor(m.Level)+"] ", m.Level))
	if len(m.Caller) > 0 {
		buf.WriteString(fmt.Sprintf("["+FileColor+"]", m.Caller))
	}
	typ := m.Type
	switch typ {
	case "DATABASE":
		printDatabaseLog(m, buf)
	case "REQ":
		printReqResponseLog(m, buf)
	default:
		buf.WriteString(" " + m.Message)
	}

}

// 简单粗暴看看是不是 {}
func isValidJSON(line string) bool {
	if (strings.HasPrefix(line, "{") && strings.HasSuffix(line, "}")) ||
		(strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")) {
		return true
	}
	return false
}

func printReqResponseLog(m Log, buf *bytes.Buffer) {
	buf.WriteString("[REQ] ")

	buf.WriteString(m.Message + " | ")
	buf.WriteString(colorMethod(m.Method) + " | ")
	if m.Status > 0 {
		buf.WriteString(colorStatus(m.Status) + " | ")
		buf.WriteString(fmt.Sprintf("["+MysqlTimeColor+"]", m.Latency) + " | ")
	}

	buf.WriteString(m.Path + " ")
	buf.WriteString(m.ClientIP)
}

func printDatabaseLog(m Log, buf *bytes.Buffer) {
	buf.WriteString(fmt.Sprintf(" "+MysqlTimeColor+" ", m.Latency))
	buf.WriteString(m.SQL + " ")
	buf.WriteString(fmt.Sprintf("RowsAffected: %d", m.RowsAffected) + " ")
	if m.SlowQuery {
		l := "Slow Query !!"
		buf.WriteString(fmt.Sprintf(SlowQueryColor, l))
	}

	if len(m.Message) > 0 {
		buf.WriteString(fmt.Sprintf(ErrorColor, m.Message))
	}
}

func levelToColor(level string) string {
	switch level {
	case "info":
		return InfoColor
	case "error":
		return ErrorColor
	case "warning", "warn":
		return WarningColor
	case "debug", "debugging":
		return DebugColor
	default:
		return NoticeColor
	}
}

func colorStatus(code int) string {
	return fmt.Sprintf("%s %d %s", colorForStatus(code), code, reset)
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return green
	case code >= 300 && code < 400:
		return white
	case code >= 400 && code < 500:
		return yellow
	default:
		return red
	}
}

func colorMethod(method string) string {
	return fmt.Sprintf("%s %s %s", colorForMethod(method), method, reset)
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}
