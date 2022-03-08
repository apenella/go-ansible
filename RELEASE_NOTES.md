## undefined

### Fixed
- Quote `Args` value on `AnsibleAdhocCmd`'s `String` method
- On default executor, set all parent process environment variables to `cmd.Env` when a custom env vars is defined
 
### Added
- `ExecutorTimeMeasurement` is a decorator defined on `github.com/apenella/go-ansible/pkg/execute`, that measures the duration of an execution, it receives an `Executor` which is measured the execution time

### Chanded
- `MockExecute` uses `github.com/stretchr/testify/mock`
- Examples' name is prefixed by `ansibleplaybook` or `ansibleadhoc`

### Removed
- `DefaultExecutor` does not measures the execution duration anymore. Instead of it, `ExecutorTimeMeasurement` must be used
