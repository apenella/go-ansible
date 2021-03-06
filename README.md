
![go-ansible-logo](docs/logo/go-ansible_logo.png "Go-ansible Logo" )


# go-ansible

Go-ansible is a package for running Ansible playbooks from Golang applications.
It supports `ansible-playbook` command with the most of its options.


<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [go-ansible](#go-ansible)
  - [Install](#install)
  - [Packages](#packages)
    - [Ansibler](#ansibler)
    - [Execute](#execute)
      - [DefaultExecute](#defaultexecute)
      - [Custom executor](#custom-executor)
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

To install the lastest stable version run the command below:
```
$ go get github.com/apenella/go-ansible@v0.8.0
```

## Packages

### Ansibler

To run an `ansible-playbook` command you could define four objects, depending on your needs:
- **AnsiblePlaybookCmd** object is the main object which defines the `ansible-playbook` command and how to execute it. `AnsiblePlaybookCmd` definition is mandatory to run any `ansible-playbook` command.
`AnsiblePlaybookCmd` has a parameter that defines the `Executor` to use, the worker who launches the execution. If no `Executor` is specified, is used `DefaultExecutor`.
`AnsiblePlaybookCmd` also has an attribute to define the stdout callback method to use. Depending on that method, `go-ansible` manages the results in a specific way. Actually all stdout callback method's results are treated such the default method instead of `json` stdout callback, which parses the json an summerizes the stats per host. If no stdout callback method is specified, is used `default` stdout callback one.
- **AnsiblePlaybookOptions** object has those parameters described on `Options` section within ansible-playbook's man page, and defines how should be the `ansible-playbook` execution behavior and where to find execution configuration.
- **AnsiblePlaybookConnectionOptions** object has those parameters described on `Connections Options` section within ansible-playbook's man page, and defines how to connect to hosts.
- **PrivilegeEscalationOptions** object has those parameters described on `Escalation Options` section within ansible-playbook's man page, and defines how to become a user.

### Execute
An executor is the component in charge to run the command and return in somehow the result received on stdout an stderr.
Go-ansible has a default executor implementation under `execute` package. That executor is named `DefaultExecute`.

Any executor must complain `Executor` interface.
```go
// Executor interface is satisfied by those types which has a Execute(context.Context,[]string,stdoutcallback.StdoutCallbackResultsFunc,...ExecuteOptions)error method
type Executor interface {
	Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error
}
```

#### DefaultExecute
`DefaultExecutor` is the executor defined on go-ansible library. 
On its most basic setup it just writes the command stdout to system stdout and the same for stderr, but its easy to extend the way of managing the command stdout and stderr.
To extend and update its behavior it comes with a bunch of `ExecuteOptions` functions which can be passed to the executor.
```go
// ExecuteOptions is a function to set executor options
type ExecuteOptions func(Executor)
```

Another way to extend how to return the results to the user is by using `transformers`, which can also be added to `DefaultExecutor` through `WithTransformers( ...results.TransformerFunc) ExecuteOptions`

Take a look to the [examples](https://github.com/apenella/go-ansible/tree/master/examples) to see how to do that.

#### Custom executor
You could write your own executor implementation and set it on `AnsiblePlaybookCmd` object, whenever `DefaultExecutor` does not fits to your needs. `AnsiblePlaybookCmd` expects an object that implements the `Executor` interface.

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

When you run the playbook using your dummy executor, the output received is the next one.
```
$ go run myexecutor-ansibleplaybook.go
[Go ansible example] I am MyExecutor and I am doing nothing
```

### Stdout Callback
It is possible to define and specific stdout callback method on `go-ansible`. To do that is needed to set `StdoutCallback` attribute on `AnsiblePlaybookCmd` object. Depending on the used method, the results are managed by one function or another. The functions to manage `ansible-playbook`'s output are defined on the package `github.com/apenella/go-ansible/stdoutcallback/results` and must be defined following the next signature:
```go
// StdoutCallbackResultsFunc defines a function which manages ansible's stdout callbacks. The function expects a context, a reader that receives the data to be wrote and a writer that defines where to write the data comming from reader, Finally a list of transformers could be passed to update the output comming from the executor.
type StdoutCallbackResultsFunc func(context.Context, io.Reader, io.Writer, ...results.TransformerFunc) error
```

### Results
Below are described the methods to manage ansible playbooks outputs:

#### Transformers
A transformer is a function which purpose is to enrich or update the output comming from the executor, and are defined by the type `TransformerFunc`. 
```go
// TransformerFunc is used to enrich or update messages before to be printed out
type TransformerFunc func(string) string
```

The output comming from executor is processed line by line and is on that step where are applied all the transformers.
`results` package provides a set of transformers ready to be used, but can also defined by your own and passed through executor. 

- [**Prepend**](https://github.com/apenella/go-ansible/blob/output_transformers/stdoutcallback/results/transformer.go#L21): Sets a prefix string to the output line
- [**Append**:](https://github.com/apenella/go-ansible/blob/output_transformers/stdoutcallback/results/transformer.go#L28) Sets a suffix string to the output line
- [**LogFormat**:](https://github.com/apenella/go-ansible/blob/output_transformers/stdoutcallback/results/transformer.go#L35) Include date time prefix to the output line
- [**IgnoreMessage**:](https://github.com/apenella/go-ansible/blob/output_transformers/stdoutcallback/results/transformer.go#L44) Ignores the output line based on the patterns it recieves as input parameters

#### Default
By default, any stdout callback results is managed by **DefaultStdoutCallbackResults** results method. 
That results method prepends the separator string `──` tho each line on stdout and prepare all the transformers before to call the worker method, which is in charge to write the output to io.Writer. 

#### JSON
When the stdout callback method is defined to be in json format, the output is managed by **JSONStdoutCallbackResults** results method. 
That results method prepares the worker output function to use the `IgnoreMessage` transformer, to ignore those non json lines. Any other transformer will be ignored but `JSONStdoutCallbackResults`

On **JSONStdoutCallbackResults** function is defined the `skipPatterns` array where are placed the matching expressions for the lines to be ignored.
```go
skipPatterns := []string{
		// This pattern skips timer's callback whitelist output
		"^[\\s\\t]*Playbook run took [0-9]+ days, [0-9]+ hours, [0-9]+ minutes, [0-9]+ seconds$",
	}
```

##### Manage JSON output
**JSONStdoutCallbackResults** method writes to io.Writer parameter the json output.
Results packages provides a **JSONParser** that returns an **AnsiblePlaybookJSONResults**, holding the unmarshalled json on it. You could manipulate AnsiblePlaybookJSONResults object to achieve and format the json output depending on your needs.

The json schema expected from `ansible-playbook` is the defined on https://github.com/ansible/ansible/blob/v2.9.11/lib/ansible/plugins/callback/json.py.

## Examples
Below you could find an step by step example of how to use `go-ansbile` but on [examples](https://github.com/apenella/go-ansible/tree/master/examples) folder there are more examples.

When is needed to run an `ansible-playbook` from your Golang application using `go-ansible` library, you must define a `AnsiblePlaybookCmd`,`AnsiblePlaybookOptions`, `AnsiblePlaybookConnectionOptions` as its shown below.

`AnsiblePlaybookConnectionOptions` where is defined how to connect to hosts.
```go
ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
	Connection: "local",
}
```

`AnsiblePlaybookOptions` where is defined which should be the `ansible-playbook` execution behavior and where to find execution configuration.
```go
ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
    Inventory: "127.0.0.1,",
}
```

`AnsiblePlaybookPrivilegeEscalationOptions` where is defined wether to become another and how to do it.
```go
privilegeEscalationOptions := &ansibler.AnsiblePlaybookPrivilegeEscalationOptions{
    Become:        true,
    BecomeMethod:  "sudo",
}
```

`AnsiblePlaybookCmd` where is defined the command execution.
```go
playbook := &ansibler.AnsiblePlaybookCmd{
    Playbook:          "site.yml",
    ConnectionOptions: ansiblePlaybookConnectionOptions,
    Options:           ansiblePlaybookOptions,
    PrivilegeEscalationOptions: privilegeEscalationOptions,
}
```

Once the `AnsiblePlaybookCmd` is already defined it could be run it using the `Run` method. Though is not defined an Executor `DefaultExecute` is used having the default parameters
```go
err := playbook.Run(context.TODO())
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
