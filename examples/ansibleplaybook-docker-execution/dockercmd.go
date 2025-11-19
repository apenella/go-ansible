package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/exec"
	"github.com/apenella/go-docker-builder/pkg/build"
	contextpath "github.com/apenella/go-docker-builder/pkg/build/context/path"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// DockerExec struct implements the Executable interface
type DockerExec struct {
	client *client.Client
}

var _ = execute.Executabler(&DockerExec{})

func NewDockerExec(client *client.Client) *DockerExec {
	return &DockerExec{
		client: client,
	}
}

// Command is a wrapper of exec.Command
func (e *DockerExec) Command(name string, arg ...string) exec.Cmder {
	return e.CommandContext(context.TODO(), name, arg...)
}

// CommandContext is a wrapper of exec.CommandContext
func (e *DockerExec) CommandContext(ctx context.Context, name string, arg ...string) exec.Cmder {

	cmd := NewDockerCmd(e.client)
	cmd.ContainerName = "ansible_playbook_executor"
	cmd.Env = os.Environ()

	cmd.Cmd = append([]string{}, name)
	cmd.Cmd = append(cmd.Cmd, arg...)

	return cmd
}

// dockerCmdOptionsFunc is a function type that modifies dockerCmd options
type dockerCmdOptionsFunc func(*dockerCmd)

// Implementation of Cmder interface that runs commands inside a Docker container
type dockerCmd struct {
	client      *client.Client
	containerID string

	ContainerName string
	Env           []string
	Cmd           []string

	AutoRemove bool

	imagePathContext string
	mounts           []mount.Mount
	workingDir       string

	stdoutPipeReader io.ReadCloser
	stdoutPipeWriter io.WriteCloser
	stderrPipeReader io.ReadCloser
	stderrPipeWriter io.WriteCloser
}

var _ = exec.Cmder(&dockerCmd{})

// NewDockerCmd creates a new dockerCmd instance
func NewDockerCmd(client *client.Client, opts ...dockerCmdOptionsFunc) *dockerCmd {

	workingDir := "/code"

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	_ = ex

	cmd := &dockerCmd{
		client: client,
		mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: filepath.Dir(ex),
				Target: workingDir,
			},
		},
		workingDir:       workingDir,
		imagePathContext: filepath.Join("docker", "ansible"),
	}

	cmd.stdoutPipeReader, cmd.stdoutPipeWriter = io.Pipe()
	cmd.stderrPipeReader, cmd.stderrPipeWriter = io.Pipe()

	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
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
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()

}

func (cmd *dockerCmd) imageBuild(ctx context.Context, imageName string) error {
	var err error

	dockerBuildContext := &contextpath.PathBuildContext{
		Path: cmd.imagePathContext,
	}

	dockerBuilder := build.NewDockerBuildCmd(cmd.client).
		WithImageName(imageName)

	err = dockerBuilder.AddBuildContext(dockerBuildContext)
	if err != nil {
		return fmt.Errorf("failed to add build context: %w", err)
	}

	err = dockerBuilder.Run(ctx)
	if err != nil {
		return fmt.Errorf("failed to run docker build: %w", err)
	}

	return nil
}

// Start starts the command but does not wait for it to complete.
func (cmd *dockerCmd) Start() (err error) {
	var containerCreateResp container.CreateResponse
	var attach types.HijackedResponse

	ctx := context.TODO()
	imageName := "ansibleplaybook-docker-executor"

	err = cmd.imageBuild(ctx, imageName)
	if err != nil {
		return fmt.Errorf("failed to build image: %w", err)
	}

	containerConfig := &container.Config{
		Image:        imageName,
		Cmd:          cmd.Cmd,
		Tty:          false,
		WorkingDir:   cmd.workingDir,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
	}

	containerCreateResp, err = cmd.client.ContainerCreate(
		ctx,
		containerConfig,
		&container.HostConfig{
			AutoRemove: cmd.AutoRemove,
			Mounts:     cmd.mounts,
		},
		&network.NetworkingConfig{},
		nil,
		cmd.ContainerName,
	)
	if err != nil {
		return fmt.Errorf("failed to create container: %w", err)
	}

	cmd.containerID = containerCreateResp.ID
	fmt.Println("Container created:", cmd.containerID)

	attach, err = cmd.client.ContainerAttach(
		ctx,
		cmd.containerID,
		container.AttachOptions{
			Stdin:  false,
			Stdout: true,
			Stderr: true,
			Stream: true,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to attach to container: %w", err)
	}

	go func() {
		defer attach.Close()
		defer cmd.stdoutPipeWriter.Close()
		defer cmd.stderrPipeWriter.Close()
		// Copying stdout and stderr from the container to the respective pipes
		_, _ = stdcopy.StdCopy(cmd.stdoutPipeWriter, cmd.stdoutPipeWriter, attach.Reader)
	}()

	err = cmd.client.ContainerStart(ctx, cmd.containerID, container.StartOptions{})
	if err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}

	return nil
}

// StderrPipe returns a pipe that will be connected to the command's standard error when the command starts.
func (cmd *dockerCmd) StderrPipe() (io.ReadCloser, error) {
	return cmd.stderrPipeReader, nil
}

// StdinPipe returns a pipe that will be connected to the command's standard input when the command starts.
func (cmd *dockerCmd) StdinPipe() (io.WriteCloser, error) {
	return nil, nil
}

// StdoutPipe returns a pipe that will be connected to the command's standard output when the command starts.
func (cmd *dockerCmd) StdoutPipe() (io.ReadCloser, error) {
	return cmd.stdoutPipeReader, nil
}

// String returns the command string.
func (cmd *dockerCmd) String() string {
	return ""
}

// Wait waits for the command to exit and waits for any copying to stdin or copying from stdout or stderr to complete.
func (cmd *dockerCmd) Wait() error {
	var err error

	statusCh, errCh := cmd.client.ContainerWait(
		context.TODO(),
		cmd.containerID,
		container.WaitConditionNotRunning,
	)

	select {
	case err = <-errCh:

	case status := <-statusCh:

		if status.StatusCode != 0 {
			err = fmt.Errorf("container exited with code %d", status.StatusCode)
		}
	}

	err = cmd.cleanup()
	if err != nil {
		return fmt.Errorf("failed to cleanup container: %w", err)
	}

	return err
}

func (cmd *dockerCmd) cleanup() error {
	err := cmd.client.ContainerRemove(context.TODO(), cmd.containerID, container.RemoveOptions{
		Force:         true,
		RemoveVolumes: true,
	})
	if err != nil {
		return fmt.Errorf("failed to remove container: %w", err)
	}
	return nil

}
