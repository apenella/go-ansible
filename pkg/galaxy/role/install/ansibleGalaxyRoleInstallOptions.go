package galaxyroleinstall

import (
	"fmt"

	errors "github.com/apenella/go-common-utils/error"
)

const (

	// APIKeyFlag represent the API key to use to authenticate against the galaxy server. Same as --token
	APIKeyFlag = "--api-key"

	// ForceFlag represents the command line flag for forcing overwriting an existing role or role file.
	ForceFlag = "--force"

	// ForceWithDepsFlag represents the command line flag for forcing overwriting an existing role, role file, or dependencies.
	ForceWithDepsFlag = "--force-with-deps"

	// IgnoreCertsFlag represent the flag to ignore SSL certificate validation errors
	IgnoreCertsFlag = "--ignore-certs"

	// IgnoreErrorsFlag represents the command line flag for continuing processing even if a role fails to install.
	IgnoreErrorsFlag = "--ignore-errors"

	// KeepSCMMetaFlag represent the flag to use tar instead of the scm archive option when packaging the role.
	KeepSCMMetaFlag = "--keep-scm-meta"

	// NoDepsFlag represents the command line flag for not installing dependencies.
	NoDepsFlag = "--no-deps"

	// RoleFileFlag represents the command line flag for specifying the path to a file containing a list of roles to install.
	RoleFileFlag = "--role-file"

	// RolesPathFlag represents the command line flag for specifying where roles should be installed on the local filesystem.
	RolesPathFlag = "--roles-path"

	// ServerFlag represent the flag to specify the galaxy server to use
	ServerFlag = "--server"

	// TimeoutFlag represent the time to wait for operations against the galaxy server, defaults to 60s
	TimeoutFlag = "--timeout"

	// TokenFlag represent the token to use to authenticate against the galaxy server. Same as --api-key
	TokenFlag = "--token"

	// VerboseFlag verbose mode enabled
	VerboseFlag = "-vvvv"

	// VerboseVFlag verbose with -v is enabled
	VerboseVFlag = "-v"

	// VerboseVVFlag verbose with -vv is enabled
	VerboseVVFlag = "-vv"

	// VerboseVVVFlag verbose with -vvv is enabled
	VerboseVVVFlag = "-vvv"

	// VerboseVVVVFlag verbose with -vvvv is enabled
	VerboseVVVVFlag = "-vvvv"

	// VersionFlag show program's version number, config file location, configured module search path, module location, executable location and exit
	VersionFlag = "--version"
)

// AnsibleGalaxyRoleInstallOptions represents the options that can be passed to the ansible-galaxy role install command.
type AnsibleGalaxyRoleInstallOptions struct {

	// ApiKey represent the API key to use to authenticate against the galaxy server. Same as --token
	ApiKey string

	// Force represents whether to force overwriting an existing role or role file.
	Force bool

	// ForceWithDeps represents whether to force overwriting an existing role, role file, or dependencies.
	ForceWithDeps bool

	// IgnoreCerts represent the flag to ignore SSL certificate validation errors
	IgnoreCerts bool

	// IgnoreErrors represents whether to continue processing even if a role fails to install.
	IgnoreErrors bool

	// KeepSCMMeta represent the flag to use tar instead of the scm archive option when packaging the role.
	KeepSCMMeta bool

	// NoDeps represents whether to install dependencies.
	NoDeps bool

	// RoleFile represents the path to a file containing a list of roles to install.
	RoleFile string

	// RolesPath represents the path where roles should be installed on the local filesystem.
	RolesPath string

	// Server represent the flag to specify the galaxy server to use
	Server string

	// Timeout represent the time to wait for operations against the galaxy server, defaults to 60s
	Timeout string

	// Token represent the token to use to authenticate against the galaxy server. Same as --api-key
	Token string

	// Verbose verbose mode enabled
	Verbose bool

	// Verbose verbose mode -v enabled
	VerboseV bool

	// Verbose verbose mode -vv enabled
	VerboseVV bool

	// Verbose verbose mode -vvv enabled
	VerboseVVV bool

	// Verbose verbose mode -vvvv enabled
	VerboseVVVV bool

	// Version show program's version number, config file location, configured module search path, module location, executable location and exit
	Version bool
}

