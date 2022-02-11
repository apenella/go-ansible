## v1.1.5

### Added
- New function `WithEnvVar` on `github.com/apenella/go-ansible/pkg/execute` package that adds environment variables to `DefaultExecutor` command.

### Fixed
- Include missing attributes on `AnsiblePlaybookJSONResultsPlayTaskHostsItem`. Those attributes are `cmd`, `skipped`, `skip_reason`, `failed`, and `failed_when_result`
 