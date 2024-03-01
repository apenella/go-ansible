package configuration

import (
	"fmt"
	"testing"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/stretchr/testify/assert"
)

func TestWithExecutor(t *testing.T) {
	executor := execute.NewDefaultExecute()
	e := NewAnsibleWithConfigurationSettingsExecute(executor)

	assert.Equal(t, executor, e.executor)
}

// TestWithAnsibleActionWarnings tests the method that sets ANSIBLE_ACTION_WARNINGS to true
func TestWithAnsibleActionWarnings(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleActionWarnings(),
	)
	setting := exec.configurationSettings[AnsibleActionWarnings]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleActionWarnings tests the method that sets ANSIBLE_ACTION_WARNINGS to false
func TestWithoutAnsibleActionWarnings(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleActionWarnings(),
	)
	setting := exec.configurationSettings[AnsibleActionWarnings]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleAgnosticBecomePrompt tests the method that sets ANSIBLE_AGNOSTIC_BECOME_PROMPT to true
func TestWithAnsibleAgnosticBecomePrompt(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleAgnosticBecomePrompt(),
	)
	setting := exec.configurationSettings[AnsibleAgnosticBecomePrompt]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleAgnosticBecomePrompt tests the method that sets ANSIBLE_AGNOSTIC_BECOME_PROMPT to false
func TestWithoutAnsibleAgnosticBecomePrompt(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleAgnosticBecomePrompt(),
	)
	setting := exec.configurationSettings[AnsibleAgnosticBecomePrompt]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleConnectionPath tests the method that sets the value for ANSIBLE_CONNECTION_PATH
func TestWithAnsibleConnectionPath(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleConnectionPath(value),
	)
	setting := exec.configurationSettings[AnsibleConnectionPath]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCowAcceptlist tests the method that sets the value for ANSIBLE_COW_ACCEPTLIST
func TestWithAnsibleCowAcceptlist(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCowAcceptlist(value),
	)
	setting := exec.configurationSettings[AnsibleCowAcceptlist]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCowPath tests the method that sets the value for ANSIBLE_COW_PATH
func TestWithAnsibleCowPath(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCowPath(value),
	)
	setting := exec.configurationSettings[AnsibleCowPath]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCowSelection tests the method that sets the value for ANSIBLE_COW_SELECTION
func TestWithAnsibleCowSelection(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCowSelection(value),
	)
	setting := exec.configurationSettings[AnsibleCowSelection]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleForceColor tests the method that sets ANSIBLE_FORCE_COLOR to true
func TestWithAnsibleForceColor(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleForceColor(),
	)
	setting := exec.configurationSettings[AnsibleForceColor]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleForceColor tests the method that sets ANSIBLE_FORCE_COLOR to false
func TestWithoutAnsibleForceColor(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleForceColor(),
	)
	setting := exec.configurationSettings[AnsibleForceColor]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleHome tests the method that sets the value for ANSIBLE_HOME
func TestWithAnsibleHome(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleHome(value),
	)
	setting := exec.configurationSettings[AnsibleHome]
	assert.Equal(t, setting, value)
}

// TestWithNoColor tests the method that sets NO_COLOR to true
func TestWithNoColor(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithNoColor(),
	)
	setting := exec.configurationSettings[NoColor]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutNoColor tests the method that sets NO_COLOR to false
func TestWithoutNoColor(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutNoColor(),
	)
	setting := exec.configurationSettings[NoColor]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleNocows tests the method that sets ANSIBLE_NOCOWS to true
func TestWithAnsibleNocows(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleNocows(),
	)
	setting := exec.configurationSettings[AnsibleNocows]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleNocows tests the method that sets ANSIBLE_NOCOWS to false
func TestWithoutAnsibleNocows(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleNocows(),
	)
	setting := exec.configurationSettings[AnsibleNocows]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsiblePipelining tests the method that sets ANSIBLE_PIPELINING to true
func TestWithAnsiblePipelining(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePipelining(),
	)
	setting := exec.configurationSettings[AnsiblePipelining]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsiblePipelining tests the method that sets ANSIBLE_PIPELINING to false
func TestWithoutAnsiblePipelining(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsiblePipelining(),
	)
	setting := exec.configurationSettings[AnsiblePipelining]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleAnyErrorsFatal tests the method that sets ANSIBLE_ANY_ERRORS_FATAL to true
func TestWithAnsibleAnyErrorsFatal(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleAnyErrorsFatal(),
	)
	setting := exec.configurationSettings[AnsibleAnyErrorsFatal]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleAnyErrorsFatal tests the method that sets ANSIBLE_ANY_ERRORS_FATAL to false
func TestWithoutAnsibleAnyErrorsFatal(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleAnyErrorsFatal(),
	)
	setting := exec.configurationSettings[AnsibleAnyErrorsFatal]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleBecomeAllowSameUser tests the method that sets ANSIBLE_BECOME_ALLOW_SAME_USER to true
func TestWithAnsibleBecomeAllowSameUser(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleBecomeAllowSameUser(),
	)
	setting := exec.configurationSettings[AnsibleBecomeAllowSameUser]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleBecomeAllowSameUser tests the method that sets ANSIBLE_BECOME_ALLOW_SAME_USER to false
func TestWithoutAnsibleBecomeAllowSameUser(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleBecomeAllowSameUser(),
	)
	setting := exec.configurationSettings[AnsibleBecomeAllowSameUser]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleBecomePasswordFile tests the method that sets the value for ANSIBLE_BECOME_PASSWORD_FILE
func TestWithAnsibleBecomePasswordFile(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleBecomePasswordFile(value),
	)
	setting := exec.configurationSettings[AnsibleBecomePasswordFile]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleBecomePlugins tests the method that sets the value for ANSIBLE_BECOME_PLUGINS
func TestWithAnsibleBecomePlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleBecomePlugins(value),
	)
	setting := exec.configurationSettings[AnsibleBecomePlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCachePlugin tests the method that sets the value for ANSIBLE_CACHE_PLUGIN
func TestWithAnsibleCachePlugin(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCachePlugin(value),
	)
	setting := exec.configurationSettings[AnsibleCachePlugin]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCachePluginConnection tests the method that sets the value for ANSIBLE_CACHE_PLUGIN_CONNECTION
func TestWithAnsibleCachePluginConnection(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCachePluginConnection(value),
	)
	setting := exec.configurationSettings[AnsibleCachePluginConnection]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCachePluginPrefix tests the method that sets the value for ANSIBLE_CACHE_PLUGIN_PREFIX
func TestWithAnsibleCachePluginPrefix(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCachePluginPrefix(value),
	)
	setting := exec.configurationSettings[AnsibleCachePluginPrefix]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCachePluginTimeout tests the method that sets the value for ANSIBLE_CACHE_PLUGIN_TIMEOUT
func TestWithAnsibleCachePluginTimeout(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCachePluginTimeout(value),
	)
	setting := exec.configurationSettings[AnsibleCachePluginTimeout]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsibleCallbacksEnabled tests the method that sets the value for ANSIBLE_CALLBACKS_ENABLED
func TestWithAnsibleCallbacksEnabled(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCallbacksEnabled(value),
	)
	setting := exec.configurationSettings[AnsibleCallbacksEnabled]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCollectionsOnAnsibleVersionMismatch tests the method that sets the value for ANSIBLE_COLLECTIONS_ON_ANSIBLE_VERSION_MISMATCH
func TestWithAnsibleCollectionsOnAnsibleVersionMismatch(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCollectionsOnAnsibleVersionMismatch(value),
	)
	setting := exec.configurationSettings[AnsibleCollectionsOnAnsibleVersionMismatch]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCollectionsPaths tests the method that sets the value for ANSIBLE_COLLECTIONS_PATHS
func TestWithAnsibleCollectionsPaths(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCollectionsPaths(value),
	)
	setting := exec.configurationSettings[AnsibleCollectionsPaths]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCollectionsScanSysPath tests the method that sets ANSIBLE_COLLECTIONS_SCAN_SYS_PATH to true
func TestWithAnsibleCollectionsScanSysPath(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCollectionsScanSysPath(),
	)
	setting := exec.configurationSettings[AnsibleCollectionsScanSysPath]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleCollectionsScanSysPath tests the method that sets ANSIBLE_COLLECTIONS_SCAN_SYS_PATH to false
func TestWithoutAnsibleCollectionsScanSysPath(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleCollectionsScanSysPath(),
	)
	setting := exec.configurationSettings[AnsibleCollectionsScanSysPath]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleColorChanged tests the method that sets the value for ANSIBLE_COLOR_CHANGED
func TestWithAnsibleColorChanged(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorChanged(value),
	)
	setting := exec.configurationSettings[AnsibleColorChanged]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorConsolePrompt tests the method that sets the value for ANSIBLE_COLOR_CONSOLE_PROMPT
func TestWithAnsibleColorConsolePrompt(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorConsolePrompt(value),
	)
	setting := exec.configurationSettings[AnsibleColorConsolePrompt]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorDebug tests the method that sets the value for ANSIBLE_COLOR_DEBUG
func TestWithAnsibleColorDebug(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorDebug(value),
	)
	setting := exec.configurationSettings[AnsibleColorDebug]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorDeprecate tests the method that sets the value for ANSIBLE_COLOR_DEPRECATE
func TestWithAnsibleColorDeprecate(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorDeprecate(value),
	)
	setting := exec.configurationSettings[AnsibleColorDeprecate]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorDiffAdd tests the method that sets the value for ANSIBLE_COLOR_DIFF_ADD
func TestWithAnsibleColorDiffAdd(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorDiffAdd(value),
	)
	setting := exec.configurationSettings[AnsibleColorDiffAdd]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorDiffLines tests the method that sets the value for ANSIBLE_COLOR_DIFF_LINES
func TestWithAnsibleColorDiffLines(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorDiffLines(value),
	)
	setting := exec.configurationSettings[AnsibleColorDiffLines]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorDiffRemove tests the method that sets the value for ANSIBLE_COLOR_DIFF_REMOVE
func TestWithAnsibleColorDiffRemove(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorDiffRemove(value),
	)
	setting := exec.configurationSettings[AnsibleColorDiffRemove]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorError tests the method that sets the value for ANSIBLE_COLOR_ERROR
func TestWithAnsibleColorError(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorError(value),
	)
	setting := exec.configurationSettings[AnsibleColorError]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorHighlight tests the method that sets the value for ANSIBLE_COLOR_HIGHLIGHT
func TestWithAnsibleColorHighlight(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorHighlight(value),
	)
	setting := exec.configurationSettings[AnsibleColorHighlight]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorOk tests the method that sets the value for ANSIBLE_COLOR_OK
func TestWithAnsibleColorOk(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorOk(value),
	)
	setting := exec.configurationSettings[AnsibleColorOk]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorSkip tests the method that sets the value for ANSIBLE_COLOR_SKIP
func TestWithAnsibleColorSkip(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorSkip(value),
	)
	setting := exec.configurationSettings[AnsibleColorSkip]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorUnreachable tests the method that sets the value for ANSIBLE_COLOR_UNREACHABLE
func TestWithAnsibleColorUnreachable(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorUnreachable(value),
	)
	setting := exec.configurationSettings[AnsibleColorUnreachable]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorVerbose tests the method that sets the value for ANSIBLE_COLOR_VERBOSE
func TestWithAnsibleColorVerbose(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorVerbose(value),
	)
	setting := exec.configurationSettings[AnsibleColorVerbose]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleColorWarn tests the method that sets the value for ANSIBLE_COLOR_WARN
func TestWithAnsibleColorWarn(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleColorWarn(value),
	)
	setting := exec.configurationSettings[AnsibleColorWarn]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleConnectionPasswordFile tests the method that sets the value for ANSIBLE_CONNECTION_PASSWORD_FILE
func TestWithAnsibleConnectionPasswordFile(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleConnectionPasswordFile(value),
	)
	setting := exec.configurationSettings[AnsibleConnectionPasswordFile]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCoverageRemoteOutput tests the method that sets the value for _ANSIBLE_COVERAGE_REMOTE_OUTPUT
func TestWithAnsibleCoverageRemoteOutput(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCoverageRemoteOutput(value),
	)
	setting := exec.configurationSettings[AnsibleCoverageRemoteOutput]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCoverageRemotePathFilter tests the method that sets the value for _ANSIBLE_COVERAGE_REMOTE_PATH_FILTER
func TestWithAnsibleCoverageRemotePathFilter(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCoverageRemotePathFilter(value),
	)
	setting := exec.configurationSettings[AnsibleCoverageRemotePathFilter]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleActionPlugins tests the method that sets the value for ANSIBLE_ACTION_PLUGINS
func TestWithAnsibleActionPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleActionPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleActionPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleAskPass tests the method that sets ANSIBLE_ASK_PASS to true
func TestWithAnsibleAskPass(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleAskPass(),
	)
	setting := exec.configurationSettings[AnsibleAskPass]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleAskPass tests the method that sets ANSIBLE_ASK_PASS to false
func TestWithoutAnsibleAskPass(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleAskPass(),
	)
	setting := exec.configurationSettings[AnsibleAskPass]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleAskVaultPass tests the method that sets ANSIBLE_ASK_VAULT_PASS to true
func TestWithAnsibleAskVaultPass(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleAskVaultPass(),
	)
	setting := exec.configurationSettings[AnsibleAskVaultPass]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleAskVaultPass tests the method that sets ANSIBLE_ASK_VAULT_PASS to false
func TestWithoutAnsibleAskVaultPass(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleAskVaultPass(),
	)
	setting := exec.configurationSettings[AnsibleAskVaultPass]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleBecome tests the method that sets ANSIBLE_BECOME to true
func TestWithAnsibleBecome(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleBecome(),
	)
	setting := exec.configurationSettings[AnsibleBecome]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleBecome tests the method that sets ANSIBLE_BECOME to false
func TestWithoutAnsibleBecome(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleBecome(),
	)
	setting := exec.configurationSettings[AnsibleBecome]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleBecomeAskPass tests the method that sets ANSIBLE_BECOME_ASK_PASS to true
func TestWithAnsibleBecomeAskPass(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleBecomeAskPass(),
	)
	setting := exec.configurationSettings[AnsibleBecomeAskPass]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleBecomeAskPass tests the method that sets ANSIBLE_BECOME_ASK_PASS to false
func TestWithoutAnsibleBecomeAskPass(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleBecomeAskPass(),
	)
	setting := exec.configurationSettings[AnsibleBecomeAskPass]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleBecomeExe tests the method that sets the value for ANSIBLE_BECOME_EXE
func TestWithAnsibleBecomeExe(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleBecomeExe(value),
	)
	setting := exec.configurationSettings[AnsibleBecomeExe]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleBecomeFlags tests the method that sets the value for ANSIBLE_BECOME_FLAGS
func TestWithAnsibleBecomeFlags(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleBecomeFlags(value),
	)
	setting := exec.configurationSettings[AnsibleBecomeFlags]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleBecomeMethod tests the method that sets the value for ANSIBLE_BECOME_METHOD
func TestWithAnsibleBecomeMethod(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleBecomeMethod(value),
	)
	setting := exec.configurationSettings[AnsibleBecomeMethod]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleBecomeUser tests the method that sets the value for ANSIBLE_BECOME_USER
func TestWithAnsibleBecomeUser(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleBecomeUser(value),
	)
	setting := exec.configurationSettings[AnsibleBecomeUser]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCachePlugins tests the method that sets the value for ANSIBLE_CACHE_PLUGINS
func TestWithAnsibleCachePlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCachePlugins(value),
	)
	setting := exec.configurationSettings[AnsibleCachePlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCallbackPlugins tests the method that sets the value for ANSIBLE_CALLBACK_PLUGINS
func TestWithAnsibleCallbackPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCallbackPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleCallbackPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleCliconfPlugins tests the method that sets the value for ANSIBLE_CLICONF_PLUGINS
func TestWithAnsibleCliconfPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleCliconfPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleCliconfPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleConnectionPlugins tests the method that sets the value for ANSIBLE_CONNECTION_PLUGINS
func TestWithAnsibleConnectionPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleConnectionPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleConnectionPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleDebug tests the method that sets ANSIBLE_DEBUG to true
func TestWithAnsibleDebug(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleDebug(),
	)
	setting := exec.configurationSettings[AnsibleDebug]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleDebug tests the method that sets ANSIBLE_DEBUG to false
func TestWithoutAnsibleDebug(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleDebug(),
	)
	setting := exec.configurationSettings[AnsibleDebug]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleExecutable tests the method that sets the value for ANSIBLE_EXECUTABLE
func TestWithAnsibleExecutable(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleExecutable(value),
	)
	setting := exec.configurationSettings[AnsibleExecutable]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleFactPath tests the method that sets the value for ANSIBLE_FACT_PATH
func TestWithAnsibleFactPath(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleFactPath(value),
	)
	setting := exec.configurationSettings[AnsibleFactPath]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleFilterPlugins tests the method that sets the value for ANSIBLE_FILTER_PLUGINS
func TestWithAnsibleFilterPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleFilterPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleFilterPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleForceHandlers tests the method that sets ANSIBLE_FORCE_HANDLERS to true
func TestWithAnsibleForceHandlers(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleForceHandlers(),
	)
	setting := exec.configurationSettings[AnsibleForceHandlers]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleForceHandlers tests the method that sets ANSIBLE_FORCE_HANDLERS to false
func TestWithoutAnsibleForceHandlers(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleForceHandlers(),
	)
	setting := exec.configurationSettings[AnsibleForceHandlers]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleForks tests the method that sets the value for ANSIBLE_FORKS
func TestWithAnsibleForks(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleForks(value),
	)
	setting := exec.configurationSettings[AnsibleForks]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsibleGatherSubset tests the method that sets the value for ANSIBLE_GATHER_SUBSET
func TestWithAnsibleGatherSubset(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGatherSubset(value),
	)
	setting := exec.configurationSettings[AnsibleGatherSubset]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGatherTimeout tests the method that sets the value for ANSIBLE_GATHER_TIMEOUT
func TestWithAnsibleGatherTimeout(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGatherTimeout(value),
	)
	setting := exec.configurationSettings[AnsibleGatherTimeout]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsibleGathering tests the method that sets the value for ANSIBLE_GATHERING
func TestWithAnsibleGathering(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGathering(value),
	)
	setting := exec.configurationSettings[AnsibleGathering]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleHashBehaviour tests the method that sets the value for ANSIBLE_HASH_BEHAVIOUR
func TestWithAnsibleHashBehaviour(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleHashBehaviour(value),
	)
	setting := exec.configurationSettings[AnsibleHashBehaviour]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventory tests the method that sets the value for ANSIBLE_INVENTORY
func TestWithAnsibleInventory(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventory(value),
	)
	setting := exec.configurationSettings[AnsibleInventory]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleHttpapiPlugins tests the method that sets the value for ANSIBLE_HTTPAPI_PLUGINS
func TestWithAnsibleHttpapiPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleHttpapiPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleHttpapiPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventoryPlugins tests the method that sets the value for ANSIBLE_INVENTORY_PLUGINS
func TestWithAnsibleInventoryPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleInventoryPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleJinja2Extensions tests the method that sets the value for ANSIBLE_JINJA2_EXTENSIONS
func TestWithAnsibleJinja2Extensions(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleJinja2Extensions(value),
	)
	setting := exec.configurationSettings[AnsibleJinja2Extensions]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleJinja2Native tests the method that sets ANSIBLE_JINJA2_NATIVE to true
func TestWithAnsibleJinja2Native(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleJinja2Native(),
	)
	setting := exec.configurationSettings[AnsibleJinja2Native]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleJinja2Native tests the method that sets ANSIBLE_JINJA2_NATIVE to false
func TestWithoutAnsibleJinja2Native(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleJinja2Native(),
	)
	setting := exec.configurationSettings[AnsibleJinja2Native]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleKeepRemoteFiles tests the method that sets ANSIBLE_KEEP_REMOTE_FILES to true
func TestWithAnsibleKeepRemoteFiles(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleKeepRemoteFiles(),
	)
	setting := exec.configurationSettings[AnsibleKeepRemoteFiles]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleKeepRemoteFiles tests the method that sets ANSIBLE_KEEP_REMOTE_FILES to false
func TestWithoutAnsibleKeepRemoteFiles(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleKeepRemoteFiles(),
	)
	setting := exec.configurationSettings[AnsibleKeepRemoteFiles]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleLibvirtLxcNoseclabel tests the method that sets ANSIBLE_LIBVIRT_LXC_NOSECLABEL to true
func TestWithAnsibleLibvirtLxcNoseclabel(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleLibvirtLxcNoseclabel(),
	)
	setting := exec.configurationSettings[AnsibleLibvirtLxcNoseclabel]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleLibvirtLxcNoseclabel tests the method that sets ANSIBLE_LIBVIRT_LXC_NOSECLABEL to false
func TestWithoutAnsibleLibvirtLxcNoseclabel(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleLibvirtLxcNoseclabel(),
	)
	setting := exec.configurationSettings[AnsibleLibvirtLxcNoseclabel]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleLoadCallbackPlugins tests the method that sets ANSIBLE_LOAD_CALLBACK_PLUGINS to true
func TestWithAnsibleLoadCallbackPlugins(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleLoadCallbackPlugins(),
	)
	setting := exec.configurationSettings[AnsibleLoadCallbackPlugins]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleLoadCallbackPlugins tests the method that sets ANSIBLE_LOAD_CALLBACK_PLUGINS to false
func TestWithoutAnsibleLoadCallbackPlugins(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleLoadCallbackPlugins(),
	)
	setting := exec.configurationSettings[AnsibleLoadCallbackPlugins]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleLocalTemp tests the method that sets the value for ANSIBLE_LOCAL_TEMP
func TestWithAnsibleLocalTemp(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleLocalTemp(value),
	)
	setting := exec.configurationSettings[AnsibleLocalTemp]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleLogFilter tests the method that sets the value for ANSIBLE_LOG_FILTER
func TestWithAnsibleLogFilter(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleLogFilter(value),
	)
	setting := exec.configurationSettings[AnsibleLogFilter]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleLogPath tests the method that sets the value for ANSIBLE_LOG_PATH
func TestWithAnsibleLogPath(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleLogPath(value),
	)
	setting := exec.configurationSettings[AnsibleLogPath]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleLookupPlugins tests the method that sets the value for ANSIBLE_LOOKUP_PLUGINS
func TestWithAnsibleLookupPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleLookupPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleLookupPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleModuleArgs tests the method that sets the value for ANSIBLE_MODULE_ARGS
func TestWithAnsibleModuleArgs(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleModuleArgs(value),
	)
	setting := exec.configurationSettings[AnsibleModuleArgs]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleLibrary tests the method that sets the value for ANSIBLE_LIBRARY
func TestWithAnsibleLibrary(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleLibrary(value),
	)
	setting := exec.configurationSettings[AnsibleLibrary]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleModuleUtils tests the method that sets the value for ANSIBLE_MODULE_UTILS
func TestWithAnsibleModuleUtils(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleModuleUtils(value),
	)
	setting := exec.configurationSettings[AnsibleModuleUtils]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleNetconfPlugins tests the method that sets the value for ANSIBLE_NETCONF_PLUGINS
func TestWithAnsibleNetconfPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleNetconfPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleNetconfPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleNoLog tests the method that sets ANSIBLE_NO_LOG to true
func TestWithAnsibleNoLog(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleNoLog(),
	)
	setting := exec.configurationSettings[AnsibleNoLog]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleNoLog tests the method that sets ANSIBLE_NO_LOG to false
func TestWithoutAnsibleNoLog(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleNoLog(),
	)
	setting := exec.configurationSettings[AnsibleNoLog]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleNoTargetSyslog tests the method that sets ANSIBLE_NO_TARGET_SYSLOG to true
func TestWithAnsibleNoTargetSyslog(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleNoTargetSyslog(),
	)
	setting := exec.configurationSettings[AnsibleNoTargetSyslog]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleNoTargetSyslog tests the method that sets ANSIBLE_NO_TARGET_SYSLOG to false
func TestWithoutAnsibleNoTargetSyslog(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleNoTargetSyslog(),
	)
	setting := exec.configurationSettings[AnsibleNoTargetSyslog]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleNullRepresentation tests the method that sets the value for ANSIBLE_NULL_REPRESENTATION
func TestWithAnsibleNullRepresentation(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleNullRepresentation(value),
	)
	setting := exec.configurationSettings[AnsibleNullRepresentation]
	assert.Equal(t, setting, value)
}

