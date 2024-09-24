# Release notes

## [unreleased]

### Fixed

- In the example _ansibleplaybook-embed-python_ upgrade the Ansible version to 2.17.4, which fixes an Ansible vulnerability. (https://github.com/apenella/go-ansible/security/dependabot/7)

### Added

- New example _ansibleplaybook-ssh-become-root-with-password/_, showcasing how to execute a playbook that requires to become root user and set the user password through the variable _ansible_sudo_pass_

### Changed

- The internal package `internal/executable/os/exec` has been moved to `pkg/execute/exec`, making it public. Along with this change, the `Exec` struct has been renamed to `OsExec`.
