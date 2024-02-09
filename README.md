
# go-ansible

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) ![Test](https://github.com/apenella/go-ansible/actions/workflows/testing.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/apenella/go-ansible)](https://goreportcard.com/report/github.com/apenella/go-ansible) [![Go Reference](https://pkg.go.dev/badge/github.com/apenella/go-ansible.svg)](https://pkg.go.dev/github.com/apenella/go-ansible)

![go-ansible-logo](docs/logo/go-ansible_logo.png "Go-ansible Logo" )

Go-ansible is a Go package that allows executing `ansible-playbook` or `ansible` commands directly from Golang applications. It offers a variety of options for each command, facilitating seamless integration of Ansible functionality into your projects. Let's dive in and explore the capabilities of `go-ansible` together.

_**Important:** The master branch may contain unreleased or prereleased features. Exercise caution when utilizing the library in your projects._

> **Note**
> The latest major version of _go-ansible_, version 2.x, introduced significant changes and breaking changes. If you are currently using a version prior to 2.x, please refer to the [upgrade guide](https://github.com/apenella/go-ansible/blob/master/docs/upgrade_guide_to_2.x.md) for detailed information on how to migrate to version 2.x.
> The most relevant change is that command structs no longer execute commands. So, `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` do not require an `Executor` anymore. Instead, the `Executor` is responsible for the command execution. To achieve that, the `Executor` depends on the command structs to generate the commands to execute.



- [go-ansible](#go-ansible)
  - [Install](#install)
    - [Upgrade to 1.x](#upgrade-to-1x)
    - [Upgrade to 2.x](#upgrade-to-2x)
  - [Concepts](#concepts)
    - [Executor](#executor)
    - [Command Generator](#command-generator)
    - [Results Outputer](#results-outputer)
  - [Getting Started](#getting-started)
    - [Create the _AnsiblePlaybookCmd_ struct](#create-the-ansibleplaybookcmd-struct)
    - [Create the _DefaultExecute_ executor](#create-the-defaultexecute-executor)
    - [Manage the output of the command execution](#manage-the-output-of-the-command-execution)
  - [Usage Reference](#usage-reference)
    - [Execute](#execute)
      - [Executor interface](#executor-interface)
      - [Commander interface](#commander-interface)
      - [Executabler interface](#executabler-interface)
      - [DefaultExecute](#defaultexecute)
      - [Custom executor](#custom-executor)
      - [Measurements](#measurements)
    - [Adhoc](#adhoc)
    - [Playbook](#playbook)
    - [Options](#options)
      - [Ansible ad-hoc and ansible-playbook Common Options](#ansible-ad-hoc-and-ansible-playbook-common-options)
    - [Stdout Callback](#stdout-callback)
    - [Results](#results)
      - [Transformers](#transformers)
      - [Default](#default)
      - [JSON](#json)
        - [Manage JSON Output](#manage-json-output)
    - [Vault](#vault)
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

To install the latest stable version of _go-ansible_, use the following command:

```sh
go get github.com/apenella/go-ansible@v2.0.0
```

Use this command to fetch and install the latest version of _go-ansible_, ensuring you have the most up-to-date and stable release.

### Upgrade to 1.x

If you are currently using a _go-ansible_ version prior to _1.0.0_, note that there have been significant breaking changes introduced in version _1.0.0_ and beyond. Before proceeding with the upgrade, we highly recommend reading the [changelog](https://github.com/apenella/go-ansible/blob/master/CHANGELOG.md) and the [upgrade guide](https://github.com/apenella/go-ansible/blob/master/docs/upgrade_guide_to_1.x.md) carefully. These resources provide detailed information on the changes and steps required for a smooth transition to the new version.

### Upgrade to 2.x

Version _2.0.0_ introduced notable changes since the major version 1, including several breaking changes. The [upgrade guide](https://github.com/apenella/go-ansible/blob/master/docs/upgrade_guide_to_2.x.md) conveys the necessary information to migrate to version _2.x_. Thoroughly read that document and the [changelog](https://github.com/apenella/go-ansible/blob/master/CHANGELOG.md) before upgrading from version _1.x_ to _2.x_.

## Concepts

There are a few concepts that you need to understand before using the _go-ansible_ library. These concepts are essential to effectively use the library and to understand the examples and usage references provided in this document.

### Executor

An _executor_ is a component that executes the command and retrieves the results from the execution output. The library includes the `DefaultExecute` executor, which is the default executor implementation. However, you can also create a custom executor to meet your specific requirements.
To know more about the executor, refer to the [DefaultExecute](#defaultexecute) section.

### Command Generator

A _command generator_ or a _commander_ is responsible for generating the command to be executed. The `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs are examples of command generators. The [DefaultExecute](#defaultexecute) requires a _command generator_ to create the command to be executed.

### Results Outputer

A _results outputer_ or an _outputer_ is responsible for managing the output of the command execution. The library includes two output mechanisms: the `DefaultResults` and the `JSONResults`.

// TODO pending to complete

## Getting Started

This section will guide you through the step-by-step process of using the _go-ansible_ library. Follow these instructions to create an application that utilizes the `ansible-playbook` utility. However, the same guidelines can be applied to the `ansible` command as well.

Before proceeding, ensure you have installed the latest version of the go-ansible library. If not, please refer to the [Installation section](#install) for instructions.

To create an application that uses the `ansible-playbook` command you need to create an `AnsiblePlaybookCmd` struct. This struct defines the command to execute the `ansible-playbook` command. Then, you need to execute the command using an [executor](#executor). In that case, we will use the `DefaultExecute` executor, which is the default executor provided by the _go-ansible_ library.

### Create the _AnsiblePlaybookCmd_ struct

To execute `ansible-playbook` commands, first, define the necessary connection, playbook, and privilege escalation options.

Start by creating the `AnsiblePlaybookConnectionOptions` struct:

```go
ansiblePlaybookConnectionOptions := &options.AnsiblePlaybookConnectionOptions{
  Connection: "local",
}
```

Next, define the playbook options using the `AnsiblePlaybookOptions` struct:

```go
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
  Inventory: "127.0.0.1,",
}
```

Then, use the `AnsiblePlaybookPrivilegeEscalationOptions` struct to define the privilege escalation options:

```go
privilegeEscalationOptions := &options.AnsiblePlaybookPrivilegeEscalationOptions{
  Become:        true,
  BecomeMethod:  "sudo",
}
```

Finally, create the `AnsiblePlaybookCmd` struct that generates the command to execute the `ansible-playbook` command:

```go
playbookCmd := &playbook.AnsiblePlaybookCmd{
  Playbook:          "site.yml",
  ConnectionOptions: ansiblePlaybookConnectionOptions,
  Options:           ansiblePlaybookOptions,
  PrivilegeEscalationOptions: privilegeEscalationOptions,
}
```

Once the `AnsiblePlaybookCmd` is defined, provide the command to an executor to run the `ansible-playbook` command.

### Create the _DefaultExecute_ executor

The `DefaultExecute` struct, provided by the _go-ansible_ library, executes external commands. It requires a `Commander` responsible for generating the command to be executed. In this guided example, we will use the `AnsiblePlaybookCmd` previously defined as the `Commander` to generate the command to execute the `ansible-playbook` command.

```go
// playbookCmd is the Commander responsible for generating the command to execute
exec := execute.NewDefaultExecute(
  execute.WithCmd(playbookCmd),
)
```

Once you have defined the `DefaultExecute`, execute the `ansible-playbook` command using the following code:

```go
err := exec.Execute(context.Background())
if err != nil {
  panic(err)
}
```

### Manage the output of the command execution

By default, the `DefaultExecute` uses the `DefaultResults` struct to manage the output of the command execution. The `DefaultResults` struct handles the output as plain text.

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

The development reference is a comprehensive resource that provides everything you need to know to effectively use the library _go-ansible_.

This section describes the different packages and their main resources available in the _go-ansible_ library. You can find the complete reference of structs, methods and functions [here](https://pkg.go.dev/github.com/apenella/go-ansible).

### Execute

The execute package located in `github.com/apenella/go-ansible/pkg/execute` provides an [executor](#executor) implementation. That executor is the `DefaultExecute`. You can use it to execute any commands and handle the output of the command execution.
The same package also provides the `Executor` interface, which can be implemented to create custom executors.

#### Executor interface

The `Executor` interface requires the implementation of the `Execute` method with the following signature:

```go
type Executor interface {
  Execute(ctx context.Context) error
}
```

#### Commander interface

The `Commander` interface is used to define the components that generate the command to be executed. The `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs implement the `Commander` interface.

```go
type Commander interface {
  Command() ([]string, error)
}
```

#### Executabler interface

// TODO

The `Executabler` interface is used to define the components that can be executed. The `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` structs implement the `Executabler` interface.

```go
type Executabler interface {
  Command(name string, arg ...string) exec.Cmder
  CommandContext(ctx context.Context, name string, arg ...string) exec.Cmder
}
```

#### DefaultExecute

The `DefaultExecute` executor, provided by `go-ansible`, writes the command's stdout and stderr to the system stdout by default. However, it can be easily customized and extended to handle stdout and stderr differently. This can be achieved using the `ExecuteOptions` functions, which can be passed to the executor.

```go
// ExecuteOptions is a function to set executor options
type ExecuteOptions func(Executor)
```

The `DefaultExecute` executor also allows the configuration of Ansible through environment variables. This can be done using the `WithEnvVar` function, which injects environment variables into the command before its execution. To apply Ansible parameters as environment variables, pass `WithEnvVar` as an `ExecuteOption` to `NewDefaultExecute`.

To further customize the way results are returned to the user, `go-ansible` provides transformers. These transformers can be added to the `DefaultExecute` using the `WithTransformers(...results.TransformerFunc)` function, which can be included as an `ExecuteOption`.

For more examples and practical use cases, refer to the [examples](https://github.com/apenella/go-ansible/tree/master/examples) directory in the `go-ansible` repository.

#### Custom executor

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

func (e *MyExecutor) Execute(ctx context.Context, command []string, options ...execute.ExecuteOptions) error {
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

#### Measurements

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

### Adhoc

This section provides an overview of the `adhoc` package in the _go-ansible_ library, outlining its key components and functionalities.

The `github.com/apenella/go-ansible/pkg/adhoc` package facilitates the generation of Ansible ad-hoc commands. It does not execute the commands directly, but instead provides the necessary structs to generate the command to be executed by an executor. The `adhoc` package includes the following essential structs for executing ad-hoc commands:

- **AnsibleAdhocCmd**: This struct enables the generation of _Ansible ad-hoc_ commands. It implements the `Commander` interface, so its method `Command` returns an array of strings that represents the command to be executed. An executor can use it to create the command to be executed.
- **AnsibleAdhocOptions**: With this struct, you can define parameters described in Ansible's manual page's `Options` section.

Additionally, users can set privilege escalation options or connection options to the `AnsibleAdhocCmd`. These options are defined in the `github.com/apenella/go-ansible/pkg/options` package. Refer to the [options](#options) section for further details.

### Playbook

This section provides an overview of the `playbook` package in the _go-ansible_ library. Here are described its main components and functionalities.

The `github.com/apenella/go-ansible/pkg/playbook` package facilitates the generation of _ansible-playbook_ commands. It does not execute the commands directly, but instead provides the necessary structs to generate the command to be executed by an executor. The `playbook` package includes the following essential structs for executing ad-hoc commands:

- **AnsibleAdhocCmd**: This struct enables the generation of _ansible-playbook_ commands. It implements the `Commander` interface, so its method `Command` returns an array of strings that represents the command to be executed. An executor can use it to create the command to be executed.
- **AnsibleAdhocOptions**: With this struct, you can define parameters described in Ansible's manual page's `Options` section.

Additionally, users can set privilege escalation options or connection options to the `AnsibleAdhocCmd`. These options are defined in the `github.com/apenella/go-ansible/pkg/options` package. Refer to the [options](#options) section for further details.

### Options

These options can be used to customize the behaviour of `ansible` and `ansible-playbook` commands executions.
The `go-ansible` library provides types for defining command execution options in the `github.com/apenella/go-ansible/pkg/options` package.

#### Ansible ad-hoc and ansible-playbook Common Options

- **AnsibleConnectionOptions**: This struct includes parameters described in the Connections Options section within the ansible or ansible-playbook's manual page. It defines how to connect to hosts when executing Ansible commands.
- **AnsiblePrivilegeEscalationOptions**: This struct includes parameters described in the Escalation Options section within the ansible or ansible-playbook's manual page. It defines how to escalate privileges and become a user during ansible execution.

### Stdout Callback

In `go-ansible`, you can define a specific stdout callback method by setting the `StdoutCallback` attribute on the `AnsiblePlaybookCmd` or `AnsibleAdhocCmd` structs. This allows you to customize the output of the commands. The output is managed by a function that adheres to the following signature:

```go
// StdoutCallbackResultsFunc defines a function which manages ansible's stdout callbacks. The function expects a context, a reader that receives the data to be wrote and a writer that defines where to write the data coming from reader, Finally a list of transformers could be passed to update the output coming from the executor.
type StdoutCallbackResultsFunc func(context.Context, io.Reader, io.Writer, ...results.TransformerFunc) error
```

The functions to manage the output are defined in the `github.com/apenella/go-ansible/pkg/stdoutcallback/results` package. By utilizing these functions and defining a custom stdout callback, you can customize the output of the execution.

### Results

In the `github.com/apenella/go-ansible/pkg/execute/result` package, there are different methods available to manage the outputs of Ansible commands.

#### Transformers

In `go-ansible`, a transformer is a function that enriches or updates the output received from the executor, according to your needs. The `TransformerFunc` type defines the signature of a transformer function:

```go
// TransformerFunc is used to enrich or update messages before to be printed out
type TransformerFunc func(string) string
```

When the output is received from the executor, it is processed line by line, and the transformers are applied to each line. The `github.com/apenella/go-ansible/pkg/execute/result/transformer` package provides a set of ready-to-use transformers, but you can also write custom transformers and set them through the executor.

Here are some examples of transformers available in the results package:

- **Prepend**: Adds a prefix string to each output line.
- **Append**: Adds a suffix string to each output line.
- **LogFormat**: Includes a date-time prefix to each output line.
- **IgnoreMessage**: Ignores output lines based on the patterns provided as input parameters.

#### Default

By default, the execution results are managed by the `DefaultResults` struct, defined in the package `github.com/apenella/go-ansible/pkg/execute/result/default`. Its `Print` method handles the output by prepending the separator string `──` to each line when no transformer is defined. It also prepares all the transformers before invoking the worker function responsible for writing the output to the `io.Writer`.

The `Print` method ensures that the output is formatted correctly and provides basic handling of the results when no specific transformers are applied.

#### JSON

When the stdout callback method is set to `JSON` format, the output is managed by the [JSONStdoutCallbackResults](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/JSONResults.go#L151) method. This method prepares the worker output function to use the [IgnoreMessage](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L44) transformer, which ignores any non-JSON lines. Other transformers are ignored, except for those specific to [JSONStdoutCallbackResults](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/JSONResults.go#L151).

Within the [JSONStdoutCallbackResults](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/JSONResults.go#L151) function, there is an array called `skipPatterns` that contains matching expressions for lines that should be ignored. These patterns are used to skip specific lines that may not be relevant to the JSON output.

Here is an example of the `skipPatterns` array:

```go
skipPatterns := []string{
    // This pattern skips timer's callback whitelist output
    "^[\\s\\t]*Playbook run took [0-9]+ days, [0-9]+ hours, [0-9]+ minutes, [0-9]+ seconds$",
  }
```

##### Manage JSON Output

The [JSONStdoutCallbackResults](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/JSONResults.go#L151) method writes the `JSON` output to the provided `io.Writer` parameter. The `github.com/apenella/go-ansible/pkg/stdoutcallback/results` package includes the `ParseJSONResultsStream` function, which can be used to decode the JSON output into an `AnsiblePlaybookJSONResults` data structure. You can manipulate this data structure to format the JSON output according to your specific needs.

The expected JSON schema from `ansible-playbook` is defined [here](https://github.com/ansible/ansible/blob/v2.9.11/lib/ansible/plugins/callback/json.py) file within the Ansible repository.

### Vault

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

The `go-ansible` library provides a set of packages that can be used as `PasswordReader` to read the password for encryption. The following sections describe these packages and how they can be used.

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

The `go-ansible` library includes a variety of examples that demonstrate how to use the library in different scenarios. These examples can be found in the [examples](https://github.com/apenella/go-ansible/tree/master/examples) directory of the `go-ansible` repository.

The examples cover various use cases and provide practical demonstrations of utilizing different features and functionalities offered by `go-ansible`. They serve as a valuable resource to understand and learn how to integrate `go-ansible` into your applications.

Feel free to explore the [examples](https://github.com/apenella/go-ansible/tree/master/examples) directory to gain insights and ideas on how to leverage the `go-ansible` library in your projects.

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

The `go-ansible` project is committed to providing a friendly, safe and welcoming environment for all, regardless of gender, sexual orientation, disability, ethnicity, religion, or similar personal characteristics.

We expect all contributors, users, and community members to follow this code of conduct. This includes all interactions within the `go-ansible` community, whether online, in person, or otherwise.

Please to know more about the code of conduct refer [here](https://github.com/apenella/go-ansible/blob/master/CODE-OF-CONDUCT.md).

## License

The `go-ansible` library is available under [MIT](https://github.com/apenella/go-ansible/blob/master/LICENSE) license.
