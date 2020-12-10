package main

// TO BE DONE

/*
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// Log request log msg
type Log struct {
	Message  string `json:"message"`
	ReqID    string `json:"x-reqid"`
	Type     string `json:"type"`
	Level    string `json:"level"`
	Time     string `json:"time"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	Status   int    `json:"status"`
	Latency  string `json:"latency"`
	ClientIP string `json:"client_ip"`
	Caller   string `json:"caller"`
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

}*/
