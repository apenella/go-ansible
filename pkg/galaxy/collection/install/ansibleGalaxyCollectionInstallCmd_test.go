package galaxycollectioninstall

import (
	"testing"

	"github.com/apenella/go-ansible/v2/pkg/galaxy"
	galaxycollection "github.com/apenella/go-ansible/v2/pkg/galaxy/collection"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestNewAnsibleGalaxyCollectionInstallCmd(t *testing.T) {
	cmd := NewAnsibleGalaxyCollectionInstallCmd(
		WithBinary("ansible-galaxy-binary"),
		WithCollectionNames("collection-name-1", "collection-name-2"),
		WithGalaxyCollectionInstallOptions(&AnsibleGalaxyCollectionInstallOptions{
			Force: true,
		}),
	)

	expect := &AnsibleGalaxyCollectionInstallCmd{
		Binary:          "ansible-galaxy-binary",
		CollectionNames: []string{"collection-name-1", "collection-name-2"},
		GalaxyCollectionInstallOptions: &AnsibleGalaxyCollectionInstallOptions{
			Force: true,
		},
	}

	assert.Equal(t, expect, cmd)
}

func TestAnsibleGalaxyCollectionInstallCmdCommand(t *testing.T) {

	tests := []struct {
		desc    string
		cmd     *AnsibleGalaxyCollectionInstallCmd
		command []string
		err     error
	}{
		{
			desc: "Testing generate a command for AnsibleGalaxyCollectionInstallCmd with all flags using default binary",
			cmd: NewAnsibleGalaxyCollectionInstallCmd(
				WithCollectionNames("collection-name"),
				WithGalaxyCollectionInstallOptions(&AnsibleGalaxyCollectionInstallOptions{
					ClearResponseCache:          true,
					DisableGPGVerify:            true,
					ForceWithDeps:               true,
					IgnoreSignatureStatusCode:   true,
					IgnoreSignatureStatusCodes:  "ignore_codes",
					Keyring:                     "keyring",
					NoCache:                     true,
					Offline:                     true,
					Pre:                         true,
					RequiredValidSignatureCount: 1,
					Signature:                   "signature",
					Timeout:                     "10",
					Token:                       "token",
					APIKey:                      "apikey",
					Upgrade:                     true,
					IgnoreCerts:                 true,
					Force:                       true,
					IgnoreErrors:                true,
					NoDeps:                      true,
					CollectionsPath:             "path",
					RequirementsFile:            "requirements",
					Server:                      "server",
					Verbose:                     true,
					Version:                     true,
				}),
			),
			err: &errors.Error{},
			command: []string{
				galaxy.DefaultAnsibleGalaxyBinary,
				galaxycollection.AnsibleGalaxyCollectionSubCommand,
				AnsibleGalaxyCollectionInstallSubCommand,
				APIKeyFlag, "apikey",
				ClearResponseCacheFlag,
				DisableGPGVerifyFlag,
				ForceWithDepsFlag,
				IgnoreSignatureStatusCodeFlag,
				IgnoreSignatureStatusCodesFlag, "ignore_codes",
				KeyringFlag, "keyring",
				NoCacheFlag,
				OfflineFlag,
				PreFlag,
				RequiredValidSignatureCountFlag, "1",
				SignatureFlag, "signature",
				TimeoutFlag, "10",
				TokenFlag, "token",
				UpgradeFlag,
				IgnoreCertsFlag,
				ForceFlag,
				IgnoreErrorsFlag,
				NoDepsFlag,
				CollectionsPathFlag, "path",
				RequirementsFileFlag, "requirements",
				ServerFlag, "server",
				VerboseFlag,
				VersionFlag,
				"collection-name",
			},
		},
		{
			desc: "Testing generate a command for AnsibleGalaxyCollectionInstallCmd with all flags",
			cmd: NewAnsibleGalaxyCollectionInstallCmd(
				WithBinary("ansible-galaxy-binary"),
				WithCollectionNames("collection-name"),
				WithGalaxyCollectionInstallOptions(&AnsibleGalaxyCollectionInstallOptions{
					ClearResponseCache:          true,
					DisableGPGVerify:            true,
					ForceWithDeps:               true,
					IgnoreSignatureStatusCode:   true,
					IgnoreSignatureStatusCodes:  "ignore_codes",
					Keyring:                     "keyring",
					NoCache:                     true,
					Offline:                     true,
					Pre:                         true,
					RequiredValidSignatureCount: 1,
					Signature:                   "signature",
					Timeout:                     "10",
					Token:                       "token",
					APIKey:                      "apikey",
					Upgrade:                     true,
					IgnoreCerts:                 true,
					Force:                       true,
					IgnoreErrors:                true,
					NoDeps:                      true,
					CollectionsPath:             "path",
					RequirementsFile:            "requirements",
					Server:                      "server",
					Verbose:                     true,
					Version:                     true,
				}),
			),
			err: &errors.Error{},
			command: []string{
				"ansible-galaxy-binary",
				galaxycollection.AnsibleGalaxyCollectionSubCommand,
				AnsibleGalaxyCollectionInstallSubCommand,
				APIKeyFlag, "apikey",
				ClearResponseCacheFlag,
				DisableGPGVerifyFlag,
				ForceWithDepsFlag,
				IgnoreSignatureStatusCodeFlag,
				IgnoreSignatureStatusCodesFlag, "ignore_codes",
				KeyringFlag, "keyring",
				NoCacheFlag,
				OfflineFlag,
				PreFlag,
				RequiredValidSignatureCountFlag, "1",
				SignatureFlag, "signature",
				TimeoutFlag, "10",
				TokenFlag, "token",
				UpgradeFlag,
				IgnoreCertsFlag,
				ForceFlag,
				IgnoreErrorsFlag,
				NoDepsFlag,
				CollectionsPathFlag, "path",
				RequirementsFileFlag, "requirements",
				ServerFlag, "server",
				VerboseFlag,
				VersionFlag,
				"collection-name",
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

func TestAnsibleGalaxyCollectionInstallCmdString(t *testing.T) {
	tests := []struct {
		desc    string
		cmd     *AnsibleGalaxyCollectionInstallCmd
		command string
	}{

		{
			desc: "Testing generate a command for AnsibleGalaxyCollectionInstallCmd with all flags using default binary",
			cmd: NewAnsibleGalaxyCollectionInstallCmd(
				WithCollectionNames("collection-name"),
				WithGalaxyCollectionInstallOptions(&AnsibleGalaxyCollectionInstallOptions{
					ClearResponseCache:          true,
					DisableGPGVerify:            true,
					ForceWithDeps:               true,
					IgnoreSignatureStatusCode:   true,
					IgnoreSignatureStatusCodes:  "ignore_codes",
					Keyring:                     "keyring",
					NoCache:                     true,
					Offline:                     true,
					Pre:                         true,
					RequiredValidSignatureCount: 1,
					Signature:                   "signature",
					Timeout:                     "10",
					Token:                       "token",
					APIKey:                      "apikey",
					Upgrade:                     true,
					IgnoreCerts:                 true,
					Force:                       true,
					IgnoreErrors:                true,
					NoDeps:                      true,
					CollectionsPath:             "path",
					RequirementsFile:            "requirements",
					Server:                      "server",
					Verbose:                     true,
					Version:                     true,
				}),
			),
			command: "ansible-galaxy collection install --api-key apikey --clear-response-cache --disable-gpg-verify --force-with-deps --ignore-signature-status-code --ignore-signature-status-codes ignore_codes --keyring keyring --no-cache --offline --pre --required-valid-signature-count 1 --signature signature --timeout 10 --token token --upgrade --ignore-certs --force --ignore-errors --no-deps --collections-path path --requirements-file requirements --server server --verbose --version collection-name",
		},
		{
			desc: "Testing generate a command for AnsibleGalaxyCollectionInstallCmd with all flags",
			cmd: NewAnsibleGalaxyCollectionInstallCmd(
				WithBinary("ansible-galaxy-binary"),
				WithCollectionNames("collection-name"),
				WithGalaxyCollectionInstallOptions(&AnsibleGalaxyCollectionInstallOptions{
					ClearResponseCache:          true,
					DisableGPGVerify:            true,
					ForceWithDeps:               true,
					IgnoreSignatureStatusCode:   true,
					IgnoreSignatureStatusCodes:  "ignore_codes",
					Keyring:                     "keyring",
					NoCache:                     true,
					Offline:                     true,
					Pre:                         true,
					RequiredValidSignatureCount: 1,
					Signature:                   "signature",
					Timeout:                     "10",
					Token:                       "token",
					APIKey:                      "apikey",
					Upgrade:                     true,
					IgnoreCerts:                 true,
					Force:                       true,
					IgnoreErrors:                true,
					NoDeps:                      true,
					CollectionsPath:             "path",
					RequirementsFile:            "requirements",
					Server:                      "server",
					Verbose:                     true,
					Version:                     true,
				}),
			),
			command: "ansible-galaxy-binary collection install --api-key apikey --clear-response-cache --disable-gpg-verify --force-with-deps --ignore-signature-status-code --ignore-signature-status-codes ignore_codes --keyring keyring --no-cache --offline --pre --required-valid-signature-count 1 --signature signature --timeout 10 --token token --upgrade --ignore-certs --force --ignore-errors --no-deps --collections-path path --requirements-file requirements --server server --verbose --version collection-name",
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
