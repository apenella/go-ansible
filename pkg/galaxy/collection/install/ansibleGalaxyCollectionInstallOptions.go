package galaxycollectioninstall

import (
	"fmt"
	"strconv"
	"strings"

	errors "github.com/apenella/go-common-utils/error"
)

const (

	// APIKeyFlag represent the API key to use to authenticate against the galaxy server. Same as --token
	APIKeyFlag = "--api-key"

	// ClearResponseCacheFlag clears the existing server response cache.
	ClearResponseCacheFlag = "--clear-response-cache"

	// CollectionsPathFlag is the path to the directory containing your collections.
	CollectionsPathFlag = "--collections-path"

	// DisableGPGVerifyFlag disables GPG signature verification when installing collections from a Galaxy server.
	DisableGPGVerifyFlag = "--disable-gpg-verify"

	// ForceFlag forces overwriting an existing role or collection.
	ForceFlag = "--force"

	// ForceWithDepsFlag forces overwriting an existing collection and its dependencies.
	ForceWithDepsFlag = "--force-with-deps"

	// IgnoreCertsFlag ignores SSL certificate validation errors.
	IgnoreCertsFlag = "--ignore-certs"

	// IgnoreErrorsFlag ignores errors during installation and continue with the next specified collection.
	IgnoreErrorsFlag = "--ignore-errors"

	// IgnoreSignatureStatusCodeFlag suppresses this argument. It may be specified multiple times.
	IgnoreSignatureStatusCodeFlag = "--ignore-signature-status-code"

	// IgnoreSignatureStatusCodesFlag is a space separated list of status codes to ignore during signature verification.
	IgnoreSignatureStatusCodesFlag = "--ignore-signature-status-codes"

	// KeyringFlag is the keyring used during signature verification.
	KeyringFlag = "--keyring"

	// NoCacheFlag does not use the server response cache.
	NoCacheFlag = "--no-cache"

	// NoDepsFlag doesn’t download collections listed as dependencies.
	NoDepsFlag = "--no-deps"

	// OfflineFlag installs collection artifacts (tarballs) without contacting any distribution servers.
	OfflineFlag = "--offline"

	// PreFlag includes pre-release versions. Semantic versioning pre-releases are ignored by default.
	PreFlag = "--pre"

	// RequiredValidSignatureCountFlag is the number of signatures that must successfully verify the collection.
	RequiredValidSignatureCountFlag = "--required-valid-signature-count"

	// RequirementsFileFlag is a file containing a list of collections to be installed.
	RequirementsFileFlag = "--requirements-file"

	// ServerFlag is the Galaxy API server URL.
	ServerFlag = "--server"

	// SignatureFlag is an additional signature source to verify the authenticity of the MANIFEST.json.
	SignatureFlag = "--signature"

	// TimeoutFlag is the time to wait for operations against the galaxy server, defaults to 60s.
	TimeoutFlag = "--timeout"

	// TokenFlag represent the token to use to authenticate against the galaxy server. Same as --api-key
	TokenFlag = "--token"

	// UpgradeFlag upgrades installed collection artifacts. This will also update dependencies unless –no-deps is provided.
	UpgradeFlag = "--upgrade"

	// VerboseFlag verbose mode enabled
	VerboseFlag = "--verbose"

	// VersionFlag show program's version number, config file location, configured module search path, module location, executable location and exit
	VersionFlag = "--version"
)

