# Upgrade Guide to go-ansible 2.x

## Overview

This document offers guidance for upgrading from _go-ansible_ _v1.x_ to _v2.x_. It also presents the changes introduced in _go-ansible v2.0.0_ since the major version _1.x_. Some of these are breaking changes.

The most relevant change is that the package name has been changed from `github.com/apenella/go-ansible` to `github.com/apenella/go-ansible/v2`. So, you need to update your import paths to use the new module name.

Another important change to highlight is that command structs no longer execute commands. So, `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` do not require an `Executor` anymore. Instead, the `Executor` is responsible for the command execution. To achieve that, the `Executor` depends on the command structs to generate the commands to execute.
That change is motivated by the need to segregate the command generation from the command execution. Having the `Executor` as the central component of the command execution process allows the `Executor` to be more flexible and customizable. The _go-ansible_ library provides a set of decorator structs to configure the `Executor` with different features, such as stdout callback management, and Ansible configuration settings.

Proceed through the following sections to understand the changes in version _2.x_ and learn how to adapt your code accordingly.

- [Upgrade Guide to go-ansible 2.x](#upgrade-guide-to-go-ansible-2x)
  - [Overview](#overview)
  - [Changes on the interfaces](#changes-on-the-interfaces)
    - [Added _Cmder_ interface](#added-cmder-interface)
    - [Added _Commander_ interface](#added-commander-interface)
    - [Added _ErrorEnricher_ interface](#added-errorenricher-interface)
    - [Added _Executabler_ interface](#added-executabler-interface)
    - [Added _ExecutorEnvVarSetter_ interface](#added-executorenvvarsetter-interface)
    - [Added _ExecutorQuietStdoutCallbackSetter_ interface](#added-executorquietstdoutcallbacksetter-interface)
    - [Added _ExecutorStdoutCallbackSetter_ interface](#added-executorstdoutcallbacksetter-interface)
    - [Added _ResultsOutputer_ interface](#added-resultsoutputer-interface)
    - [Updated _Executor_ interface](#updated-executor-interface)
  - [Changes on the _Executor_ interface](#changes-on-the-executor-interface)
    - [Replacing the _command_ argument](#replacing-the-command-argument)
    - [Replacing the _resultsFunc_ argument](#replacing-the-resultsfunc-argument)
      - [DefaultResults struct](#defaultresults-struct)
      - [JSONStdoutCallbackResults struct](#jsonstdoutcallbackresults-struct)
    - [Replacing the _options_ argument](#replacing-the-options-argument)
  - [Changes on the _DefaultExecute_ struct](#changes-on-the-defaultexecute-struct)
    - [Adding _Cmd_ attribute to generate commands](#adding-cmd-attribute-to-generate-commands)
    - [Adding _ErrorEnrich_ attribute to enrich error messages](#adding-errorenrich-attribute-to-enrich-error-messages)
    - [Adding _Exec_ attribute for running external commands](#adding-exec-attribute-for-running-external-commands)
    - [Adding _Output_ attribute for printing execution results](#adding-output-attribute-for-printing-execution-results)
    - [Removing the _ShowDuration_ attribute](#removing-the-showduration-attribute)
    - [Removing the error enrichment for ansible-playbook commands](#removing-the-error-enrichment-for-ansible-playbook-commands)
    - [Changing the _Transformer_ location](#changing-the-transformer-location)
  - [Changes on the _AnsiblePlaybookCmd_ struct](#changes-on-the-ansibleplaybookcmd-struct)
    - [Renaming the _Options_ attribute](#renaming-the-options-attribute)
    - [Removing the _Exec_ attribute and _Run_ method](#removing-the-exec-attribute-and-run-method)
    - [Removing the _StdoutCallback_ attribute](#removing-the-stdoutcallback-attribute)
    - [Using the _AnsiblePlaybookExecute_ struct as executor](#using-the-ansibleplaybookexecute-struct-as-executor)
  - [Changes on the _AnsibleAdhocCmd_ struct](#changes-on-the-ansibleadhoccmd-struct)
  - [Changes on the _AnsibleInventoryCmd_ struct](#changes-on-the-ansibleinventorycmd-struct)
  - [Changes on the _Transformer_ functions](#changes-on-the-transformer-functions)
  - [Managing Ansible Stdout Callback](#managing-ansible-stdout-callback)
  - [Managing Ansible configuration settings](#managing-ansible-configuration-settings)
    - [Removing configuration functions](#removing-configuration-functions)
      - [Replacing the _AnsibleForceColor_ function](#replacing-the-ansibleforcecolor-function)
      - [Replacing the _AnsibleAvoidHostKeyChecking_ function](#replacing-the-ansibleavoidhostkeychecking-function)
      - [Replacing the _AnsibleSetEnv_ function](#replacing-the-ansiblesetenv-function)

## Changes on the interfaces

The version _v2.x_ introduces several changes in the interfaces used by the _go-ansible_ library. Throughout this document, you will find references to these interfaces and this section presents the new interfaces and where to find them.

### Added _Cmder_ interface

The `Cmder` interface is defined in _github.com/apenella/go-ansible/v2/pkg/execute/exec_ and it is used to run external commands. The `os/exec` package implements the `Cmder` interface. The [Executabler](#added-executabler-interface)'s `Command` and `CommandContext` methods return a `Cmder` interface.
You can find the definition of the `Cmder` interface below:

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

### Added _Commander_ interface

The `Commander` interface defined in _github.com/apenella/go-ansible/v2/pkg/execute_ is used to generate the commands to be executed. It is required by `DefaultExecute` struct, but you can also use it to implement your custom executor.

The `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs implement the `Commander` interface. You can find the definition of the `Commander` interface below:

```go
// Commander generates commands to be executed
type Commander interface {
  Command() ([]string, error)
  String() string
}
```

### Added _ErrorEnricher_ interface

The `ErrorEnricher` interface defined in _github.com/apenella/go-ansible/v2/pkg/execute_ is used to enrich the error message. The `DefaultExecute` struct uses that enable you to append additional information to the error message when an error occurs.

```go
// ErrorEnricher interface to enrich and customize errors
type ErrorEnricher interface {
  Enrich(err error) error
}
```

### Added _Executabler_ interface

The `Executabler` interface defined in _github.com/apenella/go-ansible/v2/pkg/execute_ is used to run external commands. It is required by `DefaultExecute` struct, but you can also use it to implement your custom executor.

The `OsExec` struct implements the `Executabler` interface. You can find the definition of the `Executabler` interface below:

```go
// Executabler is an interface to run commands
type Executabler interface {
  Command(name string, arg ...string) exec.Cmder
  CommandContext(ctx context.Context, name string, arg ...string) exec.Cmder
}
```

### Added _ExecutorEnvVarSetter_ interface

The `ExecutorEnvVarSetter` interface defined in _github.com/apenella/go-ansible/v2/pkg/execute/configuration_ defines an executor interface to which you can set environment variables. It is required by `AnsibleWithConfigurationSettingsExecute` decorator struct.

```go
// ExecutorEnvVarSetter extends the executor interface by adding methods to configure environment variables
type ExecutorEnvVarSetter interface {
  // executor interface defined in github.com/apenella/go-ansible/v2/pkg/execute
  execute.Executor
  // AddEnvVar adds an environment variable to the executor
  AddEnvVar(key, value string)
}
```

### Added _ExecutorQuietStdoutCallbackSetter_ interface

The `ExecutorQuietStdoutCallbackSetter` interface defined in _github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback_ extends the [ExecutorStdoutCallbackSetter](#added-executorstdoutcallbacksetter-interface) interface by adding the `Quiet` method to remove the verbosity of the command execution.

```go
// ExecutorQuietStdoutCallbackSetter extends the ExecutorStdoutCallbackSetter interface by adding a method to force the non-verbose mode in the Stdout Callback configuration
type ExecutorQuietStdoutCallbackSetter interface {
  // ExecutorStdoutCallbackSetter interface defined in github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback
  ExecutorStdoutCallbackSetter
  // Quiet removes the verbosity of the command execution
  Quiet()
}
```

### Added _ExecutorStdoutCallbackSetter_ interface

The `ExecutorStdoutCallbackSetter` interface defined in _github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback_ is used to set the stdout callback method to the executor. It is required by the stdout callback decorator structs defined in the same package.

```go
// ExecutorStdoutCallbackSetter extends the executor interface by adding methods to configure the Stdout Callback configuration
type ExecutorStdoutCallbackSetter interface {
  // executor interface defined in github.com/apenella/go-ansible/v2/pkg/execute
  execute.Executor
  // AddEnvVar adds an environment variable to the executor
  AddEnvVar(key, value string)
  // WithOutput sets the output mechanism to print the execution results to the executor
  WithOutput(output result.ResultsOutputer)
}
```

### Added _ResultsOutputer_ interface

The `ResultsOutputer` interface defined in _github.com/apenella/go-ansible/v2/pkg/execute/result_ is used to print the execution results. It is required by `DefaultExecute`, but you can also use it to implement your custom executor. The `DefaultResults` and `JSONResults` structs implement the `ResultsOutputer` interface.

```go
// OptionsFunc is a function to configure a ResultsOutputer struct
type OptionsFunc func(ResultsOutputer)

// ResultsOutputer is the interface that must implements an struct to print the execution results
type ResultsOutputer interface {
 Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...OptionsFunc) error
}
```

### Updated _Executor_ interface

Read the section [Changes on the _Executor_ interface](#changes-on-the-executor-interface) to learn about the changes on the `Executor` interface.

## Changes on the _Executor_ interface

> **Note**
> The modifications to the `Executor` interface in _go-ansible_ involve breaking changes that impact various packages and structs. This section provides guidance on adapting your custom executor implementation.
> Refer to corresponding sections for insights into how these changes affect other components.

The `Executor` interface has undergone significant breaking changes. It removes the `command`, `resultsFunc`, and `options` arguments from the `Execute` method.

 The revised interface is now:

```go
type Executor interface {
  Execute(ctx context.Context) error
}
```

To align with these changes, adjust your custom executor by removing the `command`, `resultsFunc`, and `options` arguments from its `Execute` method. The following points outline how to replace each of these arguments.

### Replacing the _command_ argument

Instead of utilizing the _command_ argument, the `Executor` now relies on a `Commander` to generate the command for execution. Consequently, your executor should have an attribute of type `Commander`. For more details about the `Commander` interface, refer [here](#added-commander-interface).

The `Command` method, part of the `Commander` interface, returns an array of strings that represents the command to execute. This array should be handed over to the component responsible for executing external commands, an `Executabler`.

Both the `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs implement the `Commander` interface. For insights into how `DefaultExecute` has been adapted to use the `Commander` for generating the command, review the changes [here](#adding-cmd-attribute-to-generate-commands).

### Replacing the _resultsFunc_ argument

Previously, the _resultsFunc_ managed the results output from command execution. With its removal, your executor now requires a new component to handle this responsibility. This component should be of type `ResultsOutputer`. The definition of the `ResultsOutputer` interface is available [here](#added-resultsoutputer-interface).

The _go-ansible_ library provides two implementations of the `ResultsOutputer` interface:

#### DefaultResults struct

Found in the package _github.com/apenella/go-ansible/v2/pkg/execute/result/default_, the `DefaultResults` struct handles Ansible's results in plain text.

#### JSONStdoutCallbackResults struct

Defined in the package _github.com/apenella/go-ansible/v2/pkg/execute/json_, the `JSONStdoutCallbackResults` struct manages Ansible's results in JSON format.

Select the appropriate mechanism based on the stdout callback plugin you are using.

To replace the _resultsFunc_, introduce an attribute of type `ResultsOutputer` in your executor. Utilize this attribute to print the results output from command execution. For an example of how the `DefaultExecute` struct has been adapted to use a `ResultsOutputer` for printing execution results, refer [Here](#adding-output-attribute-for-printing-execution-results).

You can also read the section [Managing Ansible Stdout Callback](#managing-ansible-stdout-callback) to learn how to benefit from the stdout callback management structs provided by the _go-ansible_ library.

### Replacing the _options_ argument

With the removal of the _options_ argument, the ability to overwrite the `Executor` struct attributes in the `Execute` method is no longer available. To configure your executor, ensure that necessary settings are established during the instantiation of the struct.

This signifies that any customization or configuration of the executor should be done at the time of creating the instance, and the `Execute` method should execute the command using the predefined settings.

## Changes on the _DefaultExecute_ struct

The `DefaultExecute` struct is a ready-to-go component provided by the _go-ansible_ library for executing external commands. You can find its definition in the _github.com/apenella/go-ansible/v2/pkg/execute_ package.
Changes on the `Executor` interface impact the `DefaultExecute` struct. You can read more about the changes on the `Executor` interface [here](#changes-on-the-executor-interface).

In version _v2.x_ you need to instantiate the `DefaultExecute` struct to execute the Ansible commands, as is shown in the following code snippet.

```go
// Define the AnsiblePlaybookCmd and the required options.
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Connection: "local",
  Inventory: "all,",
}

playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml", "site2.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

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

The `NewDefaultExecute` uses the [options design pattern](https://dev.to/kittipat1413/understanding-the-options-pattern-in-go-390c) to configure the `DefaultExecute` struct. So, it can receive multiple functions to configure the `DefaultExecute` instance you create.

If you already configured a `DefaultExecute` struct in your code, you should adapt it to the new version. Follow the coming sections to learn how to adapt your code to these changes.

### Adding _Cmd_ attribute to generate commands

The `DefaultExecute` now requires a `Commander` to generate external commands. Consequently, it includes the `Cmd` attribute of type `Commander`. Both the `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs implement the `Commander` interface.

When instantiating the `DefaultExecute` struct, provide the `Cmd` attribute with a `Commander` to generate the commands. The following example demonstrates how to instantiate the `DefaultExecute` struct using an `AnsiblePlaybookCmd` as the `Commander`:

```go
// Define the AnsiblePlaybookCmd and the required options.
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Connection: "local",
  Inventory: "all,",
}

playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml", "site2.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

// Instanciate a DefaultExecutoe by providing 'playbookCmd' as the Commander.
exec := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
)
```

In the above example, `playbookCmd` is of type `Commander`. The `Cmd` value is set to `playbookCmd` using the `WithCmd` function when instantiating a new `DefaultExecute`. The `DefaultExecute` will then use `playbookCmd` to generate the command for execution.

### Adding _ErrorEnrich_ attribute to enrich error messages

The `ErrorEnrich` attribute provides the component responsible for enriching error messages. The `DefaultExecute` struct uses the `ErrorEnricher` interface to append additional information to the error message when an error occurs.

You can set that attribute when you instantiate the `DefaultExecute` struct. The following code snippet demonstrates how to instantiate a `DefaultExecute` struct with a custom `ErrorEnricher`:

```go
exec := execute.NewDefaultExecute(
  execute.WithCmd(cmd),
  execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
)
```

That is related to the [Removing the error enrichment for ansible-playbook commands](#removing-the-error-enrichment-for-ansible-playbook-commands).

### Adding _Exec_ attribute for running external commands

The `DefaultExecute` now includes the `Exec` attribute of type `Executabler`. The `Exec` component is responsible for executing external commands. If you do not define the `Exec` attribute, it defaults to using the `OsExec` struct, which wraps the `os/exec` package.

If you need a custom _executabler_, it must implement the `Executabler` interface. Learn more about the `Executabler` interface [here](#added-executabler-interface).

The example below illustrates how to instantiate a `DefaultExecute` struct with a custom `Executabler`:

```go
// Define the AnsiblePlaybookCmd and the required options.
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Connection: "local",
  Inventory: "all,",
}

playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml", "site2.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

// Define a custom Executabler
executable := &myCustomExecutabler{}

// Instanciate a DefaultExecutoe by providing 'playbookCmd' and 'executabler' as the Commander and Executabler respectively.
executor := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
  execute.WithExecutable(executable),
)
```

In the example, `executable` implements the `Executabler` interface. When creating a new `DefaultExecute`, set the value of `Exec` through the `WithExecutable` function. The `DefaultExecute` will then use the `executable` to execute the command.

### Adding _Output_ attribute for printing execution results

To handle the output of Ansible commands, the `DefaultExecute` now includes the `Output` attribute of type `ResultsOutputer`. This component manages the execution results' output, and if not specified, it uses the `DefaultResults` struct as a fallback mechanism. You can find the definition of the `ResultsOutputer` interface [here](#added-resultsoutputer-interface).

Use the `WithOutput` function from the _github.com/apenella/go-ansible/v2/pkg/execute_ package to configure the `Output` attribute during the instantiation of the `DefaultExecute` struct.

The example below demonstrates how to instantiate a `DefaultExecute` struct with a custom `ResultsOutputer`:

```go
// Define the AnsiblePlaybookCmd and the required options.
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Connection: "local",
  Inventory: "all,",
}

playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml", "site2.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

// Define a custom ResultsOutputer
output := &myCustomResultsOutputer{}

// Instanciate a DefaultExecutoe by providing 'playbookCmd' and 'outputer' as the Commander and ResultsOutputer respectively.
executor := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
  execute.WithOutput(output),
)
```

In the example above, `output` is of type `ResultsOutputer`. When creating a new `DefaultExecute`, set the `Output` value to `output` by the `WithOutput` function. The `DefaultExecute` will then use `output` to print the execution results.

### Removing the _ShowDuration_ attribute

The `DefaultExecute` has removed the `ShowDuration` attribute in version _v2.0.0_. To measure execution duration, use the `ExecutorTimeMeasurement` struct. This struct acts as a decorator over the `Executor` and is available in the _github.com/apenella/go-ansible/v2/pkg/execute/measure_ package.

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

### Removing the error enrichment for ansible-playbook commands

The `DefaultExecute` struct used to enrich the error message based on the exit code. Those enrichments are no longer available by default. The main reason is because those enrichments were based on the _ansible-playbook_ command exit code. However, the `DefaultExecute` is provided by the attribute `ErrorEnrich` to allow you to enrich the error messages.

That is related to [Adding _ErrorEnrich_ attribute to enrich error messages](#adding-errorenrich-attribute-to-enrich-error-messages).

### Changing the _Transformer_ location

If you configure transformers to modify the output of the execution's results, note that the _transformer_ package in the _go-ansible_ library has been relocated. It was moved from _github.com/apenella/go-ansible/pkg/stdoutcallback/results_ to _github.com/apenella/go-ansible/v2/pkg/execute/result/transformer_. Therefore, ensure that your code is adapted to this updated location.

Refer to the section [Changes on the _Transformer_ functions](#changes-on-the-transformer-functions) for more details on how to adapt your code to these changes.

## Changes on the _AnsiblePlaybookCmd_ struct

The `AnsiblePlaybookCmd` struct has undergone significant changes. It no longer executes commands, instead, it now implements the `Commander` interface, which is responsible for generating commands for execution. This section guides adapting your code to these changes.

### Renaming the _Options_ attribute

The `Options` attribute has been renamed to `PlaybookOptions` to better reflect its purpose. The `PlaybookOptions` attribute is of type `*AnsiblePlaybookOptions` and is used to configure the playbook execution. The `AnsiblePlaybookOptions` struct is defined in the _github.com/apenella/go-ansible/v2/pkg/playbook_ package.

### Removing the _Exec_ attribute and _Run_ method

The `AnsiblePlaybookCmd` struct no longer handles command execution, therefore, the `Exec` attribute and `Run` method have been removed. To execute a command, you should use an `Executor`. The `Executor` should receive an `AnsiblePlaybookCmd` struct to generate the command to execute.

Here's a basic example of running an ansible-playbook command through the `DefaultExecute`:

```go
// Define the AnsiblePlaybookCmd and the required options.
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Connection: "local",
  Inventory: "127.0.0.1,",
}

playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml", "site2.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

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

In the section [Using the _AnsiblePlaybookExecute_ as executor](#using-the-ansibleplaybookexecute-as-executor), you can find a more straightforward alternative to execute the `ansible-playbook` command for simple use cases.

### Removing the _StdoutCallback_ attribute

The responsibility to set the stdout callback method is no longer part of the `AnsiblePlaybookCmd` struct, therefore, the `StdoutCallback` attribute has been removed.

Configuring the stdout callback involves setting the environment variable `ANSIBLE_STDOUT_CALLBACK` and the component to handle results output from command execution. The executor is now responsible for this setup. Adapt your code to this change using the provided decorator structs. For example, setting up the JSON stdout callback method:

```go
// Define the AnsiblePlaybookCmd and the required options.
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Connection: "local",
  Inventory: "127.0.0.1,",
}

playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml", "site2.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

// Use the DefaultExecute struct to execute the command
exec := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
)

// Use the JSONStdoutCallbackExecute struct to set the JSON stdout callback method
jsonexec := stdoutcallback.NewJSONStdoutCallbackExecute(exec)

// Execute the external command through the executor
err := jsonexec.Execute(context.TODO())
if err != nil {
  panic(err)
}
```

For more details on managing Ansible Stdout Callback, refer to the [Managing Ansible Stdout Callback](#managing-ansible-stdout-callback) section.

### Using the _AnsiblePlaybookExecute_ struct as executor

The usage of `AnsiblePlaybookExecute` as an _executor_ simplifies the process of running `ansible-playbook` commands, especially for straightforward use cases. It encapsulates the instantiation of `AnsiblePlaybookCmd` and creates a `DefaultExecute` to run the command.

Here's the code snippet demonstrating how to use `AnsiblePlaybookExecute`:

```go
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Connection: "local",
  Inventory: "127.0.0.1,",
}

err := playbook.NewAnsiblePlaybookExecute("site.yml", "site2.yml").
  WithPlaybookOptions(ansiblePlaybookOptions).
  Execute(context.TODO())

if err != nil {
  fmt.Println(err.Error())
  os.Exit(1)
}
```

Please refer to the [ansibleplaybook-simple](https://github.com/apenella/go-ansible/blob/master/examples/ansibleplaybook-simple/ansibleplaybook-simple.go) example to see the complete code.

## Changes on the _AnsibleAdhocCmd_ struct

Similar to the changes made to the `AnsiblePlaybookCmd` struct, the `AnsibleAdhocCmd` struct no longer executes commands. Instead, it now implements the `Commander` interface, responsible for generating commands for execution. To adapt your code to these changes, refer to the guidelines provided in the section [Changes on the _AnsiblePlaybookCmd_ struct](#changes-on-the-ansibleplaybookcmd-struct).

## Changes on the _AnsibleInventoryCmd_ struct

The `AnsibleInventoryCmd` has undergone significant changes. It no longer executes commands, instead, it now implements the `Commander` interface, which is responsible for generating commands for execution. To adapt your code to these changes, refer to the guidelines provided in the section [Changes on the _AnsiblePlaybookCmd_ struct](#changes-on-the-ansibleplaybookcmd-struct).

## Changes on the _Transformer_ functions

In version _v2.0.0_, the _github.com/apenella/go-ansible/pkg/stdoutcallback/results_ package has been removed. This package previously contained the transformer functions responsible for modifying the output lines of the execution's results. This section guides how to adapt your code to these changes.

To adapt your code, you should update the imported package to _github.com/apenella/go-ansible/v2/pkg/execute/result/transformer_. This is the new location of the transformer functions.

The available transformer functions are still the same and you invoke them in the same way. The following is a list of available transformer functions:

- Prepend
- Append
- LogFormat
- IgnoreMessage

## Managing Ansible Stdout Callback

The latest version of _go-ansible_ introduces new features for managing stdout callback methods. This section does not delve into adapting your code to these changes but focuses on presenting the new features and how to use them. If you are seeking guidance on adapting your code, refer to the section [Changes on the _AnsiblePlaybookCmd_ struct](#changes-on-the-ansibleplaybookcmd-struct).

Configuring the StdoutCallback method involves two steps:

- Set the `ANSIBLE_STDOUT_CALLBACK` environment variable to the desired stdout callback plugin name.
- Set the method responsible for handling the results output from command execution. The responsibility of setting the StdoutCallback method has shifted to the `Executor` struct, necessitating an adjustment in your code.

To simplify the stdout callback configuration, the _go-ansible_ library provides a set of structs dedicated to setting the stdout callback method and results output mechanism. Each struct corresponds to a stdout callback plugin and is available in the _github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback_ package. The following is a list of available structs:

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

These structs serve as decorators over the `ExecutorStdoutCallbackSetter` interface, defined in _github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback_. The `ExecutorStdoutCallbackSetter` interface is defined [here](#added-executorstdoutcallbacksetter-interface).

For each stdout callback struct, there is a corresponding constructor function that takes an `ExecutorStdoutCallbackSetter` as an argument. The `DefaultExecute` struct implements the `ExecutorStdoutCallbackSetter` interface, allowing you to set the stdout callback method using the constructor functions.
The following code snippet demonstrates how to instantiate the `JSONStdoutCallbackExecute` struct:

```go
// the NewJSONStdoutCallbackExecute struct to set the JSON stdout callback method
jsonexec := stdoutcallback.NewJSONStdoutCallbackExecute(
  execute.NewDefaultExecute(
    execute.WithCmd(playbookCmd),
  )
)
```

With these new mechanisms for configuring the stdout callback method, the _github.com/apenella/go-ansible/pkg/stdoutcallback/results_ package, which defined [results functions](#replacing-the-resultsfunc-argument) used in previous _go-ansible_ versions, has been removed. Therefore, you need to adapt your code accordingly.

## Managing Ansible configuration settings

Version _v2.0.0_ introduces a new capability allowing you to configure Ansible settings for your executor. A new decorator struct, `AnsibleWithConfigurationSettingsExecute`, has been added for this purpose. To instantiate this struct, you can use the `NewAnsibleWithConfigurationSettingsExecute` function, available in the _github.com/apenella/go-ansible/v2/pkg/execute/configuration_ package. This function requires an `ExecutorEnvVarSetter` as an argument and a list of functions to configure Ansible settings. The package also provides individual functions to configure each Ansible setting, and you can find one function per Ansible setting here.

Refer to the [Added _ExecutorEnvVarSetter_ interface](#added-executorenvvarsetter-interface) section for more information about the `ExecutorEnvVarSetter` interface. The `DefaultExecute` struct implements the `ExecutorEnvVarSetter` interface, allowing you to set environment variables for the executor.

Here's an example illustrating how to prepare an executor to set Ansible configuration settings:

```go
// Define the AnsiblePlaybookCmd and the required options.
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Connection: "local",
  Inventory: "127.0.0.1,",
}

playbookCmd := playbook.NewAnsiblePlaybookCmd(
  playbook.WithPlaybooks("site.yml", "site2.yml"),
  playbook.WithPlaybookOptions(ansiblePlaybookOptions),
)

// Instanciate a DefaultExecutoe by providing 'playbookCmd' as the Commander and enabling the Ansible Force Color setting
exec := measure.NewExecutorTimeMeasurement(
  configuration.NewAnsibleWithConfigurationSettingsExecute(
    execute.NewDefaultExecute(
      execute.WithCmd(playbookCmd),
    ),
    configuration.WithAnsibleForceColor(),
  )
)
```

### Removing configuration functions

With the new capability to configure Ansible settings described [here](#managing-ansible-configuration-settings), the functions _AnsibleForceColor_, _AnsibleAvoidHostKeyChecking_, and _AnsibleSetEnv_ have been removed. You should adapt your code to these changes. The following sections explain how to do that.

#### Replacing the _AnsibleForceColor_ function

To enable the _AnsibleForceColor_ setting, the `AnsibleWithConfigurationSettingsExecute` should receive the `WithAnsibleForceColor` function as an argument. The `WithAnsibleForceColor` function is available in the _github.com/apenella/go-ansible/v2/pkg/execute/configuration_ package.

```go
// import "github.com/apenella/go-ansible/v2/pkg/execute/configuration"

// Instantiate a DefaultExecutoe by providing 'playbookCmd' as the Commander and enabling the Ansible Force Color setting
exec := measure.NewExecutorTimeMeasurement(
  configuration.NewAnsibleWithConfigurationSettingsExecute(
    execute.NewDefaultExecute(
      execute.WithCmd(playbookCmd),
    ),
    configuration.WithAnsibleForceColor(),
  )
)
```

#### Replacing the _AnsibleAvoidHostKeyChecking_ function

To disable the _AnsibleHostKeyChecking_ setting, the `AnsibleWithConfigurationSettingsExecute` should receive the `WithoutAnsibleHostKeyChecking` function as an argument. The `WithoutAnsibleHostKeyChecking` function is available in the _github.com/apenella/go-ansible/v2/pkg/execute/configuration_ package.

```go
// import "github.com/apenella/go-ansible/v2/pkg/execute/configuration"

// Instantiate a DefaultExecutoe by providing 'playbookCmd' as the Commander and enabling the Ansible Avoid Host Key Checking setting
exec := measure.NewExecutorTimeMeasurement(
  configuration.NewAnsibleWithConfigurationSettingsExecute(
    execute.NewDefaultExecute(
      execute.WithCmd(playbookCmd),
    ),
    configuration.WithoutAnsibleHostKeyChecking(),
  )
)
```

In case you need to enable the _AnsibleHostKeyChecking_ setting, you should use the `WithAnsibleHostKeyChecking` function.

#### Replacing the _AnsibleSetEnv_ function

If you used the `AnsibleSetEnv` function to set environment variables for the _executor_, you should replace it by the `AddEnvVar` method.
In case you were using the `AnsibleSetEnv` to set an _Ansible_ configuration setting, it is recommended to use the `AnsibleWithConfigurationSettingsExecute` _executor_ instead. The `AnsibleWithConfigurationSettingsExecute` struct is available in the _github.com/apenella/go-ansible/v2/pkg/execute/configuration_ package and provides you with multiple functions to configure _Ansible_ settings.

If you were previously using the `AnsibleSetEnv` function to set environment variables for the _executor_, you should replace it with the `AddEnvVar` method.

Additionally, if you were using `AnsibleSetEnv` to configure an Ansible setting, it's recommended to use the `AnsibleWithConfigurationSettingsExecute` executor instead. This struct, available in the _github.com/apenella/go-ansible/v2/pkg/execute/configuration_ package, offers multiple functions to configure Ansible settings more effectively.

Here's how you can make these replacements from the following example:

```go
// Previous usage of AnsibleSetEnv:
options.AnsibleSetEnv("ANSIBLE_LOG_PATH", "/path/to/logfile")
```

- Replacement using the `AddEnvVar` method:

```go
executor := execute.NewDefaultExecute()
executor.AddEnvVar("ANSIBLE_LOG_PATH", "/path/to/logfile")
```

- Using the `AnsibleWithConfigurationSettingsExecute` _executor_:

```go
executor := configuration.NewAnsibleWithConfigurationSettingsExecute(
    execute.NewDefaultExecute(
      execute.WithCmd(playbookCmd),
    ),
    configuration.WithAnsibleLogPath("/path/to/logfile"),
)
```
