
![go-ansible-logo](docs/logo/go-ansible_logo.png "Go-ansible Logo" )


# go-ansible

Go-ansible is a package for running `ansible-playbook` or `ansible` commands from Golang applications.
It supports most of its options for each command.

> **Disclaimer**: master branch could contain unreleased features.

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [go-ansible](#go-ansible)
	- [Install](#install)
		- [Upgrade to 1.x](#upgrade-to-1x)
	- [Packages](#packages)
		- [Adhoc](#adhoc)
		- [Playbook](#playbook)
		- [Execute](#execute)
			- [DefaultExecute](#defaultexecute)
			- [Custom executor](#custom-executor)
			- [Measurements](#measurements)
		- [Options](#options)
			- [ansible adhoc and ansible-playbook common options](#ansible-adhoc-and-ansible-playbook-common-options)
		- [Stdout Callback](#stdout-callback)
		- [Results](#results)
			- [Transformers](#transformers)
			- [Default](#default)
			- [JSON](#json)
				- [Manage JSON output](#manage-json-output)
	- [Examples](#examples)
	- [License](#license)

<!-- /code_chunk_output -->

## Install

To install the latest stable version run the command below:
```
$ go get github.com/apenella/go-ansible@v1.1.7
```

### Upgrade to 1.x
Since `go-ansible` v1.0.0 has introduced many breaking changes read the [changelog](https://github.com/apenella/go-ansible/blob/master/CHANGELOG.md) and the [upgrade guide](https://github.com/apenella/go-ansible/blob/master/docs/upgrade_guide_to_1.0.0.md) carefully before proceeding to the upgrade.

## Packages

### Adhoc

`github.com/apenella/go-ansible/pkg/adhoc` package lets you to run `ansible` adhoc commands. You can use these `adhoc` types to run ansible commands:

- **AnsibleAdhocCmd** is the main object type which defines the `ansible` adhoc command and how to execute it. `AnsibleAdhocCmd` definition is mandatory to run any `ansible` adhoc command.
`AnsibleAdhocCmd` has a parameter that defines the `Executor` to use, the worker that launches the execution. If no `Executor` is specified, is used a bare `DefaultExecutor`.
- **AnsibleAdhocOptions** type has those parameters described on `Options` section within ansible's man page and defines how should be the `ansible` execution behaviour and where to find execution configuration.

You could also provide to `AnsiblePlaybookCmd` privilege escalation options or connection options, defined in `github.com/apenella/go-ansible/pkg/options`

### Playbook

`github.com/apenella/go-ansible/pkg/playbook` package lets you run `ansible-playbook` commands. You can use these `playbook` types to run ansible playbooks:

- **AnsiblePlaybookCmd** is the main object type which defines the `ansible-playbook` command and how to execute it. `AnsiblePlaybookCmd` definition is mandatory to run any `ansible-playbook` command.
`AnsiblePlaybookCmd` has a parameter that defines the `Executor` to use, the worker that launches the execution. If no `Executor` is specified, a bare `DefaultExecutor` is used.
- **AnsiblePlaybookOptions** type has those parameters described on `Options` section within ansible-playbook's man page and defines how should be the `ansible-playbook` execution behaviour and where to find execution configuration.

You could also provide to `AnsiblePlaybookCmd` escalation privileged options or connection options, defined in `github.com/apenella/go-ansible/pkg/options`.

### Execute
An executor is a component in charge to run the command and return the result received on stdout and stderr.
Go-ansible has a default executor implementation under the `execute` package. That executor is named `DefaultExecute`.

Any executor must comply with the `Executor` interface.
```go
// Executor interface is satisfied by those types which has a Execute(context.Context,[]string,stdoutcallback.StdoutCallbackResultsFunc,...ExecuteOptions)error method
type Executor interface {
	Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error
}
```

#### DefaultExecute
`DefaultExecute` is the executor defined in the go-ansible library.
On its most basic setup, it just writes the command stdout to system stdout and the same for stderr, but it is easy to extend the way of managing the command stdout and stderr.
To extend and update `DefaultExecute` behaviour, it comes with a bunch of `ExecuteOptions` functions which can be passed to the executor.
```go
// ExecuteOptions is a function to set executor options
type ExecuteOptions func(Executor)
```

`DefaultExecute` also allows configuring ansible through environment variables, and it should be done by `WithEnvVar` function, which injects environment variables to the command, before its execution. It should be passed to `NewDefaultExecute` as an `ExecuteOption` for each ansible parameter.

Another way to extend how to return the results to the user is by using `transformers`, which can also be added to `DefaultExecutor` through `WithTransformers( ...results.TransformerFunc) ExecuteOptions`

Take a look at the [examples](https://github.com/apenella/go-ansible/tree/master/examples)](https://github.com/apenella/go-ansible/tree/master/examples) to see how to do that.

#### Custom executor
You could write your executor implementation and set it on `AnsiblePlaybookCmd` object, whenever `DefaultExecutor` does not accomplish your requirements or expectations. `AnsiblePlaybookCmd` expects an object that implements the `Executor` interface.

Below there is an example of a custom executor which could be configured by `ExecuteOptions` functions.
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

Finally, on the next snipped is executed the `ansible-playbook` using the custom executor
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
```
$ go run myexecutor-ansibleplaybook.go
[Go ansible example] I am MyExecutor and I am doing nothing
```

#### Measurements
For the sake of taking some measurements, go-ansible includes the package `github.com/apenella/go-ansible/pkg/execute/measure`. At this moment, there is available `ExecutorTimeMeasurement` that acts as an `Executor` decorator, which could measure ansible or ansible-playbook commands' execution time.

To use the time measurement, an `ExecutorTimeMeasurement` must be created.
```go
executorTimeMeasurement := measure.NewExecutorTimeMeasurement(
		execute.NewDefaultExecute(
			execute.WithWrite(io.Writer(buff)),
		),
	)
```

Then, pass the created `ExecutorTimeMeasurement`, through the `Exec` attribute, to `AnsiblePlaybookCmd` or `AnsibleAdhocCmd`.
```go
playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         playbooksList,
		Exec:              executorTimeMeasurement,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		StdoutCallback:    "json",
	}
```

A measurement example is found in [ansibleplaybook-time-measurement](https://github.com/apenella/go-ansible/blob/master/examples/ansibleplaybook-time-measurement/ansibleplaybook-time-measurement.go).

### Options

The types to define command execution options can be found in `github.com/apenella/go-ansible/pkg/options`.

#### ansible adhoc and ansible-playbook common options
- **AnsibleConnectionOptions** object has those parameters described on `Connections Options` section within ansible-playbook's man page and defines how to connect to hosts.
- **AnsiblePrivilegeEscalationOptions** object has those parameters described on `Escalation Options` section within ansible-playbook's man page and defines how to become a user.

### Stdout Callback
It is possible to define and specific stdout callback method on `go-ansible`. You can set the `StdoutCallback` attribute on `AnsiblePlaybookCmd` object. Depending on the used method, the results are managed by one function or another. The functions to manage `ansible-playbook`'s output are defined in the package `github.com/apenella/go-ansible/pkg/stdoutcallback/results` and must be defined following the next signature:
```go
// StdoutCallbackResultsFunc defines a function which manages ansible's stdout callbacks. The function expects a context, a reader that receives the data to be wrote and a writer that defines where to write the data coming from reader, Finally a list of transformers could be passed to update the output coming from the executor.
type StdoutCallbackResultsFunc func(context.Context, io.Reader, io.Writer, ...results.TransformerFunc) error
```

### Results
Below are described the methods to manage ansible playbooks outputs:

#### Transformers
A transformer is a function whose purpose is to enrich or update the output coming from the executor and is defined by the type `TransformerFunc`.
```go
// TransformerFunc is used to enrich or update messages before to be printed out
type TransformerFunc func(string) string
```

The output coming from an executor is processed line by line and is on that step where are applied all the transformers.
`results` package provides a set of transformers ready to be used, but can also write a custom transformer and set it through the executor.

- [**Prepend**](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L21): Sets a prefix string to the output line
- [**Append**:](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L28) Sets a suffix string to the output line
- [**LogFormat**:](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L35) Include date time prefix to the output line
- [**IgnoreMessage**:](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L44) Ignores the output line based on the patterns it receives as input parameters

#### Default
By default, any stdout callback results are managed by **DefaultStdoutCallbackResults** results method.
The results method prepends the separator string `──` tho each line on stdout when no transformer is defined and prepares all the transformers before calling the worker function, which is in charge to write the output to io.Writer.

#### JSON
When the stdout callback method is defined to be in JSON format, the output is managed by **JSONStdoutCallbackResults** results method.
The results method prepares the worker output function to use the `IgnoreMessage` transformer, to ignore those non JSON lines. Any other transformer will be ignored but `JSONStdoutCallbackResults`

On **JSONStdoutCallbackResults** function is defined the `skipPatterns` array where are placed the matching expressions for the lines to be ignored.
```go
skipPatterns := []string{
		// This pattern skips timer's callback whitelist output
		"^[\\s\\t]*Playbook run took [0-9]+ days, [0-9]+ hours, [0-9]+ minutes, [0-9]+ seconds$",
	}
```

##### Manage JSON output
**JSONStdoutCallbackResults** method writes to io.Writer parameter the JSON output.
`"github.com/apenella/go-ansible/pkg/stdoutcallback/results"` package provides **ParseJSONResultsStream** function that returns an **AnsiblePlaybookJSONResults** data structure, holding the JSON output decoded on it. You could manipulate AnsiblePlaybookJSONResults data structure to achieve and format the JSON output depending on your needs.

The JSON schema expected from `ansible-playbook` is the defined one in https://github.com/ansible/ansible/blob/v2.9.11/lib/ansible/plugins/callback/json.py.

## Examples
Below you could find a step-by-step example of how to use `go-ansible` but on [examples](https://github.com/apenella/go-ansible/tree/master/examples) folder there are more examples.

When it needs to run an `ansible-playbook` from your Golang application using `go-ansible` library, you must define a `AnsiblePlaybookCmd`,`AnsiblePlaybookOptions`, `AnsiblePlaybookConnectionOptions` as it is shown below.

`AnsiblePlaybookConnectionOptions` where are defined connection options to the hosts.
```go
ansiblePlaybookConnectionOptions := &options.AnsiblePlaybookConnectionOptions{
	Connection: "local",
}
```

`AnsiblePlaybookOptions` where is defined which should be the `ansible-playbook` execution behaviour and where to find execution configuration.
```go
ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
    Inventory: "127.0.0.1,",
}
```

`AnsiblePlaybookPrivilegeEscalationOptions` where is defined if is required to become another and how to do it.
```go
privilegeEscalationOptions := &options.AnsiblePlaybookPrivilegeEscalationOptions{
    Become:        true,
    BecomeMethod:  "sudo",
}
```

`AnsiblePlaybookCmd` defines the command execution.
```go
cmd := &playbook.AnsiblePlaybookCmd{
    Playbook:          "site.yml",
    ConnectionOptions: ansiblePlaybookConnectionOptions,
    Options:           ansiblePlaybookOptions,
    PrivilegeEscalationOptions: privilegeEscalationOptions,
}
```

Once the `AnsiblePlaybookCmd` is already defined it could be run it using the `Run` method. Though is not defined an Executor `DefaultExecute` is used having the default parameters
```go
err := cmd.Run(context.TODO())
if err != nil {
    panic(err)
}
```

The result of the `ansible-playbook` execution is shown below.
```
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

## License
go-ansible is available under [MIT](https://github.com/apenella/go-ansible/blob/master/LICENSE) license.
