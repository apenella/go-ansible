# Release notes

## [2.4.0] (2025-11-08)

### Changed

- Use [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) to manage the concurrent operations in the defaultExecutor and results structs.

### Fixed

- Fixed a deadlock in the defaultExecutor that occurred while handling the stdout and stderr output messages. #176
