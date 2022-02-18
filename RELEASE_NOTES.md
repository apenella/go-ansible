## undefined

### Added
- `ExecutorTimeMeasurement` is a decorator or middleware defined on `github.com/apenella/go-ansible/pkg/execute`, that measures the duration of an execution. It receives an `Executor` and in the same time implements the `Executor` interface

### Removed
- `DefaultExecutor` does not measures the execution duration anymore. Instead of it, `ExecutorTimeMeasurement` must be used. 