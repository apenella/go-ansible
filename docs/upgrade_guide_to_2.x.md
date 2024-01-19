# Upgrade Guide to go-ansible 2.x

- [Upgrade Guide to go-ansible 2.x](#upgrade-guide-to-go-ansible-2x)
  - [Overview](#overview)
  - [Changes on the _Executor_ interface](#changes-on-the-executor-interface)
    - [Replacing the _command_ argument](#replacing-the-command-argument)
    - [Replacing the _resultsFunc_ argument](#replacing-the-resultsfunc-argument)
    - [Replacing the _options_ argument](#replacing-the-options-argument)
  - [Changes on the _DefaultExecute_ struct](#changes-on-the-defaultexecute-struct)
    - [Adding _Cmd_ attribute to generate commands](#adding-cmd-attribute-to-generate-commands)
    - [Adding _Exec_ attribute for running external commands](#adding-exec-attribute-for-running-external-commands)
    - [Adding _Output_ attribute for printing execution results](#adding-output-attribute-for-printing-execution-results)
    - [Removing the _ShowDuration_ attribute](#removing-the-showduration-attribute)
    - [Adpating your code to new the _Transformer_ location](#adpating-your-code-to-new-the-transformer-location)
  - [Changes on the _AnsiblePlaybookCmd_ struct](#changes-on-the-ansibleplaybookcmd-struct)
    - [Removing the _Exec_ attribute and _Run_ method](#removing-the-exec-attribute-and-run-method)
    - [Removing the _StdoutCallback_ attribute](#removing-the-stdoutcallback-attribute)
  - [Changes on the _AnsibleAdhocCmd_ struct](#changes-on-the-ansibleadhoccmd-struct)
  - [Packages Reorganization](#packages-reorganization)
    - [github.com/apenella/go-ansible/pkg/stdoutcallback](#githubcomapenellago-ansiblepkgstdoutcallback)

## Overview

This document offers guidance for upgrading from _go-ansible_ _v1.x_ to _v2.x_. It also presents the changes introduced in _go-ansible v2.0.0_ since the major version _1.x_. Some of those changes are breaking changes.

The most relevant change is that command structs no longer execute commands. So, `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` do not require an `Executor` anymore. Instead, the `Executor` is responsible for the command execution. To achieve that, the `Executor` depends on the command structs to generate the commands to execute.

Go through the following sections to learn about the changes introduced in version _2.x_ and how to adapt your code to those changes.

## Changes on the _Executor_ interface

> Changes on the _Executor_ interface are breaking changes. These changes affects several packages or structs in _go-ansible_. In this section, you will find guidance on how to adapt your custom implentation of an executor.
> To know how these changes impact other packages or structs, refer to the corresponding sections in this document.

The `Executor` interface has undergone significant breaking changes. It removes the `command`, `resultsFunc`, and `options` arguments from the `Execute` method.

Here is the updated `Executor` interface:

```go
type Executor interface {
  Execute(ctx context.Context) error
}
```

To align with the updated `Executor` interface, you need to adapt your custom executor by removing the `command`, `resultsFunc`, and `options` arguments from its `Execute` method. The following points detail how to replace each of these arguments.

### Replacing the _command_ argument

Instead of using the _command_ argument, the `Executor` should expect a `Commander` to generates the command to execute. Therefore, your executor should include an attribute of type `Commander`. The `Commander` is interface defined in _github.com/apenella/go-ansible/pkg/execute_ as follows:

```go
// Commander generates commands to be executed
type Commander interface {
  Command() ([]string, error)
}
```

The `Command` method returns an array of strings representing the command to execute. You should provide the component responsible for executing external commands with this array.
Both the `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs implement the `Commander` interface.

You can review changes on `DefaultExecute` [here](#adding-cmd-attribute-to-generate-commands) and see how it has been adapted to use the `Commander` to generate the command to execute.

### Replacing the _resultsFunc_ argument

The _resultsFunc_ previously managed the results output from command execution. With its removal, your executor should expect a new component to assume this responsibility. That component for handling the results output should be of type `ResultsOutputer`.
A `ResultsOutputer` is an interface defined in _github.com/apenella/go-ansible/pkg/execute/result_ as follows:

```go
// OptionsFunc is a function that can be used to configure a ResultsOutputer struct
type OptionsFunc func(ResultsOutputer)

// ResultsOutputer is the interface that must implements an struct to print the execution results
type ResultsOutputer interface {
 Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...OptionsFunc) error
}
```

The _go-ansible_ library provides two implementations of the `ResultsOutputer` interface:

- **DefaultResults Struct**
Located in the package _github.com/apenella/go-ansible/pkg/execute/result/default_, the `DefaultResults` struct handles Ansible's results in plain text.

- **JSONResults Struct**
Defined in the package _github.com/apenella/go-ansible/pkg/execute/json_, the `JSONResults` struct manages Ansible's results in JSON format.

Choose between these mechanisms based on the stdout callback plugin you use.

To sum it up, to replace the _resultsFunc_, introduce an attribute of type `ResultsOutputer` in your executor, and utilize this attribute to print the results output from command execution.

[Here](#adding-output-attribute-for-printing-execution-results) you can find how the `DefaultExecute` struct has been adapted to use a `ResultsOutputer` to print the execution results.

### Replacing the _options_ argument

By removing the options argument, the ability to overwrite the `Executor` struct attributes in the `Execute` method is no longer available. To configure your executor, you must set it up during the instantiation of the struct.

## Changes on the _DefaultExecute_ struct

The `DefaultExecute` struct is a ready-to-go component provided by the _go-ansible_ library for executing external commands. You can find its definition in the _github.com/apenella/go-ansible/pkg/execute_ package.
Changes on the `Executor` interface impacts the `DefaultExecute` struct. You can read more about the changes on the `Executor` interface [here](#changes-on-the-executor-interface).

In version _v2.x_ you need to instantiate the `DefaultExecute` struct to execute the Ansible commands, as is shown in the following code snippet. 

```go
// playbookCmd is the Commander responsible for generating the command to execute
exec := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
)

// Execute the Ansible command
err := exec.Execute(context.TODO())
if err != nil {
  panic(err)
}
```

If you already configured the `DefaultExecute` struct in your code, you should adapt it to the new version. Follow the coming sections to learn how to adapt your code to these changes.

### Adding _Cmd_ attribute to generate commands

The `DefaultExecute` requires a `Commander` to generate the external ommand to execute. For that reason, it includes the `Cmd` attribute of type `Commander`. Both the `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs implement the `Commander` interface.

When you instantiate the `DefaultExecute` struct, you should provide the `Cmd` attribute with a `Commander` to generate the commands. The following code shows how to instantiate the `DefaultExecute` struct using a `AnsiblePlaybookCmd` as the `Commander`.

```go
// Define the AnsiblePlaybookCmd and the required options.
playbookCmd := &playbook.AnsiblePlaybookCmd{
  Playbooks:         []string{"site.yml", "site2.yml"},
  ConnectionOptions: &options.AnsibleConnectionOptions{
    Connection: "local",
  },
  Options:           &playbook.AnsiblePlaybookOptions{
    Inventory: "all,",
  },
}
// Instanciate a DefaultExecutoe by providing 'playbookCmd' as the Commander.
exec := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
)
```

In the example above, the `playbookCmd` is of type `Commander`. You set the `Cmd` value as `playbookCmd` through the function `WithCmd` when you instantiate a new `DefaultExecute`. So, the `DefaultExecute` utilizes the `playbookCmd` to generate the command to execute.

### Adding _Exec_ attribute for running external commands

In the latest _go-ansible_ version, the `DefaultExecute` struct includes the `Exec` attribute of type `Executabler`. The `Exec` component is responsible for executing external commands.
By default, if you do not define the `Exec` attribute, it uses the `OsExec` struct. The `OsExec` implementation is found in the _github.com/apenella/go-ansible/internal/execute/executable/os/exec_ package. This struct wraps the `os/exec` package.

If you need to implement a custom executabler, you should implement the `Executabler` interface. The interface is defined in _github.com/apenella/go-ansible/pkg/execute_ as follows:

```go
// Executabler is an interface to run commands
type Executabler interface {
  Command(name string, arg ...string) exec.Cmder
  CommandContext(ctx context.Context, name string, arg ...string) exec.Cmder
}
```

Below, you can find an example of how to instantiate a `DefaultExecute` struct with a custom executabler.

```go
// Define the AnsiblePlaybookCmd and the required options.
playbookCmd := &playbook.AnsiblePlaybookCmd{
  Playbooks:         []string{"site.yml", "site2.yml"},
  ConnectionOptions: &options.AnsibleConnectionOptions{
    Connection: "local",
  },
  Options:           &playbook.AnsiblePlaybookOptions{
    Inventory: "all,",
  },
}

// Define a custom Executabler
executable := &myCustomExecutabler{}

// Instanciate a DefaultExecutoe by providing 'playbookCmd' and 'executabler' as the Commander and Executabler respectively.
executor := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
  execute.WithExecutable(executable),
)
```

In the example above, _executable_ implements the `Executabler` interface. When you instantiate a new `DefaultExecute`, you set the `Exec` value to _executable_ using the function `WithExecutable`. So, the `DefaultExecute` will use the _executable_ to execute the command.

### Adding _Output_ attribute for printing execution results

To align with the new `Executor` interface, the `DefaultExecute` struct includes the `Output` attribute of type `ResultsOutputer`. It manages the output of Ansible commands. You can find the definition for `ResultsOutputer` in _github.com/apenella/go-ansible/pkg/execute/result_ as follows:

```go
// OptionsFunc is a function that can be used to configure a ResultsOutputer struct
type OptionsFunc func(ResultsOutputer)

// ResultsOutputer is the interface that must implements an struct to print the execution results
type ResultsOutputer interface {
 Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...OptionsFunc) error
}
```

When you do not specify the `Output` attribute, it uses as a fallback mechanism the `DefaultResults` struct for output. You can find this struct in the _github.com/apenella/go-ansible/pkg/execute/result/default_ package.

You can use the `WithOutput` function defined in the _github.com/apenella/go-ansible/pkg/execute_ package to configure the `Output` attribute during the instantiation of the `DefaultExecute` struct.

Below, you can find an example of how to instantiate a `DefaultExecute` struct with a custom output mechanism.

```go
// Define the AnsiblePlaybookCmd and the required options.
playbookCmd := &playbook.AnsiblePlaybookCmd{
  Playbooks:         []string{"site.yml", "site2.yml"},
  ConnectionOptions: &options.AnsibleConnectionOptions{
    Connection: "local",
  },
  Options:           &playbook.AnsiblePlaybookOptions{
    Inventory: "all,",
  },
}
// Define a custom ResultsOutputer
output := &myCustomResultsOutputer{}

// Instanciate a DefaultExecutoe by providing 'playbookCmd' and 'outputer' as the Commander and ResultsOutputer respectively.
executor := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
  execute.WithOutput(output),
)
```

The example above shows how to instantiate a `DefaultExecute` struct with a custom `ResultsOutputer`. The _output_ is of type `ResultsOutputer`. When you instantiate a new `DefaultExecute`, you set the `Output` attribute with the _output_. So, the `DefaultExecute` uses _output_ to print the execution results.

### Removing the _ShowDuration_ attribute

As announced in prior go-ansible versions, the `DefaultExecute` has removed the ShowDuration attribute.

Starting from version _v2.0.0_, to measure the duration of the execution, you should use the `ExecutorTimeMeasurement` struct. This struct acts as a decorator over the `Executor` and is available in the _github.com/apenella/go-ansible/pkg/execute/measure_ package.

For guidance on how to use the ExecutorTimeMeasurement, please refer to the [ansibleplaybook-time-measurement](https://github.com/apenella/go-ansible/blob/master/examples/ansibleplaybook-time-measurement/ansibleplaybook-time-measurement.go) example. However, the following code snippet shows how to use the `ExecutorTimeMeasurement` struct.

```go
exec := measure.NewExecutorTimeMeasurement(
    execute.NewDefaultExecute(
      execute.WithCmd(playbook),
    ),
)

err := exec.Execute(context.TODO())
if err != nil {
  fmt.Println(err.Error())
}

fmt.Println("Duration: ", exec.Duration().String())
```

### Adpating your code to new the _Transformer_ location

You can configure a set of transformers to modify the output of the execution's results. The _go-ansible_ library has moved the `transformer` package from _github.com/apenella/go-ansible/pkg/stdoutcallback/results_ to _github.com/apenella/go-ansible/pkg/execute/result/transformer_. So, you should adapt your code to this change.

## Changes on the _AnsiblePlaybookCmd_ struct

The `AnsiblePlaybookCmd` struct has undergone significant changes. It changed its responsibilities and no longer executes commands. Instead, it implements the `Commander` interface, which generates commands for execution. So, you need to adapt your code to these changes. This section outlines the necessary steps to migrate from the older version to the new one.

### Removing the _Exec_ attribute and _Run_ method

The `AnsiblePlaygookCmd` struct is not responsible for executing commands anymore. For that reason, the `Exec` attribute has been removed.

Along with the `Exec` attribute, the `Run` method is not available anymore. To execute a command, you should use an `Executor`. Then, the `Executor` should receive an `AnsiblePlaybookCmd` struct to generate the command to execute.

```go
// Define the AnsiblePlaybookCmd and the required options.
playbookCmd := &playbook.AnsiblePlaybookCmd{
  Playbooks:         []string{"site.yml", "site2.yml"},
  ConnectionOptions: &options.AnsibleConnectionOptions{
    Connection: "local",
  },
  Options:           &playbook.AnsiblePlaybookOptions{
    Inventory: "127.0.0.1,",
  },
}

// Instanciate a DefaultExecutoe by providing 'playbookCmd' as the Commander
exec := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
)

// Execute the external command through the executor
err := exec.Execute(context.TODO())
if err != nil {
  panic(err)
}
```

### Removing the _StdoutCallback_ attribute

## Changes on the _AnsibleAdhocCmd_ struct

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
