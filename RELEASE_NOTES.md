## undefined

### Added
- `ExecutorTimeMeasurement` is a decorator defined on `github.com/apenella/go-ansible/pkg/execute`, that measures the duration of an execution, it receives an `Executor` which is measured the execution time.

### Chanded
- `MockExecute` uses `github.com/stretchr/testify/mock`

### Removed
- `DefaultExecutor` does not measures the execution duration anymore. Instead of it, `ExecutorTimeMeasurement` must be used.
