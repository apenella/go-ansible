## v1.1.6

### Fixed
- Quote `Args` value on `AnsibleAdhocCmd`'s `String` method #91
- On default executor, set all parent process environment variables to `cmd.Env` when a custom env vars is defined #94
- Fix parsing of long lines in output #101

### Added
- `ExecutorTimeMeasurement` is a decorator defined on `github.com/apenella/go-ansible/pkg/execute`, that measures the duration of an execution, it receives an `Executor` which is measured the execution time #92
- Add `unreachable` state on task play results struct `AnsiblePlaybookJSONResultsPlayTaskHostsItem` #100

### Chanded
- `MockExecute` uses `github.com/stretchr/testify/mock` #92
- Examples' name are prefixed by `ansibleplaybook` or `ansibleadhoc`

### Removed
- `DefaultExecutor` does not measures the execution duration anymore. Instead of it, `ExecutorTimeMeasurement` must be used #92
