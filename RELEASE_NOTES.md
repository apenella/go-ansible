## [v1.0.0]

## Added
- Included `ansible-playbook` version `2.10.6` options on `AnsiblePlaybookOptions`
- Included `github.com/apenella/go-ansible/pkg/adhoc` package to interact to `ansible` adhoc command
- New function type `ExecuteOptions` to provide options to executor instances
- New `DefaultExecute` constructor `NewDefaultExecute` that accepts a list of `ExecuteOptions`
- New component to customize ansible output lines. That component is named *transformer*
- Include a bunch of transformers that can be already used:
    - Prepend(string): Prepends and string to the output line
    - Append(string): Appends and string to the output line
    - LogFormat(string): Prepends date time to the output line
    - IgnoreMessage([]string): Ignores the output lines based on input strings
- New private method `output` on `results` package to manage how to write the output lines and that can be used by any `StdoutCallbackResultsFunc`

## Changed
- **BREAKING CHANGE**: `ansibler` has been restructured and splitted to multiple packages:
  - Type `AnsiblePlaybookConnectionOptions` is renamed to `AnsibleConnectionOptions` and placed to `github.com/apenella/go-ansible/pkg/options`
  - Type `AnsiblePlaybookPrivilegeEscalationOptions` is renamed to `AnsiblePrivilegeEscalationOptions` and placed to `github.com/apenella/go-ansible/pkg/options`
  - All constants regarding connection options and privileged escalations options has been placed to `github.com/apenella/go-ansible/pkg/options`
  - `AnsiblePlaybookCmd` and `AnsiblePlaybookOptions` has been placed to `github.com/apenella/go-ansible/pkg/playbook`
  - All constants regarding ansible-playbook command interaction has been placed to `github.com/apenella/go-ansible/pkg/playbook`
- **BREAKING CHANGE**: `Playbook` attribute on `AnsiblePlaybookCmd` has been replaced to `Playbooks` attribut which accept multiple playbooks to be run
- **BREAKING CHANGE**: `Executor` interface has been moved from `ansibler` package to `github.com/apenella/go-ansible/pkg/execute` package
- **BREAKING CHANGE**: `Executor` interface is changed to `Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error`
- **BREAKING CHANGE**: `DefaultExecute` has been updated to use options pattern design, and includes a bunch of `WithXXX` methods to set its attributes
- **BREAKING CHANGE**: `StdoutCallbackResultsFunc` signature has been updated to `func(context.Context, io.Reader, io.Writer, ...results.TransformerFunc) error`. Prefix argument has been removed and a list of transformers could be passed to the function
- `DefaultStdoutCallbackResults` and `JSONStdoutCallbackResults` prepares default transformers for default output an calls `output`, instead of managing the output by its own

## Removed
- **BREAKING CHANGE**: Remove `ExecPrefix` from `AnsiblePlaybookCmd`
- **BREAKING CHANGE**: Remove `CmdRunDir` from `AnsiblePlaybookCmd`
- **BREAKING CHANGE**: Remove `Writer` from `AnsiblePlaybookCmd`
- **BREAKING CHANGE**: Remove `ResultsFunc` from `DefaultExecute`
- **BREAKING CHANGE**: Remove `Prefix` from `DefaultExecute`. Prefix is not manatory any more and could be added using the `Prepend` transformer.
- `skipLine` method has been removed. Replaced by `IgnoreMessage` transformer