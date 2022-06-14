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
	var transformers []TransformerFunc

	errContext := "(results::DefaultStdoutCallbackResults)"

	if len(trans) > 0 {
		transformers = append(transformers, Prepend(PrefixTokenSeparator))
	}

	transformers = append(transformers, trans...)

	err := output(ctx, r, w, transformers...)
	if err != nil {
		return errors.New(errContext, "Error processing execution output", err)
	}

	return nil
}

// output process the output data with the transformers comming from the execution an writes it to the input writer
func output(ctx context.Context, r io.Reader, w io.Writer, trans ...TransformerFunc) error {
	printChan := make(chan string)
	errChan := make(chan error)
	done := make(chan struct{})

	errContext := "(results::output)"

	if r == nil {
		return errors.New(errContext, "Reader is not defined")
	}

	if w == nil {
		w = os.Stdout
	}

	if trans == nil {
		trans = []TransformerFunc{}
	}

	go func() {
		defer close(done)
		defer close(errChan)
		defer close(printChan)

		reader := bufio.NewReader(r)
		for {
			line, err := readLine(reader)
			if err != nil {
				if err != io.EOF {
					errChan <- err
				}

				break
			}

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
			_, err := fmt.Fprintln(w, line)
			if err != nil {
				return err
			}
		case err := <-errChan:
			return err
		case <-done:
			return nil
		case <-ctx.Done():
			return nil
		}
	}
}

func readLine(r *bufio.Reader) (string, error) {
	var line []byte
	for {
		l, more, err := r.ReadLine()
		if err != nil {
			return "", err
		}

		// Avoid the copy if the first call produced a full line.
		if line == nil && !more {
			return string(l), nil
		}

		line = append(line, l...)
		if !more {
			break
		}
	}

	return string(line), nil
}
