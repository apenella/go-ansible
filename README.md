
# go-ansible

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) ![Test](https://github.com/apenella/go-ansible/actions/workflows/ci.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/apenella/go-ansible/v2)](https://goreportcard.com/report/github.com/apenella/go-ansible/v2) [![Go Reference](https://pkg.go.dev/badge/github.com/apenella/go-ansible/v2.svg)](https://pkg.go.dev/github.com/apenella/go-ansible/v2)[![Static Badge](https://img.shields.io/badge/Changelog-CHANGELOG.md-blue)
](https://github.com/apenella/go-ansible/blob/master/CHANGELOG.md)

![go-ansible-logo](docs/logo/go-ansible_logo.png "Go-ansible Logo" )

Go-ansible is a Go package that allows executing _Ansible_ commands, such as `ansible-playbook`, `ansible-inventory`, or `ansible`, directly from Golang applications. It offers a variety of options for each command, facilitating seamless integration of _Ansible_ functionality into your projects. It is important to highlight that _go-ansible_ is not an alternative implementation of _Ansible_, but rather a wrapper around the _Ansible_ commands.
Let's dive in and explore the capabilities of _go-ansible_ together.

_**Important:** The master branch may contain unreleased or pre-released features. Be cautious when using that branch in your projects. It is recommended to use the stable releases available in the [releases](https://github.com/apenella/go-ansible/releases)._

> **Note**
> The latest major version of _go-ansible_, version _2.x_, introduced significant and breaking changes. The first change is that the package name has been changed from `github.com/apenella/go-ansible` to `github.com/apenella/go-ansible/v2`. So, you need to update your import paths to use the new module name.
> To migrate your code from prior version to _2.x_, please refer to the [upgrade guide](https://github.com/apenella/go-ansible/blob/master/docs/upgrade_guide_to_2.x.md) for detailed information on how to migrate to version _2.x_.
> The most relevant change is that [command structs](#command-generator) no longer execute commands. So, [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct), [AnsibleInventoryCmd](#ansibleinventorycmd-struct) and [AnsibleAdhocCmd](#ansibleadhoccmd-struct) do not require an [executor](#executor) anymore. Instead, the _executor_ is responsible for receiving the command to execute and executing it.

- [go-ansible](#go-ansible)
  - [Install](#install)
    - [Upgrade to 1.x](#upgrade-to-1x)
    - [Upgrade to 2.x](#upgrade-to-2x)
  - [Concepts](#concepts)
    - [Executor](#executor)
    - [Command Generator](#command-generator)
    - [Results Handler](#results-handler)
  - [Considerations](#considerations)
    - [Execute go-ansible inside a container](#execute-go-ansible-inside-a-container)
    - [Disable pseudo-terminal allocation](#disable-pseudo-terminal-allocation)
  - [Getting Started](#getting-started)
    - [Create the _AnsiblePlaybookCmd_ struct](#create-the-ansibleplaybookcmd-struct)
    - [Create the _DefaultExecute_ executor](#create-the-defaultexecute-executor)
    - [Manage the output of the command execution](#manage-the-output-of-the-command-execution)
  - [Usage Reference](#usage-reference)
    - [Adhoc package](#adhoc-package)
      - [AnsibleAdhocCmd struct](#ansibleadhoccmd-struct)
      - [AnsibleAdhocExecute struct](#ansibleadhocexecute-struct)
      - [AnsibleAdhocOptions struct](#ansibleadhocoptions-struct)
    - [Execute package](#execute-package)
      - [Executor interface](#executor-interface)
      - [Commander interface](#commander-interface)
      - [ErrorEnricher interface](#errorenricher-interface)
      - [Executabler interface](#executabler-interface)
      - [DefaultExecute struct](#defaultexecute-struct)
      - [Defining a Custom Executor](#defining-a-custom-executor)
      - [Customizing the Execution](#customizing-the-execution)
        - [Configuration package](#configuration-package)
          - [ExecutorEnvVarSetter interface](#executorenvvarsetter-interface)
          - [Ansible Configuration functions](#ansible-configuration-functions)
          - [AnsibleWithConfigurationSettingsExecute struct](#ansiblewithconfigurationsettingsexecute-struct)
        - [Exec package](#exec-package)
          - [Cmder interface](#cmder-interface)
          - [Cmd struct](#cmd-struct)
          - [OsExec struct](#osexec-struct)
        - [Measure package](#measure-package)
        - [Result package](#result-package)
          - [ResultsOutputer interface](#resultsoutputer-interface)
          - [DefaultResults struct](#defaultresults-struct)
          - [JSONStdoutCallbackResults struct](#jsonstdoutcallbackresults-struct)
          - [Transformer functions](#transformer-functions)
        - [Stdoutcallback package](#stdoutcallback-package)
          - [ExecutorStdoutCallbackSetter interface](#executorstdoutcallbacksetter-interface)
          - [ExecutorQuietStdoutCallbackSetter interface](#executorquietstdoutcallbacksetter-interface)
          - [Stdout Callback Execute structs](#stdout-callback-execute-structs)
        - [Workflow package](#workflow-package)
          - [WorkflowExecute struct](#workflowexecute-struct)
    - [Galaxy package](#galaxy-package)
      - [Galaxy Collection Install package](#galaxy-collection-install-package)
        - [AnsibleGalaxyCollectionInstallCmd struct](#ansiblegalaxycollectioninstallcmd-struct)
        - [AnsibleGalaxyCollectionInstallOptions struct](#ansiblegalaxycollectioninstalloptions-struct)
      - [Galaxy Role Install package](#galaxy-role-install-package)
        - [AnsibleGalaxyRoleInstallCmd struct](#ansiblegalaxyroleinstallcmd-struct)
        - [AnsibleGalaxyRoleInstallOptions struct](#ansiblegalaxyroleinstalloptions-struct)
    - [Inventory package](#inventory-package)
      - [AnsibleInventoryCmd struct](#ansibleinventorycmd-struct)
      - [AnsibleInventoryExecute struct](#ansibleinventoryexecute-struct)
      - [AnsibleInventoryOptions struct](#ansibleinventoryoptions-struct)
    - [Playbook package](#playbook-package)
      - [AnsiblePlaybookCmd struct](#ansibleplaybookcmd-struct)
      - [AnsiblePlaybookErrorEnrich struct](#ansibleplaybookerrorenrich-struct)
      - [AnsiblePlaybookExecute struct](#ansibleplaybookexecute-struct)
      - [AnsiblePlaybookOptions struct](#ansibleplaybookoptions-struct)
    - [Vault package](#vault-package)
      - [Encrypt](#encrypt)
      - [Password](#password)
        - [Envvars](#envvars)
        - [File](#file)
        - [Resolve](#resolve)
        - [Text](#text)
  - [Examples](#examples)
  - [Development Reference](#development-reference)
    - [Development Environment](#development-environment)
      - [Testing](#testing)
      - [Static Analysis](#static-analysis)
    - [Contributing](#contributing)
    - [Code Of Conduct](#code-of-conduct)
  - [License](#license)

## Install

Use this command to fetch and install the _go-ansible_ module. You can install the release candidate version by executing the following command:

```sh
go get github.com/apenella/go-ansible/v2@v2.1.0
```

You can also install the previous stable version by executing the following command:

```sh
go get github.com/apenella/go-ansible
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

A _results handler_ or a _results outputer_ is responsible for managing the output of the command execution. The library includes two output mechanisms: the [DefaultResults](#defaultresults-struct) and the [JSONStdoutCallbackResults](#jsonstdoutcallbackresults-struct) structs.

## Considerations

Before you proceed further, please take note of the following considerations to ensure optimal usage of the go-ansible library.

### Execute go-ansible inside a container

When executing _Ansible_ commands using the _go-ansible_ library inside a container, ensure that the container has configured an init system. The init system is necessary to manage the child processes created by the _Ansible_ commands. If the container does not have an init system, the child processes may not be correctly managed, leading to unexpected behavior such as zombie processes.

You can read more about that in the issue [139](https://github.com/apenella/go-ansible/issues/139) and [here](https://github.com/ansible/ansible/issues/49270#issuecomment-462306244).

### Disable pseudo-terminal allocation

_Ansible_ commands force the pseudo-terminal allocation when executed in a terminal. That configuration can cause the SSH connection leave zombie processes when the command finished. If you are experiencing this issue, you can disable the pseudo-terminal allocation by setting the `-T` to the SSH extra arguments, which will disable the pseudo-terminal allocation.

You can read more about that in the issue [139](https://github.com/apenella/go-ansible/issues/139) and [here](https://groups.google.com/g/ansible-project/c/IQoTNwDBIiA/m/qiHUTgg31lkJ).

## Getting Started

This section will guide you through the step-by-step process of using the _go-ansible_ library. Follow these instructions to create an application that utilizes the `ansible-playbook` utility. The same guidelines can be applied to other _Ansible_ commands, such as the `ansible` or `ansible-inventory` command.

> **Note**
> The following example will guide you through a complete process of creating all the components necessary to execute an `ansible-playbook` command. For a simpler example utilizing the [AnsiblePlaybookExecute](#ansibleplaybookexecute-struct) struct, please refer to the [ansibleplaybook-simple](https://github.com/apenella/go-ansible/blob/master/examples/ansibleplaybook-simple/ansibleplaybook-simple.go) example in the repository.

Before proceeding, ensure you have installed the latest version of the _go-ansible_ library. If not, please refer to the [Installation section](#install) for instructions.

To create an application that launches the `ansible-playbook` command you need to create an [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) struct. This struct generates the _Ansible_ command to be run. Then, you need to execute the command using an [executor](#executor)](#executor). In that guided example, you will use the [DefaultExecute](#defaultexecute-struct) executor, an _executor_ provided by the _go-ansible_ library.

### Create the _AnsiblePlaybookCmd_ struct

To execute `ansible-playbook` commands, first, define the necessary connection, playbook, and privilege escalation options.

Start by creating the [AnsiblePlaybookOptions](#ansibleplaybookoptions-struct) struct:

```go
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Become:     true,
  Connection: "local",
  Inventory:  "127.0.0.1,",
}
```

Finally, create the [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) struct that generates the command to execute the playbook `site.yml` using the `ansible-playbook` command:

```go
playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml", "site2.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

```

Once the `AnsiblePlaybookCmd` is defined, provide the command to an [executor](#executor) to run the command.

### Create the _DefaultExecute_ executor

We will use the [DefaultExecute](#defaultexecute-struct) struct, provided by the _go-ansible_ library, to execute the `ansible-playbook` command. It requires a [Commander](#commander-interface) responsible for generating the command to be executed. In that example, you will use the [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) previously defined.

```go
// PlaybookCmd is the Commander responsible for generating the command to execute
exec := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
  execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
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
For detailed information on the library's packages, structs, methods, and functions, please refer to the complete reference available [here](https://pkg.go.dev/github.com/apenella/go-ansible/v2).

### Adhoc package

This section provides an overview of the `adhoc` package in the _go-ansible_ library, outlining its key components and functionalities.

The `github.com/apenella/go-ansible/v2/pkg/adhoc` package facilitates the generation of _Ansible_ ad-hoc commands. It does not execute the commands directly, but instead provides the necessary structs to generate the command to be executed by an executor. The `adhoc` package includes the following essential structs for executing ad-hoc commands:

#### AnsibleAdhocCmd struct

The `AnsibleAdhocCmd` struct enables the generation of _Ansible ad-hoc_ commands. It implements the [Commander](#commander-interface) interface, so its method `Command` returns an array of strings that represents the command to be executed. An executor can use it to create the command to be executed.

The package provides the `NewAnsibleAdhocCmd` function to create a new instance of the `AnsibleAdhocCmd` struct. The function accepts a list of options to customize the ad-hoc command. The following functions are available:

- `WithAdhocOptions(options *AnsibleAdhocOptions) AdhocOptionsFunc`: Set the ad-hoc options for the command.
- `WithBinary(binary string) AdhocOptionsFunc`: Set the binary for the ad-hoc command.
- `WithPattern(pattern string) AdhocOptionsFunc`: Set the pattern for the ad-hoc command.

The following code snippet demonstrates how to use the `AnsibleAdhocCmd` struct to generate an ad-hoc command:

```go
ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
  Inventory:  "127.0.0.1,",
  ModuleName: "debug",
  Args:       "msg={{ arg }}",
  ExtraVars:  map[string]interface{}{
      "arg": "value",
  }
}

adhocCmd := adhoc.NewAnsibleAdhocCmd(
  adhoc.WithPattern("all"),
  adhoc.WithAdhocOptions(ansibleAdhocOptions),
)

// Generate the command to be executed
cmd, err := adhocCmd.Command()
if err != nil {
  // Manage the error
}
```

#### AnsibleAdhocExecute struct

The `AnsibleAdhocExecute` struct serves as a streamlined [executor](#executor) for running `ansible` command. It encapsulates the setup process for both the [command generator](#command-generator) and _executor_. This _executor_ is particularly useful when no additional configuration or customization is required.

The following methods are available to set attributes for the [AnsibleAdhocCmd](#ansibleadhoccmd-struct) struct:

- `WithBinary(binary string) *AnsibleAdhocExecute`: The method sets the `Binary` attribute.
- `WithAdhocOptions(options *AnsibleAdhocOptions) *AnsibleAdhocExecute`: The method sets the `AdhocOptions`  attribute.

Here is an example of launching an `ansible` command using `AnsibleAdhocExecute`:

```go
err := adhoc.NewAnsibleAdhocExecute("all").
  WithAdhocOptions(ansibleAdhocOptions).
  Execute(context.TODO())

if err != nil {
  // Manage the error
}
```

#### AnsibleAdhocOptions struct

With `AnsibleAdhocOptions` struct, you can define parameters described in Ansible's manual page's `Options` section. On the same struct, you can define the connection options and privilage escalation options.

### Execute package

The _execute_ package, available at `github.com/apenella/go-ansible/v2/pkg/execute`, provides the [DefaultExecute](#defaultexecute-struct), a ready-to-use [executor](#executor). Additionally, the package defines some interfaces for managing the command execution and customizing the behavior of the _executor_.

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
  String() string
}
```

#### ErrorEnricher interface

The `ErrorEnricher` interface defines components responsible for enriching the error message. The [AnsiblePlaybookErrorEnrich](#ansibleplaybookerrorenrich-struct) struct implements this interface.
The [DefaultExecute](#defaultexecute-struct) struct uses this interface to enrich the error message. Below is the definition of the `ErrorEnricher` interface:

```go
type ErrorEnricher interface {
  Enrich(err error) error
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

The `DefaultExecute` executor is a component provided by the _go-ansible_ library for managing the execution of the commands. It offers flexibility and customization options to suit various use cases.
Think of the `DefaultExecute` executor as a pipeline that handles command execution. It consists of three main stages, each managed by a different component:

- **Commander**: Generates the command to be executed.
- **Executabler**: Executes the command.
- **ResultsOutputer**: Manages the output of the command execution.

By default, the `DefaultExecute` executor uses the `OsExec` struct as the `Executabler` for executing commands, a wrapper around the `os/exec` package. It also uses the [DefaultResults](#defaultresults-struct) struct as the [ResultsOutputer](#resultsoutputer-interface) for managing the output of the command execution. However, you can customize these components to tailor the execution process to your needs.

The following functions can be provided when creating a new instance of the `DefaultExecute` to customize its behavior. All of them are available in the `github.com/apenella/go-ansible/v2/pkg/execute` package:

- `WithCmd(cmd Commander) ExecuteOptions`: Set the component responsible for generating the command.
- `WithCmdRunDir(cmdRunDir string) ExecuteOptions`: Define the directory where the command will be executed.
- `WithEnvVars(vars map[string]string) ExecuteOptions`: Set environment variables for command execution.
- `WithErrorEnricher(errEnricher ErrorEnricher) ExecuteOptions`: Define the component responsible for enriching the error message.
- `WithExecutable(executable Executabler) ExecuteOptions`: Define the component responsible for executing the command.
- `WithOutput(output result.ResultsOutputer) ExecuteOptions`: Specify the component responsible for managing command output.
- `WithTransformers(trans ...transformer.TransformerFunc) ExecuteOptions`: Add transformers to modify command output.
- `WithWrite(w io.Writer) ExecuteOptions`: Set the writer for command output.
- `WithWriteError(w io.Writer) ExecuteOptions`: Set the writer for command error output.

The snippet below shows how to customize the `DefaultExecute` executor using the `ExecuteOptions` functions:

```go
// PlaybookCmd is the Commander responsible for generating the command to execute
playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

// MyExecutabler is an hypothetical implementation of the Executabler interface
executabler := &myExecutabler{}

// MyOutputer is an hypothetical implementation of the ResultsOutputer interface
output := &myOutputer{}

// Exec is an instance of the DefaultExecute executor
exec := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
  execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
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
playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

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

The `github.com/apenella/go-ansible/v2/pkg/execute/configuration` package provides components for configuring the _Ansible_ settings during command execution. In the following sections, we will explore the available elements for customizing the execution process.

###### ExecutorEnvVarSetter interface

The `ExecutorEnvVarSetter` interface extends the [Executor](#executor-interface) interface with the capability of setting environment variables for the command execution. The [DefaultExecute](#defaultexecute-struct) struct implements this interface. Below is the definition of the `ExecutorEnvVarSetter` interface:

```go
type ExecutorEnvVarSetter interface {
  execute.Executor
  AddEnvVar(key, value string)
}
```

###### Ansible Configuration functions

The `github.com/apenella/go-ansible/v2/pkg/execute/configuration` package provides a set of functions for configuring _Ansible_ settings during command execution. Each function corresponds to a configuration setting available in [Ansible's reference guide](https://docs.ansible.com/ansible/latest/reference_appendices/config.html). The functions follow a consistent naming convention: `With<setting name>` or `Without<setting name>`, where `<setting name>` is the name of the _Ansible_ setting to be configured.

###### AnsibleWithConfigurationSettingsExecute struct

The `AnsibleWithConfigurationSettingsExecute` struct serves as a decorator over an [ExecutorEnvVarSetter](#executorenvvarsetter-interface), enabling configuration of _Ansible_ settings for execution. When instantiating a new `AnsibleWithConfigurationSettingsExecute`, you must provide an `ExecutorEnvVarSetter` and a list of functions for configuring Ansible settings.
Here you can see an example of how to use the `AnsibleWithConfigurationSettingsExecute` struct to configure _Ansible_ settings for execution:

```go
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Connection: "local",
  Inventory:  "127.0.0.1,",
}

playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

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

##### Exec package

The `github.com/apenella/go-ansible/v2/pkg/execute/exec` package abstracts the execution of external commands and serves as a wrapper around the `os/exec` package. It includes the following components:

###### Cmder interface

The `Cmder` interface defines the methods available in the [Cmd](#cmd-struct) struct, which replicate those in the `os/exec.Cmd` struct. The `Cmder` interface abstracts the execution of external commands, enabling the use of additional components to customize the execution process and manage command output. Below is the definition of the `Cmder` interface:

```go
type Cmder interface {
  CombinedOutput() ([]byte, error)
  Environ() []string
  Output() ([]byte, error)
  Run() error
  Start() error
  StderrPipe() (io.ReadCloser, error)
  StdinPipe() (io.WriteCloser, error)
  StdoutPipe() (io.ReadCloser, error)
  String() string
  Wait() error
}
```

###### Cmd struct

The `Cmd` struct acts as a wrapper for the `os/exec.Cmd` struct. It is utilized by the [OsExec](#osexec-struct) struct to execute external commands.

###### OsExec struct

The `OsExec` struct replicates the behaviour of the `Command` and `CommandContext` functions from the `os/exec` package. However, instead of returning the `*os/exec.Cmd` struct, these functions return a [Cmder](#cmder-interface) interface.

This abstraction facilitates the use of additional components for executing external commands, customizing the execution process, and managing command output. Another benefit of this abstraction is that it allows for mocking command execution in tests.

##### Measure package

The _go-ansible_ library offers a convenient mechanism for measuring the execution time of _Ansible_ commands through the `github.com/apenella/go-ansible/v2/pkg/execute/measure` package. This package includes the `ExecutorTimeMeasurement` struct, which acts as a decorator over an [Executor](#executor) to track the time taken for command execution.

To illustrate, consider the following code snippet, which demonstrates how to use the `ExecutorTimeMeasurement` struct to measure the time it takes to execute the `ansible-playbook` command:

```go
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Connection: "local",
  Inventory:  "127.0.0.1,",
}

playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

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

The `github.com/apenella/go-ansible/v2/pkg/execute/result` package provides a set of components and subpackages to manage the output of _Ansible_ commands. The following sections describe the available elements.

###### ResultsOutputer interface

The `ResultsOutputer` interface in the `github.com/apenella/go-ansible/v2/pkg/execute/result` package defines a component responsible for managing the output of command execution within the _go-ansible_ library. Both the [DefaultResults](#defaultresults-struct) and [JSONStdoutCallbackResults](#jsonstdoutcallbackresults-struct) structs implement this interface. Below is the definition of the ResultsOutputer interface:

```go
type ResultsOutputer interface {
  Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...OptionsFunc) error
}
```

###### DefaultResults struct

The `DefaultResults` struct, located in the `github.com/apenella/go-ansible/v2/pkg/execute/result/default` package, serves as the default output manager for command execution within the _go-ansible_ library. It implements the [ResultsOutputer](#resultsoutputer-interface) interface, providing functionality to handle command output as plain text.

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

The `JSONStdoutCallbackResults` struct, located in the `github.com/apenella/go-ansible/v2/pkg/execute/result/json` package, is designed to handle the output of command execution when using the `JSON` stdout callback method. It implements the [ResultsOutputer](#resultsoutputer-interface) interface, providing functionality to parse and manipulate _JSON-formatted_ output from _Ansible_ commands.

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

When the output is received from the [executor](#executor), it undergoes processing line by line, with each line being passed through the available transformers. The `github.com/apenella/go-ansible/v2/pkg/execute/result/transformer` package provides a set of ready-to-use transformers, and users can also create custom transformers as needed.

Here you have the transformer functions available in the `github.com/apenella/go-ansible/v2/pkg/execute/result/transformer` package:

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

The `github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback` package in the _go-ansible_ library facilitates the management of _Ansible_'s stdout callback method. Configuring the stdout callback method typically involves two steps:

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

###### ExecutorQuietStdoutCallbackSetter interface

The `ExecutorQuietStdoutCallbackSetter` interface in the `github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback` package extends the [ExecutorStdoutCallbackSetter](#executorstdoutcallbacksetter-interface) interface with the capability to remove the verbosity of the command execution output. The [DefaultExecute](#defaultexecute-struct) struct implements this interface, allowing you to silence the output of the command execution.

That interface is required by the [JSONStdoutCallbackExecute](#stdout-callback-execute-structs) struct to remove the verbosity of the command execution output.

The next code snippet shows the definition of the `ExecutorQuietStdoutCallbackSetter` interface:

```go
type ExecutorQuietStdoutCallbackSetter interface {
  ExecutorStdoutCallbackSetter
  Quiet()
}
```

###### Stdout Callback Execute structs

The `github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback` package provides a collection of structs designed to simplify the configuration of stdout callback methods for _Ansible_ command execution. These structs act as decorators over an [ExecutorStdoutCallbackSetter](#executorstdoutcallbacksetter-interface), allowing seamless integration of different stdout callback plugins with command execution.

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

##### Workflow package

The `github.com/apenella/go-ansible/v2/pkg/execute/workflow` package provides an executor that allows you to run a sequence of executors.

###### WorkflowExecute struct

The `WorkflowExecute` struct is responsible for managing the execution of a sequence of executors. It allows you to define a workflow for running _Ansible_ commands. Additionally, the `WorkflowExecute` implements the [Executor](#executor-interface) interface, enabling you to apply some decorators such as the [ExecutorTimeMeasurement](#measure-package) to the workflow.

By default, when an _executor_ execution in the sequence fails, the `WorkflowExecute` stops the execution of the remaining executors. However, you can customize this behaviour by setting the `ContinueOnError` attribute to `true`. In that case, if one executor fails, the `WorkflowExecute` continues with the remaining executors and raises an error at the end of the execution.

The `WorkflowExecute` struct provides the following methods to setup the execution process:

- `AppendExecutor(exec Executor) *WorkflowExecute`: Appends an executor to the sequence.
- `Execute(ctx context.Context) error`: Executes the sequence of executors.
- `WithContinueOnError() *WorkflowExecute`: Sets the `ContinueOnError` attribute to `true`.

Here is an example of how to use the `WorkflowExecute` struct to run a sequence of executors:

```go
exec1 := playbook.NewAnsiblePlaybookExecute("first.yml").
    WithPlaybookOptions(ansiblePlaybookOptions)

exec2 := playbook.NewAnsiblePlaybookExecute("second.yml").
    WithPlaybookOptions(ansiblePlaybookOptions)

err := workflow.NewWorkflowExecute(exec1, exec2).
    WithContinueOnError().
    Execute(context.TODO())

if err != nil {
  // Manage the error
}
```

### Galaxy package

The `go-ansible` library provides you with the ability to interact with the _Ansible Galaxy_ command-line tool. To do that it includes the following package:

- [github.com/apenella/go-ansible/v2/pkg/galaxy/collection/install](#galaxy-collection-install-package): Provides the functionality to install collections from the _Ansible Galaxy_.
- [github.com/apenella/go-ansible/v2/pkg/galaxy/role/install](#galaxy-role-install-package): Provides the functionality to install roles from the _Ansible Galaxy_.

#### Galaxy Collection Install package

The `github.com/apenella/go-ansible/v2/pkg/galaxy/collection/install` package allows you to install collections from the _Ansible Galaxy_ using the `ansible-galaxy` command. The package provides the following structs and functions:

##### AnsibleGalaxyCollectionInstallCmd struct

The `AnsibleGalaxyCollectionInstallCmd` struct enables the generation of `ansible-galaxy` commands to install collections. It implements the [Commander](#commander-interface) interface, so its method `Command` returns an array of strings that represents the command to be executed. An executor can use it to create the command to be executed.

The package provides the `NewAnsibleGalaxyCollectionInstallCmd` function to create a new instance of the `AnsibleGalaxyCollectionInstallCmd` struct. The function accepts a list of options to customize the `ansible-galaxy` command. The following functions are available:

- `WithBinary(binary string) AnsibleGalaxyCollectionInstallOptionsFunc`: Set the binary for the `ansible-galaxy` command.
- `WithCollectionInstallOptions(options *AnsibleGalaxyCollectionInstallOptions) AnsibleGalaxyCollectionInstallOptionsFunc`: Set the collection install options for the command.
- `WithCollectionNames(collectionNames ...string) AnsibleGalaxyCollectionInstallOptionsFunc`: Set the collection names for the `ansible-galaxy` command.

##### AnsibleGalaxyCollectionInstallOptions struct

The `AnsibleGalaxyCollectionInstallOptions` struct includes parameters described in the `Options` section of the _Ansible Galaxy_ manual page. It defines the behavior of the _Ansible Galaxy_ collection installation operations and specifies where to find the configuration settings.

#### Galaxy Role Install package

The `github.com/apenella/go-ansible/v2/pkg/galaxy/role/install` package allows you to install roles from the _Ansible Galaxy_ using the `ansible-galaxy` command. The package provides the following structs and functions:

##### AnsibleGalaxyRoleInstallCmd struct

The `AnsibleGalaxyRoleInstallCmd` struct enables the generation of `ansible-galaxy` commands to install roles. It implements the [Commander](#commander-interface) interface, so its method `Command` returns an array of strings that represents the command to be executed. An executor can use it to create the command to be executed.

The package provides the `NewAnsibleGalaxyRoleInstallCmd` function to create a new instance of the `AnsibleGalaxyRoleInstallCmd` struct. The function accepts a list of options to customize the `ansible-galaxy` command. The following functions are available:

- `WithBinary(binary string) AnsibleGalaxyRoleInstallOptionsFunc`: Set the binary for the `ansible-galaxy` command.
- `WithGalaxyRoleInstallOptions(options *AnsibleGalaxyRoleInstallOptions) AnsibleGalaxyRoleInstallOptionsFunc`: Set the role install options for the command.
- `WithRoleNames(roleNames ...string) AnsibleGalaxyRoleInstallOptionsFunc`: Set the role names for the `ansible-galaxy` command.

##### AnsibleGalaxyRoleInstallOptions struct

The `AnsibleGalaxyRoleInstallOptions` struct includes parameters described in the `Options` section of the _Ansible Galaxy_ manual page. It defines the behavior of the _Ansible Galaxy_ role installation operations and specifies where to find the configuration settings.

### Inventory package

The information provided in this section gives an overview of the `Inventory` package in `go-ansible`.

The `github.com/apenella/go-ansible/v2/pkg/inventory` package provides the functionality to execute `ansible-inventory`. To perform these tasks, you can use the following inventory structs:

#### AnsibleInventoryCmd struct

The `AnsibleInventoryCmd` struct enables the generation of `ansible-inventory` commands. It implements the [Commander](#commander-interface) interface, so its method `Command` returns an array of strings that represents the command to be executed. An executor can use it to create the command to be executed.

The package provides the `NewAnsibleInventoryCmd` function to create a new instance of the `AnsibleInventoryCmd` struct. The function accepts a list of options to customize the `ansible-inventory` command. The following functions are available:

- `WithBinary(binary string) InventoryOptionsFunc`: Set the binary for the `ansible-inventory` command.
- `WithInventoryOptions(options *AnsibleInventoryOptions) InventoryOptionsFunc`: Set the inventory options for the command.
- `WithPattern(pattern string) InventoryOptionsFunc`: Set the pattern for the `ansible-inventory` command.

> Note
> Unlike other _Ansible_ commands, the `ansible-inventory` command does not provide privilege escalation or connection options, aligning with the functionality of the command itself.

#### AnsibleInventoryExecute struct

The `AnsibleInventoryExecute` struct serves as a streamlined [executor](#executor) for running `ansible-inventory` commands. It encapsulates the setup process for both the [command generator](#command-generator) and _executor_. This _executor_ is particularly useful when no additional configuration or customization is required.

The following methods are available to set attributes for the [AnsibleInventoryCmd](#ansibleinventorycmd-struct) struct:

- `WithBinary(binary string) *AnsibleInventoryExecute`: The method sets the `Binary` attribute.
- `WithInventoryOptions(options *AnsibleAdhocOptions) *AnsibleInventoryExecute`: The method sets the `InventoryOptions`  attribute.
- `WithPattern(pattern string) *AnsibleInventoryExecute`: The method sets the `Pattern` attribute.

Here is an example of launching an `ansible-inventory` command using `AnsibleInventoryExecute`:

```go
err := inventory.NewAnsibleInventoryExecute().
  WithInventoryOptions(&ansibleInventoryOptions).
  WithPattern("all").
  Execute(context.TODO())

if err != nil {
  // Manage the error
}
```

#### AnsibleInventoryOptions struct

The `AnsibleInventoryOptions` struct includes parameters described in the `Options` section of the _Ansible_ manual page. It defines the behavior of the Ansible inventory operations and specifies where to find the configuration settings.

### Playbook package

This section provides an overview of the `playbook` package in the _go-ansible_ library. Here are described its main components and functionalities.

The `github.com/apenella/go-ansible/v2/pkg/playbook` package facilitates the generation of _ansible-playbook_ commands. It does not execute the commands directly, but instead provides the necessary structs to generate the command to be executed by an executor. The `playbook` package includes the following essential structs for executing ad-hoc commands:

#### AnsiblePlaybookCmd struct

The `AnsiblePlaybookCmd` struct enables the generation of _ansible-playbook_ commands. It implements the [Commander](#commander-interface) interface, so its method `Command` returns an array of strings that represents the command to be executed. An executor can use it to create the command to be executed.

The package provides the `NewAnsiblePlaybookCmd` function to create a new instance of the `AnsiblePlaybookCmd` struct. The function accepts a list of options to customize the _ansible-playbook_ command. The following functions are available:

- `WithBinary(binary string) PlaybookOptionsFunc`: Set the binary for the _ansible-playbook_ command.
- `WithPlaybookOptions(options *AnsiblePlaybookOptions) PlaybookOptionsFunc`: Set the playbook options for the command.
- `WithPlaybooks(playbooks ...string) PlaybookOptionsFunc`: Set the playbooks for the _ansible-playbook_ command.

Next is an example of how to use the `AnsiblePlaybookCmd` struct to generate an _ansible-playbook_ command:

```go
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Connection: "local",
  Become:     true,
  Inventory:  "127.0.0.1,",
}

playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

// Generate the command to be executed
cmd, err := playbookCmd.Command()
if err != nil {
  // Manage the error
}
```

#### AnsiblePlaybookErrorEnrich struct

The `AnsiblePlaybookErrorEnrich` struct, that implements the [ErrorEnricher](#errorenricher-interface) interface, is responsible for enriching the error message when executing an _ansible-playbook_ command. Based on the exit code of the command execution, the `AnsiblePlaybookErrorEnrich` struct appends additional information to the error message. This additional information includes the exit code, the command that was executed, and the error message.

#### AnsiblePlaybookExecute struct

The `AnsiblePlaybookExecute` struct serves as a streamlined [executor](#executor) for running `ansible-playbook` command. It encapsulates the setup process for both the [command generator](#command-generator) and _executor_. Additionally, it provides the ability to enrich the error message when an error occurs during command execution.
This _executor_ is particularly useful when no additional configuration or customization is required.

The following methods are available to set attributes for the [AnsiblePlaybookCmd](#ansibleplaybookcmd-struct) struct:

- `WithBinary(binary string) *AnsiblePlaybookExecute`: The method sets the `Binary` attribute.
- `WithPlaybookOptions(options *AnsiblePlaybookOptions) *AnsiblePlaybookExecute`: The method sets the `PlaybookOptions` attribute.

Here is an example of launching an `ansible-playbook` command using `AnsiblePlaybookExecute`:

```go
err := playbook.NewAnsiblePlaybookExecute("site.yml", "site2.yml").
  WithPlaybookOptions(ansiblePlaybookOptions).
  Execute(context.TODO())

if err != nil {
  // Manage the error
}
```

#### AnsiblePlaybookOptions struct

With `AnsiblePlaybookOptions` struct, you can define parameters described in Ansible's manual page's `Options` section. It also allows you to define the connection options and privilege escalation options.

### Vault package

The `github.com/apenella/go-ansible/v2/pkg/vault` package provides functionality to encrypt variables. It introduces the `VariableVaulter` struct, which is responsible for creating a `VaultVariableValue` from the value that you need to encrypt.

The `VaultVariableValue` can return the instantiated variable in JSON format.

To perform the encryption, the `vault` package relies on an `Encrypter` interface implementation.

```go
type Encrypter interface {
  Encrypt(plainText string) (string, error)
}
```

The encryption functionality is implemented in the `encrypt` package, which is described in the following section.

#### Encrypt

The `github.com/apenella/go-ansible/v2/pkg/vault/encrypt` package is responsible for encrypting variables. It implements the `Encrypter` interface defined in the `github.com/apenella/go-ansible/v2/pkg/vault` package.

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

The `github.com/apenella/go-ansible/v2/pkg/vault/password/envvars` package allows you to read the password from an environment variable. To use this package, you need to use the `NewReadPasswordFromEnvVar` function and provide the name of the environment variable where the password is stored using the `WithEnvVar` option:

```go
reader := NewReadPasswordFromEnvVar(
  WithEnvVar("VAULT_PASSWORD"),
)
```

In this example, the `VAULT_PASSWORD` environment variable is specified as the source of the password. The `NewReadPasswordFromEnvVar` function creates a password reader that reads the password from the specified environment variable.

Using the `envvars` package, you can conveniently read the password from an environment variable and use it for encryption.

##### File

The `github.com/apenella/go-ansible/v2/pkg/vault/password/file` package allows you to read the password from a file, using the [afero](https://github.com/spf13/afero/blob/master/README.md) file system abstraction.

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

The `github.com/apenella/go-ansible/v2/pkg/vault/password/resolve` package provides a mechanism to resolve the password by exploring multiple `PasswordReader` implementations. It returns the first password obtained from any of the `PasswordReader` instances.

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

The `github.com/apenella/go-ansible/v2/pkg/vault/password/text` package provides functionality to read the password from a text source.

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
- [ansibleinventory-graph](https://github.com/apenella/go-ansible/tree/master/examples/ansibleinventory-graph)
- [ansibleinventory-simple](https://github.com/apenella/go-ansible/tree/master/examples/ansibleinventory-simple)
- [ansibleinventory-vaulted-vars](https://github.com/apenella/go-ansible/tree/master/examples/ansibleinventory-vaulted-vars)
- [ansibleplaybook-become](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-become)
- [ansibleplaybook-cobra-cmd](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-cobra-cmd)
- [ansibleplaybook-custom-transformer](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-custom-transformer)
- [ansibleplaybook-extravars-file](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-extravars-file)
- [ansibleplaybook-extravars](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-extravars)
- [ansibleplaybook-json-stdout](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-json-stdout)
- [ansibleplaybook-myexecutor](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-myexecutor)
- [ansibleplaybook-signals-and-cancellation](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-signals-and-cancellation)
- [ansibleplaybook-simple-embedfs](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-simple-embedfs)
- [ansibleplaybook-simple-with-prompt](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-simple-with-prompt)
- [ansibleplaybook-simple](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-simple)
- [ansibleplaybook-skipping-failing](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-skipping-failing)
- [ansibleplaybook-ssh](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-ssh)
- [ansibleplaybook-ssh-become-root-with-password/](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-ssh-become-root-with-password/)
- [ansibleplaybook-time-measurement](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-time-measurement)
- [ansibleplaybook-walk-through-json-output](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-walk-through-json-output)
- [ansibleplaybook-with-executor-time-measurament](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-with-executor-time-measurament)
- [ansibleplaybook-with-timeout](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-with-timeout)
- [ansibleplaybook-with-vaulted-extravar](https://github.com/apenella/go-ansible/tree/master/examples/ansibleplaybook-with-vaulted-extravar)
- [workflowexecute-simple](https://github.com/apenella/go-ansible/tree/master/examples/workflowexecute-simple)
- [workflowexecute-time-measurament](https://github.com/apenella/go-ansible/tree/master/examples/workflowexecute-time-measurament)

## Development Reference

This section provides a reference guide for developing and contributing to the _go-ansible_ library.

### Development Environment

To set up a development environment for the _go-ansible_ library, you need to have the following tools installed on your system:

- [Docker Compose](https://docs.docker.com/compose/). The version used for development is `docker-compose version v2.26.1`.
- [Docker](https://docs.docker.com/engine/reference/commandline/cli/). The version used for development is `Docker version 26.0.1`.
- [Go](https://golang.org/). The version used for development is `1.22`.
- [make](https://www.gnu.org/software/make/) utility. The version used for development is `GNU Make 4.3`. It is used to wrap the continuous integration and development processes, such us testing or linting.

#### Testing

To run the tests, you can use the following command:

```bash
make test
```

#### Static Analysis

To run the static analysis tools, you can use the following command:

```bash
make static-analysis
```

### Contributing

Thank you for your interest in contributing to go-ansible! All contributions are welcome, whether they are bug reports, feature requests, or code contributions. Please read the contributor's guide [here](https://github.com/apenella/go-ansible/blob/master/CONTRIBUTING.md) to learn more about how to contribute.

### Code Of Conduct

The _go-ansible_ project is committed to providing a friendly, safe and welcoming environment for all, regardless of gender, sexual orientation, disability, ethnicity, religion, or similar personal characteristics.

We expect all contributors, users, and community members to follow this code of conduct. This includes all interactions within the _go-ansible_ community, whether online, in person, or otherwise.

Please to know more about the code of conduct refer [here](https://github.com/apenella/go-ansible/blob/master/CODE-OF-CONDUCT.md).

## License

The _go-ansible_ library is available under [MIT](https://github.com/apenella/go-ansible/blob/master/LICENSE) license.
