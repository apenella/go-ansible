package json

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/execute/result"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	errors "github.com/apenella/go-common-utils/error"
	"golang.org/x/sync/errgroup"
)

const (
	// PrefixTokenSeparator is and string printed between prefix and ansible output
	PrefixTokenSeparator = "\u2500\u2500"
)

type JSONStdoutCallbackResults struct {
	trans []transformer.TransformerFunc
}

func NewJSONStdoutCallbackResults(options ...result.OptionsFunc) *JSONStdoutCallbackResults {
	results := &JSONStdoutCallbackResults{}
	results.Options(options...)
	return results
}

// WithTransformers sets a transformers list to DefaultResults
func WithTransformers(trans ...transformer.TransformerFunc) result.OptionsFunc {
	return func(r result.ResultsOutputer) {
		r.(*JSONStdoutCallbackResults).trans = append(r.(*JSONStdoutCallbackResults).trans, trans...)
	}
}

// Options executes the options functions received as a parameters to set the DefaultResults attributes
func (r *JSONStdoutCallbackResults) Options(options ...result.OptionsFunc) {
	for _, opt := range options {
		opt(r)
	}
}

// Print method manges the ansible' JSON stdout callback and print the result stats
func (r *JSONStdoutCallbackResults) Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...result.OptionsFunc) error {
	var transformers []transformer.TransformerFunc

	errContext := "(result::json::JSONStdoutCallbackResults::Print)"

	if reader == nil {
		return errors.New(errContext, "JSONStdoutCallbackResults requires a reader to print the output of the execution")
	}

	if writer == nil {
		return errors.New(errContext, "JSONStdoutCallbackResults requires a writer to print the output of the execution")
	}

	skipPatterns := []string{
		// This pattern skips timer's callback whitelist output
		"^[\\s\\t]*Playbook run took [0-9]+ days, [0-9]+ hours, [0-9]+ minutes, [0-9]+ seconds$",
	}

	r.Options(options...)

	if len(r.trans) > 0 {
		transformers = append(transformers, transformer.Prepend(PrefixTokenSeparator))
	}

	transformers = append(transformers, r.trans...)
	transformers = append(transformers, transformer.IgnoreMessage(skipPatterns))

	err := output(ctx, reader, writer, transformers...)
	if err != nil {
		return errors.New(errContext, "error processing execution output", err)
	}

	return nil
}

// output processes the output data with the transformers coming from the execution an writes it to the input writer
func output(ctx context.Context, reader io.Reader, writer io.Writer, trans ...transformer.TransformerFunc) error {
	errContext := "(result::json::JSONStdoutCallbackResults::output)"
	goroutine, ctx := errgroup.WithContext(ctx)

	if reader == nil {
		return errors.New(errContext, "JSONStdoutCallbackResults requires a reader to print the output of the execution")
	}

	if writer == nil {
		writer = os.Stdout
	}

	if trans == nil {
		trans = []transformer.TransformerFunc{}
	}

	// buffered channel to decouple reader/writer
	lines := make(chan string, 10)
	goroutine.Go(func() error {

		defer close(lines)

		r := bufio.NewReader(reader)
		for {
			line, err := readLine(r)
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return errors.New(errContext, "error reading line", err)
			}

			// apply transformers
			for _, t := range trans {
				line = t(line)
			}

			select {
			case lines <- line:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	goroutine.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case line, ok := <-lines:
				if !ok {
					// channel closed => finished
					return nil
				}
				if _, err := fmt.Fprintln(writer, line); err != nil {
					return errors.New(errContext, "error writing the received data", err)
				}
			}
		}
	})

	// wait for both goroutines to complete or fail. The expected behaviour is that when there are no more lines to read, the first goroutine finishes, and the errgroup cancels the second goroutine, which handles the writing operations.
	err := goroutine.Wait()
	if err != nil {
		return errors.New(errContext, "error processing output", err)
	}

	return nil
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
