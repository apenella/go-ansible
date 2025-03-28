# Release notes

## [2.2.0] (2025-03-28)

## Added

- Support the stdout callback plugin `ansible.posix.jsonl` by adding a stdout callback executor: `AnsiblePosixJsonlStdoutCallbackExecute`.
- Include the `AnsiblePlaybookJSONLEventResults` struct into the `github.com/apenella/go-ansible/v2/pkg/execute/result/json` package that represent the `ansible.posix.jsonl` events.