// GenerateCommandOptions generates the command line options for the ansible-galaxy role install command.
func (o *AnsibleGalaxyRoleInstallOptions) GenerateCommandOptions() ([]string, error) {

	errContext := "(galaxy::AnsibleGalaxyRoleInstallOptions::GenerateCommandOptions)"
	options := []string{}

	if o == nil {
		return nil, errors.New(errContext, "AnsibleGalaxyRoleInstallOptions is nil")
	}

	if o.ApiKey != "" {
		options = append(options, APIKeyFlag, o.ApiKey)
	}

	if o.Force {
		options = append(options, ForceFlag)
	}

	if o.ForceWithDeps {
		options = append(options, ForceWithDepsFlag)
	}

	if o.IgnoreCerts {
		options = append(options, IgnoreCertsFlag)
	}

	if o.IgnoreErrors {
		options = append(options, IgnoreErrorsFlag)
	}

	if o.KeepSCMMeta {
		options = append(options, KeepSCMMetaFlag)
	}

	if o.NoDeps {
		options = append(options, NoDepsFlag)
	}

	if o.RoleFile != "" {
		options = append(options, RoleFileFlag, o.RoleFile)
	}

	if o.RolesPath != "" {
		options = append(options, RolesPathFlag, o.RolesPath)
	}

	if o.Server != "" {
		options = append(options, ServerFlag, o.Server)
	}

	if o.Timeout != "" {
		options = append(options, TimeoutFlag, o.Timeout)
	}

	if o.Token != "" {
		options = append(options, TokenFlag, o.Token)
	}

	verboseFlag, err := o.generateVerbosityFlag()
	if err != nil {
		return nil, errors.New(errContext, "", err)
	}

	if verboseFlag != "" {
		options = append(options, verboseFlag)
	}

	if o.Version {
		options = append(options, VersionFlag)
	}

	return options, nil
}

// generateVerbosityFlag return a string with the verbose flag. Higher verbosity (more v's) has precedence over lower
func (o *AnsibleGalaxyRoleInstallOptions) generateVerbosityFlag() (string, error) {
	if o.Verbose {
		return VerboseFlag, nil
	}

	if o.VerboseVVVV {
		return VerboseVVVVFlag, nil
	}

	if o.VerboseVVV {
		return VerboseVVVFlag, nil
	}

	if o.VerboseVV {
		return VerboseVVFlag, nil
	}

	if o.VerboseV {
		return VerboseVFlag, nil
	}

	return "", nil
}

// String returns a string representation of the ansible-galaxy role install options.
func (o *AnsibleGalaxyRoleInstallOptions) String() string {
	str := ""

	if o.ApiKey != "" {
		str = fmt.Sprintf("%s %s %s", str, APIKeyFlag, o.ApiKey)
	}

	if o.Force {
		str = fmt.Sprintf("%s %s", str, ForceFlag)
	}

	if o.ForceWithDeps {
		str = fmt.Sprintf("%s %s", str, ForceWithDepsFlag)
	}

	if o.IgnoreCerts {
		str = fmt.Sprintf("%s %s", str, IgnoreCertsFlag)
	}

	if o.IgnoreErrors {
		str = fmt.Sprintf("%s %s", str, IgnoreErrorsFlag)
	}

	if o.KeepSCMMeta {
		str = fmt.Sprintf("%s %s", str, KeepSCMMetaFlag)
	}

	if o.NoDeps {
		str = fmt.Sprintf("%s %s", str, NoDepsFlag)
	}

	if o.RoleFile != "" {
		str = fmt.Sprintf("%s %s %s", str, RoleFileFlag, o.RoleFile)
	}

	if o.RolesPath != "" {
		str = fmt.Sprintf("%s %s %s", str, RolesPathFlag, o.RolesPath)
	}

	if o.Server != "" {
		str = fmt.Sprintf("%s %s %s", str, ServerFlag, o.Server)
	}

	if o.Timeout != "" {
		str = fmt.Sprintf("%s %s %s", str, TimeoutFlag, o.Timeout)
	}

	if o.Token != "" {
		str = fmt.Sprintf("%s %s %s", str, TokenFlag, o.Token)
	}

	if o.Verbose {
		str = fmt.Sprintf("%s %s", str, VerboseFlag)
	}

	if o.VerboseV {
		str = fmt.Sprintf("%s %s", str, VerboseVFlag)
	}

	if o.VerboseVV {
		str = fmt.Sprintf("%s %s", str, VerboseVVFlag)
	}

	if o.VerboseVVV {
		str = fmt.Sprintf("%s %s", str, VerboseVVVFlag)
	}

	if o.VerboseVVVV {
		str = fmt.Sprintf("%s %s", str, VerboseVVVVFlag)
	}

	if o.Version {
		str = fmt.Sprintf("%s %s", str, VersionFlag)
	}

	return str
}
