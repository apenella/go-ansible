package galaxyroleinstall

import (
	"testing"

	"github.com/apenella/go-ansible/v2/pkg/galaxy"
	galaxyrole "github.com/apenella/go-ansible/v2/pkg/galaxy/role"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestNewAnsibleGalaxyRoleInstallCmd(t *testing.T) {
	cmd := NewAnsibleGalaxyRoleInstallCmd(
		WithBinary("ansible-galaxy-binary"),
		WithRoleNames("nginx"),
		WithGalaxyRoleInstallOptions(&AnsibleGalaxyRoleInstallOptions{
			Force: true,
		}),
	)

	expect := &AnsibleGalaxyRoleInstallCmd{
		Binary:    "ansible-galaxy-binary",
		RoleNames: []string{"nginx"},
		GalaxyRoleInstallOptions: &AnsibleGalaxyRoleInstallOptions{
			Force: true,
		},
	}

	assert.Equal(t, expect, cmd)
}

func TestAnsibleGalaxyRoleInstallCmdCommand(t *testing.T) {

	tests := []struct {
		desc    string
		cmd     *AnsibleGalaxyRoleInstallCmd
		command []string
		err     error
	}{
		{
			desc: "Testing generate a command for AnsibleGalaxyRoleInstallCmd with all flags using default binary",
			cmd: NewAnsibleGalaxyRoleInstallCmd(
				WithRoleNames("role-name"),
				WithGalaxyRoleInstallOptions(&AnsibleGalaxyRoleInstallOptions{
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
				}),
			),
			err: &errors.Error{},
			command: []string{
				galaxy.DefaultAnsibleGalaxyBinary,
				galaxyrole.AnsibleGalaxyRoleSubCommand,
				AnsibleGalaxyRoleInstallSubCommand,
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
				"role-name",
			},
		},
		{
			desc: "Testing generate a command for AnsibleGalaxyRoleInstallCmd with all flags",
			cmd: NewAnsibleGalaxyRoleInstallCmd(
				WithBinary("ansible-galaxy-binary"),
				WithRoleNames("role-name"),
				WithGalaxyRoleInstallOptions(&AnsibleGalaxyRoleInstallOptions{
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
				}),
			),
			err: &errors.Error{},
			command: []string{
				"ansible-galaxy-binary",
				galaxyrole.AnsibleGalaxyRoleSubCommand,
				AnsibleGalaxyRoleInstallSubCommand,
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
				"role-name",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			command, err := test.cmd.Command()

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.command, command, "Unexpected command value")
			}
		})
	}
}

func TestAnsibleGalaxyRoleInstallCmdString(t *testing.T) {
	tests := []struct {
		desc    string
		cmd     *AnsibleGalaxyRoleInstallCmd
		command string
	}{

		{
			desc: "Testing generate a command for AnsibleGalaxyRoleInstallCmd with all flags using default binary",
			cmd: NewAnsibleGalaxyRoleInstallCmd(
				WithRoleNames("role-name"),
				WithGalaxyRoleInstallOptions(&AnsibleGalaxyRoleInstallOptions{
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
				}),
			),
			command: "ansible-galaxy role install  --api-key apikey --force --force-with-deps --ignore-certs --ignore-errors --keep-scm-meta --no-deps --role-file rolefile --roles-path rolespath --server server --timeout timeout --token token -vvvv -v -vv -vvv -vvvv --version role-name",
		},
		{
			desc: "Testing generate a command for AnsibleGalaxyRoleInstallCmd with all flags",
			cmd: NewAnsibleGalaxyRoleInstallCmd(
				WithBinary("ansible-galaxy-binary"),
				WithRoleNames("role-name"),
				WithGalaxyRoleInstallOptions(&AnsibleGalaxyRoleInstallOptions{
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
				}),
			),
			command: "ansible-galaxy-binary role install  --api-key apikey --force --force-with-deps --ignore-certs --ignore-errors --keep-scm-meta --no-deps --role-file rolefile --roles-path rolespath --server server --timeout timeout --token token -vvvv -v -vv -vvv -vvvv --version role-name",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			command := test.cmd.String()

			assert.Equal(t, test.command, command, "Unexpected command value")
		})
	}
}
