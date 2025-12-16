
# Example Ansible Playbook Docker Execution

Managing Ansible playbook execution in multiple environments can be challenging due to dependency conflicts, inconsistent tool versions, and the need for isolation between runs. This example was created to address these issues by demonstrating how to run Ansible playbooks inside a Docker container, fully managed from a Go application.

This example introduces a `DockerExec` struct that implements the [`Executabler`](https://pkg.go.dev/github.com/apenella/go-ansible/v2/pkg/execute#Executabler) interface and a `DockerCmd` struct that implements the [`Cmder`](https://pkg.go.dev/github.com/apenella/go-ansible/v2/pkg/execute/exec#Cmder) interface. The example configures `DockerExec` as the executor for the DefaultExecutor component, enabling seamless execution of Ansible playbooks within a containerized environment. The Go application builds the required Docker image, creates and starts the container, attaches to its output, and cleans up resources after execution.

By executing Ansible within a container, you gain several key benefits:

**Isolation:** Each playbook run occurs in a clean, controlled environment, avoiding interference from the host system or other processes.
**Dependency consistency and reproducibility:** The container image can include all required dependencies and specific versions, avoiding conflicts with the host system.
**Easier CI/CD integration:** This approach makes it straightforward to run Ansible playbooks in automated pipelines, regardless of the underlying build agent or environment.

## How it works

This example provides a way to execute Ansible playbooks inside a Docker container, fully managed from a Go application. The process is as follows:

1. **Custom Executor and Command Implementation:**
   `DockerExec` implements the `Executabler` interface, and `dockerCmd` implements the `Cmder` interface, encapsulating the logic to build a Docker image, create and start a container, and handle its input/output streams.

    ```go
    // dockercmd.go
    type DockerExec struct {
       client *client.Client
       Env    []string
    }
    ```

    ```go
    // Implementation of Cmder interface that runs commands inside a Docker container
    type dockerCmd struct {
    	client      *client.Client
    	containerID string
        // ...other attributes...
    }
    ```

2. **Environment and Command Setup:**
   The Go application creates a `DockerExec` instance, optionally passing environment variables. When an Ansible playbook command is requested, `DockerExec` creates a `dockerCmd` instance, setting up the command, environment, and container configuration.

   ```go
   // ansibleplaybook-docker-execution.go
   executable := NewDockerExec(
	   apiClient,
	   WithEnv([]string{"ANSIBLE_FORCE_COLOR=true"}),
   )
   ```

   ```go
   // dockercmd.go
   func (e *DockerExec) CommandContext(ctx context.Context, name string, arg ...string) exec.Cmder {
	   cmd := NewDockerCmd(e.client)
	   cmd.ContainerName = "ansible_playbook_executor"
	   cmd.Env = append([]string{}, e.Env...)
	   cmd.Cmd = append([]string{}, name)
	   cmd.Cmd = append(cmd.Cmd, arg...)
	   return cmd
   }
   ```

3. **Image Build and Container Lifecycle:**
      Before running the command, the required Docker image is built (if needed). The container is created with the specified command and environment, and its output streams are attached to pipes so the Go application can capture and process stdout and stderr.

      ```go
      // dockercmd.go (Start method - main container management steps)
      err = cmd.imageBuild(ctx, imageName)
      if err != nil {
         return fmt.Errorf("failed to build image: %w", err)
      }

      containerConfig := &container.Config{
         Image: imageName,
         Cmd:   cmd.Cmd,
         Env:   cmd.Env,
         // ...other config...
      }

      containerCreateResp, err = cmd.client.ContainerCreate(
         ctx,
         containerConfig,
         // ...host config, networking, etc...
      )
      if err != nil {
         return fmt.Errorf("failed to create container: %w", err)
      }

      attach, err = cmd.client.ContainerAttach(
         ctx,
         cmd.containerID,
         // ...attach options...
      )
      if err != nil {
         return fmt.Errorf("failed to attach to container: %w", err)
      }

      // ...setup output streaming with stdcopy...

      err = cmd.client.ContainerStart(ctx, cmd.containerID, container.StartOptions{})
      if err != nil {
         return fmt.Errorf("failed to start container: %w", err)
      }
      // ...rest of method omitted for brevity...
      ```

4. **Execution and Output Handling:**

      The container is started, and the Go application streams its output in real time. Docker combines stdout and stderr into a single stream, so the Go standard library's `stdcopy.StdCopy` function is used to demultiplex this stream and write the output to separate Go pipes for stdout and stderr. You can then read from these pipes and handle the output as needed, for example, printing to your application's stdout and stderr:

      ```go
      // dockercmd.go
      go func() {
         defer attach.Close()
         defer cmd.stdoutPipeWriter.Close()
         defer cmd.stderrPipeWriter.Close()
         // stdcopy.StdCopy demultiplexes Docker's combined output stream into separate stdout and stderr writers
         _, err := stdcopy.StdCopy(cmd.stdoutPipeWriter, cmd.stderrPipeWriter, attach.Reader)
         if err != nil {
            fmt.Println("Error copying output from container:", err)
         }
      }()

5. **Integration with go-ansible:**
   The example integrates with the `go-ansible` library, allowing you to run Ansible playbooks as if they were local commands, but actually executing them inside a Docker container for isolation and reproducibility.

   ```go
   // ansibleplaybook-docker-execution.go
   ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
	   Connection: "local",
	   Inventory:  "127.0.0.1,",
   }
   playbookCmd := playbook.NewAnsiblePlaybookCmd(
	   playbook.WithPlaybooks("site.yml"),
	   playbook.WithPlaybookOptions(ansiblePlaybookOptions),
   )
   ```

   ```go
   // ansibleplaybook-docker-execution.go
   executor := execute.NewDefaultExecute(
	   execute.WithCmd(playbookCmd),
	   execute.WithExecutable(executable),
	   execute.WithTransformers(
		   transformer.Prepend("ansibleplaybook-docker-executor example"),
	   ),
   )
   err = executor.Execute(context.TODO())
   ```

## Alternative Implementations

This example provides an implementation of the `Executabler` and `Cmder` interfaces. You can execute Ansible playbooks within a Docker container by developing a custom executor, rather than using the default `DefaultExecute` implementation. This approach allows you to tailor the execution environment and behavior to your specific requirements.
