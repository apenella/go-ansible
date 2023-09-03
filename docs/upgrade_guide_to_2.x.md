# Upgrade Guide to go-ansible 2.x

- [Upgrade Guide to go-ansible 2.x](#upgrade-guide-to-go-ansible-2x)
  - [Overview](#overview)
  - [Changes on the Executor interface](#changes-on-the-executor-interface)
  - [Changes on the DefaultExecute struct](#changes-on-the-defaultexecute-struct)
    - [Removed the ShowDuration attribute](#removed-the-showduration-attribute)
    - [New attribute _Output_ for printing execution results](#new-attribute-output-for-printing-execution-results)
    - [New attribute _Exec_ for running external commands](#new-attribute-exec-for-running-external-commands)
  - [Packages Reorganization](#packages-reorganization)
    - [github.com/apenella/go-ansible/pkg/stdoutcallback](#githubcomapenellago-ansiblepkgstdoutcallback)

## Overview

This document presents the changes introduced in go-ansible v2.0.0 and offers guidance for upgrading from go-ansible v1.x to v2.x.

## Changes on the Executor interface

The `Executor` interface has undergone a significant change in version v2.0.0 by removing the `resultsFunc` argument from the `Execute` method. This change constitutes a breaking change. In prior versions, the `Execute` method accepted an argument for `resultsFunc`, which represented the function responsible for handling and printing the execution's results. With the revised signature, the mechanism for printing the execution results is now anticipated to be an attribute of the executor.

If you defined a custom executor you need to adapt it by removing the `resultsFunc` argument, and defining the component that prints the execution's output in the struct. The section [Changes on the DefaultExecute struct](#changes-to-defaultexecute-struct) describes how the `DefaultExecutor` is provided with a new attribute that is responsible for printing the execution's output.

Here is the updated `Executor` interface:

```go
type Executor interface {
  Execute(ctx context.Context, command []string, options ...ExecuteOptions) error
}
```

## Changes on the DefaultExecute struct

The `DefaultExecute` struct, functioning as the default executor in the `go-ansible` library, has introduced significant changes in version 2.0.0. The next sections describe how to adapt that component to version v2.0.0:

### Removed the ShowDuration attribute

The `DefaultExacute` has removed the attribute `ShowDuration`, as previously announced. Starting from version v2.0.0, to measure the duration of the execution you should utilize the `measure.ExecutiorTimeMeasurement` component.

For guidance on how to use the ExecutiorTimeMeasurement decorator component, please refer to the example [ansibleplaybook-time-measurement](https://github.com/apenella/go-ansible/blob/master/examples/ansibleplaybook-time-measurement/ansibleplaybook-time-measurement.go).

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

For further details on the DefaultResults struct, please refer to the section that discusses the package reorganization at [github.com/apenella/go-ansible/pkg/stdoutcallback](#githubcomapenellago-ansiblepkgstdoutcallback).

### New attribute _Exec_ for running external commands

In versions prior to v2.0.0, `DefaultResults` utilized the `os.exec` package to run external commands. However, in version v2.0.0, direct usage of `os.exec` has been replaced by the introduction of a new attribute called `Exec`. This attribute instantiates a component of type `Executabler`, responsible for executing external commands.

By default, if you do not explicitly define an `Exec` attribute, `DefaultExecute` uses the `Exec` struct to execute external commands, which is defined in the `github.com/apenella/go-ansible/pkg/execute/executable/os/exec package`.


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
