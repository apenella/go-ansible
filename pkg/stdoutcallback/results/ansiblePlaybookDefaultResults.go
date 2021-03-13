package results

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	errors "github.com/apenella/go-common-utils/error"
)

// DefaultStdoutCallbackResults is the default method to print ansible-playbook results
func DefaultStdoutCallbackResults(ctx context.Context, r io.Reader, w io.Writer, trans ...TransformerFunc) error {

	tranformers := []TransformerFunc{}

	if len(trans) > 0 {
		tranformers = append(tranformers, Prepend(PrefixTokenSeparator))
	}

	tranformers = append(tranformers, trans...)

	err := output(ctx, r, w, tranformers...)
	if err != nil {
		return errors.New("(results::DefaultStdoutCallbackResults)", "Error processing execution output", err)
	}

	return nil
}

// output process the output data with the transformers comming from the execution an writes it to the input writer
func output(ctx context.Context, r io.Reader, w io.Writer, trans ...TransformerFunc) error {

	printChan := make(chan string)
	done := make(chan struct{})

	if r == nil {
		return errors.New("(results::DefaultStdoutCallbackResults)", "Reader is not defined")
	}

	if w == nil {
		w = os.Stdout
	}

	if trans == nil {
		trans = []TransformerFunc{}
	}

	go func() {
		defer close(done)
		defer close(printChan)

		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()

			for _, t := range trans {
				line = t(line)
			}

			printChan <- line
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
