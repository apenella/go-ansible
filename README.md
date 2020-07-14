# go-ansible

Go-ansible is a package for running Ansible playbooks from Golang.
This implementation is a fork of an upstream project and allows to better manipulate ansible output
It only supports to run `ansible-playbook` with the most of its options.

To run a `ansible-playbook` command you must define three objectes:
- **PlaybookCmd** object is the main object which defines the `ansible-playbook` command and how to execute it.
- **PlaybookOptions** object has those parameters described on `Options` section within ansible-playbook's man page, and which defines how should be the `ansible-playbook` execution behavior and where to find execution configuration
- **PlaybookConnectionOptions** object has those parameters described on `Connections Options` section within ansible-playbook's man page, and which defines how to connect to hosts.

## Executor
Go-ansible package has its own and default executor implementation which runs the `ansible-playbook`command and prints its output with a prefix on each line.
As opposed to the upstream project executor is not defined as an interface and user cannot use custom executor within modifying the library code , on the other hands library provides an embedded output parser that fills up using json output from ansible runs. In the example section below is described an example on how to grab data. Below is reported informations available in PlaybookResults struct that you can retrive after running an ansible playbook
```
type PlaybookResults struct {
	RawStdout string //json output of the playbook run
	TimeElapsed string //time elapsed for the playbook run
	Changed  int64 //number of items changed
    Failures int64 //number of failures
    Ignored int64  // so on ...
    Ok int64
    Rescued int64
    Skipped int64
    Unreachable int64
}
```
To help you parse the PlaybookResults object ther is a function PlaybookResultsChecks that applies to PlaybookResults object and return error if failures > 0 or if host is unreachable.
```
func (r *PlaybookResults) PlaybookResultsChecks() error {
	if r == nil {
		return errors.New("(ansible:PlaybookResultsChecks) -> passed result is nil")
	}
	if r.Unreachable > 0 {
		return errors.New("(ansible:Run) -> host is not reachable")
	}
	if r.Failures > 0 {
		return errors.New("(ansible:Run) -> one of tasks defined in playbook is failing")
	}
	return nil
}
```

## Example

When is needed to run an `ansible-playbook` from your Golang application using `go-ansible` package, you must define a `PlaybookCmd`,`PlaybookOptions`, `PlaybookConnectionOptions` and `PlaybookResults` as its shown below.


`PlaybookConnectionOptions` where is defined how to connect to hosts.
```go
ansiblePlaybookConnectionOptions := &ansibler.PlaybookConnectionOptions{
	Connection: "local",
}
```

`PlaybookOptions` where is defined which should be the `ansible-playbook` execution behavior and where to find execution configuration.
```go
ansiblePlaybookOptions := &ansibler.PlaybookOptions{
    Inventory: "127.0.0.1,",
}
```

`PlaybookCmd` where is defined the command execution.
```go
playbook := &ansibler.PlaybookCmd{
    Playbook:          "site.yml",
    ConnectionOptions: ansiblePlaybookConnectionOptions,
    Options:           ansiblePlaybookOptions,
    ExecPrefix:        "Go-ansible example",
}
```

Once the `PlaybookCmd` is already defined it could be run it using the `Run` method.
```go

res := &PlaybookResults{}
res, err := playbook.Run()
err = res.PlaybookResultsChecks() //you can obviously use a separated err var
if err != nil {
    panic(err)
}
fmt.Println(res.RawStdout)
```
