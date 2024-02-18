# Release notes

## v2.0.0

Version 2.0.0 of *go-ansible* introduces several disruptive changes. Read the upgrade guide carefully before proceeding with the upgrade.

### BREAKING CHANGES

- The relationship between the executor and `AnsiblePlaybookCmd` / `AnsibleAdhocCmd` have undergone an important change.
  - **Inversion of responsabilities**: The executor is now responsible for executing external commands, while `AnsiblePlaybookCmd` and `AnsibleAdhocCmd` have cut down its responsibilities, primarily focusing on generating the command to be executed.
  - **Method and Attribute Removal**: The following methods and attributes have been removed on `AnsiblePlaybookCmd` and `AnsibleAdhocCmd`:
    - The `Run` method.
    - The `Exec` and `StdoutCallback` attributes.
- The `Executor` interface has undergone a significant signature change. This change entails the removal of the following arguments `resultsFunc` and `options`. The current signature is: `Execute(ctx context.Context) error`.
- The `github.com/apenella/go-ansible/pkg/stdoutcallback` package has been removed. Its responsabilities has been absorbed by two distinc packages `github.com/apenella/go-ansible/pkg/execute/result`, which manages the output of the commands, and `github.com/apenella/go-ansible/pkg/execute/stdoutcallback` that enables the setting of the stdout callback.
- The constants `AnsibleForceColorEnv` and `AnsibleHostKeyCheckingEnv` have been removed from the `github.com/apenella/go-ansible/pkg/options` package.
- The functions `AnsibleForceColor`, `AnsibleAvoidHostKeyChecking` and `AnsibleSetEnv` have been removed from the `github.com/apenella/go-ansible/pkg/options` package.
- The methods `WithWrite` and `withshowduration` has been removed from the `ExecutorTimeMeasurement` decorator. Instead, a new method named `Duration` has been introduced for obtaining the duration of the execution.

### Added

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
- An utility to generate the code for the configuration package has been introduced. This utility is located in the `utils/cmd/configGenerator.go`.

### Changed

- The `AdhocPlaybookCmd` struct has been updated to implement the `Commander` interface.
- The `AnsiblePlaybookCmd` struct has been updated to implement the `Commander` interface.
- The `DefaultExecute` struct has been updated to have a new attribute named `Exec` of type `Executabler` that is responsible for executing external commands.
- The `DefaultExecute` struct has been updated to have a new attribute named `Output` of type `ResultsOutputer` that is responsible for printing the execution's output.
- The `DefaultExecute` struct has been updated to implement the `Executor` interface.
- The `DefaultExecute` struct has been updated to implement the `ExecutorEnvVarSetter` interface.
- The `DefaultExecute` struct has been updated to implement the `ExecutorStdoutCallbackSetter` interface.
- The examples has been adapted to use executor as the component to execute Ansible commands.
- The package `github.com/apenella/go-ansible/pkg/stdoutcallback/result/transformer` has been moved to `github.com/apenella/go-ansible/pkg/execute/result/transformer`.

### Removed

- The `Exec` attribute has been removed from `AnsiblePlaybookCmd` and `AdhocPlaybookCmd`.
- The `github.com/apenella/go-ansible/pkg/stdoutcallback` package has been removed.
- The `Run` method has been removed from the `AnsiblePlaybookCmd` and `AdhocPlaybookCmd` structs.
- The `ShowDuration` attribute in the `DefaultExecute` struct has been removed.
- The `StdoutCallback` attribute has been removed from `AnsiblePlaybookCmd` and `AdhocPlaybookCmd`.
- The constants `AnsibleForceColorEnv` and `AnsibleHostKeyCheckingEnv` have been removed from the `github.com/apenella/go-ansible/pkg/options` package.
- The functions `AnsibleForceColor`, `AnsibleAvoidHostKeyChecking` and `AnsibleSetEnv` have been removed from the `github.com/apenella/go-ansible/pkg/options` package.
- The methods `WithWrite` and `withshowduration` have been removed from the `ExecutorTimeMeasurement` decorator.
