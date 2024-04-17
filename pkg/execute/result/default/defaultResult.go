package defaultresult

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/execute/result"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	errors "github.com/apenella/go-common-utils/error"
)

const (
	// PrefixTokenSeparator is and string printed between prefix and ansible output
	PrefixTokenSeparator = "\u2500\u2500"
)

// DefaultResults prints results directly to stdout
type DefaultResults struct {
	trans []transformer.TransformerFunc
}

// NewDefaultResults returns a DefaultResults instance
func NewDefaultResults(options ...result.OptionsFunc) *DefaultResults {
	results := &DefaultResults{}
	results.Options(options...)
	return results
}

// WithTransformers sets a transformers list to DefaultResults
func WithTransformers(trans ...transformer.TransformerFunc) result.OptionsFunc {
	return func(r result.ResultsOutputer) {
		r.(*DefaultResults).trans = append(r.(*DefaultResults).trans, trans...)
	}
}

// Options executes the options functions received as a parameters to set the DefaultResults attributes
func (r *DefaultResults) Options(options ...result.OptionsFunc) {
	for _, opt := range options {
		opt(r)
	}
}

// Print method prints the to the DefaultResults writer the date received as input
func (r *DefaultResults) Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...result.OptionsFunc) error {
	var transformers []transformer.TransformerFunc

	errContext := "(DefaultResults::Print)"

	r.Options(options...)

	if len(r.trans) > 0 {
		transformers = append(transformers, transformer.Prepend(PrefixTokenSeparator))
	}

	// TODO ensure r.reader and r.writer are set

	transformers = append(transformers, r.trans...)

	err := output(ctx, reader, writer, transformers...)
	if err != nil {
		return errors.New(errContext, "Error processing the execution output", err)
	}

	return nil
}

// output processes the output data with the transformers coming from the execution an writes it to the input writer
func output(ctx context.Context, r io.Reader, w io.Writer, trans ...transformer.TransformerFunc) error {
	printChan := make(chan string)
	errChan := make(chan error)
	done := make(chan struct{})

	errContext := "(DefaultResults::output)"

	if r == nil {
		return errors.New(errContext, "Reader is not defined")
	}

	if w == nil {
		w = os.Stdout
	}

	if trans == nil {
		trans = []transformer.TransformerFunc{}
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
