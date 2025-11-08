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
	"golang.org/x/sync/errgroup"
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
	errContext := "(DefaultResults::output)"
	goroutine, ctx := errgroup.WithContext(ctx)

	if r == nil {
		return errors.New(errContext, "Reader is not defined")
	}

	if w == nil {
		w = os.Stdout
	}

	if trans == nil {
		trans = []transformer.TransformerFunc{}
	}

	// buffered channel to decouple reader/writer
	dataChan := make(chan string, 10)
	goroutine.Go(func() error {

		defer close(dataChan)

		reader := bufio.NewReader(r)
		for {
			line, err := readLine(reader)
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return errors.New(errContext, "error reading line", err)
			}

			// Apply transformers
			for _, t := range trans {
				line = t(line)
			}

			select {
			case dataChan <- line:
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
			case line, ok := <-dataChan:
				if !ok {
					// channel closed => finished
					return nil
				}
				if _, err := fmt.Fprintln(w, line); err != nil {
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

		// avoid the copy if the first call produced a full line.
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