// TestWithAnsiblePollInterval tests the method that sets the value for ANSIBLE_POLL_INTERVAL
func TestWithAnsiblePollInterval(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePollInterval(value),
	)
	setting := exec.configurationSettings[AnsiblePollInterval]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsiblePrivateKeyFile tests the method that sets the value for ANSIBLE_PRIVATE_KEY_FILE
func TestWithAnsiblePrivateKeyFile(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePrivateKeyFile(value),
	)
	setting := exec.configurationSettings[AnsiblePrivateKeyFile]
	assert.Equal(t, setting, value)
}

// TestWithAnsiblePrivateRoleVars tests the method that sets ANSIBLE_PRIVATE_ROLE_VARS to true
func TestWithAnsiblePrivateRoleVars(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePrivateRoleVars(),
	)
	setting := exec.configurationSettings[AnsiblePrivateRoleVars]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsiblePrivateRoleVars tests the method that sets ANSIBLE_PRIVATE_ROLE_VARS to false
func TestWithoutAnsiblePrivateRoleVars(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsiblePrivateRoleVars(),
	)
	setting := exec.configurationSettings[AnsiblePrivateRoleVars]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleRemotePort tests the method that sets the value for ANSIBLE_REMOTE_PORT
func TestWithAnsibleRemotePort(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleRemotePort(value),
	)
	setting := exec.configurationSettings[AnsibleRemotePort]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsibleRemoteUser tests the method that sets the value for ANSIBLE_REMOTE_USER
