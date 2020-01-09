package goal

import (
	"bufio"
	"io"
	"strings"
)

func ScanLines(r io.Reader) <-chan string {
	ch := make(chan string, 100000)
	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}

			ch <- line
		}
	}()

	return ch
}
