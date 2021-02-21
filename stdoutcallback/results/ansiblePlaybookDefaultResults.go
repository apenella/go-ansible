package results

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	errors "github.com/apenella/go-common-utils/error"
)

const (
	// PrefixTokenSeparator is and string printed between prefix and ansible output
	PrefixTokenSeparator = "\u2500\u2500"
)

// DefaultStdoutCallbackResults is the default method to print ansible-playbook results
func DefaultStdoutCallbackResults(ctx context.Context, prefix string, r io.Reader, w io.Writer) error {

	printChan := make(chan string)
	done := make(chan struct{})

	if r == nil {
		return errors.New("(results::DefaultStdoutCallbackResults)", "Reader is not defined")
	}

	if w == nil {
		w = os.Stdout
	}

	go func() {
		defer close(done)
		defer close(printChan)

		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			printChan <- fmt.Sprintf("%s %s %s", prefix, PrefixTokenSeparator, scanner.Text())
		}
		done <- struct{}{}
	}()

	for {
		select {
		case line := <-printChan:
			fmt.Fprintf(w, "%s\n", line)
		case <-done:
			return nil
		case <-ctx.Done():
			return nil
		}
	}
}
