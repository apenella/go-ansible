# go-ansible

Go-ansible is a package for running Ansible playbooks from Golang.
It only supports to run `ansible-playbook` with the most of its options.

To run a `ansible-playbook` command you must define three objectes:
- **AnsiblePlaybookCmd** object is the main object which defines the `ansible-playbook` command and how to execute it.
- **AnsiblePlaybookOptions** object has those parameters described on `Options` section within ansible-playbook's man page, and which defines how should be the `ansible-playbook` execution behavior and where to find execution configuration
- **AnsiblePlaybookConnectionOptions** object has those parameters described on `Connections Options` section within ansible-playbook's man page, and which defines how to connect to hosts.

## Executor
Go-ansible package has its own and default executor implementation which runs the `ansible-playbook`command and prints its output with a prefix on each line.
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


## Example

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

`AnsiblePlaybookCmd` where is defined the command execution.
```go
playbook := &ansibler.AnsiblePlaybookCmd{
    Playbook:          "site.yml",
    ConnectionOptions: ansiblePlaybookConnectionOptions,
    Options:           ansiblePlaybookOptions,
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
