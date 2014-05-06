package main

import (
	gonx "../.."
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var format string
var logFile string

func init() {
	flag.StringVar(&format, "format", "$remote_addr [$time_local] \"$request\"", "Log format")
	flag.StringVar(&logFile, "log", "dummy", "Log file name to read. Read from STDIN if file name is '-'")
}

func main() {
	flag.Parse()

	// Read given file or from STDIN
	var file io.Reader
	if logFile == "dummy" {
		file = strings.NewReader(`89.234.89.123 [08/Nov/2013:13:39:18 +0000] "GET /api/foo/bar HTTP/1.1"`)
	} else if logFile == "-" {
		file = os.Stdin
	} else {
		file, err := os.Open(logFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	// Create reader and call Read method until EOF
	reader := gonx.NewReader(file, format)
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		// Process the record... e.g.
		fmt.Printf("%+v\n", rec)
	}
}
