# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.5.0]
### Added
- Changelog based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
- New package to manage `ansible-playbook` output
- Manage Json stdout callback results
- DefaultExecutor includes an error managemnt depending on `ansible-playbook` exit code
- Use go mod to manage dependencies

## [v0.4.1]
### Added
- start using go mod as dependencies manager

### Fixed
- fix bug ansible always showing error " error: unrecognized arguments" when use private key 

## [v0.4.0]
### Added
- Include privilege escalation options

## [v0.3.0]
### Added
- AnsiblePlaybookCmd has a Write attribute, which must be defined by user.

## [v0.2.0]
### Added
- Use package github.com/apenella/go-common-utils

### Changed
- Change package name to ansibler

