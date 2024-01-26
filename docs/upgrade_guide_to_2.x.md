# Upgrade Guide to go-ansible 2.x

- [Upgrade Guide to go-ansible 2.x](#upgrade-guide-to-go-ansible-2x)
  - [Overview](#overview)
  - [Changes on the interfaces](#changes-on-the-interfaces)
    - [Added _Cmder_ interface](#added-cmder-interface)
    - [Added _Commander_ interface](#added-commander-interface)
    - [Added _Executabler_ interface](#added-executabler-interface)
    - [Added _ExecutorEnvVarSetter_ interface](#added-executorenvvarsetter-interface)
    - [Added _ExecutorStdoutCallbackSetter_ interface](#added-executorstdoutcallbacksetter-interface)
    - [Added _ResultsOutputer_ interface](#added-resultsoutputer-interface)
    - [Updated _Executor_ interface](#updated-executor-interface)
  - [Changes on the _Executor_ interface](#changes-on-the-executor-interface)
    - [Replacing the _command_ argument](#replacing-the-command-argument)
    - [Replacing the _resultsFunc_ argument](#replacing-the-resultsfunc-argument)
    - [Replacing the _options_ argument](#replacing-the-options-argument)
  - [Changes on the _DefaultExecute_ struct](#changes-on-the-defaultexecute-struct)
    - [Adding _Cmd_ attribute to generate commands](#adding-cmd-attribute-to-generate-commands)
    - [Adding _Exec_ attribute for running external commands](#adding-exec-attribute-for-running-external-commands)
    - [Adding _Output_ attribute for printing execution results](#adding-output-attribute-for-printing-execution-results)
    - [Removing the _ShowDuration_ attribute](#removing-the-showduration-attribute)
    - [Changing the _Transformer_ location](#changing-the-transformer-location)
  - [Changes on the _AnsiblePlaybookCmd_ struct](#changes-on-the-ansibleplaybookcmd-struct)
    - [Removing the _Exec_ attribute and _Run_ method](#removing-the-exec-attribute-and-run-method)
    - [Removing the _StdoutCallback_ attribute](#removing-the-stdoutcallback-attribute)
  - [Changes on the _AnsibleAdhocCmd_ struct](#changes-on-the-ansibleadhoccmd-struct)
  - [Changes on the _Transformer_ functions](#changes-on-the-transformer-functions)
  - [Managing Ansible Stdout Callback](#managing-ansible-stdout-callback)
  - [Managing Ansible configuration settings](#managing-ansible-configuration-settings)
    - [Removing configuration functions](#removing-configuration-functions)
      - [Replacing the _AnsibleForceColor_ function](#replacing-the-ansibleforcecolor-function)
      - [Replacing the _AnsibleAvoidHostKeyChecking_ function](#replacing-the-ansibleavoidhostkeychecking-function)
      - [Replacing the _AnsibleSetEnv_ function](#replacing-the-ansiblesetenv-function)

## Overview

This document offers guidance for upgrading from _go-ansible_ _v1.x_ to _v2.x_. It also presents the changes introduced in _go-ansible v2.0.0_ since the major version _1.x_. Some of those changes are breaking changes.

The most relevant change is that command structs no longer execute commands. So, `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` do not require an `Executor` anymore. Instead, the `Executor` is responsible for the command execution. To achieve that, the `Executor` depends on the command structs to generate the commands to execute.

Go through the following sections to learn about the changes introduced in version _2.x_ and how to adapt your code to those changes.

## Changes on the interfaces

The version _v2.x_ introduces several changes in the interfaces used by the _go-ansible_ library. Throughout this document, you will find references to these interfaces. This section presents the new interfaces and where to find them.

### Added _Cmder_ interface

The `Cmder` interface defined in _github.com/apenella/go-ansible/internal/executable/os/exec_ is a wrapper over the `os/exec` package. It is used to run external commands by the [Executabler](#added-executabler-interface) interface.

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

The `Commander` interface defined in _github.com/apenella/go-ansible/pkg/execute_ is used to generate commands to be executed. It is required by `DefaultExecute`, but you can use it to implement your custom executor.

The `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs implement the `Commander` interface. You can find the definition of the `Commander` interface below:

```go
// Commander generates commands to be executed
type Commander interface {
  Command() ([]string, error)
}
```

### Added _Executabler_ interface

The `Executabler` interface defined in _github.com/apenella/go-ansible/pkg/execute_ is used to run external commands. It is required by `DefaultExecute`, but you can use it to implement your custom executor.

The `OsExec` struct implements the `Executabler` interface. You can find the definition of the `Executabler` interface below:

```go
// Executabler is an interface to run commands
type Executabler interface {
  Command(name string, arg ...string) exec.Cmder
  CommandContext(ctx context.Context, name string, arg ...string) exec.Cmder
}
```

### Added _ExecutorEnvVarSetter_ interface

The `ExecutorEnvVarSetter` interface defined in _github.com/apenella/go-ansible/pkg/execute/configuration_ is used to set environment variables to the executor. It is required by `ExecutorWithAnsibleConfigurationSettings` decorator struct.

```go
// ExecutorEnvVarSetter extends the executor interface by adding methods to configure environment variables
type ExecutorEnvVarSetter interface {
  // executor interface defined in github.com/apenella/go-ansible/pkg/execute
  execute.Executor
  // AddEnvVar adds an environment variable to the executor
  AddEnvVar(key, value string)
}
```

### Added _ExecutorStdoutCallbackSetter_ interface

The `ExecutorStdoutCallbackSetter` interface defined in _github.com/apenella/go-ansible/pkg/execute/stdoutcallback_ is used to set the stdout callback method to the executor. It is required by the stdout callback decorator structs defined in the same package.

```go
// ExecutorStdoutCallbackSetter extends the executor interface by adding methods to configure the Stdout Callback configuration
type ExecutorStdoutCallbackSetter interface {
  // executor interface defined in github.com/apenella/go-ansible/pkg/execute
  execute.Executor
  // AddEnvVar adds an environment variable to the executor
  AddEnvVar(key, value string)
  // WithOutput sets the output mechanism to print the execution results to the executor
  WithOutput(output result.ResultsOutputer)
}
```

### Added _ResultsOutputer_ interface

The `ResultsOutputer` interface defined in _github.com/apenella/go-ansible/pkg/execute/result_ is used to print the execution results. It is required by `DefaultExecute`, but you can use it to implement your custom executor. The `DefaultResults` and `JSONResults` structs implement the `ResultsOutputer` interface.

```go
// OptionsFunc is a function that can be used to configure a ResultsOutputer struct
type OptionsFunc func(ResultsOutputer)

// ResultsOutputer is the interface that must implements an struct to print the execution results
type ResultsOutputer interface {
 Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...OptionsFunc) error
}
```

### Updated _Executor_ interface

Read the section [Changes on the _Executor_ interface](#changes-on-the-executor-interface) to learn about the changes on the `Executor` interface.

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

Instead of using the _command_ argument, the `Executor` requires a `Commander` to generate the command to execute. Therefore, your executor should include an attribute of type `Commander`. 
You can know more about the `Commander` interface [here](#added-commander-interface).

The `Command` method returns an array of strings representing the command to execute. You should provide the component responsible for executing external commands with this array.
Both the `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs implement the `Commander` interface.

You can review changes on `DefaultExecute` [here](#adding-cmd-attribute-to-generate-commands) and see how it has been adapted to use the `Commander` to generate the command to execute.

### Replacing the _resultsFunc_ argument

The _resultsFunc_ previously managed the results output from command execution. With its removal, your executor requires a new component to assume this responsibility. That component for handling the results output should be of type `ResultsOutputer`. You can find the definition of the `ResultsOutputer` interface [here](#added-resultsoutputer-interface).

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
Changes on the `Executor` interface impact the `DefaultExecute` struct. You can read more about the changes on the `Executor` interface [here](#changes-on-the-executor-interface).

In version _v2.x_ you need to instantiate the `DefaultExecute` struct to execute the Ansible commands, as is shown in the following code snippet.

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

The `DefaultExecute` requires a `Commander` to generate the external command to execute. For that reason, it includes the `Cmd` attribute of type `Commander`. Both the `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs implement the `Commander` interface.

When you instantiate the `DefaultExecute` struct, you should provide the `Cmd` attribute with a `Commander` to generate the commands. The following code shows how to instantiate the `DefaultExecute` struct using an `AnsiblePlaybookCmd` as the `Commander`.

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

In case you require to implement a custom _executabler_, it needs to implement the `Executabler` interface. Learn more about the `Executabler` interface [here](#added-executabler-interface).

Below, you can find an example of how to instantiate a `DefaultExecute` struct with a custom _executabler_.

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

In the example above, the `executable` variable implements the `Executabler` interface. When you instantiate a new `DefaultExecute`, you set the value to `Exec` through the function `WithExecutable`. So, the `DefaultExecute` will use the _executable_ to execute the command.

### Adding _Output_ attribute for printing execution results

To align with the new `Executor` interface, the `DefaultExecute` struct includes the `Output` attribute of type `ResultsOutputer`. It manages the output of Ansible commands. You can find the definition of the `ResultsOutputer` interface [here](#added-resultsoutputer-interface).

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

As announced in prior _go-ansible_ versions, the `DefaultExecute` has removed the `ShowDuration` attribute.

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

### Changing the _Transformer_ location

You can configure a set of transformers to modify the output of the execution's results. The _go-ansible_ library has moved the `transformer` package from _github.com/apenella/go-ansible/pkg/stdoutcallback/results_ to _github.com/apenella/go-ansible/pkg/execute/result/transformer_. So, you should adapt your code to this change.

## Changes on the _AnsiblePlaybookCmd_ struct

The `AnsiblePlaybookCmd` struct has undergone significant changes. It changed its responsibilities and no longer executes commands. Instead, it implements the `Commander` interface, which generates commands for execution. So, you need to adapt your code to these changes. This section outlines the necessary steps to migrate from the older version to the new one.

### Removing the _Exec_ attribute and _Run_ method

The `AnsiblePlaygookCmd` struct is not responsible for executing commands anymore. For that reason, the `Exec` attribute has been removed.

Along with the `Exec` attribute, the `Run` method is not available anymore. To execute a command, you should use an `Executor`. Then, the `Executor` should receive an `AnsiblePlaybookCmd` struct to generate the command to execute.

The following snip showcases a basic example of how to run an `ansible-playbook` command through an `Executor`.

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

The responsibility to set the stdout callback method is not part of the `AnsiblePlaybookCmd` struct anymore, so the attribute `StdoutCallback` has been removed from there.
Setting the `StdoutCallback` method implies setting the environment variable `ANSIBLE_STDOUT_CALLBACK` and the component to handle the results output from command execution. From now on, the `Executor` struct is in charge of that setup. So, you should adapt your code to this change.

The library already provides a set of structs that facilitate the stdout callback configuration. These structs act as decorators over the `Executor` struct, and are responsible for setting the stdout callback method, as well as the results output mechanism.

Here you have an example of how to set up the JSON stdout callback method.

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

Delve into the section [Managing Ansible Stdout Callback](#managing-ansible-stdout-callback) to learn the mechanism to set the stdout callback method.

## Changes on the _AnsibleAdhocCmd_ struct

Regarding the `AnsibleAdhocCmd` struct, the changes are similar to those on the `AnsiblePlaybookCmd` struct. The `AnsibleAdhocCmd` struct no longer executes commands. Instead, it implements the `Commander` interface, which generates commands for execution. So, you need to adapt your code to these changes.

Review the section [Changes on the _AnsiblePlaybookCmd_ struct](#changes-on-the-ansibleplaybookcmd-struct) to learn how to adapt your code to these changes.

## Changes on the _Transformer_ functions

The `github.com/apenella/go-ansible/pkg/stdoutcallback/results` package has been removed in version v2.0.0. This package contained the transformer functions, which are responsible for modifying the output lines of the execution's results. This section explains how to adapt your code to these changes.

To adapt your code to these changes, you should update the imported package to `github.com/apenella/go-ansible/pkg/execute/result/transformer`. That is the new location of the transformer functions.

The transformer functions available are:

- Prepend
- Append
- LogFormat
- IgnoreMessage

## Managing Ansible Stdout Callback

The new version of _go-ansible_ introduces new features to manage the stdout callback methods. The section does not explain how to adapt your code to these changes. Instead, it presents the new features and how to use them. If you want to learn how to adapt your code to these changes, refer to the section [Changes on the _AnsiblePlaybookCmd_ struct](#changes-on-the-ansibleplaybookcmd-struct).

Setting the StdoutCallback method consists of the following two steps:

- Firstly, set the environment variable `ANSIBLE_STDOUT_CALLBACK` to the name of the stdout callback plugin.
- Lastly, set the method that manages the results output from command execution. The responsibility to set the StdooutCallback method has been moved to the `Executor` struct. So, you should adapt your code to this change.

To simplify the stdout callback configuration, the _go-ansible_ library provides a set of structs responsible for setting the stdout callback method, as well as the results output mechanism. The structs are available in the _github.com/apenella/go-ansible/pkg/stdoutcallback_ package, and there is one struct per stdout callback plugin. Here you have a list of the available structs:

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

Those structs act as decorators over the `ExecutorStdoutCallbackSetter` interface, which is defined in _github.com/apenella/go-ansible/pkg/stdoutcallback_. You can find the definition of the `ExecutorStdoutCallbackSetter` interface [here](#added-executorstdoutcallbacksetter-interface).

There is a constructor function for each of the stdout callback structs. Those functions receive an `ExecutorStdoutCallbackSetter` as an argument. The following code snippet shows how to instantiate the `JSONStdoutCallbackExecute` struct.

```go
// the NewJSONStdoutCallbackExecute struct to set the JSON stdout callback method
jsonexec := stdoutcallback.NewJSONStdoutCallbackExecute(
  execute.NewDefaultExecute(
    execute.WithCmd(playbookCmd),
  )
)
```

Having these new mechanisms to set up the stdout callback method, the package _github.com/apenella/go-ansible/pkg/stdoutcallback/results_, that defined the [results functions](#replacing-the-resultsfunc-argument) used on the previous version of _go-ansible_ has been removed. So, you should adapt your code to this change.

## Managing Ansible configuration settings

A new capability provided in the version v2.0.0 is the ability to set Ansible configuration settings.
To configure Ansible settings to your executor a new decorator struct has been added, the `ExecutorWithAnsibleConfigurationSettings`. You can instantiate this struct by using the `NewExecutorWithAnsibleConfigurationSettings` function, which is available in the _github.com/apenella/go-ansible/pkg/execute/configuration_ package. The function must receive an `ExecutorEnvVarSetter` as an argument as well as a list of functions to configure the Ansible settings. The package also provides a set of functions to configure the Ansible settings, and there is one function per Ansible setting defined [here](https://docs.ansible.com/ansible/latest/reference_appendices/config.html).

To know more about the `ExecutorEnvVarSetter` interface, refer to the section [Added _ExecutorEnvVarSetter_ interface](#added-executorenvvarsetter-interface).

Here you have an example of how to prepare an executor to set the Ansible configuration settings.

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

// Instanciate a DefaultExecutoe by providing 'playbookCmd' as the Commander and enabling the Ansible Force Color setting
exec := measure.NewExecutorTimeMeasurement(
  configuration.NewExecutorWithAnsibleConfigurationSettings(
    execute.NewDefaultExecute(
      execute.WithCmd(playbookCmd),
    ),
    configuration.WithAnsibleForceColor(),
  )
)
```

### Removing configuration functions

With the new capability to configure Ansible settings, the functions AnsibleForceColor, AnsibleAvoidHostKeyChecking, and AnsibleSetEnv have been removed. So, you should adapt your code to this change. The following sections explain how to adapt your code to these changes.

#### Replacing the _AnsibleForceColor_ function

To enable the AnsibleForceColor setting, the `ExecutorWithAnsibleConfigurationSettings` should receive the `WithAnsibleForceColor` function as an argument. The `WithAnsibleForceColor` function is available in the _github.com/apenella/go-ansible/pkg/execute/configuration_ package.

```go
// import "github.com/apenella/go-ansible/pkg/execute/configuration"

// Instantiate a DefaultExecutoe by providing 'playbookCmd' as the Commander and enabling the Ansible Force Color setting
exec := measure.NewExecutorTimeMeasurement(
  configuration.NewExecutorWithAnsibleConfigurationSettings(
    execute.NewDefaultExecute(
      execute.WithCmd(playbookCmd),
    ),
    configuration.WithAnsibleForceColor(),
  )
)
```

#### Replacing the _AnsibleAvoidHostKeyChecking_ function

To disable the AnsibleAvoidHostKeyChecking setting, the `ExecutorWithAnsibleConfigurationSettings` should receive the `WithoutAnsibleHostKeyChecking` function as an argument. The `WithoutAnsibleHostKeyChecking` function is available in the _github.com/apenella/go-ansible/pkg/execute/configuration_ package.

```go
// import "github.com/apenella/go-ansible/pkg/execute/configuration"

// Instantiate a DefaultExecutoe by providing 'playbookCmd' as the Commander and enabling the Ansible Avoid Host Key Checking setting
exec := measure.NewExecutorTimeMeasurement(
  configuration.NewExecutorWithAnsibleConfigurationSettings(
    execute.NewDefaultExecute(
      execute.WithCmd(playbookCmd),
    ),
    configuration.WithoutAnsibleHostKeyChecking(),
  )
)
```

In case you need to enable the Ansible Host Key Checking setting, you should use the `WithAnsibleHostKeyChecking` function.

#### Replacing the _AnsibleSetEnv_ function

If you used the `AnsibleSetEnv` function to set environment variables to the executor, you should use the `AddEnvVar` method instead. In case of using any Ansible configuration setting, you should look for its corresponding function in the _github.com/apenella/go-ansible/pkg/execute/configuration_ package.
