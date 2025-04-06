# Release notes

## [2.2.0] (2025-04-06)

## Added

- Add the examples `ansibleplaybook-posix-jsonl-stdout` and `ansibleplaybook-posix-jsonl-stdout-persistence` to demostrate the usage of the `AnsiblePosixJsonlStdoutCallbackExecute` executor.
- Include the `AnsiblePlaybookJSONLEventResults` struct into the `github.com/apenella/go-ansible/v2/pkg/execute/result/json` package that represent the `ansible.posix.jsonl` events.
- Include the `JSONLEventStdoutCallbackResults` struct into the `github.com/apenella/go-ansible/v2/pkg/execute/result/json` as a `ResultsOutputer` to handle the `ansible.posix.jsonl` stdout callback method events.
- Support the stdout callback plugin `ansible.posix.jsonl` by adding a stdout callback executor: `AnsiblePosixJsonlStdoutCallbackExecute`.
