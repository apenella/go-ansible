package main

import (
	"context"
	"fmt"
	"io"

	"golang.org/x/sync/errgroup"

	"github.com/apenella/go-ansible/v2/pkg/execute/exec"
	"github.com/moby/moby/client"
)

// dockerCmdOptionsFunc is a function type that modifies dockerCmd options
type dockerCmdOptionsFunc func(*dockerCmd)

// Implementation of Cmder interface that runs commands inside a Docker container
type dockerCmd struct {
	client *client.Client

	imageBuildOptions *client.ImageBuildOptions
}

var _ = exec.Cmder(&dockerCmd{})

// NewDockerCmd creates a new dockerCmd instance
func NewDockerCmd(client *client.Client, opts ...dockerCmdOptionsFunc) *dockerCmd {

	cmd := &dockerCmd{}

	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
}

// WithImageBuildOptions sets the image build options for the dockerCmd
func WithImageBuildOptions(options *client.ImageBuildOptions) dockerCmdOptionsFunc {
	return func(cmd *dockerCmd) {
		cmd.imageBuildOptions = options
	}
}

//
// Cmder interface implementation

// CombinedOutput runs the command and returns its combined standard output and standard error.
func (cmd *dockerCmd) CombinedOutput() ([]byte, error) {

	return nil, nil
}

// Environ returns the environment variables for the command.
func (cmd *dockerCmd) Environ() []string {

	return nil
}

// Output runs the command and returns its standard output.
func (cmd *dockerCmd) Output() ([]byte, error) {
	return nil, nil
}

// Run runs the command.
func (cmd *dockerCmd) Run() error {
	return nil
}

func (cmd *dockerCmd) imageBuild(ctx context.Context, options *client.ImageBuildOptions) error {
	var routineGroup *errgroup.Group

	routineGroup, ctx = errgroup.WithContext(ctx)

	if options == nil {
		return nil
	}

	result, err := cmd.client.ImageBuild(ctx, options.Context, *options)
	if err != nil {
		return fmt.Errorf("error building image: %w", err)
	}
	defer result.Body.Close()

	routineGroup.Go(func() error {

		return nil
	})

	routineGroup.Go(func() error {

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return nil
			}
		}
	})

	return nil
}

// Start starts the command but does not wait for it to complete.
func (cmd *dockerCmd) Start() error {

	cmd.imageBuild(context.TODO(), &client.ImageBuildOptions{})

	return nil
}

// StderrPipe returns a pipe that will be connected to the command's standard error when the command starts.
func (cmd *dockerCmd) StderrPipe() (io.ReadCloser, error) {
	return nil, nil
}

// StdinPipe returns a pipe that will be connected to the command's standard input when the command starts.
func (cmd *dockerCmd) StdinPipe() (io.WriteCloser, error) {
	return nil, nil
}

// StdoutPipe returns a pipe that will be connected to the command's standard output when the command starts.
func (cmd *dockerCmd) StdoutPipe() (io.ReadCloser, error) {
	return nil, nil
}

// String returns the command string.
func (cmd *dockerCmd) String() string {
	return ""
}

// Wait waits for the command to exit and waits for any copying to stdin or copying from stdout or stderr to complete.
func (cmd *dockerCmd) Wait() error {
	return nil
}
