
# go-ansible

![go-ansible-logo](docs/logo/go-ansible_logo.png "Go-ansible Logo" )

Go-ansible is a Go package that enables the execution of `ansible-playbook` or `ansible` commands directly from Golang applications. It supports a wide range of options for each command, enabling smooth integration of Ansible functionality into your projects.
Let's dive in and explore the capabilities of `go-ansible` together.

> **Disclaimer**: Please note that the master branch may contain unreleased features. Be aware of this when utilizing the library in your projects.

- [go-ansible](#go-ansible)
	- [Install](#install)
		- [Upgrade to 1.x](#upgrade-to-1x)
	- [Getting Started](#getting-started)
	- [Development Reference](#development-reference)
		- [Packages](#packages)
			- [Adhoc](#adhoc)
			- [Playbook](#playbook)
			- [Execute](#execute)
				- [DefaultExecute](#defaultexecute)
				- [Custom executor](#custom-executor)
				- [Measurements](#measurements)
			- [Options](#options)
				- [Ansible ad-hoc and ansible-playbook common options](#ansible-ad-hoc-and-ansible-playbook-common-options)
			- [Stdout Callback](#stdout-callback)
			- [Results](#results)
				- [Transformers](#transformers)
				- [Default](#default)
				- [JSON](#json)
					- [Manage JSON output](#manage-json-output)
	- [Examples](#examples)
	- [License](#license)

## Install

To install the latest stable version of `go-ansible`, run the following command:

```sh
go get github.com/apenella/go-ansible@v1.2.0
```

This command will fetch and install the latest version of `go-ansible`, ensuring that you have the most up-to-date and stable release.

### Upgrade to 1.x

If you are currently using a `go-ansible` version prior to 1.0.0, it's important to note that there have been significant breaking changes introduced in version 1.0.0 and beyond. Before proceeding with the upgrade, we highly recommend reading the [changelog](https://github.com/apenella/go-ansible/blob/master/CHANGELOG.md) and the [upgrade guide](https://github.com/apenella/go-ansible/blob/master/docs/upgrade_guide_to_1.0.0.md) carefully. These resources provide detailed information on the changes and steps required for a smooth transition to the new version.

## Getting Started

This section will guide you through the step-by-step process of using the `go-ansible` library. Follow these instructions to create an application that utilizes the `ansible-playbook` utility.

Before proceeding, make sure you have the latest version of the `go-ansible` library installed. If you haven't done so yet, please refer to the [Installation section](#install) for instructions on how to install the library.

At this point, you are ready to define the required structs and execute the `ansible-playbook` command using the `go-ansible` library.

First, let's define the `AnsiblePlaybookConnectionOptions` struct for connection options:

```go
ansiblePlaybookConnectionOptions := &options.AnsiblePlaybookConnectionOptions{
  Connection: "local",
}
```

Next, define the `AnsiblePlaybookOptions` struct to specify the execution behaviour and configuration location:

```go
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Inventory: "127.0.0.1,",
}
```

Then, define the `AnsiblePlaybookPrivilegeEscalationOptions` struct to specify privilege escalation requirements:

```go
privilegeEscalationOptions := &options.AnsiblePlaybookPrivilegeEscalationOptions{
  Become:        true,
  BecomeMethod:  "sudo",
}
```

Finally, create the `AnsiblePlaybookCmd` struct to define the command execution:

```go
cmd := &playbook.AnsiblePlaybookCmd{
  Playbook:          "site.yml",
  ConnectionOptions: ansiblePlaybookConnectionOptions,
  Options:           ansiblePlaybookOptions,
  PrivilegeEscalationOptions: privilegeEscalationOptions,
}
```

Once the `AnsiblePlaybookCmd` is defined, you can execute it using the Run method. If an executor is not explicitly defined, `DefaultExecute` is used with the default parameters:

```go
err := cmd.Run(context.TODO())
if err != nil {
  panic(err)
}
```

The output of the `ansible-playbook` execution will be similar to the following:

```sh
 ──
 ── PLAY [all] *********************************************************************
 ──
 ── TASK [Gathering Facts] *********************************************************
 ── ok: [127.0.0.1]
 ──
 ── TASK [simple-ansibleplaybook] **************************************************
 ── ok: [127.0.0.1] => {
 ──     "msg": "Your are running 'simple-ansibleplaybook' example"
 ── }
 ──
 ── PLAY RECAP *********************************************************************
 ── 127.0.0.1                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
 ──
 ── Playbook run took 0 days, 0 hours, 0 minutes, 0 seconds
```

## Development Reference

The development reference is a comprehensive resource that provides everything you need to know to effectively use the library `go-ansible`.

### Packages

This section describes the different packages and their resources available in the `go-ansible` library.

#### Adhoc

The information provided in this section highlights the key components and functionalities of the `adhoc` package within `go-ansible`.

The `github.com/apenella/go-ansible/pkg/adhoc` package enables you to execute Ansible ad-hoc commands. To execute these commands, you can utilize the following ad-hoc structs:

- **AnsibleAdhocCmd**: This main struct defines the Ansible ad-hoc command and specifies how to execute it. It is mandatory to define an `AnsibleAdhocCmd` to run the command. The `AnsibleAdhocCmd` requires a parameter to specify which `Executor` to use. The executor serves as the worker responsible for launching the execution. If no `Executor` is explicitly specified, the command uses the `DefaultExecute`.

- **AnsibleAdhocOptions**: This struct provides parameters described in the `Options` section within Ansible's manual page. It defines the behaviour of the Ansible execution and specifies where to find the execution configuration.

Additionally, you can provide privilege escalation options or connection options to the `AnsibleAdhocCmd`. These options are defined in the `github.com/apenella/go-ansible/pkg/options` package. Refer to the [options](#options) sections to know more about it.

#### Playbook

The information provided in this section gives an overview of the `playbook` package in `go-ansible`.

The `github.com/apenella/go-ansible/pkg/playbook` package provides the functionality to execute `ansible-playbook` commands. To run `ansible-playbook` commands, you can use the following types:

- **AnsiblePlaybookCmd**: This is the main object type that defines the `ansible-playbook` command and specifies how to execute it. It is mandatory to define an `AnsiblePlaybookCmd` to run any `ansible-playbook` command. The `AnsiblePlaybookCmd` includes a parameter that defines the `Executor` to use, which acts as the worker responsible for launching the execution. If no `Executor` is explicitly specified, the `DefaultExecute` is used by default.
- **AnsiblePlaybookOptions**: This type includes the parameters described in the `Options` section within Ansible's manual page. It defines the execution behaviour of the `ansible-playbook` and specifies where to find the execution configuration.

Additionally, you can provide privilege escalation options or connection options to the `AnsiblePlaybookCmd`. These options are defined in the `github.com/apenella/go-ansible/pkg/options` package. Refer to the [options](#options) sections to know more about it.

#### Execute

An executor in `go-ansible` is a component that executes the command and retrieves the results from stdout and stderr. The library includes a default executor implementation called `DefaultExecute`, which is located in the `github.com/apenella/go-ansible/pkg/execute` package. The `DefaultExecute` executor adheres to the `Executor` interface.

The `Executor` interface requires the implementation of the `Execute` method with the following signature:

```go
// Executor interface is satisfied by those types which has a Execute(context.Context,[]string,stdoutcallback.StdoutCallbackResultsFunc,...ExecuteOptions)error method
type Executor interface {
  Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error
}
```

##### DefaultExecute

The `DefaultExecute` executor, provided by `go-ansible`, writes the command's stdout and stderr to the system stdout by default. However, it can be easily customized and extended to handle stdout and stderr differently. This can be achieved using the `ExecuteOptions` functions, which can be passed to the executor.

```go
// ExecuteOptions is a function to set executor options
type ExecuteOptions func(Executor)
```

The `DefaultExecute` executor also allows the configuration of Ansible through environment variables. This can be done using the `WithEnvVar` function, which injects environment variables into the command before its execution. To apply Ansible parameters as environment variables, pass `WithEnvVar` as an `ExecuteOption` to `NewDefaultExecute`.

To further customize the way results are returned to the user, `go-ansible` provides transformers. These transformers can be added to the `DefaultExecute` using the `WithTransformers(...results.TransformerFunc)` function, which can be included as an `ExecuteOption`.

For more examples and practical use cases, refer to the [examples](https://github.com/apenella/go-ansible/tree/master/examples) directory in the `go-ansible` repository.

##### Custom executor

If the `DefaultExecute` executor does not meet your requirements or expectations, you have the flexibility to implement a custom executor and set it on the `AnsiblePlaybookCmd` struct. The `AnsiblePlaybookCmd` expects a struct that implements the `Executor` interface.

Here is an example of a custom executor that demonstrates how to implement a custom executor and integrate it with the `AnsiblePlaybookCmd` or `AnsibleAdhocCmd` structs to execute the playbook with your desired behaviour.

```go
type MyExecutor struct {
  Prefix string
}

// Options method is used as a helper to apply a bunch of options to executor
func (e *MyExecutor) Options(options ...execute.ExecuteOptions) {
  // apply all options to the executor
  for _, opt := range options {
    opt(e)
  }
}

// WithPrefix method is used to set the executor prefix attribute
func WithPrefix(prefix string) execute.ExecuteOptions {
  return func(e execute.Executor) {
    e.(*MyExecutor).Prefix = prefix
  }
}

func (e *MyExecutor) Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...execute.ExecuteOptions) error {
  // It is possible to apply extra options when Execute is called
  for _, opt := range options {
    opt(e)
  }
  // that's a dummy work
  fmt.Println(fmt.Sprintf("[%s] %s\n", e.Prefix, "I am MyExecutor and I am doing nothing"))
  return nil
}
```

To execute the `ansible-playbook` command using the custom executor, you can use the following code:

```go
// define an instance for the new executor and set the options
exe := &MyExecutor{}

exe.Options(
  WithPrefix("Go ansible example"),
)

playbook := &ansibler.AnsiblePlaybookCmd{
  Playbook:          "site.yml",
  ConnectionOptions: ansiblePlaybookConnectionOptions,
  Options:           ansiblePlaybookOptions,
  Exec:              exe,
}

playbook.Run(context.TODO())
```

When you run the playbook using your dummy executor, the output will be as follows:

```sh
go run myexecutor-ansibleplaybook.go
[Go ansible example] I am MyExecutor and I am doing nothing
```

##### Measurements

To facilitate taking measurements, `go-ansible` provides the `github.com/apenella/go-ansible/pkg/execute/measure` package. This package includes an `ExecutorTimeMeasurement` that acts as an `Executor` decorator, allowing you to measure the execution time that takes to finish either `ansible` or `ansible-playbook` commands.

To use the time measurement feature, you need to create an instance of `ExecutorTimeMeasurement`:

```go
executorTimeMeasurement := measure.NewExecutorTimeMeasurement(
    execute.NewDefaultExecute(),
  )
```

Next, pass the created `ExecutorTimeMeasurement` as the `Exec` attribute value to either `AnsiblePlaybookCmd` or `AnsibleAdhocCmd`:

```go
playbook := &playbook.AnsiblePlaybookCmd{
    Playbooks:         playbooksList,
    Exec:              executorTimeMeasurement,
    ConnectionOptions: ansiblePlaybookConnectionOptions,
    Options:           ansiblePlaybookOptions,
  }
```

For a detailed example showcasing how to use measurement, refer to the [ansibleplaybook-time-measurement](https://github.com/apenella/go-ansible/blob/master/examples/ansibleplaybook-time-measurement/ansibleplaybook-time-measurement.go) example in the `go-ansible repository.

#### Options

These options can be used to customize the behaviour of `ansible` and `ansible-playbook` commands executions.
The `go-ansible` library provides types for defining command execution options in the `github.com/apenella/go-ansible/pkg/options` package.

##### Ansible ad-hoc and ansible-playbook common options

- **AnsibleConnectionOptions**: This struct includes parameters described in the Connections Options section within the ansible or ansible-playbook's manual page. It defines how to connect to hosts when executing Ansible commands.
- **AnsiblePrivilegeEscalationOptions**: This struct includes parameters described in the Escalation Options section within the ansible or ansible-playbook's manual page. It defines how to escalate privileges and become a user during ansible execution.

#### Stdout Callback

In `go-ansible`, you can define a specific stdout callback method by setting the `StdoutCallback` attribute on the `AnsiblePlaybookCmd` or `AnsibleAdhocCmd` structs. This allows you to customize the output of the commands. The output is managed by a function that adheres to the following signature:

```go
// StdoutCallbackResultsFunc defines a function which manages ansible's stdout callbacks. The function expects a context, a reader that receives the data to be wrote and a writer that defines where to write the data coming from reader, Finally a list of transformers could be passed to update the output coming from the executor.
type StdoutCallbackResultsFunc func(context.Context, io.Reader, io.Writer, ...results.TransformerFunc) error
```

The functions to manage the output are defined in the `github.com/apenella/go-ansible/pkg/stdoutcallback/results` package. By utilizing these functions and defining a custom stdout callback, you can customize how the output of the execution.

#### Results

In the `github.com/apenella/go-ansible/pkg/stdoutcallback/results` package, there are different methods available to manage the outputs of Ansible commands.

##### Transformers

In `go-ansible`, a transformer is a function that enriches or updates the output received from the executor, according to your needs. The `TransformerFunc` type defines the signature of a transformer function:

```go
// TransformerFunc is used to enrich or update messages before to be printed out
type TransformerFunc func(string) string
```

When the output is received from the executor, it is processed line by line, and the transformers are applied to each line. The `github.com/apenella/go-ansible/pkg/stdoutcallback/results` package provides a set of ready-to-use transformers, but you can also write custom transformers and set them through the executor.

Here are some examples of transformers available in the results package:

- [**Prepend**](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L21): Adds a prefix string to each output line.
- [**Append**:](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L28) Adds a suffix string to each output line.
- [**LogFormat**:](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L35) Includes a date-time prefix to each output line.
- [**IgnoreMessage**:](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L44) Ignores output lines based on the patterns provided as input parameters.

##### Default

By default, the stdout callback results are managed by the [DefaultStdoutCallbackResults](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/DefaultResults.go#L14) method. This method handles the output by prepending the separator string `──` to each line when no transformer is defined. It also prepares all the transformers before invoking the worker function responsible for writing the output to the `io.Writer`.

The [DefaultStdoutCallbackResults](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/DefaultResults.go#L14) method ensures that the output is formatted correctly and provides a basic handling of the results when no specific transformers are applied.

##### JSON

When the stdout callback method is set to `JSON` format, the output is managed by the [JSONStdoutCallbackResults](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/JSONResults.go#L151) method. This method prepares the worker output function to use the [IgnoreMessage](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L44) transformer, which ignores any non-JSON lines. Other transformers are ignored, except for those specific to [JSONStdoutCallbackResults](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/JSONResults.go#L151).

Within the [JSONStdoutCallbackResults](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/JSONResults.go#L151) function, there is an array called `skipPatterns` that contains matching expressions for lines that should be ignored. These patterns are used to skip specific lines that may not be relevant to the JSON output.

Here is an example of the skipPatterns array:

```go
skipPatterns := []string{
    // This pattern skips timer's callback whitelist output
    "^[\\s\\t]*Playbook run took [0-9]+ days, [0-9]+ hours, [0-9]+ minutes, [0-9]+ seconds$",
  }
```

###### Manage JSON output

The [JSONStdoutCallbackResults](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/JSONResults.go#L151) method writes the `JSON` output to the provided `io.Writer` parameter. The `github.com/apenella/go-ansible/pkg/stdoutcallback/results` package includes the `ParseJSONResultsStream` function, which can be used to decode the JSON output into an `AnsiblePlaybookJSONResults` data structure. You can manipulate this data structure to format the JSON output according to your specific needs.

The expected JSON schema from `ansible-playbook` is defined [here](https://github.com/ansible/ansible/blob/v2.9.11/lib/ansible/plugins/callback/json.py) file within the Ansible repository.

## Examples

The `go-ansible` library includes a variety of examples that demonstrate how to use the library in different scenarios. These examples can be found in the [examples](https://github.com/apenella/go-ansible/tree/master/examples) directory of the `go-ansible` repository.

The examples cover various use cases and provide practical demonstrations of utilizing different features and functionalities offered by `go-ansible`. They serve as a valuable resource to understand and learn how to integrate go-ansible into your applications.

Feel free to explore the [examples](https://github.com/apenella/go-ansible/tree/master/examples) directory to gain insights and ideas on how to leverage the go-ansible library in your projects.

Here you have a list of the examples:

- [ansibleadhoc-command-module](https://github.com/apenella/go-ansible/tree/master/examples/ansibleadhoc-command-module)
- [ansibleadhoc-simple](https://github.com/apenella/go-ansible/tree/master/examples/ansibleadhoc-simple)
- [ansibleplaybook-become](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-become)
- [ansibleplaybook-cobra-cmd](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-cobra-cmd)
- [ansibleplaybook-custom-transformer](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-custom-transformer)
- [ansibleplaybook-extravars-file](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-extravars-file)
- [ansibleplaybook-json-stdout](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-json-stdout)
- [ansibleplaybook-myexecutor](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-myexecutor)
- [ansibleplaybook-signals-and-cancellation](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-signals-and-cancellation)
- [ansibleplaybook-simple](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-simple)
- [ansibleplaybook-simple-embedfs](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-simple-embedfs)
- [ansibleplaybook-simple-with-prompt](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-simple-with-prompt)
- [ansibleplaybook-skipping-failing](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-skipping-failing)
- [ansibleplaybook-time-measurement](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-time-measurement)
- [ansibleplaybook-walk-through-json-output](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-walk-through-json-output)
- [ansibleplaybook-with-executor-time-measurament](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-with-executor-time-measurament)
- [ansibleplaybook-with-timeout](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-with-timeout)
- [ansibleplaybook-with-vaulted-extravar](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-with-vaulted-extravar)

## License

go-ansible is available under [MIT](https://github.com/apenella/go-ansible/blob/master/LICENSE) license.
