# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2020-07-17
### Added
- Ansible output is now in json (ANSIBLE_STDOUT_CALLBACK=json)
- added PlaybookResults to parse playbook's json results
- added ExitCode on Executor in order to provode explicit playbook fails

### Changed
- Vars renamed making names more compact

## [1.0.0] - 2020-07-17
### Added
- PlaybookCmd has a Write attribute, which must be defined by user.

## [1.0.0] - 2020-07-17
### Added
- Change package name to ansibler
- Use package github.com/apenella/go-common-utils
