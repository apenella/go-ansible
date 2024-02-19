
# go-ansible

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) ![Test](https://github.com/apenella/go-ansible/actions/workflows/testing.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/apenella/go-ansible)](https://goreportcard.com/report/github.com/apenella/go-ansible) [![Go Reference](https://pkg.go.dev/badge/github.com/apenella/go-ansible.svg)](https://pkg.go.dev/github.com/apenella/go-ansible)

![go-ansible-logo](docs/logo/go-ansible_logo.png "Go-ansible Logo" )

Go-ansible is a Go package that allows executing _Ansible_ commands, such as `ansible-playbook` or `ansible`, directly from Golang applications. It offers a variety of options for each command, facilitating seamless integration of _Ansible_ functionality into your projects. It is important to highlight that _go-ansible_ is not an alternative implementation of _Ansible_, but rather a wrapper around the _Ansible_ commands.
Let's dive in and explore the capabilities of _go-ansible_ together.

_**Important:** The master branch may contain unreleased or pre-released features. Be cautious when using that branch in your projects. It is recommended to use the stable releases available in the [releases](https://github.com/apenella/go-ansible/releases)._

> **Note**
> The latest major version of _go-ansible_, version _2.x_, introduced significant and breaking changes. If you are currently using a version prior to _2.x_, please refer to the [upgrade guide](https://github.com/apenella/go-ansible/blob/master/docs/upgrade_guide_to_2.x.md) for detailed information on how to migrate to version _2.x_.
> The most relevant change is that [command structs](#command-generator) no longer execute commands. So, [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) and [AnsibleAdhocCmd](#ansibleadhoccmd-struct) do not require an [executor](#executor) anymore. Instead, the _executor_ is responsible for receiving the command to execute and executing it.

- [go-ansible](#go-ansible)
  - [Install](#install)
    - [Upgrade to 1.x](#upgrade-to-1x)
    - [Upgrade to 2.x](#upgrade-to-2x)
  - [Concepts](#concepts)
    - [Executor](#executor)
    - [Command Generator](#command-generator)
    - [Results Handler](#results-handler)
  - [Getting Started](#getting-started)
    - [Create the _AnsiblePlaybookCmd_ struct](#create-the-ansibleplaybookcmd-struct)
    - [Create the _DefaultExecute_ executor](#create-the-defaultexecute-executor)
    - [Manage the output of the command execution](#manage-the-output-of-the-command-execution)
  - [Usage Reference](#usage-reference)
    - [Adhoc package](#adhoc-package)
      - [AnsibleAdhocCmd struct](#ansibleadhoccmd-struct)
      - [AnsibleAdhocOptions struct](#ansibleadhocoptions-struct)
    - [Execute package](#execute-package)
      - [Executor interface](#executor-interface)
      - [Commander interface](#commander-interface)
      - [Executabler interface](#executabler-interface)
      - [DefaultExecute struct](#defaultexecute-struct)
      - [Defining a Custom Executor](#defining-a-custom-executor)
      - [Customizing the Execution](#customizing-the-execution)
        - [Configuration package](#configuration-package)
          - [ExecutorEnvVarSetter interface](#executorenvvarsetter-interface)
          - [Ansible Configuration functions](#ansible-configuration-functions)
          - [AnsibleWithConfigurationSettingsExecute struct](#ansiblewithconfigurationsettingsexecute-struct)
        - [Measure package](#measure-package)
        - [Result package](#result-package)
          - [ResultsOutputer interface](#resultsoutputer-interface)
          - [DefaultResults struct](#defaultresults-struct)
          - [JSONStdoutCallbackResults struct](#jsonstdoutcallbackresults-struct)
          - [Transformer functions](#transformer-functions)
        - [Stdoutcallback package](#stdoutcallback-package)
          - [ExecutorStdoutCallbackSetter interface](#executorstdoutcallbacksetter-interface)
          - [Stdout Callback Execute structs](#stdout-callback-execute-structs)
    - [Inventory package](#inventory-package)
      - [AnsibleInventoryCmd struct](#ansibleinventorycmd-struct)
      - [AnsibleInventoryOptions struct](#ansibleinventoryoptions-struct)
    - [Options package](#options-package)
      - [AnsibleConnectionOptions struct](#ansibleconnectionoptions-struct)
      - [AnsiblePrivilegeEscalationOptions struct](#ansibleprivilegeescalationoptions-struct)
    - [Playbook package](#playbook-package)
      - [AnsiblePlaybookCmd struct](#ansibleplaybookcmd-struct)
      - [AnsiblePlaybookOptions struct](#ansibleplaybookoptions-struct)
    - [Vault package](#vault-package)
      - [Encrypt](#encrypt)
      - [Password](#password)
        - [Envvars](#envvars)
        - [File](#file)
        - [Resolve](#resolve)
        - [Text](#text)
  - [Examples](#examples)
  - [Contributing](#contributing)
    - [Code Of Conduct](#code-of-conduct)
  - [License](#license)

## Install

Use this command to fetch and install the latest version of _go-ansible_, ensuring you have the most up-to-date and stable release.

```sh
go get github.com/apenella/go-ansible@v2.0.0
```

### Upgrade to 1.x

If you are currently using a _go-ansible_ version before _1.x_, note that there have been significant breaking changes introduced in version _1.0.0_ and beyond. Before proceeding with the upgrade, we highly recommend reading the [changelog](https://github.com/apenella/go-ansible/blob/master/CHANGELOG.md) and the [upgrade guide](https://github.com/apenella/go-ansible/blob/master/docs/upgrade_guide_to_1.x.md) carefully. These resources provide detailed information on the changes and steps required for a smooth transition to the new version.

### Upgrade to 2.x

Version _2.0.0_ introduced notable changes since the major version _1_, including several breaking changes. The [upgrade guide](https://github.com/apenella/go-ansible/blob/master/docs/upgrade_guide_to_2.x.md) conveys the necessary information to migrate to version _2.x_. Thoroughly read that document and the [changelog](https://github.com/apenella/go-ansible/blob/master/CHANGELOG.md) before upgrading from version _1.x_ to _2.x_.

## Concepts

There are a few concepts that you need to understand before using the _go-ansible_ library. These concepts are essential to effectively use the library and to understand the examples and usage references provided in this document.

### Executor

An _executor_ is a component that executes commands and handles the results from the execution output. The library includes the [DefaultExecute](#defaultexecute-struct) executor, which is a ready-to-go implementation of an executor. If the `DefaultExecute` does not meet your requirements, you can also create a custom executor.
To know more about the `DefaultExecute`, refer to [that](#defaultexecute-struct) section.

### Command Generator

A _command generator_ or a _commander_ is responsible for generating the command to be executed. The [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) and [AnsibleAdhocCmd](#ansibleadhoccmd-struct) structs are examples of command generators. That concept has been introduced in the major version _2.0.0_.

### Results Handler

A _results handler_ or an _results outputer_ is responsible for managing the output of the command execution. The library includes two output mechanisms: the [DefaultResults](#defaultexecute-struct) and the [JSONStdoutCallbackResults](#jsonstdoutcallbackresults-struct) structs.

## Getting Started

This section will guide you through the step-by-step process of using the _go-ansible_ library. Follow these instructions to create an application that utilizes the `ansible-playbook` utility. The same guidelines can be applied to other _Ansible_ commands, such as the `ansible` command.

Before proceeding, ensure you have installed the latest version of the go-ansible library. If not, please refer to the [Installation section](#install) for instructions.

To create an application that launches the `ansible-playbook` command you need to create an [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) struct. This struct generates the _Ansible_ command to be run. Then, you need to execute the command usingng an [executor](#executor). In that guided example, you will use the [DefaultExecute](#defaultexecute-struct) executor, an _executor_ provided by the _go-ansible_ library.

### Create the _AnsiblePlaybookCmd_ struct

To execute `ansible-playbook` commands, first, define the necessary connection, playbook, and privilege escalation options.

Start by creating the [AnsiblePlaybookConnectionOptions](#ansibleconnectionoptions-struct) struct:

```go
ansiblePlaybookConnectionOptions := &options.AnsiblePlaybookConnectionOptions{
  Connection: "local",
}
```

Next, define the playbook options using the [AnsiblePlaybookOptions](#ansibleplaybookoptions-struct) struct:

```go
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Inventory: "127.0.0.1,",
}
```

Then, use the [AnsiblePlaybookPrivilegeEscalationOptions](#ansibleprivilegeescalationoptions-struct) struct to define the privilege escalation options:

```go
privilegeEscalationOptions := &options.AnsiblePlaybookPrivilegeEscalationOptions{
  Become:        true,
}
```

Finally, create the [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) struct that generates the command to execute the playbook `site.yml` using the `ansible-playbook` command:

```go
playbookCmd := &playbook.AnsiblePlaybookCmd{
  Playbook:          "site.yml",
  ConnectionOptions: ansiblePlaybookConnectionOptions,
  Options:           ansiblePlaybookOptions,
  PrivilegeEscalationOptions: privilegeEscalationOptions,
}
```

Once the `AnsiblePlaybookCmd` is defined, provide the command to an [executor](#executor) to run the command.

### Create the _DefaultExecute_ executor

We will use the [DefaultExecute](#defaultexecute-struct) struct, provided by the _go-ansible_ library, to execute the `ansible-playbook` command. It requires a [Commander](#commander-interface) responsible for generating the command to be executed. In that example, you will use the [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) previously defined.

```go
// PlaybookCmd is the Commander responsible for generating the command to execute
exec := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
)
```

Once you have defined the [DefaultExecute](#defaultexecute-struct), execute the _Ansible_ command using the following code:

```go
err := exec.Execute(context.Background())
if err != nil {
  // Manage the error
}
```

### Manage the output of the command execution

By default, the [DefaultExecute](#defaultexecute-struct) uses the [DefaultResults](#defaultresults-struct) struct to manage the output of the command execution. The `DefaultResults` struct handles the output as plain text.

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

## Usage Reference

The Usage Reference section provides an overview of the different packages and their main resources available in the _go-ansible_ library. Here you will find the details to effectively use the library to execute _Ansible_ commands, such as `ansible-playbook` and `ansible`.
For detailed information on the library's packages, structs, methods, and functions, please refer to the complete reference available [here](https://pkg.go.dev/github.com/apenella/go-ansible).

### Adhoc package

This section provides an overview of the `adhoc` package in the _go-ansible_ library, outlining its key components and functionalities.

The `github.com/apenella/go-ansible/pkg/adhoc` package facilitates the generation of _Ansible_ ad-hoc commands. It does not execute the commands directly, but instead provides the necessary structs to generate the command to be executed by an executor. The `adhoc` package includes the following essential structs for executing ad-hoc commands:

#### AnsibleAdhocCmd struct

The `AnsibleAdhocCmd` struct enables the generation of _Ansible ad-hoc_ commands. It implements the [Commander](#commander-interface) interface, so its method `Command` returns an array of strings that represents the command to be executed. An executor can use it to create the command to be executed.

The following code snippet demonstrates how to use the `AnsibleAdhocCmd` struct to generate an ad-hoc command:

```go
ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
  Inventory : "127.0.0.1,",
  ModuleName: "degut",
  Args: "msg={{ arg }}",
  ExtraVars: map[string]interface{}{
      "arg": "value",
  }
}

adhocCmd := &adhoc.AnsibleAdhocCmd{
  Pattern:           "all",
  Options:           ansibleAdhocOptions,
}

// Generate the command to be executed
cmd, err := adhocCmd.Command()
if err != nil {
  // Manage the error
}
```

#### AnsibleAdhocOptions struct

With `AnsibleAdhocOptions` struct, you can define parameters described in Ansible's manual page's `Options` section.

Additionally, users can set privilege escalation options or connection options to the [AnsibleAdhocCmd](#ansibleadhoccmd-struct). These options are defined in the `github.com/apenella/go-ansible/pkg/options` package. Refer to the [options](#options-package) section for further details.

### Execute package

The _execute_ package, available at `github.com/apenella/go-ansible/pkg/execute`, provides the [DefaultExecute](#defaultexecute-struct), a ready-to-use [executor](#executor). Additionally, the package defines some interfaces for managing the command execution and customizing the behavior of the _executor_.

Find below the main resources available in the execute package:

#### Executor interface

The `Executor` interface defines components responsible for executing external commands. The [DefaultExecute](#defaultexecute-struct) struct implements this interface. Below is the definition of the `Executor` interface:

```go
type Executor interface {
  Execute(ctx context.Context) error
}
```

#### Commander interface

The `Commander` interface defines components responsible for generating the command to be executed. The [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) and [AnsibleAdhocCmd](#ansibleadhoccmd-struct) structs implement this interface. Below is the definition of the `Commander` interface:

```go
type Commander interface {
  Command() ([]string, error)
}
```

#### Executabler interface

The `Executabler` interface defines a component required by [DefaultExecute](#defaultexecute-struct) to execute commands. Through the `Executabler` interface, you can customize the execution of commands according to your requirements.

Below is the definition of the `Executabler` interface:

```go
type Executabler interface {
  Command(name string, arg ...string) exec.Cmder
  CommandContext(ctx context.Context, name string, arg ...string) exec.Cmder
}
```

#### DefaultExecute struct

The `DefaultExecute` executor is a component provided by the _go-ansible_ library for managing the commands execution. It offers flexibility and customization options to suit various use cases.
Think of the `DefaultExecute` executor as a pipeline that handles command execution. It consists of three main stages, each managed by a different component:

- **Commander**: Generates the command to be executed.
- **Executabler**: Executes the command.
- **ResultsOutputer**: Manages the output of the command execution.

By default, the `DefaultExecute` executor uses the `OsExec` struct as the `Executabler` for executing commands, a wrapper around the `os/exec` package. It also uses the [DefaultResults](#defaultresults-struct) struct as the [ResultsOutputer](#resultsoutputer-interface) for managing the output of the command execution. However, you can customize these components to tailor the execution process to your needs.

The following functions can be provided when creating a new instance of the `DefaultExecute` to customize its behavior. All of them are available in the `github.com/apenella/go-ansible/pkg/execute` package:

- `WithCmd(cmd Commander) ExecuteOptions`: Set the component responsible for generating the command.
- `WithCmdRunDir(cmdRunDir string) ExecuteOptions`: Define the directory where the command will be executed.
- `WithEnvVars(vars map[string]string) ExecuteOptions`: Set environment variables for command execution.
- `WithExecutable(executable Executabler) ExecuteOptions`: Define the component responsible for executing the command.
- `WithOutput(output result.ResultsOutputer) ExecuteOptions`: Specify the component responsible for managing command output.
- `WithTransformers(trans ...transformer.TransformerFunc) ExecuteOptions`: Add transformers to modify command output.
- `WithWrite(w io.Writer) ExecuteOptions`: Set the writer for command output.
- `WithWriteError(w io.Writer) ExecuteOptions`: Set the writer for command error output.

The snippet below shows how to customize the `DefaultExecute` executor using the `ExecuteOptions` functions:

```go
// PlaybookCmd is the Commander responsible for generating the command to execute
playbookCmd := &playbook.AnsiblePlaybookCmd{
  Playbook:          "site.yml",
  ConnectionOptions: ansiblePlaybookConnectionOptions,
  Options:           ansiblePlaybookOptions,
  PrivilegeEscalationOptions: privilegeEscalationOptions,
}

// MyExecutabler is an hypothetical implementation of the Executabler interface
executabler := &myExecutabler{}

// MyOutputer is an hypothetical implementation of the ResultsOutputer interface
output := &myOutputer{}

// Exec is an instance of the DefaultExecute executor
exec := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
  execute.WithExecutable(executabler),
  execute.WithOutput(output),
)

// Execute the ansible-playbook command
err := exec.Execute(context.Background())
if err != nil {
  // Manage the error
}
```

For more examples and practical use cases, refer to the [examples](https://github.com/apenella/go-ansible/tree/master/examples) directory in the _go-ansible_ repository.

#### Defining a Custom Executor

If [DefaultExecute](#defaultexecute-struct) does not meet your requirements or expectations, you have the option to implement a custom _executor_. Below is an example of a custom _executor_ that demonstrates how to integrate it with the [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) or [AnsibleAdhocCmd](#ansibleadhoccmd-struct) structs to execute the playbook with your desired behavior:

```go
type MyExecutor struct {
  Prefix string
  Cmd    Commander
}

func (e *MyExecutor) Execute(ctx context.Context) error {
  // That's a dummy work
  fmt.Println(fmt.Sprintf("[%s] %s\n", e.Prefix, "I am a lazy executor and I am doing nothing"))
  return nil
}
```

The next code snippet demonstrates how to execute the `ansible-playbook` command using the custom executor:

```go
// Define the command to execute
playbookCmd := &ansibler.AnsiblePlaybookCmd{
  Playbook:          "site.yml",
  ConnectionOptions: ansiblePlaybookConnectionOptions,
  Options:           ansiblePlaybookOptions,
}

// Define an instance for the new executor and set the options
exec := &MyExecutor{
  Prefix: "go-ansible example",
  Cmd:    playbookCmd,
}

err := exec.Execute(context.Background())
if err != nil {
  // Manage the error
}
```

When you run the playbook using the custom executor, the output will be:

```sh
 [go-ansible example] I am a lazy executor and I am doing nothing
```

#### Customizing the Execution

The _go-ansible_ library offers a range of options to configure and customize the execution of the _Ansible_ commands. These customization capabilities were introduced in [version 2.0.0](https://github.com/apenella/go-ansible/blob/master/docs/upgrade_guide_to_2.x.md), when the [executor](#executor) became the central component in the execution process.

In the following sections, we will explore the components available for customizing the execution process:

##### Configuration package

The `github.com/apenella/go-ansible/pkg/execute/configuration` package provides components for configuring the _Ansible_ settings during command execution. In the following sections, we will explore the available elements for customizing the execution process.

###### ExecutorEnvVarSetter interface

The `ExecutorEnvVarSetter` interface extends the [Executor](#executor-interface) interface with the capability of setting environment variables for the command execution. The [DefaultExecute](#defaultexecute-struct) struct implements this interface. Below is the definition of the `ExecutorEnvVarSetter` interface:

```go
type ExecutorEnvVarSetter interface {
  execute.Executor
  AddEnvVar(key, value string)
}
```

###### Ansible Configuration functions

The `github.com/apenella/go-ansible/pkg/execute/configuration` package provides a set of functions for configuring _Ansible_ settings during command execution. Each function corresponds to a configuration setting available in [Ansible's reference guide](https://docs.ansible.com/ansible/latest/reference_appendices/config.html). The functions follow a consistent naming convention: `With<setting name>` or `Without<setting name>`, where `<setting name>` is the name of the _Ansible_ setting to be configured.

###### AnsibleWithConfigurationSettingsExecute struct

The `AnsibleWithConfigurationSettingsExecute` struct serves as a decorator over an [ExecutorEnvVarSetter](#executorenvvarsetter-interface), enabling configuration of _Ansible_ settings for execution. When instantiating a new `AnsibleWithConfigurationSettingsExecute`, you must provide an `ExecutorEnvVarSetter` and a list of functions for configuring Ansible settings.
Here you can see an example of how to use the `AnsibleWithConfigurationSettingsExecute` struct to configure _Ansible_ settings for execution:

```go
ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
  Connection: "local",
}

ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Inventory: "127.0.0.1,",
}

playbookCmd := &playbook.AnsiblePlaybookCmd{
  Playbooks:         []string{"site.yml"},
  ConnectionOptions: ansiblePlaybookConnectionOptions,
  Options:           ansiblePlaybookOptions,
}

exec := configuration.NewAnsibleWithConfigurationSettingsExecute(
  execute.NewDefaultExecute(
    execute.WithCmd(playbookCmd),
  ),
  configuration.WithAnsibleForceColor(),
  configuration.WithAnsibleForks(10),
  configuration.WithAnsibleHome("/path/to/ansible/home"),
  configuration.WithAnsibleHostKeyChecking(),
  configuration.WithoutAnsibleActionWarnings(),
)

err := exec.Execute(context.Background())
if err != nil {
  // Manage the error
}
```

##### Measure package

The _go-ansible_ library offers a convenient mechanism for measuring the execution time of _Ansible_ commands through the `github.com/apenella/go-ansible/pkg/execute/measure` package. This package includes the `ExecutorTimeMeasurement` struct, which acts as a decorator over an [Executor](#executor) to track the time taken for command execution.

To illustrate, consider the following code snippet, which demonstrates how to use the `ExecutorTimeMeasurement` struct to measure the time it takes to execute the `ansible-playbook` command:

```go
ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
  Connection: "local",
}

ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Inventory: "127.0.0.1,",
}

playbookCmd := &playbook.AnsiblePlaybookCmd{
  Playbooks:         []string{"site.yml"},
  ConnectionOptions: ansiblePlaybookConnectionOptions,
  Options:           ansiblePlaybookOptions,
}

executorTimeMeasurement := measure.NewExecutorTimeMeasurement(
  execute.NewDefaultExecute(
    execute.WithCmd(playbookCmd),
  ),
)

err := executorTimeMeasurement.Execute(context.Background())
if err != nil {
  // Manage the error
}

fmt.Println("Duration: ", exec.Duration().String())
```

For a complete example showcasing how to use measurement, refer to the [ansibleplaybook-time-measurement](https://github.com/apenella/go-ansible/blob/master/examples/ansibleplaybook-time-measurement/ansibleplaybook-time-measurement.go) example in the _go-ansible_ repository.

##### Result package

The `github.com/apenella/go-ansible/pkg/execute/result` package provides a set of components and subpackages to manage the output of _Ansible_ commands. The following sections describe the available elements.

###### ResultsOutputer interface

The `ResultsOutputer` interface in the `github.com/apenella/go-ansible/pkg/execute/result` package defines a component responsible for managing the output of command execution within the _go-ansible_ library. Both the [DefaultResults](#defaultresults-struct) and [JSONStdoutCallbackResults](#jsonstdoutcallbackresults-struct) structs implement this interface. Below is the definition of the ResultsOutputer interface:

```go
type ResultsOutputer interface {
  Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...OptionsFunc) error
}
```

###### DefaultResults struct

The `DefaultResults` struct, located in the `github.com/apenella/go-ansible/pkg/execute/result/default` package, serves as the default output manager for command execution within the _go-ansible_ library. It implements the [ResultsOutputer](#resultsoutputer-interface) interface, providing functionality to handle command output as plain text.

The `DefaultResults` struct reads the execution output line by line and applies a set of [transformers](#transformer-functions) to each line. A _transformer_ allows you to enrich or modify of the output according to your specific requirements. You can specify the _transformers_ during the instantiation of a new `DefaultResults` instance or when configuring the [DefaultExecute](#defaultexecute-struct) executor. Here you have an example of how to specify them when creating a new instance of [DefaultExecute](#defaultexecute-struct):

```go
exec := execute.NewDefaultExecute(
  execute.WithCmd(playbook),
  execute.WithTransformers(
    transformer.Prepend("Go-ansible example"),
    transformer.LogFormat(transformer.DefaultLogFormatLayout, transformer.Now),
  ),
)
```

###### JSONStdoutCallbackResults struct

The `JSONStdoutCallbackResults` struct, located in the `github.com/apenella/go-ansible/pkg/execute/result/json` package, is designed to handle the output of command execution when using the `JSON` stdout callback method. It implements the [ResultsOutputer](#resultsoutputer-interface) interface, providing functionality to parse and manipulate _JSON-formatted_ output from _Ansible_ commands.

The package also includes the `ParseJSONResultsStream` function, which decodes the `JSON` output into an `AnsiblePlaybookJSONResults` data structure. This structure can be further manipulated to format the `JSON` output according to specific requirements. The expected `JSON` schema from _Ansible_ output is defined in the [json.py](https://github.com/ansible/ansible/blob/v2.9.11/lib/ansible/plugins/callback/json.py) file within the _Ansible_ repository.

When creating a new instance of `JSONStdoutCallbackResults`, you can specify a list of [transformers](#transformer-functions) to be applied to the output. These transformers enrich or update the output as needed. By default, the `JSONStdoutCallbackResults` struct applies the `IgnoreMessage` transformer to ignore any `non-JSON` lines defined in the `skipPatterns` array.

Here you have the `skipPatterns` definition within the `JSONStdoutCallbackResults` code:

```go
skipPatterns := []string{
    // This pattern skips timer's callback whitelist output
    "^[\\s\\t]*Playbook run took [0-9]+ days, [0-9]+ hours, [0-9]+ minutes, [0-9]+ seconds$",
  }
```

The following code snippet demonstrates how to use the `JSONStdoutCallbackResults` struct to manage the output of command execution:

```go
executorTimeMeasurement := stdoutcallback.NewJSONStdoutCallbackExecute(
  execute.NewDefaultExecute(
    execute.WithCmd(playbookCmd),
    // The default executor writes the output to the buffer to be parsed by the JSONStdoutCallbackResults
    execute.WithWrite(io.Writer(buff)),
  ),
)

err = exec.Execute(context.TODO())
if err != nil {
  // Manage the error
}

// Parse the JSON output from the buffer
res, err = results.ParseJSONResultsStream(io.Reader(buff))
if err != nil {
  // Manage the error
}

fmt.Println(res.String())
```

For a detailed example showcasing how to use measurement, refer to the [ansibleplaybook-json-stdout](https://github.com/apenella/go-ansible/blob/master/examples/ansibleplaybook-json-stdout/ansibleplaybook-json-stdout.go) example in the _go-ansible_ repository.

###### Transformer functions

In _go-ansible_, transformer functions are essential components that enrich or update the output received from the [executor](#executor), allowing users to customize the output according to their specific requirements. Each transformer function follows the signature defined by the `TransformerFunc` type:

```go
// TransformerFunc is used to enrich or update messages before to be printed out
type TransformerFunc func(string) string
```

When the output is received from the [executor](#executor), it undergoes processing line by line, with each line being passed through the available transformers. The `github.com/apenella/go-ansible/pkg/execute/result/transformer` package provides a set of ready-to-use transformers, and users can also create custom transformers as needed.

Here you have the transformer functions available in the `github.com/apenella/go-ansible/pkg/execute/result/transformer` package:

- **Prepend**: Adds a specified prefix string to each output line.

```go
transformer.Prepend("Prefix: ")
```

- **Append**: Adds a specified suffix string to each output line.

```go
transformer.Append(" [suffix]")
```

- **LogFormat**: Prepends each output line with a date-time prefix.

```go
transformer.LogFormat(transformer.DefaultLogFormatLayout, transformer.Now)
```

- **IgnoreMessage**: Filters out output lines based on specified patterns. It uses the [regexp.MatchString](https://pkg.go.dev/regexp#MatchString) function to match the output lines with the specified patterns.

```go
skipPatterns := []string{"regexp-pattern1", "regexp-pattern2"}
transformer.IgnoreMessage(skipPatterns...)
```

##### Stdoutcallback package

The `github.com/apenella/go-ansible/pkg/execute/stdoutcallback` package in the _go-ansible_ library facilitates the management of _Ansible_'s stdout callback method. Configuring the stdout callback method typically involves two steps:

- Setting the `ANSIBLE_STDOUT_CALLBACK` environment variable to specify the desired stdout callback plugin name.
- Defining the method responsible for handling the results output from the command execution.

To streamline the configuration process, the _go-ansible_ library provides a collection of structs dedicated to setting each stdout callback method, along with the [ExecutorStdoutCallbackSetter](#executorstdoutcallbacksetter-interface) interface.

The following sections detail the available components for configuring the stdout callback method:

###### ExecutorStdoutCallbackSetter interface

The `ExecutorStdoutCallbackSetter` interface extends the capabilities of the [Executor](#executor-interface) interface by enabling the setting of environment variables and specifying the results output mechanism for command execution. The [DefaultExecute](#defaultexecute-struct) struct implements this interface, providing flexibility in configuring the stdout callback method.

Below is the definition of the `ExecutorStdoutCallbackSetter` interface:

```go
type ExecutorStdoutCallbackSetter interface {
  execute.Executor
  AddEnvVar(key, value string)
  WithOutput(output result.ResultsOutputer)
}
```

###### Stdout Callback Execute structs

The `github.com/apenella/go-ansible/pkg/execute/stdoutcallback` package provides a collection of structs designed to simplify the configuration of stdout callback methods for _Ansible_ command execution. These structs act as decorators over an [ExecutorStdoutCallbackSetter](#executorstdoutcallbacksetter-interface), allowing seamless integration of different stdout callback plugins with command execution.

Each stdout callback method in _Ansible_ corresponds to a specific struct in _go-ansible_, making it easy to select and configure the desired method. Here are the available structs:

- DebugStdoutCallbackExecute
- DefaultStdoutCallbackExecute
- DenseStdoutCallbackExecute
- JSONStdoutCallbackExecute
- MinimalStdoutCallbackExecute
- NullStdoutCallbackExecute
- OnelineStdoutCallbackExecute
- StderrStdoutCallbackExecute
- TimerStdoutCallbackExecute
- YamlStdoutCallbackExecute

For example, to configure the `JSON` stdout callback method for command execution, you can use the `JSONStdoutCallbackExecute` struct as shown in the following code snippet:

```go
execJson := stdoutcallback.NewJSONStdoutCallbackExecute(
  execute.NewDefaultExecute(
    execute.WithCmd(playbookCmd),
  )
)

err := execJson.Execute(context.Background())
if err != nil {
  // Manage the error
}
```

### Inventory package

The information provided in this section gives an overview of the `Inventory` package in `go-ansible`.

The `github.com/apenella/go-ansible/pkg/inventory` package provides the functionality to execute `ansible-inventory`. To perform these tasks, you can use the following inventory structs:

#### AnsibleInventoryCmd struct

The `AnsibleInventoryCmd` struct enables the generation of `ansible-inventory` commands. It implements the [Commander](#commander-interface) interface, so its method `Command` returns an array of strings that represents the command to be executed. An executor can use it to create the command to be executed.

> Note
> Unlike other _Ansible_ commands, the `ansible-inventory` command does not provide privilege escalation or connection options, aligning with the functionality of the command itself.

#### AnsibleInventoryOptions struct

The `AnsibleInventoryOptions` struct includes parameters described in the `Options` section of the _Ansible_ manual page. It defines the behavior of the Ansible inventory operations and specifies where to find the configuration settings.

### Options package

These options can be used to customize the behaviour of `ansible` and `ansible-playbook` commands executions.
The _go-ansible_ library provides types for defining command execution options in the `github.com/apenella/go-ansible/pkg/options` package.

#### AnsibleConnectionOptions struct

The `AnsibleConnectionOptions` includes parameters described in the Connections Options section within the _ansible_ or _ansible-playbook_'s manual page. It defines how to connect to hosts when executing _Ansible_ commands.

The struct can be used either in the [AnsibleAdhocCmd](#ansibleadhoccmd-struct) or [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) to define the connection options.

#### AnsiblePrivilegeEscalationOptions struct

The `AnsiblePrivilegeEscalationOptions` includes parameters described in the Privilege Escalation Options section within the _ansible_ or _ansible-playbook_'s manual page. It defines how to escalate privileges and become a user during _ansible_ execution.

The struct can be used either in the [AnsibleAdhocCmd](#ansibleadhoccmd-struct) or [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) to define the connection options.

### Playbook package

This section provides an overview of the `playbook` package in the _go-ansible_ library. Here are described its main components and functionalities.

The `github.com/apenella/go-ansible/pkg/playbook` package facilitates the generation of _ansible-playbook_ commands. It does not execute the commands directly, but instead provides the necessary structs to generate the command to be executed by an executor. The `playbook` package includes the following essential structs for executing ad-hoc commands:

#### AnsiblePlaybookCmd struct

The `AnsiblePlaybookCmd` struct enables the generation of _ansible-playbook_ commands. It implements the [Commander](#commander-interface) interface, so its method `Command` returns an array of strings that represents the command to be executed. An executor can use it to create the command to be executed.

Next is an example of how to use the `AnsiblePlaybookCmd` struct to generate a _ansible-playbook_ command:

```go
ansiblePlaybookConnectionOptions := &options.AnsiblePlaybookConnectionOptions{
  Connection: "local",
}

ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Inventory: "127.0.0.1,",
}

privilegeEscalationOptions := &options.AnsiblePlaybookPrivilegeEscalationOptions{
  Become:        true,
}

playbookCmd := &playbook.AnsiblePlaybookCmd{
  Playbook:          "site.yml",
  ConnectionOptions: ansiblePlaybookConnectionOptions,
  Options:           ansiblePlaybookOptions,
  PrivilegeEscalationOptions: privilegeEscalationOptions,
}

// Generate the command to be executed
cmd, err := playbookCmd.Command()
if err != nil {
  // Manage the error
}
```

#### AnsiblePlaybookOptions struct

With `AnsiblePlaybookOptions` struct, you can define parameters described in Ansible's manual page's `Options` section.

Additionally, users can set privilege escalation options or connection options to the [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct). These options are defined in the `github.com/apenella/go-ansible/pkg/options` package. Refer to the [options](#options-package) section for further details.

### Vault package

The `github.com/apenella/go-ansible/pkg/vault` package provides functionality to encrypt variables. It introduces the `VariableVaulter` struct, which is responsible for creating a `VaultVariableValue` from the value that you need to encrypt.

The `VaultVariableValue` can return the instantiated variable in JSON format.

To perform the encryption, the `vault` package relies on an `Encrypter` interface implementation.

```go
type Encrypter interface {
  Encrypt(plainText string) (string, error)
}
```

The encryption functionality is implemented in the `encrypt` package, which is described in the following section.

#### Encrypt

The `github.com/apenella/go-ansible/pkg/vault/encrypt` package is responsible for encrypting variables. It implements the `Encrypter` interface defined in the `github.com/apenella/go-ansible/pkg/vault` package.

Currently, the package provides the `EncryptString` struct, which supports the encryption of string variables. It utilizes the `github.com/sosedoff/ansible-vault-go` library for encryption.

To use the `EncryptString` struct, you need to instantiate it with a password reader. The password reader is responsible for providing the password used for encryption and it should implement the `PasswordReader` interface:

```go
type PasswordReader interface {
  Read() (string, error)
}
```

Here's an example of how to instantiate the `EncryptString` struct:

```go
encrypt := NewEncryptString(
  WithReader(
    text.NewReadPasswordFromText(
      text.WithText("secret"),
    ),
  ),
)
```

In this example, the `text.NewReadPasswordFromText` function is used to create a password reader that reads the password from a text source. The `WithText` option is used to specify the actual password value.

#### Password

The _go-ansible_ library provides a set of packages that can be used as `PasswordReader` to read the password for encryption. The following sections describe these packages and how they can be used.

##### Envvars

The `github.com/apenella/go-ansible/pkg/vault/password/envvars` package allows you to read the password from an environment variable. To use this package, you need to use the `NewReadPasswordFromEnvVar` function and provide the name of the environment variable where the password is stored using the `WithEnvVar` option:

```go
reader := NewReadPasswordFromEnvVar(
  WithEnvVar("VAULT_PASSWORD"),
)
```

In this example, the `VAULT_PASSWORD` environment variable is specified as the source of the password. The `NewReadPasswordFromEnvVar` function creates a password reader that reads the password from the specified environment variable.

Using the `envvars` package, you can conveniently read the password from an environment variable and use it for encryption.

##### File

The `github.com/apenella/go-ansible/pkg/vault/password/file` package allows you to read the password from a file, using the [afero](https://github.com/spf13/afero/blob/master/README.md) file system abstraction.

To use this package, you need to instantiate the `NewReadPasswordFromFile` function and provide the necessary options. The `WithFs` option is used to specify the file system, and the `WithFile` option is used to specify the path to the password file.

If you don't explicitly define a file system, the package uses the default file system, which is the [OsFs](https://pkg.go.dev/github.com/spf13/afero#OsFs) from the `github.com/spf13/afero` package. The OsFs represents the file system of your host machine.

Therefore, if you don't provide a specific file system using the `WithFs` option when instantiating the password reader, the file package will automatically use the [OsFs](https://pkg.go.dev/github.com/spf13/afero#OsFs) as the file system to read the password from a file.

Here's an example without specifying the file system:

```go
reader := NewReadPasswordFromFile(
  WithFile("/password"),
)
```

In this case, the [OsFs](https://pkg.go.dev/github.com/spf13/afero#OsFs) will be used to access the `/password` file on your host file system.

##### Resolve

The `github.com/apenella/go-ansible/pkg/vault/password/resolve` package provides a mechanism to resolve the password by exploring multiple `PasswordReader` implementations. It returns the first password obtained from any of the `PasswordReader` instances.

To use this package, you need to create a `NewReadPasswordResolve` instance and provide a list of `PasswordReader` implementations as arguments to the `WithReader` option:

```go
reader := NewReadPasswordResolve(
  WithReader(
    envvars.NewReadPasswordFromEnvVar(
      envvars.WithEnvVar("VAULT_PASSWORD"),
    ),
    file.NewReadPasswordFromFile(
      file.WithFs(testFs),
      file.WithFile("/password"),
    ),
  ),
)
```

In this example, the `ReadPasswordResolve` instance is created with two `PasswordReader` implementations: one that reads the password from an environment variable (`envvars.NewReadPasswordFromEnvVar`), and another that reads the password from a file (`file.NewReadPasswordFromFile`).

The `ReadPasswordResolve` will attempt to obtain the password from each `PasswordReader` in the provided order. The first successful password read will be returned. It returns an error when no password is achieved.

Using the `resolve` package, you can explore multiple `PasswordReader` implementations to resolve the password for encryption.

##### Text

The `github.com/apenella/go-ansible/pkg/vault/password/text` package provides functionality to read the password from a text source.

To use this package, you need to instantiate the `NewReadPasswordFromText` function and provide the password as a text value using the `WithText` option:

```go
reader := NewReadPasswordFromText(
  WithText("ThatIsAPassword"),
)
```

In this example, the password is directly specified as the text value "ThatIsAPassword" using the `WithText` option.

## Examples

The _go-ansible_ library includes a variety of examples that demonstrate how to use the library in different scenarios. These examples can be found in the [examples](https://github.com/apenella/go-ansible/tree/master/examples) directory of the _go-ansible_ repository.

The examples cover various use cases and provide practical demonstrations of utilizing different features and functionalities offered by _go-ansible_. They serve as a valuable resource to understand and learn how to integrate _go-ansible_ into your applications.

Feel free to explore the [examples](https://github.com/apenella/go-ansible/tree/master/examples) directory to gain insights and ideas on how to leverage the _go-ansible_ library in your projects.

Here you have a list of examples:

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

## Contributing

Thank you for your interest in contributing to go-ansible! All contributions are welcome, whether they are bug reports, feature requests, or code contributions. Please read the contributor's guide [here](https://github.com/apenella/go-ansible/blob/master/CONTRIBUTING.md) to learn more about how to contribute.

### Code Of Conduct

The _go-ansible_ project is committed to providing a friendly, safe and welcoming environment for all, regardless of gender, sexual orientation, disability, ethnicity, religion, or similar personal characteristics.

We expect all contributors, users, and community members to follow this code of conduct. This includes all interactions within the _go-ansible_ community, whether online, in person, or otherwise.

Please to know more about the code of conduct refer [here](https://github.com/apenella/go-ansible/blob/master/CODE-OF-CONDUCT.md).

## License

The _go-ansible_ library is available under [MIT](https://github.com/apenella/go-ansible/blob/master/LICENSE) license.
