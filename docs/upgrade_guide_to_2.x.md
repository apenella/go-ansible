# Upgrade Guide to go-ansible 2.x

- [Upgrade Guide to go-ansible 2.x](#upgrade-guide-to-go-ansible-2x)
  - [Overview](#overview)
  - [Changes on the _Executor_ interface](#changes-on-the-executor-interface)
    - [How to adapt your code to the new _Executor_ interface](#how-to-adapt-your-code-to-the-new-executor-interface)
      - [How to replace the _command_ argument](#how-to-replace-the-command-argument)
      - [How to replace the _resultsFunc_ argument](#how-to-replace-the-resultsfunc-argument)
      - [How to replace the _options_ argument](#how-to-replace-the-options-argument)
  - [Changes on the _DefaultExecute_ struct](#changes-on-the-defaultexecute-struct)
    - [Removed the _ShowDuration_ attribute](#removed-the-showduration-attribute)
    - [New attribute _Output_ for printing execution results](#new-attribute-output-for-printing-execution-results)
    - [New attribute _Exec_ for running external commands](#new-attribute-exec-for-running-external-commands)
  - [Packages Reorganization](#packages-reorganization)
    - [github.com/apenella/go-ansible/pkg/stdoutcallback](#githubcomapenellago-ansiblepkgstdoutcallback)

## Overview

This document offers guidance for upgrading from go-ansible _v1.x_ to _v2.x_. It also presents the changes introduced in _go-ansible v2.0.0_ since the major version _1.x_. Some those changes are breaking changes.

The most relevant change is that command structs no longer execute commands. So, `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` do not requiere an `Executor` anymore. Instead, the `Executor` is responsible of the command execution. To achieve that, the `Executor` depends on the command structs to generates the commands to execute.

Go through the following sections to learn about the changes introduced in version 2.x and how to adapt your code to those changes.

## Changes on the _Executor_ interface

The `Executor` interface has undergone significant breaking changes. It removes the `command`, `resultsFunc`, and `options` arguments from the `Execute` method.
That changes only affects your code if you have defined a custom _Executor_ struct.

Here is the updated `Executor` interface:

```go
type Executor interface {
  Execute(ctx context.Context) error
}
```

### How to adapt your code to the new _Executor_ interface

To align with the updated `Executor` interface, you need to adapt your custom _Executor_ by removing the `command`, `resultsFunc`, and `options` arguments from its `Execute` method. The following points detail how to replace each of these arguments.

#### How to replace the _command_ argument

Instead of using the command argument, the new `Executor` utilizes a `Commander` to generate commands for execution. Therefore, the `Executor` should include an attribute of type `Commander`.

Both the `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs already implement the `Commander` interface.

The `Commander` interface definition is as follows:

```go
// Commander generates commands to be executed
type Commander interface {
  Command() ([]string, error)
}
```

The `Command` method returns an array of strings representing the command to execute. You should use this array for the component responsible for executing external commands.

The `DefaultExecute` in _go-ansible_ utilizes an `Executabler` to execute external commands. The `Executabler` is an interface defined as:

```go
// Executabler is an interface to run commands
type Executabler interface {
  Command(name string, arg ...string) exec.Cmder
  CommandContext(ctx context.Context, name string, arg ...string) exec.Cmder
}
```

The `os/exec`'s `Cmd` struct implements the `Executabler` interface.

The following code showcases how the `DefaultExecute` uses the `Commander` and the `Executabler`.

```go
// e.Cmd is a Commander
command, err := e.Cmd.Command()
if err != nil {
  return errors.New("(DefaultExecute::Execute)", "Error creating command", err)
}

// e.Exec is an Executabler
cmd := e.Exec.CommandContext(ctx, command[0], command[1:]...)
```

By incorporating these changes, your code will align with the updated `Executor` interface in _go-ansible v2.x_.

#### How to replace the _resultsFunc_ argument

The _resultsFunc_ previously managed the results output from command execution. With its removal, a new component within the Executor must assume this responsibility.

To handle the results output, _go-ansible_ provides two mechanisms:

- **DefaultResults Struct**
Located in the package _github.com/apenella/go-ansible/pkg/execute/result/default_, the `DefaultResults` struct handles Ansible's results in plain text.

- **JSONResults Struct**
Defined in the package github.com/apenella/go-ansible/pkg/execute/json, the JSONResults struct manages Ansible's results in JSON format.

Choose between these mechanisms based on the stdout callback plugin you use.

Both components implement the `ResultsOutputer` interface, defined in _github.com/apenella/go-ansible/pkg/execute/result_ as follows:

```go
// OptionsFunc is a function that can be used to configure a ResultsOutputer struct
type OptionsFunc func(ResultsOutputer)

// ResultsOutputer is the interface that must implements an struct to print the execution results
type ResultsOutputer interface {
 Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...OptionsFunc) error
}
```

To replace the _resultsFunc_, introduce an attribute of type `ResultsOutputer` in your `Executor` struct. Utilize this attribute to print the results output from command execution.

The `DefaultExecute` includes a default `ResultsOutputer` attribute named `Output`, which uses the `DefaultResults` struct to print execution results.

#### How to replace the _options_ argument

By removing the _options_ argument, you will not be able to overwrite the _Executor_ configuration when you executing the command. You must set up the Executor when instanciating the struct.

## Changes on the _DefaultExecute_ struct

The `DefaultExecute` struct, functioning as the default executor in the `go-ansible` library, has introduced significant changes in version 2.0.0. The next sections describe how to adapt your code to that component to version v2.0.0:

### Removed the _ShowDuration_ attribute

The `DefaultExacute` has removed the attribute `ShowDuration`, as previously announced. Starting from version v2.0.0, to measure the duration of the execution you should utilize the `measure.ExecutiorTimeMeasurement` component.

For guidance on how to use the `ExecutiorTimeMeasurement` decorator component, please refer to the example [ansibleplaybook-time-measurement](https://github.com/apenella/go-ansible/blob/master/examples/ansibleplaybook-time-measurement/ansibleplaybook-time-measurement.go) to know how to use it.

### New attribute _Output_ for printing execution results

Inside the `DefaultExecute` struct, the `Execute` method has undergone an update that removes the `resultsFunc` attribute. This attribute was previously of type `stdoutcallback.StdoutCallbackResultsFunc` and was responsible for printing the execution's output. Instead, a new attribute named `Output` has been introduced to the struct.

The `Output` attribute is of type `results.ResultsOutputer`, which represents an interface defined as follows:

```go
// OptionsFunc is a function that can be used to configure a ResultsOutputer struct
type OptionsFunc func(ResultsOutputer)

// ResultsOutputer is the interface that must implements an struct to print the execution results
type ResultsOutputer interface {
  Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...OptionsFunc) error
}
```

If the `Output` attribute is not explicitly specified within the `DefaultExecute` struct, the default behaviour uses the `DefaultResults` struct as the output mechanism. You can find this struct in the `github.com/apenella/go-ansible/pkg/execute/result/default` package, and it is responsible for handling the printing of execution output.

For further details on the `DefaultResults` struct, please refer to the section that discusses the package reorganization at [github.com/apenella/go-ansible/pkg/stdoutcallback](#githubcomapenellago-ansiblepkgstdoutcallback).

### New attribute _Exec_ for running external commands

In versions prior to v2.0.0, `DefaultResults` utilized the `os.exec` package to run external commands. However, in version v2.0.0, direct usage of `os.exec` has been replaced by defining a new attribute called `Exec`. This attribute instantiates a component of type `Executabler`, responsible for executing external commands.

By default, if you do not explicitly define an `Exec` attribute, `Exec` uses the `DefaultExecute` struct to execute external commands, which is defined in the `github.com/apenella/go-ansible/pkg/execute/executable/os/exec` package.

The `Executabler` interface is defined as follows:

```go
// Executabler is an interface to run commands
type Executabler interface {
  Command(name string, arg ...string) exec.Cmder
  CommandContext(ctx context.Context, name string, arg ...string) exec.Cmder
}
```

The `exec.Cmder` type used on the `Executabler` interface is defined as:

```go
// Cmder is an interface to run a command
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



if you have configured the executor

```go
execute := execute.NewDefaultExecute(
  execute.WithWrite(io.Writer(buff)),
)
```

```go
exec := measure.NewExecutorTimeMeasurement(
  stdoutcallback.NewJSONStdoutCallbackExecute(
    execute.NewDefaultExecute(
      execute.WithCmd(playbook),
      execute.WithWrite(io.Writer(buff)),
    ),
  ),
)
```

## Packages Reorganization

Version v2.0.0 introduces changes to the package structure, including some reorganization and removals. This section outlines the necessary steps to migrate from the older packages to the new ones.

### github.com/apenella/go-ansible/pkg/stdoutcallback

The `github.com/apenella/go-ansible/pkg/stdoutcallback/results` package suffered several changes in version v2.0.0. This section explains how to adapt your code to these changes.

Previously, this package contained various structs and functions, which have now been split and moved into new packages based on their responsibilities:

- The functions for transforming the output lines of the execution's results are now available in the `github.com/apenella/go-ansible/pkg/execute/result/transformer` package. To utilize these functions, import the `github.com/apenella/go-ansible/pkg/execute/result/transformer` package and update the corresponding functions (Prepend, Append, LogFormat, and IgnoreMessage) in your code to use the `transformer` package.

```go
// import "github.com/apenella/go-ansible/pkg/execute/result/transformer"
transformer.Prepend("Go-ansible example")
```

- The `github.com/apenella/go-ansible/pkg/execute/result/default` package introduces the `DefaultResults` struct, which takes over the functionality previously provided by the `DefaultStdoutCallbackResults` function, defined in the `github.com/apenella/go-ansible/pkg/stdoutcallback/results` package. As the `DefaultStdoutCallbackResults` function is no longer available, you should use the `DefaultResults` struct as the default mechanism for printing the output of the execution's results. Furthermore, the `DefaultExecutor` now employs the `DefaultResults` struct as the default component for printing the execution results.