type AnsibleGalaxyCollectionInstallOptions struct {

	// APIKey is the Ansible Galaxy API key.
	APIKey string

	// ClearResponseCache clears the existing server response cache.
	ClearResponseCache bool

	// DisableGPGVerify disables GPG signature verification when installing collections from a Galaxy server.
	DisableGPGVerify bool

	// ForceWithDeps forces overwriting an existing collection and its dependencies.
	ForceWithDeps bool

	// IgnoreSignatureStatusCode suppresses this argument. It may be specified multiple times.
	IgnoreSignatureStatusCode bool

	// IgnoreSignatureStatusCodes is a space separated list of status codes to ignore during signature verification.
	IgnoreSignatureStatusCodes string

	// Keyring is the keyring used during signature verification.
	Keyring string

	// NoCache does not use the server response cache.
	NoCache bool

	// Offline installs collection artifacts (tarballs) without contacting any distribution servers.
	Offline bool

	// Pre includes pre-release versions. Semantic versioning pre-releases are ignored by default.
	Pre bool

	// RequiredValidSignatureCount is the number of signatures that must successfully verify the collection.
	RequiredValidSignatureCount int

	// Signature is an additional signature source to verify the authenticity of the MANIFEST.json.
	Signature string

	// Timeout is the time to wait for operations against the galaxy server, defaults to 60s.
	Timeout string

	// Token is the Ansible Galaxy API key.
	Token string

	// Upgrade upgrades installed collection artifacts. This will also update dependencies unless –no-deps is provided.
	Upgrade bool

	// IgnoreCerts ignores SSL certificate validation errors.
	IgnoreCerts bool

	// Force forces overwriting an existing role or collection.
	Force bool

	// IgnoreErrors ignores errors during installation and continue with the next specified collection.
	IgnoreErrors bool

	// NoDeps doesn’t download collections listed as dependencies.
	NoDeps bool

	// CollectionsPath is the path to the directory containing your collections.
	CollectionsPath string

	// RequirementsFile is a file containing a list of collections to be installed.
	RequirementsFile string

	// Server is the Galaxy API server URL.
	Server string

	// Verbose verbose mode enabled
	Verbose bool

	// Version show program's version number, config file location, configured module search path, module location, executable location and exit
	Version bool
}

func (o *AnsibleGalaxyCollectionInstallOptions) GenerateCommandOptions() ([]string, error) {
	errContext := "(galaxy::AnsibleGalaxyCollectionInstallOptions::GenerateCommandOptions)"
	options := []string{}

	if o == nil {
		return nil, errors.New(errContext, "AnsibleGalaxyCollectionInstallOptions is nil")
	}

	if o.APIKey != "" {
		options = append(options, APIKeyFlag, o.APIKey)
	}

	if o.ClearResponseCache {
		options = append(options, ClearResponseCacheFlag)
	}

	if o.DisableGPGVerify {
		options = append(options, DisableGPGVerifyFlag)
	}

	if o.ForceWithDeps {
		options = append(options, ForceWithDepsFlag)
	}

	if o.IgnoreSignatureStatusCode {
		options = append(options, IgnoreSignatureStatusCodeFlag)
	}

	if o.IgnoreSignatureStatusCodes != "" {
		options = append(options, IgnoreSignatureStatusCodesFlag, o.IgnoreSignatureStatusCodes)
	}

	if o.Keyring != "" {
		options = append(options, KeyringFlag, o.Keyring)
	}

	if o.NoCache {
		options = append(options, NoCacheFlag)
	}

	if o.Offline {
		options = append(options, OfflineFlag)
	}

	if o.Pre {
		options = append(options, PreFlag)
	}

	if o.RequiredValidSignatureCount > 0 {
		options = append(options, RequiredValidSignatureCountFlag, strconv.Itoa(o.RequiredValidSignatureCount))
	}

	if o.Signature != "" {
		options = append(options, SignatureFlag, o.Signature)
	}

	if o.Timeout != "" {
		options = append(options, TimeoutFlag, o.Timeout)
	}

	if o.Token != "" {
		options = append(options, TokenFlag, o.Token)
	}

	if o.Upgrade {
		options = append(options, UpgradeFlag)
	}

	if o.IgnoreCerts {
		options = append(options, IgnoreCertsFlag)
	}

	if o.Force {
		options = append(options, ForceFlag)
	}

	if o.IgnoreErrors {
		options = append(options, IgnoreErrorsFlag)
	}

	if o.NoDeps {
		options = append(options, NoDepsFlag)
	}

	if o.CollectionsPath != "" {
		options = append(options, CollectionsPathFlag, o.CollectionsPath)
	}

	if o.RequirementsFile != "" {
		options = append(options, RequirementsFileFlag, o.RequirementsFile)
	}

	if o.Server != "" {
		options = append(options, ServerFlag, o.Server)
	}

	if o.Verbose {
		options = append(options, VerboseFlag)
	}

	if o.Version {
		options = append(options, VersionFlag)
	}

	return options, nil
}

