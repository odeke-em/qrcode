package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/odeke-em/rsc/qr"
)

func main() {
	var strChan chan string
	if len(os.Args) < 2 {
		strChan = freadLines(os.Stdin)
	} else {
		strChan = linesThroughChan(os.Args[1:]...)
	}

	for url := range strChan {
		url = trimStripLine(url)
		code, err := qr.Encode(url, qr.Q)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s %v\n", url, err)
			continue
		}

		pngImage := code.PNG()

		base := path.Base(url)

		rawPath := fmt.Sprintf("%s.png", base)
		f, err := os.Create(rawPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s %v\n", rawPath, err)
			continue
		}

		fmt.Fprintf(f, "%s\n", pngImage)
		f.Close()
	}
}

func freadLines(f io.Reader) chan string {
	strChan := make(chan string)
	go func() {
		defer close(strChan)

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			strChan <- scanner.Text()
		}
	}()

	return strChan
}

func linesThroughChan(lines ...string) chan string {
	strChan := make(chan string)
	go func() {
		defer close(strChan)

		for _, line := range lines {
			strChan <- line
		}
	}()

	return strChan
}

func trimStripLine(l string) string {
	for _, tok := range []string{" ", "\n"} {
		l = strings.Trim(l, tok)
	}

	return l
}
