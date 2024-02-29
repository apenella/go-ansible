# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v2.0.0-rc.1

Version 2.0.0 of *go-ansible* introduces several disruptive changes. Read the upgrade guide carefully before proceeding with the upgrade.

### BREAKING CHANGES

> **Note**
> The latest major version of _go-ansible_, version _2.x_, introduced significant and breaking changes. If you are currently using a version prior to _2.x_, please refer to the [upgrade guide](https://github.com/apenella/go-ansible/blob/master/docs/upgrade_guide_to_2.x.md) for detailed information on how to migrate to version _2.x_.

- The relationship between the executor and `AnsiblePlaybookCmd` / `AnsibleAdhocCmd` / `AnsibleInvetoryCmd` have undergone important changes.
  - **Inversion of responsabilities**: The executor is now responsible for executing external commands, while `AnsiblePlaybookCmd`, `AnsibleInventoryCmd` and `AnsibleAdhocCmd` have cut down its responsibilities, primarily focusing on generating the command to be executed.
  - **Method and Attribute Removal**: The following methods and attributes have been removed on `AnsiblePlaybookCmd`, `AnsibleInventoryCmd` and `AnsibleAdhocCmd`:
    - The `Run` method.
    - The `Exec` and `StdoutCallback` attributes.
  - **Attributes Renaming**: The `Options` attribute has been renamed to `PlaybookOptions` in `AnsiblePlaybookCmd`, `AdhocOptions` in `AnsibleAdhocCmd` and `InventoryOptions` in `AnsibleInventoryCmd`.
- The `Executor` interface has undergone a significant signature change. This change entails the removal of the following arguments `resultsFunc` and `options`. The current signature is: `Execute(ctx context.Context) error`.
- The `github.com/apenella/go-ansible/pkg/options` package has been removed. After that deletion the attributes from `AnsibleConnectionOptions` and `AnsiblePrivilegeEscalationOptions` attributes have been moved to the `PlaybookOptions`, `AdhocOptions` and `InventoryOptions` structs.
- The `github.com/apenella/go-ansible/pkg/stdoutcallback` package has been removed. Its responsabilities has been absorbed by two distinc packages `github.com/apenella/go-ansible/pkg/execute/result`, which manages the output of the commands, and `github.com/apenella/go-ansible/pkg/execute/stdoutcallback` that enables the setting of the stdout callback.
- The constants `AnsibleForceColorEnv` and `AnsibleHostKeyCheckingEnv` have been removed from the `github.com/apenella/go-ansible/pkg/options` package.
- The functions `AnsibleForceColor`, `AnsibleAvoidHostKeyChecking` and `AnsibleSetEnv` have been removed from the `github.com/apenella/go-ansible/pkg/options` package.
- The methods `WithWrite` and `withshowduration` has been removed from the `ExecutorTimeMeasurement` decorator. Instead, a new method named `Duration` has been introduced for obtaining the duration of the execution.

### Added

- A new _executor_ `AnsibleAdhocExecute` has been introduced. That _executor_ allows you to create an executor to run `ansible` commands using the default settings of `DefaultExecute`. This _executor_ is located in the `github.com/apenella/go-ansible/pkg/execute/adhoc` package.
- A new _executor_ `AnsibleInventoryExecute` has been introduced. That _executor_ allows you to create an executor to run `ansible-inventory` commands using the default settings of `DefaultExecute`. This _executor_ is located in the `github.com/apenella/go-ansible/pkg/execute/inventory` package.
- A new _executor_ `AnsiblePlaybookExecute` has been introduced. That _executor_ allows you to create an executor to run `ansible-playbook` commands using the default settings of `DefaultExecute`. This _executor_ is located in the `github.com/apenella/go-ansible/pkg/execute/playbook` package.
- A new interface `Commander` has been introduced in the `github.com/apenella/go-ansible/pkg/execute` package. This interface defines the criteria for a struct to be compliant in generating execution commands.
- A new interface `Executabler` has been introduced in the `github.com/apenella/go-ansible/pkg/execute` package. This interface defines the criteria for a struct to be compliant in executing external commands.
- A new interface `ExecutorEnvVarSetter` in `github.com/apenella/go-ansible/pkg/execute/configuration` that defines the criteria for a struct to be compliant in setting Ansible configuration.
- A new interface `ExecutorStdoutCallbackSetter` has been introduced in the `github.com/apenella/go-ansible/pkg/execute/stdoutcallback` package. This interface defines the criteria for a struct to be compliant in setting an executor that accepts the stdout callback configuration for Ansible executions.
- A new interface named `ResultsOutputer` has been introduced in the `github.com/apenella/go-ansible/pkg/execute/result` pacakge.  This interface defines the criteria for a struct to be compliant in printing execution results.
- A new package `github.com/apenella/go-ansible/internal/executable/os/exec` has been introduced. This package serves as a wrapper for `os.exec`.
- A new package `github.com/apenella/go-ansible/pkg/execute/configuration` that incldues the `ExecutorWithAnsibleConfigurationSettings` struct, which acts as a decorator that facilitates the configuration of Ansible settings within the executor.
- A new package `github.com/apenella/go-ansible/pkg/execute/result/default` has been introduced. This package offers the default component for printing execution results. It supersedes the `DefaultStdoutCallbackResults` function that was previously defined in the `github.com/apenella/go-ansible/pkg/stdoutcallback` package.
- A new package `github.com/apenella/go-ansible/pkg/execute/result/json` has been introduced. This package offers the component for printing execution results from the JSON stdout callback. It supersedes the `JSONStdoutCallbackResults` function that was previously defined in the `github.com/apenella/go-ansible/pkg/stdoutcallback` package.
- A new package `github.com/apenella/go-ansible/pkg/execute/stdoutcallback`. This package offers multiple decorators designed to set the stdout callback for Ansible executions.
- A new package `github.com/apenella/go-ansible/pkg/execute/workflow` has been introduced. This package allows you to define a workflow for executing multiple commands in a sequence.
- An utility to generate the code for the configuration package has been introduced. This utility is located in the `utils/cmd/configGenerator.go`.

### Changed

- The `AnsibleAdhocCmd` struct has been updated to implement the `Commander` interface.
- The `AnsibleInventoryCmd` struct has been updated to implement the `Commander` interface.
- The `AnsiblePlaybookCmd` struct has been updated to implement the `Commander` interface.
- The `AnsiblePlaybookOptions` and `AnsibleAdhocOptions` structs have been updated to include the attributes from `AnsibleConnectionOptions` and `AnsiblePrivilegeEscalationOptions`.
- The `DefaultExecute` struct has been updated to have a new attribute named `Exec` of type `Executabler` that is responsible for executing external commands.
- The `DefaultExecute` struct has been updated to have a new attribute named `Output` of type `ResultsOutputer` that is responsible for printing the execution's output.
- The `DefaultExecute` struct has been updated to implement the `Executor` interface.
- The `DefaultExecute` struct has been updated to implement the `ExecutorEnvVarSetter` interface.
- The `DefaultExecute` struct has been updated to implement the `ExecutorStdoutCallbackSetter` interface.
- The `Options` attribute in `AnsibleAdhocCmd` struct has been renamed to `AdhocOptions`.
- The `Options` attribute in `AnsibleInventoryCmd` struct has been renamed to `InventoryOptions`.
- The `Options` attribute in `AnsiblePlaybookCmd` struct has been renamed to `PlaybookOptions`.
- The examples has been adapted to use executor as the component to execute Ansible commands.
- The package `github.com/apenella/go-ansible/pkg/stdoutcallback/result/transformer` has been moved to `github.com/apenella/go-ansible/pkg/execute/result/transformer`.

### Removed

- The `Exec` attribute has been removed from `AnsiblePlaybookCmd` and `AdhocPlaybookCmd`.
- The `github.com/apenella/go-ansible/pkg/options` package has been removed. After the `AnsibleConnectionOptions` and `AnsiblePrivilegeEscalationOptions` structs are not available anymore.
- The `github.com/apenella/go-ansible/pkg/stdoutcallback` package has been removed.
- The `Run` method has been removed from the `AnsiblePlaybookCmd` and `AdhocPlaybookCmd` structs.
- The `ShowDuration` attribute in the `DefaultExecute` struct has been removed.
- The `StdoutCallback` attribute has been removed from `AnsiblePlaybookCmd` and `AdhocPlaybookCmd`.
- The constants `AnsibleForceColorEnv` and `AnsibleHostKeyCheckingEnv` have been removed from the `github.com/apenella/go-ansible/pkg/options` package.
- The functions `AnsibleForceColor`, `AnsibleAvoidHostKeyChecking` and `AnsibleSetEnv` have been removed from the `github.com/apenella/go-ansible/pkg/options` package.
- The methods `WithWrite` and `withshowduration` have been removed from the `ExecutorTimeMeasurement` decorator.

## v1.3.0

### Added

- New feature to execute the Ansible inventory command. [#132](https://github.com/apenella/go-ansible/issues/132)

## v1.2.2

### Changed

- Bump golang.org/x/crypto from 0.8.0 to 0.17.0

## v1.2.1

### Fixed

- In `AnsibleConnectionOptions`, add quotes to ssh, sftp, and scp arguments when generating the command

## v1.2.0

### Added

- Introducing the `github.com/apenella/go-ansible/pkg/vault` package, which enables variable encryption.
- Added the `github.com/apenella/go-ansible/pkg/vault/password/text` package for reading encryption passwords as plain text.
- Introduced the `github.com/apenella/go-ansible/pkg/vault/password/resolve` package, which helps in resolving an encryption password.
- Added the `github.com/apenella/go-ansible/pkg/vault/password/file` package for reading encryption passwords from a file.
- Introduced the `github.com/apenella/go-ansible/pkg/vault/password/envvars` package, allowing the reading of encryption passwords from an environment variable.
- Added the `github.com/apenella/go-ansible/pkg/vault/encrypt` package, which provides the ability to encrypt strings using the `https://github.com/sosedoff/ansible-vault-go` package.
- Included an example using `embed.FS`.

## 1.1.7

### Changed

- On `AnsiblePlaybookJSONResultsPlayTaskHostsItem`, attributes `Stdout` and `Stderr` has been changed from `string` to `interface{}` #109

### Fixed

- On `AnsiblePlaybookJSONResultsPlayTaskHostsItem`, fix `Unreachable` attribute type to `bool` #103

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

## v1.1.5

### Added

- New function `WithEnvVar` on `github.com/apenella/go-ansible/pkg/execute` package that adds environment variables to `DefaultExecutor` command.

### Fixed

- Include missing attributes on `AnsiblePlaybookJSONResultsPlayTaskHostsItem`. Those attributes are `cmd`, `skipped`, `skip_reason`, `failed`, and `failed_when_result`

## v1.1.4

### Added

- New function `ParseJSONResultsStream` on `"github.com/apenella/go-ansible/pkg/stdoutcallback/results"` that allow to parse ansible stdout json output as a stream. That method supports to parse json output when multiple playbooks are executed.

## v1.1.3

### Fixed

- New attribute `ExtraVarsFile` on `AnsiblePlaybookOptions` that allows to use YAML/JSON files to define extra-vars
- New attribute `ExtraVarsFile` on `AnsibleAdhocOptions` that allows to use YAML/JSON files to define extra-vars

## v1.1.2

### Fixed

- Include `stdout` and `stdout_lines` to `AnsiblePlaybookJSONResultsPlayTaskHostsItem`
- Include `stderr` and `stderr_lines` to `AnsiblePlaybookJSONResultsPlayTaskHostsItem`

## v1.1.1

### Changed

- update dependency package github.com/apenella/go-common-utils/error
- update dependency package github.com/apenella/go-common-utils/data

### Fixed

- Fixed(#57) typos and language mistakes on Readme file
- Fixed(#64) update `Msg` type on `AnsiblePlaybookJSONResultsPlayTaskHostsItem` from `string` to `interface{}`

## v1.1.0

### Added

- support for stdin on `DefaultExecute` Execute method

## v1.0.0

### Added

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

### Changed

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

### Removed

- **BREAKING CHANGE**: Remove `ExecPrefix` from `AnsiblePlaybookCmd`
- **BREAKING CHANGE**: Remove `CmdRunDir` from `AnsiblePlaybookCmd`
- **BREAKING CHANGE**: Remove `Writer` from `AnsiblePlaybookCmd`
- **BREAKING CHANGE**: Remove `ResultsFunc` from `DefaultExecute`
- **BREAKING CHANGE**: Remove `Prefix` from `DefaultExecute`. Prefix is not manatory any more and could be added using the `Prepend` transformer.
- `skipLine` method has been removed. Replaced by `IgnoreMessage` transformer

## v0.8.0

### Added

- Include attribute CmdRunDir on AnsiblePlaybookCmd which defines the playbook run directory
- Include attribute CmdRunDir on DefaultExecutor

## v0.7.1

### Fixed

- fix to do not use a multireader for stdout and stderr on DefaultExecutor

## v0.7.0

### Added

- Add Binary attribute to AnsiblePlaybookCmd
- Add VaultPasswordFile to AnsiblePlaybookOptions

## v0.6.1

### Changed

- On error, write to output writer either stdout and stderr

### Fixed

- Quote extravars when return command as string

## v0.6.0

### Added

- New method CheckStats on results package that validates AnsiblePlaybookJSONResults stats

### Changed

- __JSONStdoutCallbackResults__ on results package does not manipulates ansible JSON output, writes output as is into a writer
- __JSONParser__ on results package has changed its signature to _JSONParse(data []byte) (*AnsiblePlaybookJSONResults, error)_
- __simple-ansibleplaybook-json__ example has been modified to use a custom executor to manipulate the JSON output.
- Use github.com/apenella/go-common-utils/error to manage errors

## v0.5.1

### Fixed

- [#12](https://github.com/apenella/go-ansible/pull/12): Fix the concurrency issue in the defaultExecute.go

## v0.5.0

### Added

- Changelog based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
- New package to manage `ansible-playbook` output
- Manage Json stdout callback results
- DefaultExecutor includes an error managemnt depending on `ansible-playbook` exit code
- Use go mod to manage dependencies

## v0.4.1

### Added

- start using go mod as dependencies manager

### Fixed

- fix bug ansible always showing error " error: unrecognized arguments" when use private key 

## v0.4.0

### Added

- Include privilege escalation options

## v0.3.0

### Added

- AnsiblePlaybookCmd has a Write attribute, which must be defined by user.

## v0.2.0

### Added

- Use package github.com/apenella/go-common-utils

### Changed

- Change package name to ansibler