// String return a string representation of the AnsibleGalaxyCollectionInstallOptions
func (o *AnsibleGalaxyCollectionInstallOptions) String() string {
	str := ""

	if o.APIKey != "" {
		str = fmt.Sprintf("%s %s %s", str, APIKeyFlag, o.APIKey)
	}

	if o.ClearResponseCache {
		str = fmt.Sprintf("%s %s", str, ClearResponseCacheFlag)
	}

	if o.DisableGPGVerify {
		str = fmt.Sprintf("%s %s", str, DisableGPGVerifyFlag)
	}

	if o.ForceWithDeps {
		str = fmt.Sprintf("%s %s", str, ForceWithDepsFlag)
	}

	if o.IgnoreSignatureStatusCode {
		str = fmt.Sprintf("%s %s", str, IgnoreSignatureStatusCodeFlag)
	}

	if o.IgnoreSignatureStatusCodes != "" {
		str = fmt.Sprintf("%s %s %s", str, IgnoreSignatureStatusCodesFlag, o.IgnoreSignatureStatusCodes)
	}

	if o.Keyring != "" {
		str = fmt.Sprintf("%s %s %s", str, KeyringFlag, o.Keyring)
	}

	if o.NoCache {
		str = fmt.Sprintf("%s %s", str, NoCacheFlag)
	}

	if o.Offline {
		str = fmt.Sprintf("%s %s", str, OfflineFlag)
	}

	if o.Pre {
		str = fmt.Sprintf("%s %s", str, PreFlag)
	}

	if o.RequiredValidSignatureCount > 0 {
		str = fmt.Sprintf("%s %s %d", str, RequiredValidSignatureCountFlag, o.RequiredValidSignatureCount)
	}

	if o.Signature != "" {
		str = fmt.Sprintf("%s %s %s", str, SignatureFlag, o.Signature)
	}

	if o.Timeout != "" {
		str = fmt.Sprintf("%s %s %s", str, TimeoutFlag, o.Timeout)
	}

	if o.Token != "" {
		str = fmt.Sprintf("%s %s %s", str, TokenFlag, o.Token)
	}

	if o.Upgrade {
		str = fmt.Sprintf("%s %s", str, UpgradeFlag)
	}

	if o.IgnoreCerts {
		str = fmt.Sprintf("%s %s", str, IgnoreCertsFlag)
	}

	if o.Force {
		str = fmt.Sprintf("%s %s", str, ForceFlag)
	}

	if o.IgnoreErrors {
		str = fmt.Sprintf("%s %s", str, IgnoreErrorsFlag)
	}

	if o.NoDeps {
		str = fmt.Sprintf("%s %s", str, NoDepsFlag)
	}

	if o.CollectionsPath != "" {
		str = fmt.Sprintf("%s %s %s", str, CollectionsPathFlag, o.CollectionsPath)
	}

	if o.RequirementsFile != "" {
		str = fmt.Sprintf("%s %s %s", str, RequirementsFileFlag, o.RequirementsFile)
	}

	if o.Server != "" {
		str = fmt.Sprintf("%s %s %s", str, ServerFlag, o.Server)
	}

	if o.Verbose {
		str = fmt.Sprintf("%s %s", str, VerboseFlag)
	}

	if o.Version {
		str = fmt.Sprintf("%s %s", str, VersionFlag)
	}

	return strings.TrimSpace(str)
}
