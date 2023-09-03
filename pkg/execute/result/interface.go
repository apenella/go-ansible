package result

import (
	"context"
	"io"
)

// OptionsFunc is a function that can be used to configure a ResultsOutputer struct
type OptionsFunc func(ResultsOutputer)

// ResultsOutputer is the interface that must implements an struct to print the execution results
type ResultsOutputer interface {
	Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...OptionsFunc) error
}
