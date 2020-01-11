package goal

import (
	"bufio"
	"io"
	"strings"
)

//ScanLines read all data from reader and split to lines, return readonly channel for iterator.
//
//Every line has been trimed and empty line removed from result.
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
