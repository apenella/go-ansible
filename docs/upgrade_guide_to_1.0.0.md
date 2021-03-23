Go-ansible upgrade guide
====


<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [Upgrade from 0.x to 1.x](#upgrade-from-0x-to-1x)
  - [Package stdoutcallback and results](#package-stdoutcallback-and-results)
  - [Package execute](#package-execute)
    - [Update DefaultExecute](#update-defaultexecute)
      - [Prefixing output lines](#prefixing-output-lines)
      - [ResultsFunc](#resultsfunc)
    - [Update custom executors](#update-custom-executors)
  - [Options](#options)
    - [AnsiblePlaybookPrivilegeEscalationOptions type](#ansibleplaybookprivilegeescalationoptions-type)
    - [AnsiblePlaybookConnectionOptions type](#ansibleplaybookconnectionoptions-type)
    - [AnsiblePlaybookOptions type](#ansibleplaybookoptions-type)
  - [Package ansibler](#package-ansibler)
    - [Playbook attribute](#playbook-attribute)
    - [ExecPrefix attribute](#execprefix-attribute)
    - [CmdRunDir attribute](#cmdrundir-attribute)
    - [Writer attribute](#writer-attribute)

<!-- /code_chunk_output -->


# Upgrade from 0.x to 1.x

## Package stdoutcallback and results

Package **stdoutcallback** has been moved from `github.com/apenella/go-ansible/stdoutcallback` to `github.com/apenella/go-ansible/pkg/stdoutcallback` and package **results** from `github.com/apenella/go-ansible/stdoutcallback/results` to `github.com/apenella/go-ansible/pkg/stdoutcallback/results`

`StdoutCallbackResultsFunc` signature has been updated to `func(context.Context, io.Reader, io.Writer, ...results.TransformerFunc) error`. As first argument it requires a context and a list of transformers has been included instead of the string used as output prefix. 

Transformers gives the chance to anyone to customize the executor output. `go-ansible` is provided by four transformers but you could write your own transformers.
- [**Prepend**](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L21): Sets a prefix string to the output line
- [**Append**:](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L28) Sets a suffix string to the output line
- [**LogFormat**:](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L35) Include date time prefix to the output line
- [**IgnoreMessage**:](https://github.com/apenella/go-ansible/blob/master/pkg/stdoutcallback/results/transformer.go#L44) Ignores the output line based on the patterns it recieves as input parameters

The example [**custom-transformer-ansibleplaybook**](https://github.com/apenella/go-ansible/blob/master/examples/custom-transformer-ansibleplaybook/custom-trasnformer-ansible-playbook.go) show how to write your own transformer.

## Package execute
To make the **execute** package more flexible and customizable it has suffered many changes. 
First of all, it has been moved from `github.com/apenella/go-ansible/execute` to `github.com/apenella/go-ansible/pkg/execute`.

One the most important changes is that executors' `Execute` method has changed its signature to:
`Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error`

Since that change, `Executor` interface definition looks like as shown below:
```go
// Executor interface is satisfied by those types which has a Execute(context.Context,[]string,stdoutcallback.StdoutCallbackResultsFunc,...ExecuteOptions)error method
type Executor interface {
	Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error
}
```

Finally, `Executor` interface has been moved to package `github.com/apenella/go-ansible/pkg/execute` and will not be found on `github.com/apenella/go-ansible/ansible`.

### Update DefaultExecute

`DefaultExecute` has been also updated. Its `ResultsFunc` and `Prefix` attributes has been removed.

The current `DefaultExecute` definition is:
```go
// DefaultExecute is a simple definition of an executor
type DefaultExecute struct {
	// Writer is where is written the command stdout
	Write io.Writer
	// WriterError is where is written the command stderr
	WriterError io.Writer
	// ShowDuration enables to show the execution duration time after the command finishes
	ShowDuration bool
	// CmdRunDir specifies the working directory of the command.
	CmdRunDir string
	// OutputFormat
	Transformers []results.TransformerFunc
}
```

#### Prefixing output lines
Since prefix is not present anymore, in case you need to prefix the executor output is required to define a new `DefaultExecute` instance and passing a prepend transformer as shown below:
```go
playbook := &playbook.AnsiblePlaybookCmd{
    Playbooks:         []string{"site.yml"},
    ConnectionOptions: ansiblePlaybookConnectionOptions,
    Options:           ansiblePlaybookOptions,
    Exec: execute.NewDefaultExecute(
        execute.WithTransformers(
            // That is the prefix
            results.Prepend("Go-ansible example"),
        ),
    ),
}
```

#### ResultsFunc
Regarding `ResultsFunc` it will not be customizable.

### Update custom executors

As is said before, `Executor` interface has changed its signature and to adapt any custom executor `Execute` method must follow the signature:
```go
Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error
```
- **ctx**: It is a new parameters to specify a `context.Context` to manage the execution flow.
- **command**:  It is an `[]string` that contains either the binary to be executed and its parameters.
- **resutlsFunc**: It is a `stdoutcallback.StdoutCallbackResultsFunc` function used to write the execution output. 
- **options**:  List of `ExecuteOptions` functions which can be used to configure the executor. 
An example of `ExecuteOptions` could be any `WithAttribute` function defined on `DefaultExecute`. Here there is a function to set the writer to `DefaultExecutor`:
```go
// WithWrite set the writer to be used by DefaultExecutor
func WithWrite(w io.Writer) ExecuteOptions {
	return func(e Executor) {
		e.(*DefaultExecute).Write = w
	}
}
```

## Options
### AnsiblePlaybookPrivilegeEscalationOptions type
`AnsiblePlaybookPrivilegeEscalationOptions` type has been moved to `github.com/apenella/go-ansible/pkg/options` and renamed to `AnsiblePrivilegeEscalationOptions`.

In case you are using `AnsiblePlaybookPrivilegeEscalationOptions`, is needed to import `github.com/apenella/go-ansible/pkg/options` and rename the instance type to `options.AnsiblePrivilegeEscalationOptions`.
```go
import (
    ...
    "github.com/apenella/go-ansible/pkg/options"
    ...
)

...
    ansiblePlaybookPrivilegeEscalationOptions := &options.AnsiblePrivilegeEscalationOptions{
		Become:        true,
		AskBecomePass: true,
	}
...

```

### AnsiblePlaybookConnectionOptions type
`AnsiblePlaybookConnectionOptions` type has been moved to `github.com/apenella/go-ansible/pkg/options` and renamed to `AnsibleConnectionOptions`.

In case you are using `AnsiblePlaybookConnectionOptions`, is needed to import `github.com/apenella/go-ansible/pkg/options` and rename the instance type to `options.AnsibleConnectionOptions`.
```go
import (
    ...
    "github.com/apenella/go-ansible/pkg/options"
    ...
)

...
    ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}
...

```

### AnsiblePlaybookOptions type
`AnsiblePlaybookOptions` type has been moved to `github.com/apenella/go-ansible/pkg/playbook`

## Package ansibler
Package ansibler will not longer exist anymore and its components has been moved from `github.com/apenella/go-ansible` to `github.com/apenella/go-ansible/pkg/playbook`.

In case you are using a prior version and you want to upgrade to 1.0.0, `AnsiblePlaybookCmd` requires some updates. Below there is a list with the updates:

### Playbook attribute
`ansible-playbook` accepts a list of playbook files to be run. To make `go-ansible` fully compatible to `ansible-playbook`, `Playbook` (*string type*) attribute has been supersed by `Playbooks` (*[]string type*).

```go
playbook := &playbook.AnsiblePlaybookCmd{
	Playbooks:         []string{"site.yml"},
	Exec:              execute,
	ConnectionOptions: ansiblePlaybookConnectionOptions,
	Options:           ansiblePlaybookOptions,
}
```


### ExecPrefix attribute
Attribute has been removed an it is configured on the executor.
`execute.NewDefaultExecute` can be used to create a new `DefaultExecute`executor with a list of configurations.

```go
ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
	Connection: "local",
}

ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
	Inventory: "127.0.0.1,",
}

execute := execute.NewDefaultExecute(
  execute.WithTransformers(
    results.Prepend("Go-ansible example"),
  ),
)

playbook := &playbook.AnsiblePlaybookCmd{
	Playbooks:         []string{"site.yml"},
	Exec:              execute,
	ConnectionOptions: ansiblePlaybookConnectionOptions,
	Options:           ansiblePlaybookOptions,
}
```

### CmdRunDir attribute
Attribute has been removed an it is configured on the executor.
`execute.NewDefaultExecute` can be used to create a new `DefaultExecute`executor with a list of configurations.

```go
ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
	Connection: "local",
}

ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
	Inventory: "127.0.0.1,",
}

execute := execute.NewDefaultExecute(
  execute.WithCmdRunDir(dir)
)

playbook := &playbook.AnsiblePlaybookCmd{
	Playbooks:         []string{"site.yml"},
	Exec:              execute,
	ConnectionOptions: ansiblePlaybookConnectionOptions,
	Options:           ansiblePlaybookOptions,
}
```

### Writer attribute
Attribute has been removed an it is configured on the executor.
`execute.NewDefaultExecute` can be used to create a new `DefaultExecute`executor with a list of configurations.

```go
buff := new(bytes.Buffer)

ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
	Connection: "local",
}

ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
	Inventory: "127.0.0.1,",
}

execute := execute.NewDefaultExecute(
	execute.WithWrite(io.Writer(buff)),
)

playbook := &playbook.AnsiblePlaybookCmd{
	Playbooks:         []string{"site.yml"},
	Exec:              execute,
	ConnectionOptions: ansiblePlaybookConnectionOptions,
	Options:           ansiblePlaybookOptions,
	StdoutCallback:    "json",
}
```