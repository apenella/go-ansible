package galaxycollectioninstall

import (
	"testing"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestAnsibleGalaxyCollectionInstallOptionsGenerateCommandOptions(t *testing.T) {

	errContext := "(galaxy::AnsibleGalaxyCollectionInstallOptions::GenerateCommandOptions)"

	tests := []struct {
		desc    string
		options *AnsibleGalaxyCollectionInstallOptions
		err     error
		expect  []string
	}{
		{
			desc:    "Testing nil AnsibleGalaxyCollectionInstallOptions definition",
			options: nil,
			err:     errors.New(errContext, "AnsibleGalaxyCollectionInstallOptions is nil"),
			expect:  []string{},
		},
		{
			desc:    "Testing an empty AnsibleGalaxyCollectionInstallOptions definition",
			options: &AnsibleGalaxyCollectionInstallOptions{},
			err:     nil,
			expect:  []string{},
		},
		{
			desc: "Testing AnsibleGalaxyCollectionInstallOptions with all flags",
			options: &AnsibleGalaxyCollectionInstallOptions{
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
			},
			err: nil,
			expect: []string{
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

func TestAnsibleGalaxyCollectionInstallOptionsString(t *testing.T) {
	tests := []struct {
		desc    string
		options *AnsibleGalaxyCollectionInstallOptions
		expect  string
	}{
		{
			desc:    "Testing generate string from an empty AnsibleGalaxyCollectionInstallOptions",
			options: &AnsibleGalaxyCollectionInstallOptions{},
			expect:  "",
		},
		{
			desc: "Testing generate string from an AnsibleGalaxyCollectionInstallOptions with all flags",
			options: &AnsibleGalaxyCollectionInstallOptions{
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
			},
			expect: "--api-key apikey --clear-response-cache --disable-gpg-verify --force-with-deps --ignore-signature-status-code --ignore-signature-status-codes ignore_codes --keyring keyring --no-cache --offline --pre --required-valid-signature-count 1 --signature signature --timeout 10 --token token --upgrade --ignore-certs --force --ignore-errors --no-deps --collections-path path --requirements-file requirements --server server --verbose --version",
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			res := test.options.String()

			assert.Equal(t, test.expect, res, "Unexpected options value for AnsibleGalaxyCollectionInstallOptions.String()")

		})
	}
}