func TestWithAnsibleRemoteUser(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleRemoteUser(value),
	)
	setting := exec.configurationSettings[AnsibleRemoteUser]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleRolesPath tests the method that sets the value for ANSIBLE_ROLES_PATH
func TestWithAnsibleRolesPath(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleRolesPath(value),
	)
	setting := exec.configurationSettings[AnsibleRolesPath]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleSelinuxSpecialFs tests the method that sets the value for ANSIBLE_SELINUX_SPECIAL_FS
func TestWithAnsibleSelinuxSpecialFs(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleSelinuxSpecialFs(value),
	)
	setting := exec.configurationSettings[AnsibleSelinuxSpecialFs]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleStdoutCallback tests the method that sets the value for ANSIBLE_STDOUT_CALLBACK
func TestWithAnsibleStdoutCallback(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleStdoutCallback(value),
	)
	setting := exec.configurationSettings[AnsibleStdoutCallback]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleStrategy tests the method that sets the value for ANSIBLE_STRATEGY
func TestWithAnsibleStrategy(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleStrategy(value),
	)
	setting := exec.configurationSettings[AnsibleStrategy]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleStrategyPlugins tests the method that sets the value for ANSIBLE_STRATEGY_PLUGINS
func TestWithAnsibleStrategyPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleStrategyPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleStrategyPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleSu tests the method that sets ANSIBLE_SU to true
func TestWithAnsibleSu(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleSu(),
	)
	setting := exec.configurationSettings[AnsibleSu]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleSu tests the method that sets ANSIBLE_SU to false
func TestWithoutAnsibleSu(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleSu(),
	)
	setting := exec.configurationSettings[AnsibleSu]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleSyslogFacility tests the method that sets the value for ANSIBLE_SYSLOG_FACILITY
func TestWithAnsibleSyslogFacility(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleSyslogFacility(value),
	)
	setting := exec.configurationSettings[AnsibleSyslogFacility]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleTerminalPlugins tests the method that sets the value for ANSIBLE_TERMINAL_PLUGINS
func TestWithAnsibleTerminalPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleTerminalPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleTerminalPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleTestPlugins tests the method that sets the value for ANSIBLE_TEST_PLUGINS
func TestWithAnsibleTestPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleTestPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleTestPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleTimeout tests the method that sets the value for ANSIBLE_TIMEOUT
func TestWithAnsibleTimeout(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleTimeout(value),
	)
	setting := exec.configurationSettings[AnsibleTimeout]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsibleTransport tests the method that sets the value for ANSIBLE_TRANSPORT
func TestWithAnsibleTransport(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleTransport(value),
	)
	setting := exec.configurationSettings[AnsibleTransport]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleErrorOnUndefinedVars tests the method that sets ANSIBLE_ERROR_ON_UNDEFINED_VARS to true
func TestWithAnsibleErrorOnUndefinedVars(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleErrorOnUndefinedVars(),
	)
	setting := exec.configurationSettings[AnsibleErrorOnUndefinedVars]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleErrorOnUndefinedVars tests the method that sets ANSIBLE_ERROR_ON_UNDEFINED_VARS to false
func TestWithoutAnsibleErrorOnUndefinedVars(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleErrorOnUndefinedVars(),
	)
	setting := exec.configurationSettings[AnsibleErrorOnUndefinedVars]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleVarsPlugins tests the method that sets the value for ANSIBLE_VARS_PLUGINS
func TestWithAnsibleVarsPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleVarsPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleVarsPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleVaultEncryptIdentity tests the method that sets the value for ANSIBLE_VAULT_ENCRYPT_IDENTITY
func TestWithAnsibleVaultEncryptIdentity(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleVaultEncryptIdentity(value),
	)
	setting := exec.configurationSettings[AnsibleVaultEncryptIdentity]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleVaultIdMatch tests the method that sets the value for ANSIBLE_VAULT_ID_MATCH
func TestWithAnsibleVaultIdMatch(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleVaultIdMatch(value),
	)
	setting := exec.configurationSettings[AnsibleVaultIdMatch]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleVaultIdentity tests the method that sets the value for ANSIBLE_VAULT_IDENTITY
func TestWithAnsibleVaultIdentity(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleVaultIdentity(value),
	)
	setting := exec.configurationSettings[AnsibleVaultIdentity]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleVaultIdentityList tests the method that sets the value for ANSIBLE_VAULT_IDENTITY_LIST
