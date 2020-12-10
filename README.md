
![go-ansible-logo](docs/logo/go-ansible_logo.png "Go-ansible Logo" )


# go-ansible

Go-ansible is a package for running Ansible playbooks from Golang applications.
It supports `ansible-playbook` command with the most of its options.

## Table of Contents
- [Packages](#packages)
  - [Ansibler](#ansibler)
  - [Execute](#execute)
  - [Stdout Callback](#stdout-callback)
  - [Results](#results)
    - [Default](#default)
    - [JSON](#json)
      - [Ansible-playbook output skipped lines](#ansible-playbook-output-skipped-lines)
      - [Manage JSON output](#manage-json-output)
- [Examples](#examples)
- [License](#license)

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
Go-ansible package has its own and default executor implementation which runs the `ansible-playbook` command and prints its output with a prefix on each line.
Whenever is required, you could write your own executor implementation and set it on `AnsiblePlaybookCmd` object, it will expect that the executor implements `Executor` interface.
```go
type Executor interface {
	Execute(command string, args []string, prefix string) error
}
```

Its possible to define your own executor and set it on `AnsiblePlaybookCmd`.
```go
type MyExecutor struct {}
func (e *MyExecutor) Execute(command string, args []string, prefix string) error {
    fmt.Println("I am doing nothing")

    return nil
}

playbook := &ansibler.AnsiblePlaybookCmd{
    Playbook:          "site.yml",
    ConnectionOptions: ansiblePlaybookConnectionOptions,
    Options:           ansiblePlaybookOptions,
    Exec:              &MyExecutor{},
}
```

When you run the playbook using your dummy executor, the output received is the next one.
```
$ go run myexecutor-ansibleplaybook.go
I am doing nothing
```

### Stdout Callback
It is possible to define and specific stdout callback method on `go-ansible`. To do that is needed to set `StdoutCallback` attribute on `AnsiblePlaybookCmd` object. Depending on the used method, the results are managed by one function or another. The functions to manage `ansible-playbook`'s output are defined on the package `github.com/apenella/go-ansible/stdoutcallback/results` and must be defined following the next signature:
```go
// StdoutCallbackResultsFunc defines a function which manages ansible's stdout callbacks. The function expects and string for prefixing output lines, a reader that receives the data to be wrote and a writer that defines where to write the data comming from reader
type StdoutCallbackResultsFunc func(string, io.Reader, io.Writer) error
```

### Results
Below are defined the methods to manage ansible playbooks outputs:

#### Default
By default, any stdout callback results is managed by **DefaultStdoutCallbackResults** results method, which writes to io.Writer `ansible-playbook`'s output, without manipulates it.

#### JSON
When the stdout callback method is defined to be in json format, the output is managed by **JSONStdoutCallbackResults** results method. This method parses the output json received from `ansible-playbook`'s output skipping the unrequired lines from the output, and writes result into io.Writer.

##### Ansible-playbook output skipped lines
Those lines from `ansible-playbook`'s output which do not belong to json are skipped and are not wrote to io.Writer.

Skip lines matching regexp are:
- "^[\\s\\t]*Playbook run took [0-9]+ days, [0-9]+ hours, [0-9]+ minutes, [0-9]+ seconds$",

##### Manage JSON output
**JSONStdoutCallbackResults** method writes to io.Writer parameter the json output.
Results packages provides a **JSONParser** that returns an **AnsiblePlaybookJSONResults**, holding the unmarshalled json on it. You could manipulate AnsiblePlaybookJSONResults object to achieve and format the json output depending on your needs.

The json schema expected from `ansible-playbook` is the defined on https://github.com/ansible/ansible/blob/v2.9.11/lib/ansible/plugins/callback/json.py.

## Examples
Below you could find a simple example of how to use `go-ansbile` but on [examples](https://github.com/apenella/go-ansible/tree/master/examples) folder there are more examples.

When is needed to run an `ansible-playbook` from your Golang application using `go-ansible` package, you must define a `AnsiblePlaybookCmd`,`AnsiblePlaybookOptions`, `AnsiblePlaybookConnectionOptions` as its shown below.

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
    ExecPrefix:        "Go-ansible example",
}
```

Once the `AnsiblePlaybookCmd` is already defined it could be run it using the `Run` method.
```go
err := playbook.Run()
if err != nil {
    panic(err)
}
```

The result of the `ansible-playbook` execution is shown below.
```
Go-ansible example =>
Go-ansible example =>  PLAY [all] *********************************************************************
Go-ansible example =>
Go-ansible example =>  TASK [Gathering Facts] *********************************************************
Go-ansible example =>  ok: [127.0.0.1]
Go-ansible example =>
Go-ansible example =>  TASK [simple-ansibleplaybook] **************************************************
Go-ansible example =>  ok: [127.0.0.1] =>
Go-ansible example =>    msg: Your are running 'simple-ansibleplaybook' example
Go-ansible example =>
Go-ansible example =>  PLAY RECAP *********************************************************************
Go-ansible example =>  127.0.0.1                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
Go-ansible example =>
Go-ansible example =>  Playbook run took 0 days, 0 hours, 0 minutes, 1 seconds
Duration: 1.816272213s
```

## License
go-ansible is available under [MIT](https://github.com/apenella/go-ansible/blob/master/LICENSE) license.
