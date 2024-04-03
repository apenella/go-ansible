package galaxyroleinstall

import (
	"testing"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestAnsibleGalaxyRoleInstallOptionsGenerateCommandOptions(t *testing.T) {

	errContext := "(galaxy::AnsibleGalaxyRoleInstallOptions::GenerateCommandOptions)"

	tests := []struct {
		desc    string
		options *AnsibleGalaxyRoleInstallOptions
		err     error
		expect  []string
	}{
		{
			desc:    "Testing nil AnsibleGalaxyRoleInstallOptions definition",
			options: nil,
			err:     errors.New(errContext, "AnsibleGalaxyRoleInstallOptions is nil"),
			expect:  []string{},
		},
		{
			desc:    "Testing an empty AnsibleGalaxyRoleInstallOptions definition",
			options: &AnsibleGalaxyRoleInstallOptions{},
			err:     nil,
			expect:  []string{},
		},
		{
			desc: "Testing AnsibleGalaxyRoleInstallOptions with all flags",
			options: &AnsibleGalaxyRoleInstallOptions{
				ApiKey:        "apikey",
				Force:         true,
				ForceWithDeps: true,
				IgnoreCerts:   true,
				IgnoreErrors:  true,
				KeepSCMMeta:   true,
				NoDeps:        true,
				RoleFile:      "rolefile",
				RolesPath:     "rolespath",
				Server:        "server",
				Timeout:       "timeout",
				Token:         "token",
				Verbose:       true,
				VerboseV:      true,
				VerboseVV:     true,
				VerboseVVV:    true,
				VerboseVVVV:   true,
				Version:       true,
			},
			err: nil,
			expect: []string{
				APIKeyFlag, "apikey",
				ForceFlag,
				ForceWithDepsFlag,
				IgnoreCertsFlag,
				IgnoreErrorsFlag,
				KeepSCMMetaFlag,
				NoDepsFlag,
				RoleFileFlag, "rolefile",
				RolesPathFlag, "rolespath",
				ServerFlag, "server",
				TimeoutFlag, "timeout",
				TokenFlag, "token",
				VerboseVVVVFlag,
				VersionFlag,
			},
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			options, err := test.options.GenerateCommandOptions()

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.expect, options, "Unexpected options value")
			}
		})
	}
}

func TestAnsibleGalaxyRoleInstallOptionsGenerateVerbosityFlag(t *testing.T) {
	tests := []struct {
		desc    string
		options *AnsibleGalaxyRoleInstallOptions
		res     string
		err     error
	}{
		{
			desc: "Testing generate verbosity flag",
			options: &AnsibleGalaxyRoleInstallOptions{
				Verbose: true,
			},
			res: "-vvvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag V",
			options: &AnsibleGalaxyRoleInstallOptions{
				VerboseV: true,
			},
			res: "-v",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VV",
			options: &AnsibleGalaxyRoleInstallOptions{
				VerboseVV: true,
			},
			res: "-vv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VVV",
			options: &AnsibleGalaxyRoleInstallOptions{
				VerboseVVV: true,
			},
			res: "-vvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VVVV",
			options: &AnsibleGalaxyRoleInstallOptions{
				VerboseVVVV: true,
			},
			res: "-vvvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VV has precedence over V",
			options: &AnsibleGalaxyRoleInstallOptions{
				VerboseVV: true,
				VerboseV:  true,
			},
			res: "-vv",
			err: &errors.Error{},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			res, err := test.options.generateVerbosityFlag()
			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, res)
			}
		})
	}
}

func TestAnsibleGalaxyRoleInstallOptionsString(t *testing.T) {
	tests := []struct {
		desc    string
		options *AnsibleGalaxyRoleInstallOptions
		expect  string
	}{
		{
			desc:    "Testing generate string from an empty AnsibleGalaxyRoleInstallOptions",
			options: &AnsibleGalaxyRoleInstallOptions{},
			expect:  "",
		},
		{
			desc: "Testing generate string from an AnsibleGalaxyRoleInstallOptions with all flags",
			options: &AnsibleGalaxyRoleInstallOptions{
				ApiKey:        "apikey",
				Force:         true,
				ForceWithDeps: true,
				IgnoreCerts:   true,
				IgnoreErrors:  true,
				KeepSCMMeta:   true,
				NoDeps:        true,
				RoleFile:      "rolefile",
				RolesPath:     "rolespath",
				Server:        "server",
				Timeout:       "timeout",
				Token:         "token",
				Verbose:       true,
				VerboseV:      true,
				VerboseVV:     true,
				VerboseVVV:    true,
				VerboseVVVV:   true,
				Version:       true,
			},
			expect: " --api-key apikey --force --force-with-deps --ignore-certs --ignore-errors --keep-scm-meta --no-deps --role-file rolefile --roles-path rolespath --server server --timeout timeout --token token -vvvv -v -vv -vvv -vvvv --version",
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			res := test.options.String()

			assert.Equal(t, test.expect, res, "Unexpected options value for AnsibleGalaxyRoleInstallOptions.String()")

		})
	}
}
