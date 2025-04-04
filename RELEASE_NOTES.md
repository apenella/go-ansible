# Release notes

## [2.2.0] (2025-04-04)

## Added

- Include the `AnsiblePlaybookJSONLEventResults` struct into the `github.com/apenella/go-ansible/v2/pkg/execute/result/json` package that represent the `ansible.posix.jsonl` events.
- Include the `JSONLEventStdoutCallbackResults` struct into the `github.com/apenella/go-ansible/v2/pkg/execute/result/json` as a `ResultsOutputer` to handle the `ansible.posix.jsonl` stdout callback method events.
- Support the stdout callback plugin `ansible.posix.jsonl` by adding a stdout callback executor: `AnsiblePosixJsonlStdoutCallbackExecute`.