func TestWithAnsibleVaultIdentityList(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleVaultIdentityList(value),
	)
	setting := exec.configurationSettings[AnsibleVaultIdentityList]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleVaultPasswordFile tests the method that sets the value for ANSIBLE_VAULT_PASSWORD_FILE
func TestWithAnsibleVaultPasswordFile(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleVaultPasswordFile(value),
	)
	setting := exec.configurationSettings[AnsibleVaultPasswordFile]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleVerbosity tests the method that sets the value for ANSIBLE_VERBOSITY
func TestWithAnsibleVerbosity(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleVerbosity(value),
	)
	setting := exec.configurationSettings[AnsibleVerbosity]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsibleDeprecationWarnings tests the method that sets ANSIBLE_DEPRECATION_WARNINGS to true
func TestWithAnsibleDeprecationWarnings(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleDeprecationWarnings(),
	)
	setting := exec.configurationSettings[AnsibleDeprecationWarnings]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleDeprecationWarnings tests the method that sets ANSIBLE_DEPRECATION_WARNINGS to false
func TestWithoutAnsibleDeprecationWarnings(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleDeprecationWarnings(),
	)
	setting := exec.configurationSettings[AnsibleDeprecationWarnings]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleDevelWarning tests the method that sets ANSIBLE_DEVEL_WARNING to true
func TestWithAnsibleDevelWarning(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleDevelWarning(),
	)
	setting := exec.configurationSettings[AnsibleDevelWarning]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleDevelWarning tests the method that sets ANSIBLE_DEVEL_WARNING to false
func TestWithoutAnsibleDevelWarning(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleDevelWarning(),
	)
	setting := exec.configurationSettings[AnsibleDevelWarning]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleDiffAlways tests the method that sets the value for ANSIBLE_DIFF_ALWAYS
func TestWithAnsibleDiffAlways(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleDiffAlways(value),
	)
	setting := exec.configurationSettings[AnsibleDiffAlways]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleDiffContext tests the method that sets the value for ANSIBLE_DIFF_CONTEXT
func TestWithAnsibleDiffContext(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleDiffContext(value),
	)
	setting := exec.configurationSettings[AnsibleDiffContext]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsibleDisplayArgsToStdout tests the method that sets ANSIBLE_DISPLAY_ARGS_TO_STDOUT to true
func TestWithAnsibleDisplayArgsToStdout(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleDisplayArgsToStdout(),
	)
	setting := exec.configurationSettings[AnsibleDisplayArgsToStdout]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleDisplayArgsToStdout tests the method that sets ANSIBLE_DISPLAY_ARGS_TO_STDOUT to false
func TestWithoutAnsibleDisplayArgsToStdout(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleDisplayArgsToStdout(),
	)
	setting := exec.configurationSettings[AnsibleDisplayArgsToStdout]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleDisplaySkippedHosts tests the method that sets ANSIBLE_DISPLAY_SKIPPED_HOSTS to true
func TestWithAnsibleDisplaySkippedHosts(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleDisplaySkippedHosts(),
	)
	setting := exec.configurationSettings[AnsibleDisplaySkippedHosts]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleDisplaySkippedHosts tests the method that sets ANSIBLE_DISPLAY_SKIPPED_HOSTS to false
func TestWithoutAnsibleDisplaySkippedHosts(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleDisplaySkippedHosts(),
	)
	setting := exec.configurationSettings[AnsibleDisplaySkippedHosts]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleDocFragmentPlugins tests the method that sets the value for ANSIBLE_DOC_FRAGMENT_PLUGINS
func TestWithAnsibleDocFragmentPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleDocFragmentPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleDocFragmentPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleDuplicateYamlDictKey tests the method that sets the value for ANSIBLE_DUPLICATE_YAML_DICT_KEY
func TestWithAnsibleDuplicateYamlDictKey(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleDuplicateYamlDictKey(value),
	)
	setting := exec.configurationSettings[AnsibleDuplicateYamlDictKey]
	assert.Equal(t, setting, value)
}

// TestWithEditor tests the method that sets the value for EDITOR
func TestWithEditor(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithEditor(value),
	)
	setting := exec.configurationSettings[Editor]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleEnableTaskDebugger tests the method that sets ANSIBLE_ENABLE_TASK_DEBUGGER to true
func TestWithAnsibleEnableTaskDebugger(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleEnableTaskDebugger(),
	)
	setting := exec.configurationSettings[AnsibleEnableTaskDebugger]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleEnableTaskDebugger tests the method that sets ANSIBLE_ENABLE_TASK_DEBUGGER to false
func TestWithoutAnsibleEnableTaskDebugger(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleEnableTaskDebugger(),
	)
	setting := exec.configurationSettings[AnsibleEnableTaskDebugger]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleErrorOnMissingHandler tests the method that sets ANSIBLE_ERROR_ON_MISSING_HANDLER to true
func TestWithAnsibleErrorOnMissingHandler(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleErrorOnMissingHandler(),
	)
	setting := exec.configurationSettings[AnsibleErrorOnMissingHandler]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleErrorOnMissingHandler tests the method that sets ANSIBLE_ERROR_ON_MISSING_HANDLER to false
func TestWithoutAnsibleErrorOnMissingHandler(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleErrorOnMissingHandler(),
	)
	setting := exec.configurationSettings[AnsibleErrorOnMissingHandler]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleFactsModules tests the method that sets the value for ANSIBLE_FACTS_MODULES
func TestWithAnsibleFactsModules(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleFactsModules(value),
	)
	setting := exec.configurationSettings[AnsibleFactsModules]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyCacheDir tests the method that sets the value for ANSIBLE_GALAXY_CACHE_DIR
