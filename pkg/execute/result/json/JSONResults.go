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

// JSONStdoutCallbackResults method manges the ansible' JSON stdout callback and print the result stats
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
		return errors.New(errContext, "Error processing execution output", err)
	}

	return nil
}

// output processes the output data with the transformers coming from the execution an writes it to the input writer
func output(ctx context.Context, reader io.Reader, writer io.Writer, trans ...transformer.TransformerFunc) error {
	printChan := make(chan string)
	errChan := make(chan error)
	done := make(chan struct{})

	errContext := "(result::json::JSONStdoutCallbackResults::output)"

	if reader == nil {
		return errors.New(errContext, "JSONStdoutCallbackResults requires a reader to print the output of the execution")
	}

	if writer == nil {
		writer = os.Stdout
	}

	if trans == nil {
		trans = []transformer.TransformerFunc{}
	}

	go func() {
		defer close(done)
		defer close(errChan)
		defer close(printChan)

		r := bufio.NewReader(reader)
		for {
			line, err := readLine(r)
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
			_, err := fmt.Fprintln(writer, line)
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
