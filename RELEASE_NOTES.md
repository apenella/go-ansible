# Release notes

## v2.0.0

### Fix

- New package `github.com/apenella/go-ansible/pkg/execute/executable/os/exec` to run external commands, which acts as a wrapper of `os.exec`.
- New package `github.com/apenella/go-ansible/pkg/execute/result/default` which provides the default component to print the execution results.
- New struct `ResultsOutputer` to print the execution results.

### Changed

- BREAKING CHANGES `Executor` interface has removed the `resultsFunc` argument.
- The package `github.com/apenella/go-ansible/pkg/stdoutcallback/result/transformer` has been moved to `github.com/apenella/go-ansible/pkg/execute/result/transformer`.

### Removed

- Removed the attribute `ShowDuration` in the struct DefaultExecute