func TestWithAnsibleGalaxyCacheDir(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyCacheDir(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyCacheDir]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyCollectionSkeleton tests the method that sets the value for ANSIBLE_GALAXY_COLLECTION_SKELETON
func TestWithAnsibleGalaxyCollectionSkeleton(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyCollectionSkeleton(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyCollectionSkeleton]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyCollectionSkeletonIgnore tests the method that sets the value for ANSIBLE_GALAXY_COLLECTION_SKELETON_IGNORE
func TestWithAnsibleGalaxyCollectionSkeletonIgnore(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyCollectionSkeletonIgnore(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyCollectionSkeletonIgnore]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyCollectionsPathWarning tests the method that sets the value for ANSIBLE_GALAXY_COLLECTIONS_PATH_WARNING
func TestWithAnsibleGalaxyCollectionsPathWarning(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyCollectionsPathWarning(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyCollectionsPathWarning]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyDisableGpgVerify tests the method that sets the value for ANSIBLE_GALAXY_DISABLE_GPG_VERIFY
func TestWithAnsibleGalaxyDisableGpgVerify(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyDisableGpgVerify(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyDisableGpgVerify]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyDisplayProgress tests the method that sets the value for ANSIBLE_GALAXY_DISPLAY_PROGRESS
func TestWithAnsibleGalaxyDisplayProgress(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyDisplayProgress(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyDisplayProgress]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyGpgKeyring tests the method that sets the value for ANSIBLE_GALAXY_GPG_KEYRING
func TestWithAnsibleGalaxyGpgKeyring(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyGpgKeyring(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyGpgKeyring]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyIgnore tests the method that sets ANSIBLE_GALAXY_IGNORE to true
func TestWithAnsibleGalaxyIgnore(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyIgnore(),
	)
	setting := exec.configurationSettings[AnsibleGalaxyIgnore]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleGalaxyIgnore tests the method that sets ANSIBLE_GALAXY_IGNORE to false
func TestWithoutAnsibleGalaxyIgnore(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleGalaxyIgnore(),
	)
	setting := exec.configurationSettings[AnsibleGalaxyIgnore]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleGalaxyIgnoreSignatureStatusCodes tests the method that sets the value for ANSIBLE_GALAXY_IGNORE_SIGNATURE_STATUS_CODES
func TestWithAnsibleGalaxyIgnoreSignatureStatusCodes(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyIgnoreSignatureStatusCodes(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyIgnoreSignatureStatusCodes]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyRequiredValidSignatureCount tests the method that sets the value for ANSIBLE_GALAXY_REQUIRED_VALID_SIGNATURE_COUNT
func TestWithAnsibleGalaxyRequiredValidSignatureCount(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyRequiredValidSignatureCount(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyRequiredValidSignatureCount]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyRoleSkeleton tests the method that sets the value for ANSIBLE_GALAXY_ROLE_SKELETON
func TestWithAnsibleGalaxyRoleSkeleton(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyRoleSkeleton(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyRoleSkeleton]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyRoleSkeletonIgnore tests the method that sets the value for ANSIBLE_GALAXY_ROLE_SKELETON_IGNORE
func TestWithAnsibleGalaxyRoleSkeletonIgnore(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyRoleSkeletonIgnore(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyRoleSkeletonIgnore]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyServer tests the method that sets the value for ANSIBLE_GALAXY_SERVER
func TestWithAnsibleGalaxyServer(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyServer(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyServer]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyServerList tests the method that sets the value for ANSIBLE_GALAXY_SERVER_LIST
func TestWithAnsibleGalaxyServerList(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyServerList(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyServerList]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyServerTimeout tests the method that sets the value for ANSIBLE_GALAXY_SERVER_TIMEOUT
func TestWithAnsibleGalaxyServerTimeout(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyServerTimeout(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyServerTimeout]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleGalaxyTokenPath tests the method that sets the value for ANSIBLE_GALAXY_TOKEN_PATH
func TestWithAnsibleGalaxyTokenPath(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleGalaxyTokenPath(value),
	)
	setting := exec.configurationSettings[AnsibleGalaxyTokenPath]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleHostKeyChecking tests the method that sets ANSIBLE_HOST_KEY_CHECKING to true
func TestWithAnsibleHostKeyChecking(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleHostKeyChecking(),
	)
	setting := exec.configurationSettings[AnsibleHostKeyChecking]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleHostKeyChecking tests the method that sets ANSIBLE_HOST_KEY_CHECKING to false
func TestWithoutAnsibleHostKeyChecking(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleHostKeyChecking(),
	)
	setting := exec.configurationSettings[AnsibleHostKeyChecking]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleHostPatternMismatch tests the method that sets the value for ANSIBLE_HOST_PATTERN_MISMATCH
func TestWithAnsibleHostPatternMismatch(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleHostPatternMismatch(value),
	)
	setting := exec.configurationSettings[AnsibleHostPatternMismatch]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInjectFactVars tests the method that sets ANSIBLE_INJECT_FACT_VARS to true
func TestWithAnsibleInjectFactVars(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInjectFactVars(),
	)
	setting := exec.configurationSettings[AnsibleInjectFactVars]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleInjectFactVars tests the method that sets ANSIBLE_INJECT_FACT_VARS to false
func TestWithoutAnsibleInjectFactVars(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleInjectFactVars(),
	)
	setting := exec.configurationSettings[AnsibleInjectFactVars]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsiblePythonInterpreter tests the method that sets the value for ANSIBLE_PYTHON_INTERPRETER
func TestWithAnsiblePythonInterpreter(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePythonInterpreter(value),
	)
	setting := exec.configurationSettings[AnsiblePythonInterpreter]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInvalidTaskAttributeFailed tests the method that sets ANSIBLE_INVALID_TASK_ATTRIBUTE_FAILED to true
func TestWithAnsibleInvalidTaskAttributeFailed(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInvalidTaskAttributeFailed(),
	)
	setting := exec.configurationSettings[AnsibleInvalidTaskAttributeFailed]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleInvalidTaskAttributeFailed tests the method that sets ANSIBLE_INVALID_TASK_ATTRIBUTE_FAILED to false
func TestWithoutAnsibleInvalidTaskAttributeFailed(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleInvalidTaskAttributeFailed(),
	)
	setting := exec.configurationSettings[AnsibleInvalidTaskAttributeFailed]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleInventoryAnyUnparsedIsFailed tests the method that sets ANSIBLE_INVENTORY_ANY_UNPARSED_IS_FAILED to true
func TestWithAnsibleInventoryAnyUnparsedIsFailed(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryAnyUnparsedIsFailed(),
	)
	setting := exec.configurationSettings[AnsibleInventoryAnyUnparsedIsFailed]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleInventoryAnyUnparsedIsFailed tests the method that sets ANSIBLE_INVENTORY_ANY_UNPARSED_IS_FAILED to false
func TestWithoutAnsibleInventoryAnyUnparsedIsFailed(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleInventoryAnyUnparsedIsFailed(),
	)
	setting := exec.configurationSettings[AnsibleInventoryAnyUnparsedIsFailed]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleInventoryCache tests the method that sets the value for ANSIBLE_INVENTORY_CACHE
func TestWithAnsibleInventoryCache(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryCache(value),
	)
	setting := exec.configurationSettings[AnsibleInventoryCache]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventoryCachePlugin tests the method that sets the value for ANSIBLE_INVENTORY_CACHE_PLUGIN
func TestWithAnsibleInventoryCachePlugin(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryCachePlugin(value),
	)
	setting := exec.configurationSettings[AnsibleInventoryCachePlugin]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventoryCacheConnection tests the method that sets the value for ANSIBLE_INVENTORY_CACHE_CONNECTION
func TestWithAnsibleInventoryCacheConnection(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryCacheConnection(value),
	)
	setting := exec.configurationSettings[AnsibleInventoryCacheConnection]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventoryCachePluginPrefix tests the method that sets the value for ANSIBLE_INVENTORY_CACHE_PLUGIN_PREFIX
func TestWithAnsibleInventoryCachePluginPrefix(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryCachePluginPrefix(value),
	)
	setting := exec.configurationSettings[AnsibleInventoryCachePluginPrefix]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventoryCacheTimeout tests the method that sets the value for ANSIBLE_INVENTORY_CACHE_TIMEOUT
func TestWithAnsibleInventoryCacheTimeout(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryCacheTimeout(value),
	)
	setting := exec.configurationSettings[AnsibleInventoryCacheTimeout]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventoryEnabled tests the method that sets the value for ANSIBLE_INVENTORY_ENABLED
func TestWithAnsibleInventoryEnabled(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryEnabled(value),
	)
	setting := exec.configurationSettings[AnsibleInventoryEnabled]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventoryExport tests the method that sets the value for ANSIBLE_INVENTORY_EXPORT
func TestWithAnsibleInventoryExport(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryExport(value),
	)
	setting := exec.configurationSettings[AnsibleInventoryExport]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventoryIgnore tests the method that sets the value for ANSIBLE_INVENTORY_IGNORE
func TestWithAnsibleInventoryIgnore(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryIgnore(value),
	)
	setting := exec.configurationSettings[AnsibleInventoryIgnore]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventoryIgnoreRegex tests the method that sets the value for ANSIBLE_INVENTORY_IGNORE_REGEX
func TestWithAnsibleInventoryIgnoreRegex(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryIgnoreRegex(value),
	)
	setting := exec.configurationSettings[AnsibleInventoryIgnoreRegex]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventoryUnparsedFailed tests the method that sets the value for ANSIBLE_INVENTORY_UNPARSED_FAILED
func TestWithAnsibleInventoryUnparsedFailed(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryUnparsedFailed(value),
	)
	setting := exec.configurationSettings[AnsibleInventoryUnparsedFailed]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleInventoryUnparsedWarning tests the method that sets ANSIBLE_INVENTORY_UNPARSED_WARNING to true
func TestWithAnsibleInventoryUnparsedWarning(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleInventoryUnparsedWarning(),
	)
	setting := exec.configurationSettings[AnsibleInventoryUnparsedWarning]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleInventoryUnparsedWarning tests the method that sets ANSIBLE_INVENTORY_UNPARSED_WARNING to false
func TestWithoutAnsibleInventoryUnparsedWarning(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleInventoryUnparsedWarning(),
	)
	setting := exec.configurationSettings[AnsibleInventoryUnparsedWarning]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleJinja2NativeWarning tests the method that sets ANSIBLE_JINJA2_NATIVE_WARNING to true
func TestWithAnsibleJinja2NativeWarning(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleJinja2NativeWarning(),
	)
	setting := exec.configurationSettings[AnsibleJinja2NativeWarning]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleJinja2NativeWarning tests the method that sets ANSIBLE_JINJA2_NATIVE_WARNING to false
func TestWithoutAnsibleJinja2NativeWarning(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleJinja2NativeWarning(),
	)
	setting := exec.configurationSettings[AnsibleJinja2NativeWarning]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleLocalhostWarning tests the method that sets ANSIBLE_LOCALHOST_WARNING to true
func TestWithAnsibleLocalhostWarning(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleLocalhostWarning(),
	)
	setting := exec.configurationSettings[AnsibleLocalhostWarning]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleLocalhostWarning tests the method that sets ANSIBLE_LOCALHOST_WARNING to false
func TestWithoutAnsibleLocalhostWarning(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleLocalhostWarning(),
	)
	setting := exec.configurationSettings[AnsibleLocalhostWarning]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleMaxDiffSize tests the method that sets the value for ANSIBLE_MAX_DIFF_SIZE
func TestWithAnsibleMaxDiffSize(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleMaxDiffSize(value),
	)
	setting := exec.configurationSettings[AnsibleMaxDiffSize]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleModuleIgnoreExts tests the method that sets the value for ANSIBLE_MODULE_IGNORE_EXTS
func TestWithAnsibleModuleIgnoreExts(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleModuleIgnoreExts(value),
	)
	setting := exec.configurationSettings[AnsibleModuleIgnoreExts]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleModuleStrictUtf8Response tests the method that sets the value for ANSIBLE_MODULE_STRICT_UTF8_RESPONSE
func TestWithAnsibleModuleStrictUtf8Response(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleModuleStrictUtf8Response(value),
	)
	setting := exec.configurationSettings[AnsibleModuleStrictUtf8Response]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleNetconfSshConfig tests the method that sets the value for ANSIBLE_NETCONF_SSH_CONFIG
func TestWithAnsibleNetconfSshConfig(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleNetconfSshConfig(value),
	)
	setting := exec.configurationSettings[AnsibleNetconfSshConfig]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleNetworkGroupModules tests the method that sets the value for ANSIBLE_NETWORK_GROUP_MODULES
func TestWithAnsibleNetworkGroupModules(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleNetworkGroupModules(value),
	)
	setting := exec.configurationSettings[AnsibleNetworkGroupModules]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleOldPluginCacheClear tests the method that sets ANSIBLE_OLD_PLUGIN_CACHE_CLEAR to true
func TestWithAnsibleOldPluginCacheClear(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleOldPluginCacheClear(),
	)
	setting := exec.configurationSettings[AnsibleOldPluginCacheClear]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleOldPluginCacheClear tests the method that sets ANSIBLE_OLD_PLUGIN_CACHE_CLEAR to false
func TestWithoutAnsibleOldPluginCacheClear(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleOldPluginCacheClear(),
	)
	setting := exec.configurationSettings[AnsibleOldPluginCacheClear]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithPager tests the method that sets the value for PAGER
func TestWithPager(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithPager(value),
	)
	setting := exec.configurationSettings[Pager]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleParamikoHostKeyAutoAdd tests the method that sets ANSIBLE_PARAMIKO_HOST_KEY_AUTO_ADD to true
func TestWithAnsibleParamikoHostKeyAutoAdd(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleParamikoHostKeyAutoAdd(),
	)
	setting := exec.configurationSettings[AnsibleParamikoHostKeyAutoAdd]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleParamikoHostKeyAutoAdd tests the method that sets ANSIBLE_PARAMIKO_HOST_KEY_AUTO_ADD to false
func TestWithoutAnsibleParamikoHostKeyAutoAdd(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleParamikoHostKeyAutoAdd(),
	)
	setting := exec.configurationSettings[AnsibleParamikoHostKeyAutoAdd]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleParamikoLookForKeys tests the method that sets ANSIBLE_PARAMIKO_LOOK_FOR_KEYS to true
func TestWithAnsibleParamikoLookForKeys(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleParamikoLookForKeys(),
	)
	setting := exec.configurationSettings[AnsibleParamikoLookForKeys]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleParamikoLookForKeys tests the method that sets ANSIBLE_PARAMIKO_LOOK_FOR_KEYS to false
func TestWithoutAnsibleParamikoLookForKeys(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleParamikoLookForKeys(),
	)
	setting := exec.configurationSettings[AnsibleParamikoLookForKeys]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsiblePersistentCommandTimeout tests the method that sets the value for ANSIBLE_PERSISTENT_COMMAND_TIMEOUT
func TestWithAnsiblePersistentCommandTimeout(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePersistentCommandTimeout(value),
	)
	setting := exec.configurationSettings[AnsiblePersistentCommandTimeout]
	assert.Equal(t, setting, value)
}

// TestWithAnsiblePersistentConnectRetryTimeout tests the method that sets the value for ANSIBLE_PERSISTENT_CONNECT_RETRY_TIMEOUT
func TestWithAnsiblePersistentConnectRetryTimeout(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePersistentConnectRetryTimeout(value),
	)
	setting := exec.configurationSettings[AnsiblePersistentConnectRetryTimeout]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsiblePersistentConnectTimeout tests the method that sets the value for ANSIBLE_PERSISTENT_CONNECT_TIMEOUT
func TestWithAnsiblePersistentConnectTimeout(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePersistentConnectTimeout(value),
	)
	setting := exec.configurationSettings[AnsiblePersistentConnectTimeout]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsiblePersistentControlPathDir tests the method that sets the value for ANSIBLE_PERSISTENT_CONTROL_PATH_DIR
func TestWithAnsiblePersistentControlPathDir(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePersistentControlPathDir(value),
	)
	setting := exec.configurationSettings[AnsiblePersistentControlPathDir]
	assert.Equal(t, setting, value)
}

// TestWithAnsiblePlaybookDir tests the method that sets the value for ANSIBLE_PLAYBOOK_DIR
func TestWithAnsiblePlaybookDir(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePlaybookDir(value),
	)
	setting := exec.configurationSettings[AnsiblePlaybookDir]
	assert.Equal(t, setting, value)
}

// TestWithAnsiblePlaybookVarsRoot tests the method that sets the value for ANSIBLE_PLAYBOOK_VARS_ROOT
func TestWithAnsiblePlaybookVarsRoot(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePlaybookVarsRoot(value),
	)
	setting := exec.configurationSettings[AnsiblePlaybookVarsRoot]
	assert.Equal(t, setting, value)
}

// TestWithAnsiblePythonModuleRlimitNofile tests the method that sets the value for ANSIBLE_PYTHON_MODULE_RLIMIT_NOFILE
func TestWithAnsiblePythonModuleRlimitNofile(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePythonModuleRlimitNofile(value),
	)
	setting := exec.configurationSettings[AnsiblePythonModuleRlimitNofile]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleRetryFilesEnabled tests the method that sets the value for ANSIBLE_RETRY_FILES_ENABLED
func TestWithAnsibleRetryFilesEnabled(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleRetryFilesEnabled(value),
	)
	setting := exec.configurationSettings[AnsibleRetryFilesEnabled]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleRetryFilesSavePath tests the method that sets the value for ANSIBLE_RETRY_FILES_SAVE_PATH
func TestWithAnsibleRetryFilesSavePath(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleRetryFilesSavePath(value),
	)
	setting := exec.configurationSettings[AnsibleRetryFilesSavePath]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleRunVarsPlugins tests the method that sets the value for ANSIBLE_RUN_VARS_PLUGINS
func TestWithAnsibleRunVarsPlugins(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleRunVarsPlugins(value),
	)
	setting := exec.configurationSettings[AnsibleRunVarsPlugins]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleShowCustomStats tests the method that sets the value for ANSIBLE_SHOW_CUSTOM_STATS
func TestWithAnsibleShowCustomStats(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleShowCustomStats(value),
	)
	setting := exec.configurationSettings[AnsibleShowCustomStats]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleStringConversionAction tests the method that sets the value for ANSIBLE_STRING_CONVERSION_ACTION
func TestWithAnsibleStringConversionAction(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleStringConversionAction(value),
	)
	setting := exec.configurationSettings[AnsibleStringConversionAction]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleStringTypeFilters tests the method that sets the value for ANSIBLE_STRING_TYPE_FILTERS
func TestWithAnsibleStringTypeFilters(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleStringTypeFilters(value),
	)
	setting := exec.configurationSettings[AnsibleStringTypeFilters]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleSystemWarnings tests the method that sets ANSIBLE_SYSTEM_WARNINGS to true
func TestWithAnsibleSystemWarnings(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleSystemWarnings(),
	)
	setting := exec.configurationSettings[AnsibleSystemWarnings]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleSystemWarnings tests the method that sets ANSIBLE_SYSTEM_WARNINGS to false
func TestWithoutAnsibleSystemWarnings(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleSystemWarnings(),
	)
	setting := exec.configurationSettings[AnsibleSystemWarnings]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleRunTags tests the method that sets the value for ANSIBLE_RUN_TAGS
func TestWithAnsibleRunTags(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleRunTags(value),
	)
	setting := exec.configurationSettings[AnsibleRunTags]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleSkipTags tests the method that sets the value for ANSIBLE_SKIP_TAGS
func TestWithAnsibleSkipTags(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleSkipTags(value),
	)
	setting := exec.configurationSettings[AnsibleSkipTags]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleTaskDebuggerIgnoreErrors tests the method that sets ANSIBLE_TASK_DEBUGGER_IGNORE_ERRORS to true
func TestWithAnsibleTaskDebuggerIgnoreErrors(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleTaskDebuggerIgnoreErrors(),
	)
	setting := exec.configurationSettings[AnsibleTaskDebuggerIgnoreErrors]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleTaskDebuggerIgnoreErrors tests the method that sets ANSIBLE_TASK_DEBUGGER_IGNORE_ERRORS to false
func TestWithoutAnsibleTaskDebuggerIgnoreErrors(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleTaskDebuggerIgnoreErrors(),
	)
	setting := exec.configurationSettings[AnsibleTaskDebuggerIgnoreErrors]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleTaskTimeout tests the method that sets the value for ANSIBLE_TASK_TIMEOUT
func TestWithAnsibleTaskTimeout(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleTaskTimeout(value),
	)
	setting := exec.configurationSettings[AnsibleTaskTimeout]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsibleTransformInvalidGroupChars tests the method that sets the value for ANSIBLE_TRANSFORM_INVALID_GROUP_CHARS
func TestWithAnsibleTransformInvalidGroupChars(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleTransformInvalidGroupChars(value),
	)
	setting := exec.configurationSettings[AnsibleTransformInvalidGroupChars]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleUsePersistentConnections tests the method that sets ANSIBLE_USE_PERSISTENT_CONNECTIONS to true
func TestWithAnsibleUsePersistentConnections(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleUsePersistentConnections(),
	)
	setting := exec.configurationSettings[AnsibleUsePersistentConnections]
	expected := "true"
	assert.Equal(t, setting, expected)
}

// TestWithoutAnsibleUsePersistentConnections tests the method that sets ANSIBLE_USE_PERSISTENT_CONNECTIONS to false
func TestWithoutAnsibleUsePersistentConnections(t *testing.T) {
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithoutAnsibleUsePersistentConnections(),
	)
	setting := exec.configurationSettings[AnsibleUsePersistentConnections]
	expected := "false"
	assert.Equal(t, setting, expected)
}

// TestWithAnsibleValidateActionGroupMetadata tests the method that sets the value for ANSIBLE_VALIDATE_ACTION_GROUP_METADATA
func TestWithAnsibleValidateActionGroupMetadata(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleValidateActionGroupMetadata(value),
	)
	setting := exec.configurationSettings[AnsibleValidateActionGroupMetadata]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleVarsEnabled tests the method that sets the value for ANSIBLE_VARS_ENABLED
func TestWithAnsibleVarsEnabled(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleVarsEnabled(value),
	)
	setting := exec.configurationSettings[AnsibleVarsEnabled]
	assert.Equal(t, setting, value)
}

// TestWithAnsiblePrecedence tests the method that sets the value for ANSIBLE_PRECEDENCE
func TestWithAnsiblePrecedence(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsiblePrecedence(value),
	)
	setting := exec.configurationSettings[AnsiblePrecedence]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleVaultEncryptSalt tests the method that sets the value for ANSIBLE_VAULT_ENCRYPT_SALT
func TestWithAnsibleVaultEncryptSalt(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleVaultEncryptSalt(value),
	)
	setting := exec.configurationSettings[AnsibleVaultEncryptSalt]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleVerboseToStderr tests the method that sets the value for ANSIBLE_VERBOSE_TO_STDERR
func TestWithAnsibleVerboseToStderr(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleVerboseToStderr(value),
	)
	setting := exec.configurationSettings[AnsibleVerboseToStderr]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleWinAsyncStartupTimeout tests the method that sets the value for ANSIBLE_WIN_ASYNC_STARTUP_TIMEOUT
func TestWithAnsibleWinAsyncStartupTimeout(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleWinAsyncStartupTimeout(value),
	)
	setting := exec.configurationSettings[AnsibleWinAsyncStartupTimeout]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsibleWorkerShutdownPollCount tests the method that sets the value for ANSIBLE_WORKER_SHUTDOWN_POLL_COUNT
func TestWithAnsibleWorkerShutdownPollCount(t *testing.T) {
	value := 10
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleWorkerShutdownPollCount(value),
	)
	setting := exec.configurationSettings[AnsibleWorkerShutdownPollCount]
	assert.Equal(t, setting, fmt.Sprint(value))
}

// TestWithAnsibleWorkerShutdownPollDelay tests the method that sets the value for ANSIBLE_WORKER_SHUTDOWN_POLL_DELAY
func TestWithAnsibleWorkerShutdownPollDelay(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleWorkerShutdownPollDelay(value),
	)
	setting := exec.configurationSettings[AnsibleWorkerShutdownPollDelay]
	assert.Equal(t, setting, value)
}

// TestWithAnsibleYamlFilenameExt tests the method that sets the value for ANSIBLE_YAML_FILENAME_EXT
func TestWithAnsibleYamlFilenameExt(t *testing.T) {
	value := "testvalue"
	exec := NewAnsibleWithConfigurationSettingsExecute(nil,
		WithAnsibleYamlFilenameExt(value),
	)
	setting := exec.configurationSettings[AnsibleYamlFilenameExt]
	assert.Equal(t, setting, value)
}
