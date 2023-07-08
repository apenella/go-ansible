# Release notes

## v1.2.0

### Added

- Introducing the `github.com/apenella/go-ansible/pkg/vault` package, which enables variable encryption.
- Added the `github.com/apenella/go-ansible/pkg/vault/password/text` package for reading encryption passwords as plain text.
- Introduced the `github.com/apenella/go-ansible/pkg/vault/password/resolve` package, which helps in resolving an encryption password.
- Added the `github.com/apenella/go-ansible/pkg/vault/password/file` package for reading encryption passwords from a file.
- Introduced the `github.com/apenella/go-ansible/pkg/vault/password/envvars` package, allowing the reading of encryption passwords from an environment variable.
- Added the `github.com/apenella/go-ansible/pkg/vault/encrypt` package, which provides the ability to encrypt strings using the `https://github.com/sosedoff/ansible-vault-go` package.
