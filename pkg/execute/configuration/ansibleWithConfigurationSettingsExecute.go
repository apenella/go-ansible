package configuration

import (
	"context"
	"fmt"
)

type configurationSettings map[string]string

// AnsibleWithConfigurationSettingsExecute is a builder for Ansible Cmd
type AnsibleWithConfigurationSettingsExecute struct {
	executor              ExecutorEnvVarSetter
	configurationSettings configurationSettings
}

// ithAnsibleWithConfigurationSettingsExecute return a new AnsibleWithConfigurationSettingsExecute
func NewAnsibleWithConfigurationSettingsExecute(executor ExecutorEnvVarSetter) *AnsibleWithConfigurationSettingsExecute {
	return &AnsibleWithConfigurationSettingsExecute{
		executor:              executor,
		configurationSettings: make(configurationSettings),
	}
}

func (e *AnsibleWithConfigurationSettingsExecute) WithExecutor(exec ExecutorEnvVarSetter) *AnsibleWithConfigurationSettingsExecute {
	e.executor = exec
	return e
}

func (e *AnsibleWithConfigurationSettingsExecute) Execute(ctx context.Context) error {
	if e.executor == nil {
		return fmt.Errorf("AnsibleWithConfigurationSettingsExecute executor requires an executor")
	}

	for key, value := range e.configurationSettings {
		e.executor.AddEnvVar(key, value)
	}

	err := e.executor.Execute(ctx)
	if err != nil {
		return fmt.Errorf("error executing command: %s", err.Error())
	}

	return nil
}

// WithAnsibleActionWarnings sets the option ANSIBLE_ACTION_WARNINGS to true (By default Ansible will issue a warning when received from a task action (module or action plugin) These warnings can be silenced by adjusting this setting to False.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleActionWarnings() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleActionWarnings] = "true"
	return e
}

// WithoutAnsibleActionWarnings sets the option ANSIBLE_ACTION_WARNINGS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleActionWarnings() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleActionWarnings] = "false"
	return e
}

// WithAnsibleAgnosticBecomePrompt sets the option ANSIBLE_AGNOSTIC_BECOME_PROMPT to true (Display an agnostic become prompt instead of displaying a prompt containing the command line supplied become method)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleAgnosticBecomePrompt() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleAgnosticBecomePrompt] = "true"
	return e
}

// WithoutAnsibleAgnosticBecomePrompt sets the option ANSIBLE_AGNOSTIC_BECOME_PROMPT to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleAgnosticBecomePrompt() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleAgnosticBecomePrompt] = "false"
	return e
}

// WithAnsibleConnectionPath sets the value for the configuraion ANSIBLE_CONNECTION_PATH (Specify where to look for the ansible-connection script. This location will be checked before searching $PATH. If null, ansible will start with the same directory as the ansible script.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleConnectionPath(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleConnectionPath] = value
	return e
}

// WithAnsibleCowAcceptlist sets the value for the configuraion ANSIBLE_COW_ACCEPTLIST (Accept list of cowsay templates that are ‘safe’ to use, set to empty list if you want to enable all installed templates. [:Version Added: 2.11])
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCowAcceptlist(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCowAcceptlist] = value
	return e
}

// WithAnsibleCowPath sets the value for the configuraion ANSIBLE_COW_PATH (Specify a custom cowsay path or swap in your cowsay implementation of choice)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCowPath(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCowPath] = value
	return e
}

// WithAnsibleCowSelection sets the value for the configuraion ANSIBLE_COW_SELECTION (This allows you to chose a specific cowsay stencil for the banners or use ‘random’ to cycle through them.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCowSelection(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCowSelection] = value
	return e
}

// WithAnsibleForceColor sets the option ANSIBLE_FORCE_COLOR to true (This option forces color mode even when running without a TTY or the “nocolor” setting is True.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleForceColor() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleForceColor] = "true"
	return e
}

// WithoutAnsibleForceColor sets the option ANSIBLE_FORCE_COLOR to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleForceColor() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleForceColor] = "false"
	return e
}

// WithAnsibleHome sets the value for the configuraion ANSIBLE_HOME (The default root path for Ansible config files on the controller.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleHome(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleHome] = value
	return e
}

// WithNoColor sets the option NO_COLOR to true (This setting allows suppressing colorizing output, which is used to give a better indication of failure and status information.)
func (e *AnsibleWithConfigurationSettingsExecute) WithNoColor() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[NoColor] = "true"
	return e
}

// WithoutNoColor sets the option NO_COLOR to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutNoColor() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[NoColor] = "false"
	return e
}

// WithAnsibleNocows sets the option ANSIBLE_NOCOWS to true (If you have cowsay installed but want to avoid the ‘cows’ (why????), use this.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleNocows() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleNocows] = "true"
	return e
}

// WithoutAnsibleNocows sets the option ANSIBLE_NOCOWS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleNocows() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleNocows] = "false"
	return e
}

// WithAnsiblePipelining sets the option ANSIBLE_PIPELINING to true (This is a global option, each connection plugin can override either by having more specific options or not supporting pipelining at all. Pipelining, if supported by the connection plugin, reduces the number of network operations required to execute a module on the remote server, by executing many Ansible modules without actual file transfer. It can result in a very significant performance improvement when enabled. However this conflicts with privilege escalation (become). For example, when using ‘sudo:’ operations you must first disable ‘requiretty’ in /etc/sudoers on all managed hosts, which is why it is disabled by default. This setting will be disabled if ANSIBLE_KEEP_REMOTE_FILES is enabled.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePipelining() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePipelining] = "true"
	return e
}

// WithoutAnsiblePipelining sets the option ANSIBLE_PIPELINING to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsiblePipelining() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePipelining] = "false"
	return e
}

// WithAnsibleAnyErrorsFatal sets the option ANSIBLE_ANY_ERRORS_FATAL to true (Sets the default value for the any_errors_fatal keyword, if True, Task failures will be considered fatal errors.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleAnyErrorsFatal() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleAnyErrorsFatal] = "true"
	return e
}

// WithoutAnsibleAnyErrorsFatal sets the option ANSIBLE_ANY_ERRORS_FATAL to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleAnyErrorsFatal() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleAnyErrorsFatal] = "false"
	return e
}

// WithAnsibleBecomeAllowSameUser sets the option ANSIBLE_BECOME_ALLOW_SAME_USER to true (This setting controls if become is skipped when remote user and become user are the same. I.E root sudo to root. If executable, it will be run and the resulting stdout will be used as the password.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleBecomeAllowSameUser() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecomeAllowSameUser] = "true"
	return e
}

// WithoutAnsibleBecomeAllowSameUser sets the option ANSIBLE_BECOME_ALLOW_SAME_USER to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleBecomeAllowSameUser() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecomeAllowSameUser] = "false"
	return e
}

// WithAnsibleBecomePasswordFile sets the value for the configuraion ANSIBLE_BECOME_PASSWORD_FILE (The password file to use for the become plugin. –become-password-file. If executable, it will be run and the resulting stdout will be used as the password.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleBecomePasswordFile(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecomePasswordFile] = value
	return e
}

// WithAnsibleBecomePlugins sets the value for the configuraion ANSIBLE_BECOME_PLUGINS (Colon separated paths in which Ansible will search for Become Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleBecomePlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecomePlugins] = value
	return e
}

// WithAnsibleCachePlugin sets the value for the configuraion ANSIBLE_CACHE_PLUGIN (Chooses which cache plugin to use, the default ‘memory’ is ephemeral.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCachePlugin(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCachePlugin] = value
	return e
}

// WithAnsibleCachePluginConnection sets the value for the configuraion ANSIBLE_CACHE_PLUGIN_CONNECTION (Defines connection or path information for the cache plugin)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCachePluginConnection(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCachePluginConnection] = value
	return e
}

// WithAnsibleCachePluginPrefix sets the value for the configuraion ANSIBLE_CACHE_PLUGIN_PREFIX (Prefix to use for cache plugin files/tables)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCachePluginPrefix(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCachePluginPrefix] = value
	return e
}

// WithAnsibleCachePluginTimeout sets the value for the configuraion ANSIBLE_CACHE_PLUGIN_TIMEOUT (Expiration timeout for the cache plugin data)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCachePluginTimeout(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCachePluginTimeout] = fmt.Sprint(value)
	return e
}

// WithAnsibleCallbacksEnabled sets the value for the configuraion ANSIBLE_CALLBACKS_ENABLED (List of enabled callbacks, not all callbacks need enabling, but many of those shipped with Ansible do as we don’t want them activated by default. [:Version Added: 2.11])
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCallbacksEnabled(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCallbacksEnabled] = value
	return e
}

// WithAnsibleCollectionsOnAnsibleVersionMismatch sets the value for the configuraion ANSIBLE_COLLECTIONS_ON_ANSIBLE_VERSION_MISMATCH (When a collection is loaded that does not support the running Ansible version (with the collection metadata key requires_ansible).)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCollectionsOnAnsibleVersionMismatch(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCollectionsOnAnsibleVersionMismatch] = value
	return e
}

// WithAnsibleCollectionsPaths sets the value for the configuraion ANSIBLE_COLLECTIONS_PATHS (Colon separated paths in which Ansible will search for collections content. Collections must be in nested subdirectories, not directly in these directories. For example, if COLLECTIONS_PATHS includes '{{ ANSIBLE_HOME ~ "/collections" }}', and you want to add my.collection to that directory, it must be saved as '{{ ANSIBLE_HOME} ~ "/collections/ansible_collections/my/collection" }}'.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCollectionsPaths(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCollectionsPaths] = value
	return e
}

// WithAnsibleCollectionsScanSysPath sets the option ANSIBLE_COLLECTIONS_SCAN_SYS_PATH to true (A boolean to enable or disable scanning the sys.path for installed collections)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCollectionsScanSysPath() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCollectionsScanSysPath] = "true"
	return e
}

// WithoutAnsibleCollectionsScanSysPath sets the option ANSIBLE_COLLECTIONS_SCAN_SYS_PATH to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleCollectionsScanSysPath() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCollectionsScanSysPath] = "false"
	return e
}

// WithAnsibleColorChanged sets the value for the configuraion ANSIBLE_COLOR_CHANGED (Defines the color to use on ‘Changed’ task status)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorChanged(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorChanged] = value
	return e
}

// WithAnsibleColorConsolePrompt sets the value for the configuraion ANSIBLE_COLOR_CONSOLE_PROMPT (Defines the default color to use for ansible-console)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorConsolePrompt(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorConsolePrompt] = value
	return e
}

// WithAnsibleColorDebug sets the value for the configuraion ANSIBLE_COLOR_DEBUG (Defines the color to use when emitting debug messages)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorDebug(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorDebug] = value
	return e
}

// WithAnsibleColorDeprecate sets the value for the configuraion ANSIBLE_COLOR_DEPRECATE (Defines the color to use when emitting deprecation messages)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorDeprecate(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorDeprecate] = value
	return e
}

// WithAnsibleColorDiffAdd sets the value for the configuraion ANSIBLE_COLOR_DIFF_ADD (Defines the color to use when showing added lines in diffs)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorDiffAdd(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorDiffAdd] = value
	return e
}

// WithAnsibleColorDiffLines sets the value for the configuraion ANSIBLE_COLOR_DIFF_LINES (Defines the color to use when showing diffs)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorDiffLines(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorDiffLines] = value
	return e
}

// WithAnsibleColorDiffRemove sets the value for the configuraion ANSIBLE_COLOR_DIFF_REMOVE (Defines the color to use when showing removed lines in diffs)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorDiffRemove(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorDiffRemove] = value
	return e
}

// WithAnsibleColorError sets the value for the configuraion ANSIBLE_COLOR_ERROR (Defines the color to use when emitting error messages)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorError(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorError] = value
	return e
}

// WithAnsibleColorHighlight sets the value for the configuraion ANSIBLE_COLOR_HIGHLIGHT (Defines the color to use for highlighting)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorHighlight(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorHighlight] = value
	return e
}

// WithAnsibleColorOk sets the value for the configuraion ANSIBLE_COLOR_OK (Defines the color to use when showing ‘OK’ task status)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorOk(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorOk] = value
	return e
}

// WithAnsibleColorSkip sets the value for the configuraion ANSIBLE_COLOR_SKIP (Defines the color to use when showing ‘Skipped’ task status)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorSkip(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorSkip] = value
	return e
}

// WithAnsibleColorUnreachable sets the value for the configuraion ANSIBLE_COLOR_UNREACHABLE (Defines the color to use on ‘Unreachable’ status)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorUnreachable(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorUnreachable] = value
	return e
}

// WithAnsibleColorVerbose sets the value for the configuraion ANSIBLE_COLOR_VERBOSE (Defines the color to use when emitting verbose messages. i.e those that show with ‘-v’s.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorVerbose(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorVerbose] = value
	return e
}

// WithAnsibleColorWarn sets the value for the configuraion ANSIBLE_COLOR_WARN (Defines the color to use when emitting warning messages)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleColorWarn(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleColorWarn] = value
	return e
}

// WithAnsibleConnectionPasswordFile sets the value for the configuraion ANSIBLE_CONNECTION_PASSWORD_FILE (The password file to use for the connection plugin. –connection-password-file.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleConnectionPasswordFile(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleConnectionPasswordFile] = value
	return e
}

// WithAnsibleCoverageRemoteOutput sets the value for the configuraion _ANSIBLE_COVERAGE_REMOTE_OUTPUT (Sets the output directory on the remote host to generate coverage reports to. Currently only used for remote coverage on PowerShell modules. This is for internal use only.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCoverageRemoteOutput(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCoverageRemoteOutput] = value
	return e
}

// WithAnsibleCoverageRemotePathFilter sets the value for the configuraion _ANSIBLE_COVERAGE_REMOTE_PATH_FILTER (A list of paths for files on the Ansible controller to run coverage for when executing on the remote host. Only files that match the path glob will have its coverage collected. Multiple path globs can be specified and are separated by :. Currently only used for remote coverage on PowerShell modules. This is for internal use only.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCoverageRemotePathFilter(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCoverageRemotePathFilter] = value
	return e
}

// WithAnsibleActionPlugins sets the value for the configuraion ANSIBLE_ACTION_PLUGINS (Colon separated paths in which Ansible will search for Action Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleActionPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleActionPlugins] = value
	return e
}

// WithAnsibleAskPass sets the option ANSIBLE_ASK_PASS to true (This controls whether an Ansible playbook should prompt for a login password. If using SSH keys for authentication, you probably do not need to change this setting.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleAskPass() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleAskPass] = "true"
	return e
}

// WithoutAnsibleAskPass sets the option ANSIBLE_ASK_PASS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleAskPass() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleAskPass] = "false"
	return e
}

// WithAnsibleAskVaultPass sets the option ANSIBLE_ASK_VAULT_PASS to true (This controls whether an Ansible playbook should prompt for a vault password.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleAskVaultPass() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleAskVaultPass] = "true"
	return e
}

// WithoutAnsibleAskVaultPass sets the option ANSIBLE_ASK_VAULT_PASS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleAskVaultPass() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleAskVaultPass] = "false"
	return e
}

// WithAnsibleBecome sets the option ANSIBLE_BECOME to true (Toggles the use of privilege escalation, allowing you to ‘become’ another user after login.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleBecome() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecome] = "true"
	return e
}

// WithoutAnsibleBecome sets the option ANSIBLE_BECOME to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleBecome() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecome] = "false"
	return e
}

// WithAnsibleBecomeAskPass sets the option ANSIBLE_BECOME_ASK_PASS to true (Toggle to prompt for privilege escalation password.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleBecomeAskPass() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecomeAskPass] = "true"
	return e
}

// WithoutAnsibleBecomeAskPass sets the option ANSIBLE_BECOME_ASK_PASS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleBecomeAskPass() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecomeAskPass] = "false"
	return e
}

// WithAnsibleBecomeExe sets the value for the configuraion ANSIBLE_BECOME_EXE (executable to use for privilege escalation, otherwise Ansible will depend on PATH)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleBecomeExe(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecomeExe] = value
	return e
}

// WithAnsibleBecomeFlags sets the value for the configuraion ANSIBLE_BECOME_FLAGS (Flags to pass to the privilege escalation executable.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleBecomeFlags(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecomeFlags] = value
	return e
}

// WithAnsibleBecomeMethod sets the value for the configuraion ANSIBLE_BECOME_METHOD (Privilege escalation method to use when become is enabled.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleBecomeMethod(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecomeMethod] = value
	return e
}

// WithAnsibleBecomeUser sets the value for the configuraion ANSIBLE_BECOME_USER (The user your login/remote user ‘becomes’ when using privilege escalation, most systems will use ‘root’ when no user is specified.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleBecomeUser(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleBecomeUser] = value
	return e
}

// WithAnsibleCachePlugins sets the value for the configuraion ANSIBLE_CACHE_PLUGINS (Colon separated paths in which Ansible will search for Cache Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCachePlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCachePlugins] = value
	return e
}

// WithAnsibleCallbackPlugins sets the value for the configuraion ANSIBLE_CALLBACK_PLUGINS (Colon separated paths in which Ansible will search for Callback Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCallbackPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCallbackPlugins] = value
	return e
}

// WithAnsibleCliconfPlugins sets the value for the configuraion ANSIBLE_CLICONF_PLUGINS (Colon separated paths in which Ansible will search for Cliconf Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleCliconfPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleCliconfPlugins] = value
	return e
}

// WithAnsibleConnectionPlugins sets the value for the configuraion ANSIBLE_CONNECTION_PLUGINS (Colon separated paths in which Ansible will search for Connection Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleConnectionPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleConnectionPlugins] = value
	return e
}

// WithAnsibleDebug sets the option ANSIBLE_DEBUG to true (Toggles debug output in Ansible. This is very verbose and can hinder multiprocessing.  Debug output can also include secret information despite no_log settings being enabled, which means debug mode should not be used in production.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleDebug() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDebug] = "true"
	return e
}

// WithoutAnsibleDebug sets the option ANSIBLE_DEBUG to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleDebug() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDebug] = "false"
	return e
}

// WithAnsibleExecutable sets the value for the configuraion ANSIBLE_EXECUTABLE (This indicates the command to use to spawn a shell under for Ansible’s execution needs on a target. Users may need to change this in rare instances when shell usage is constrained, but in most cases it may be left as is.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleExecutable(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleExecutable] = value
	return e
}

// WithAnsibleFactPath sets the value for the configuraion ANSIBLE_FACT_PATH (This option allows you to globally configure a custom path for ‘local_facts’ for the implied ansible_collections.ansible.builtin.setup_module task when using fact gathering. If not set, it will fallback to the default from the ansible.builtin.setup module: /etc/ansible/facts.d. This does not affect  user defined tasks that use the ansible.builtin.setup module. The real action being created by the implicit task is currently    ansible.legacy.gather_facts module, which then calls the configured fact modules, by default this will be ansible.builtin.setup for POSIX systems but other platforms might have different defaults.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleFactPath(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleFactPath] = value
	return e
}

// WithAnsibleFilterPlugins sets the value for the configuraion ANSIBLE_FILTER_PLUGINS (Colon separated paths in which Ansible will search for Jinja2 Filter Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleFilterPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleFilterPlugins] = value
	return e
}

// WithAnsibleForceHandlers sets the option ANSIBLE_FORCE_HANDLERS to true (This option controls if notified handlers run on a host even if a failure occurs on that host. When false, the handlers will not run if a failure has occurred on a host. This can also be set per play or on the command line. See Handlers and Failure for more details.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleForceHandlers() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleForceHandlers] = "true"
	return e
}

// WithoutAnsibleForceHandlers sets the option ANSIBLE_FORCE_HANDLERS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleForceHandlers() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleForceHandlers] = "false"
	return e
}

// WithAnsibleForks sets the value for the configuraion ANSIBLE_FORKS (Maximum number of forks Ansible will use to execute tasks on target hosts.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleForks(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleForks] = fmt.Sprint(value)
	return e
}

// WithAnsibleGatherSubset sets the value for the configuraion ANSIBLE_GATHER_SUBSET (Set the gather_subset option for the ansible_collections.ansible.builtin.setup_module task in the implicit fact gathering. See the module documentation for specifics. It does not apply to user defined ansible.builtin.setup tasks.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGatherSubset(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGatherSubset] = value
	return e
}

// WithAnsibleGatherTimeout sets the value for the configuraion ANSIBLE_GATHER_TIMEOUT (Set the timeout in seconds for the implicit fact gathering, see the module documentation for specifics. It does not apply to user defined ansible_collections.ansible.builtin.setup_module tasks.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGatherTimeout(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGatherTimeout] = fmt.Sprint(value)
	return e
}

// WithAnsibleGathering sets the value for the configuraion ANSIBLE_GATHERING (This setting controls the default policy of fact gathering (facts discovered about remote systems). This option can be useful for those wishing to save fact gathering time. Both ‘smart’ and ‘explicit’ will use the cache plugin.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGathering(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGathering] = value
	return e
}

// WithAnsibleHashBehaviour sets the value for the configuraion ANSIBLE_HASH_BEHAVIOUR (This setting controls how duplicate definitions of dictionary variables (aka hash, map, associative array) are handled in Ansible. This does not affect variables whose values are scalars (integers, strings) or arrays. WARNING, changing this setting is not recommended as this is fragile and makes your content (plays, roles, collections) non portable, leading to continual confusion and misuse. Don’t change this setting unless you think you have an absolute need for it. We recommend avoiding reusing variable names and relying on the combine filter and vars and varnames lookups to create merged versions of the individual variables. In our experience this is rarely really needed and a sign that too much complexity has been introduced into the data structures and plays. For some uses you can also look into custom vars_plugins to merge on input, even substituting the default host_group_vars that is in charge of parsing the host_vars/ and group_vars/ directories. Most users of this setting are only interested in inventory scope, but the setting itself affects all sources and makes debugging even harder. All playbooks and roles in the official examples repos assume the default for this setting. Changing the setting to merge applies across variable sources, but many sources will internally still overwrite the variables. For example include_vars will dedupe variables internally before updating Ansible, with ‘last defined’ overwriting previous definitions in same file. The Ansible project recommends you avoid “merge“ for new projects. It is the intention of the Ansible developers to eventually deprecate and remove this setting, but it is being kept as some users do heavily rely on it. New projects should avoid ‘merge’.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleHashBehaviour(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleHashBehaviour] = value
	return e
}

// WithAnsibleInventory sets the value for the configuraion ANSIBLE_INVENTORY (Comma separated list of Ansible inventory sources)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventory(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventory] = value
	return e
}

// WithAnsibleHttpapiPlugins sets the value for the configuraion ANSIBLE_HTTPAPI_PLUGINS (Colon separated paths in which Ansible will search for HttpApi Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleHttpapiPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleHttpapiPlugins] = value
	return e
}

// WithAnsibleInventoryPlugins sets the value for the configuraion ANSIBLE_INVENTORY_PLUGINS (Colon separated paths in which Ansible will search for Inventory Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryPlugins] = value
	return e
}

// WithAnsibleJinja2Extensions sets the value for the configuraion ANSIBLE_JINJA2_EXTENSIONS (This is a developer-specific feature that allows enabling additional Jinja2 extensions. See the Jinja2 documentation for details. If you do not know what these do, you probably don’t need to change this setting :))
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleJinja2Extensions(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleJinja2Extensions] = value
	return e
}

// WithAnsibleJinja2Native sets the option ANSIBLE_JINJA2_NATIVE to true (This option preserves variable types during template operations.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleJinja2Native() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleJinja2Native] = "true"
	return e
}

// WithoutAnsibleJinja2Native sets the option ANSIBLE_JINJA2_NATIVE to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleJinja2Native() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleJinja2Native] = "false"
	return e
}

// WithAnsibleKeepRemoteFiles sets the option ANSIBLE_KEEP_REMOTE_FILES to true (Enables/disables the cleaning up of the temporary files Ansible used to execute the tasks on the remote. If this option is enabled it will disable ANSIBLE_PIPELINING.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleKeepRemoteFiles() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleKeepRemoteFiles] = "true"
	return e
}

// WithoutAnsibleKeepRemoteFiles sets the option ANSIBLE_KEEP_REMOTE_FILES to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleKeepRemoteFiles() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleKeepRemoteFiles] = "false"
	return e
}

// WithAnsibleLibvirtLxcNoseclabel sets the option ANSIBLE_LIBVIRT_LXC_NOSECLABEL to true (This setting causes libvirt to connect to lxc containers by passing –noseclabel to virsh. This is necessary when running on systems which do not have SELinux.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleLibvirtLxcNoseclabel() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleLibvirtLxcNoseclabel] = "true"
	return e
}

// WithoutAnsibleLibvirtLxcNoseclabel sets the option ANSIBLE_LIBVIRT_LXC_NOSECLABEL to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleLibvirtLxcNoseclabel() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleLibvirtLxcNoseclabel] = "false"
	return e
}

// WithAnsibleLoadCallbackPlugins sets the option ANSIBLE_LOAD_CALLBACK_PLUGINS to true (Controls whether callback plugins are loaded when running /usr/bin/ansible. This may be used to log activity from the command line, send notifications, and so on. Callback plugins are always loaded for ansible-playbook.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleLoadCallbackPlugins() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleLoadCallbackPlugins] = "true"
	return e
}

// WithoutAnsibleLoadCallbackPlugins sets the option ANSIBLE_LOAD_CALLBACK_PLUGINS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleLoadCallbackPlugins() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleLoadCallbackPlugins] = "false"
	return e
}

// WithAnsibleLocalTemp sets the value for the configuraion ANSIBLE_LOCAL_TEMP (Temporary directory for Ansible to use on the controller.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleLocalTemp(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleLocalTemp] = value
	return e
}

// WithAnsibleLogFilter sets the value for the configuraion ANSIBLE_LOG_FILTER (List of logger names to filter out of the log file)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleLogFilter(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleLogFilter] = value
	return e
}

// WithAnsibleLogPath sets the value for the configuraion ANSIBLE_LOG_PATH (File to which Ansible will log on the controller. When empty logging is disabled.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleLogPath(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleLogPath] = value
	return e
}

// WithAnsibleLookupPlugins sets the value for the configuraion ANSIBLE_LOOKUP_PLUGINS (Colon separated paths in which Ansible will search for Lookup Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleLookupPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleLookupPlugins] = value
	return e
}

// WithAnsibleModuleArgs sets the value for the configuraion ANSIBLE_MODULE_ARGS (This sets the default arguments to pass to the ansible adhoc binary if no -a is specified.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleModuleArgs(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleModuleArgs] = value
	return e
}

// WithAnsibleLibrary sets the value for the configuraion ANSIBLE_LIBRARY (Colon separated paths in which Ansible will search for Modules.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleLibrary(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleLibrary] = value
	return e
}

// WithAnsibleModuleUtils sets the value for the configuraion ANSIBLE_MODULE_UTILS (Colon separated paths in which Ansible will search for Module utils files, which are shared by modules.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleModuleUtils(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleModuleUtils] = value
	return e
}

// WithAnsibleNetconfPlugins sets the value for the configuraion ANSIBLE_NETCONF_PLUGINS (Colon separated paths in which Ansible will search for Netconf Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleNetconfPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleNetconfPlugins] = value
	return e
}

// WithAnsibleNoLog sets the option ANSIBLE_NO_LOG to true (Toggle Ansible’s display and logging of task details, mainly used to avoid security disclosures.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleNoLog() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleNoLog] = "true"
	return e
}

// WithoutAnsibleNoLog sets the option ANSIBLE_NO_LOG to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleNoLog() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleNoLog] = "false"
	return e
}

// WithAnsibleNoTargetSyslog sets the option ANSIBLE_NO_TARGET_SYSLOG to true (Toggle Ansible logging to syslog on the target when it executes tasks. On Windows hosts this will disable a newer style PowerShell modules from writing to the event log.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleNoTargetSyslog() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleNoTargetSyslog] = "true"
	return e
}

// WithoutAnsibleNoTargetSyslog sets the option ANSIBLE_NO_TARGET_SYSLOG to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleNoTargetSyslog() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleNoTargetSyslog] = "false"
	return e
}

// WithAnsibleNullRepresentation sets the value for the configuraion ANSIBLE_NULL_REPRESENTATION (What templating should return as a ‘null’ value. When not set it will let Jinja2 decide.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleNullRepresentation(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleNullRepresentation] = value
	return e
}

// WithAnsiblePollInterval sets the value for the configuraion ANSIBLE_POLL_INTERVAL (For asynchronous tasks in Ansible (covered in Asynchronous Actions and Polling), this is how often to check back on the status of those tasks when an explicit poll interval is not supplied. The default is a reasonably moderate 15 seconds which is a tradeoff between checking in frequently and providing a quick turnaround when something may have completed.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePollInterval(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePollInterval] = fmt.Sprint(value)
	return e
}

// WithAnsiblePrivateKeyFile sets the value for the configuraion ANSIBLE_PRIVATE_KEY_FILE (Option for connections using a certificate or key file to authenticate, rather than an agent or passwords, you can set the default value here to avoid re-specifying –private-key with every invocation.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePrivateKeyFile(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePrivateKeyFile] = value
	return e
}

// WithAnsiblePrivateRoleVars sets the option ANSIBLE_PRIVATE_ROLE_VARS to true (Makes role variables inaccessible from other roles. This was introduced as a way to reset role variables to default values if a role is used more than once in a playbook.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePrivateRoleVars() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePrivateRoleVars] = "true"
	return e
}

// WithoutAnsiblePrivateRoleVars sets the option ANSIBLE_PRIVATE_ROLE_VARS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsiblePrivateRoleVars() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePrivateRoleVars] = "false"
	return e
}

// WithAnsibleRemotePort sets the value for the configuraion ANSIBLE_REMOTE_PORT (Port to use in remote connections, when blank it will use the connection plugin default.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleRemotePort(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleRemotePort] = fmt.Sprint(value)
	return e
}

// WithAnsibleRemoteUser sets the value for the configuraion ANSIBLE_REMOTE_USER (Sets the login user for the target machines When blank it uses the connection plugin’s default, normally the user currently executing Ansible.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleRemoteUser(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleRemoteUser] = value
	return e
}

// WithAnsibleRolesPath sets the value for the configuraion ANSIBLE_ROLES_PATH (Colon separated paths in which Ansible will search for Roles.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleRolesPath(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleRolesPath] = value
	return e
}

// WithAnsibleSelinuxSpecialFs sets the value for the configuraion ANSIBLE_SELINUX_SPECIAL_FS (Some filesystems do not support safe operations and/or return inconsistent errors, this setting makes Ansible ‘tolerate’ those in the list w/o causing fatal errors. Data corruption may occur and writes are not always verified when a filesystem is in the list. [:Version Added: 2.9])
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleSelinuxSpecialFs(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleSelinuxSpecialFs] = value
	return e
}

// WithAnsibleStdoutCallback sets the value for the configuraion ANSIBLE_STDOUT_CALLBACK (Set the main callback used to display Ansible output. You can only have one at a time. You can have many other callbacks, but just one can be in charge of stdout. See Callback plugins for a list of available options.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleStdoutCallback(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleStdoutCallback] = value
	return e
}

// WithAnsibleStrategy sets the value for the configuraion ANSIBLE_STRATEGY (Set the default strategy used for plays.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleStrategy(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleStrategy] = value
	return e
}

// WithAnsibleStrategyPlugins sets the value for the configuraion ANSIBLE_STRATEGY_PLUGINS (Colon separated paths in which Ansible will search for Strategy Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleStrategyPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleStrategyPlugins] = value
	return e
}

// WithAnsibleSu sets the option ANSIBLE_SU to true (Toggle the use of “su” for tasks.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleSu() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleSu] = "true"
	return e
}

// WithoutAnsibleSu sets the option ANSIBLE_SU to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleSu() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleSu] = "false"
	return e
}

// WithAnsibleSyslogFacility sets the value for the configuraion ANSIBLE_SYSLOG_FACILITY (Syslog facility to use when Ansible logs to the remote target)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleSyslogFacility(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleSyslogFacility] = value
	return e
}

// WithAnsibleTerminalPlugins sets the value for the configuraion ANSIBLE_TERMINAL_PLUGINS (Colon separated paths in which Ansible will search for Terminal Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleTerminalPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleTerminalPlugins] = value
	return e
}

// WithAnsibleTestPlugins sets the value for the configuraion ANSIBLE_TEST_PLUGINS (Colon separated paths in which Ansible will search for Jinja2 Test Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleTestPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleTestPlugins] = value
	return e
}

// WithAnsibleTimeout sets the value for the configuraion ANSIBLE_TIMEOUT (This is the default timeout for connection plugins to use.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleTimeout(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleTimeout] = fmt.Sprint(value)
	return e
}

// WithAnsibleTransport sets the value for the configuraion ANSIBLE_TRANSPORT (Default connection plugin to use, the ‘smart’ option will toggle between ‘ssh’ and ‘paramiko’ depending on controller OS and ssh versions)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleTransport(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleTransport] = value
	return e
}

// WithAnsibleErrorOnUndefinedVars sets the option ANSIBLE_ERROR_ON_UNDEFINED_VARS to true (When True, this causes ansible templating to fail steps that reference variable names that are likely typoed. Otherwise, any ‘{{ template_expression }}’ that contains undefined variables will be rendered in a template or ansible action line exactly as written.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleErrorOnUndefinedVars() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleErrorOnUndefinedVars] = "true"
	return e
}

// WithoutAnsibleErrorOnUndefinedVars sets the option ANSIBLE_ERROR_ON_UNDEFINED_VARS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleErrorOnUndefinedVars() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleErrorOnUndefinedVars] = "false"
	return e
}

// WithAnsibleVarsPlugins sets the value for the configuraion ANSIBLE_VARS_PLUGINS (Colon separated paths in which Ansible will search for Vars Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleVarsPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleVarsPlugins] = value
	return e
}

// WithAnsibleVaultEncryptIdentity sets the value for the configuraion ANSIBLE_VAULT_ENCRYPT_IDENTITY (The vault_id to use for encrypting by default. If multiple vault_ids are provided, this specifies which to use for encryption. The –encrypt-vault-id cli option overrides the configured value.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleVaultEncryptIdentity(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleVaultEncryptIdentity] = value
	return e
}

// WithAnsibleVaultIdMatch sets the value for the configuraion ANSIBLE_VAULT_ID_MATCH (If true, decrypting vaults with a vault id will only try the password from the matching vault-id)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleVaultIdMatch(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleVaultIdMatch] = value
	return e
}

// WithAnsibleVaultIdentity sets the value for the configuraion ANSIBLE_VAULT_IDENTITY (The label to use for the default vault id label in cases where a vault id label is not provided)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleVaultIdentity(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleVaultIdentity] = value
	return e
}

// WithAnsibleVaultIdentityList sets the value for the configuraion ANSIBLE_VAULT_IDENTITY_LIST (A list of vault-ids to use by default. Equivalent to multiple –vault-id args. Vault-ids are tried in order.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleVaultIdentityList(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleVaultIdentityList] = value
	return e
}

// WithAnsibleVaultPasswordFile sets the value for the configuraion ANSIBLE_VAULT_PASSWORD_FILE (The vault password file to use. Equivalent to –vault-password-file or –vault-id If executable, it will be run and the resulting stdout will be used as the password.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleVaultPasswordFile(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleVaultPasswordFile] = value
	return e
}

// WithAnsibleVerbosity sets the value for the configuraion ANSIBLE_VERBOSITY (Sets the default verbosity, equivalent to the number of -v passed in the command line.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleVerbosity(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleVerbosity] = fmt.Sprint(value)
	return e
}

// WithAnsibleDeprecationWarnings sets the option ANSIBLE_DEPRECATION_WARNINGS to true (Toggle to control the showing of deprecation warnings)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleDeprecationWarnings() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDeprecationWarnings] = "true"
	return e
}

// WithoutAnsibleDeprecationWarnings sets the option ANSIBLE_DEPRECATION_WARNINGS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleDeprecationWarnings() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDeprecationWarnings] = "false"
	return e
}

// WithAnsibleDevelWarning sets the option ANSIBLE_DEVEL_WARNING to true (Toggle to control showing warnings related to running devel)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleDevelWarning() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDevelWarning] = "true"
	return e
}

// WithoutAnsibleDevelWarning sets the option ANSIBLE_DEVEL_WARNING to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleDevelWarning() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDevelWarning] = "false"
	return e
}

// WithAnsibleDiffAlways sets the value for the configuraion ANSIBLE_DIFF_ALWAYS (Configuration toggle to tell modules to show differences when in ‘changed’ status, equivalent to --diff.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleDiffAlways(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDiffAlways] = value
	return e
}

// WithAnsibleDiffContext sets the value for the configuraion ANSIBLE_DIFF_CONTEXT (How many lines of context to show when displaying the differences between files.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleDiffContext(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDiffContext] = fmt.Sprint(value)
	return e
}

// WithAnsibleDisplayArgsToStdout sets the option ANSIBLE_DISPLAY_ARGS_TO_STDOUT to true (Normally ansible-playbook will print a header for each task that is run. These headers will contain the name: field from the task if you specified one. If you didn’t then ansible-playbook uses the task’s action to help you tell which task is presently running. Sometimes you run many of the same action and so you want more information about the task to differentiate it from others of the same action. If you set this variable to True in the config then ansible-playbook will also include the task’s arguments in the header. This setting defaults to False because there is a chance that you have sensitive values in your parameters and you do not want those to be printed. If you set this to True you should be sure that you have secured your environment’s stdout (no one can shoulder surf your screen and you aren’t saving stdout to an insecure file) or made sure that all of your playbooks explicitly added the no_log: True parameter to tasks which have sensitive values See How do I keep secret data in my playbook? for more information.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleDisplayArgsToStdout() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDisplayArgsToStdout] = "true"
	return e
}

// WithoutAnsibleDisplayArgsToStdout sets the option ANSIBLE_DISPLAY_ARGS_TO_STDOUT to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleDisplayArgsToStdout() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDisplayArgsToStdout] = "false"
	return e
}

// WithAnsibleDisplaySkippedHosts sets the option ANSIBLE_DISPLAY_SKIPPED_HOSTS to true (Toggle to control displaying skipped task/host entries in a task in the default callback)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleDisplaySkippedHosts() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDisplaySkippedHosts] = "true"
	return e
}

// WithoutAnsibleDisplaySkippedHosts sets the option ANSIBLE_DISPLAY_SKIPPED_HOSTS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleDisplaySkippedHosts() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDisplaySkippedHosts] = "false"
	return e
}

// WithAnsibleDocFragmentPlugins sets the value for the configuraion ANSIBLE_DOC_FRAGMENT_PLUGINS (Colon separated paths in which Ansible will search for Documentation Fragments Plugins.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleDocFragmentPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDocFragmentPlugins] = value
	return e
}

// WithAnsibleDuplicateYamlDictKey sets the value for the configuraion ANSIBLE_DUPLICATE_YAML_DICT_KEY (By default Ansible will issue a warning when a duplicate dict key is encountered in YAML. These warnings can be silenced by adjusting this setting to False.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleDuplicateYamlDictKey(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleDuplicateYamlDictKey] = value
	return e
}

// WithEditor sets the value for the configuraion EDITOR ()
func (e *AnsibleWithConfigurationSettingsExecute) WithEditor(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[Editor] = value
	return e
}

// WithAnsibleEnableTaskDebugger sets the option ANSIBLE_ENABLE_TASK_DEBUGGER to true (Whether or not to enable the task debugger, this previously was done as a strategy plugin. Now all strategy plugins can inherit this behavior. The debugger defaults to activating when a task is failed on unreachable. Use the debugger keyword for more flexibility.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleEnableTaskDebugger() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleEnableTaskDebugger] = "true"
	return e
}

// WithoutAnsibleEnableTaskDebugger sets the option ANSIBLE_ENABLE_TASK_DEBUGGER to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleEnableTaskDebugger() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleEnableTaskDebugger] = "false"
	return e
}

// WithAnsibleErrorOnMissingHandler sets the option ANSIBLE_ERROR_ON_MISSING_HANDLER to true (Toggle to allow missing handlers to become a warning instead of an error when notifying.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleErrorOnMissingHandler() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleErrorOnMissingHandler] = "true"
	return e
}

// WithoutAnsibleErrorOnMissingHandler sets the option ANSIBLE_ERROR_ON_MISSING_HANDLER to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleErrorOnMissingHandler() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleErrorOnMissingHandler] = "false"
	return e
}

// WithAnsibleFactsModules sets the value for the configuraion ANSIBLE_FACTS_MODULES (Which modules to run during a play’s fact gathering stage, using the default of ‘smart’ will try to figure it out based on connection type. If adding your own modules but you still want to use the default Ansible facts, you will want to include ‘setup’ or corresponding network module to the list (if you add ‘smart’, Ansible will also figure it out). This does not affect explicit calls to the ‘setup’ module, but does always affect the ‘gather_facts’ action (implicit or explicit).)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleFactsModules(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleFactsModules] = value
	return e
}

// WithAnsibleGalaxyCacheDir sets the value for the configuraion ANSIBLE_GALAXY_CACHE_DIR (The directory that stores cached responses from a Galaxy server. This is only used by the ansible-galaxy collection install and download commands. Cache files inside this dir will be ignored if they are world writable.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyCacheDir(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyCacheDir] = value
	return e
}

// WithAnsibleGalaxyCollectionSkeleton sets the value for the configuraion ANSIBLE_GALAXY_COLLECTION_SKELETON (Collection skeleton directory to use as a template for the init action in ansible-galaxy collection, same as --collection-skeleton.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyCollectionSkeleton(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyCollectionSkeleton] = value
	return e
}

// WithAnsibleGalaxyCollectionSkeletonIgnore sets the value for the configuraion ANSIBLE_GALAXY_COLLECTION_SKELETON_IGNORE (patterns of files to ignore inside a Galaxy collection skeleton directory)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyCollectionSkeletonIgnore(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyCollectionSkeletonIgnore] = value
	return e
}

// WithAnsibleGalaxyDisableGpgVerify sets the value for the configuraion ANSIBLE_GALAXY_DISABLE_GPG_VERIFY (Disable GPG signature verification during collection installation.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyDisableGpgVerify(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyDisableGpgVerify] = value
	return e
}

// WithAnsibleGalaxyDisplayProgress sets the value for the configuraion ANSIBLE_GALAXY_DISPLAY_PROGRESS (Some steps in ansible-galaxy display a progress wheel which can cause issues on certain displays or when outputting the stdout to a file. This config option controls whether the display wheel is shown or not. The default is to show the display wheel if stdout has a tty.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyDisplayProgress(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyDisplayProgress] = value
	return e
}

// WithAnsibleGalaxyGpgKeyring sets the value for the configuraion ANSIBLE_GALAXY_GPG_KEYRING (Configure the keyring used for GPG signature verification during collection installation and verification.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyGpgKeyring(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyGpgKeyring] = value
	return e
}

// WithAnsibleGalaxyIgnore sets the option ANSIBLE_GALAXY_IGNORE to true (If set to yes, ansible-galaxy will not validate TLS certificates. This can be useful for testing against a server with a self-signed certificate.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyIgnore() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyIgnore] = "true"
	return e
}

// WithoutAnsibleGalaxyIgnore sets the option ANSIBLE_GALAXY_IGNORE to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleGalaxyIgnore() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyIgnore] = "false"
	return e
}

// WithAnsibleGalaxyIgnoreSignatureStatusCodes sets the value for the configuraion ANSIBLE_GALAXY_IGNORE_SIGNATURE_STATUS_CODES (A list of GPG status codes to ignore during GPG signature verification. See L(https://github.com/gpg/gnupg/blob/master/doc/DETAILS#general-status-codes) for status code descriptions. If fewer signatures successfully verify the collection than GALAXY_REQUIRED_VALID_SIGNATURE_COUNT, signature verification will fail even if all error codes are ignored.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyIgnoreSignatureStatusCodes(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyIgnoreSignatureStatusCodes] = value
	return e
}

// WithAnsibleGalaxyRequiredValidSignatureCount sets the value for the configuraion ANSIBLE_GALAXY_REQUIRED_VALID_SIGNATURE_COUNT (The number of signatures that must be successful during GPG signature verification while installing or verifying collections. This should be a positive integer or all to indicate all signatures must successfully validate the collection. Prepend + to the value to fail if no valid signatures are found for the collection.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyRequiredValidSignatureCount(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyRequiredValidSignatureCount] = value
	return e
}

// WithAnsibleGalaxyRoleSkeleton sets the value for the configuraion ANSIBLE_GALAXY_ROLE_SKELETON (Role skeleton directory to use as a template for the init action in ansible-galaxy/ansible-galaxy role, same as --role-skeleton.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyRoleSkeleton(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyRoleSkeleton] = value
	return e
}

// WithAnsibleGalaxyRoleSkeletonIgnore sets the value for the configuraion ANSIBLE_GALAXY_ROLE_SKELETON_IGNORE (patterns of files to ignore inside a Galaxy role or collection skeleton directory)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyRoleSkeletonIgnore(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyRoleSkeletonIgnore] = value
	return e
}

// WithAnsibleGalaxyServer sets the value for the configuraion ANSIBLE_GALAXY_SERVER (URL to prepend when roles don’t specify the full URI, assume they are referencing this server as the source.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyServer(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyServer] = value
	return e
}

// WithAnsibleGalaxyServerList sets the value for the configuraion ANSIBLE_GALAXY_SERVER_LIST (A list of Galaxy servers to use when installing a collection. The value corresponds to the config ini header [galaxy_server.{{item}}] which defines the server details. See Configuring the ansible-galaxy client for more details on how to define a Galaxy server. The order of servers in this list is used to as the order in which a collection is resolved. Setting this config option will ignore the GALAXY_SERVER config option.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyServerList(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyServerList] = value
	return e
}

// WithAnsibleGalaxyTokenPath sets the value for the configuraion ANSIBLE_GALAXY_TOKEN_PATH (Local path to galaxy access token file)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleGalaxyTokenPath(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleGalaxyTokenPath] = value
	return e
}

// WithAnsibleHostKeyChecking sets the option ANSIBLE_HOST_KEY_CHECKING to true (Set this to “False” if you want to avoid host key checking by the underlying tools Ansible uses to connect to the host)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleHostKeyChecking() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleHostKeyChecking] = "true"
	return e
}

// WithoutAnsibleHostKeyChecking sets the option ANSIBLE_HOST_KEY_CHECKING to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleHostKeyChecking() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleHostKeyChecking] = "false"
	return e
}

// WithAnsibleHostPatternMismatch sets the value for the configuraion ANSIBLE_HOST_PATTERN_MISMATCH (This setting changes the behaviour of mismatched host patterns, it allows you to force a fatal error, a warning or just ignore it)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleHostPatternMismatch(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleHostPatternMismatch] = value
	return e
}

// WithAnsibleInjectFactVars sets the option ANSIBLE_INJECT_FACT_VARS to true (Facts are available inside the ansible_facts variable, this setting also pushes them as their own vars in the main namespace. Unlike inside the ansible_facts dictionary, these will have an ansible_ prefix.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInjectFactVars() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInjectFactVars] = "true"
	return e
}

// WithoutAnsibleInjectFactVars sets the option ANSIBLE_INJECT_FACT_VARS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleInjectFactVars() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInjectFactVars] = "false"
	return e
}

// WithAnsiblePythonInterpreter sets the value for the configuraion ANSIBLE_PYTHON_INTERPRETER (Path to the Python interpreter to be used for module execution on remote targets, or an automatic discovery mode. Supported discovery modes are auto (the default), auto_silent, auto_legacy, and auto_legacy_silent. All discovery modes employ a lookup table to use the included system Python (on distributions known to include one), falling back to a fixed ordered list of well-known Python interpreter locations if a platform-specific default is not available. The fallback behavior will issue a warning that the interpreter should be set explicitly (since interpreters installed later may change which one is used). This warning behavior can be disabled by setting auto_silent or auto_legacy_silent. The value of auto_legacy provides all the same behavior, but for backwards-compatibility with older Ansible releases that always defaulted to /usr/bin/python, will use that interpreter if present.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePythonInterpreter(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePythonInterpreter] = value
	return e
}

// WithAnsibleInvalidTaskAttributeFailed sets the option ANSIBLE_INVALID_TASK_ATTRIBUTE_FAILED to true (If ‘false’, invalid attributes for a task will result in warnings instead of errors)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInvalidTaskAttributeFailed() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInvalidTaskAttributeFailed] = "true"
	return e
}

// WithoutAnsibleInvalidTaskAttributeFailed sets the option ANSIBLE_INVALID_TASK_ATTRIBUTE_FAILED to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleInvalidTaskAttributeFailed() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInvalidTaskAttributeFailed] = "false"
	return e
}

// WithAnsibleInventoryAnyUnparsedIsFailed sets the option ANSIBLE_INVENTORY_ANY_UNPARSED_IS_FAILED to true (If ‘true’, it is a fatal error when any given inventory source cannot be successfully parsed by any available inventory plugin; otherwise, this situation only attracts a warning.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryAnyUnparsedIsFailed() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryAnyUnparsedIsFailed] = "true"
	return e
}

// WithoutAnsibleInventoryAnyUnparsedIsFailed sets the option ANSIBLE_INVENTORY_ANY_UNPARSED_IS_FAILED to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleInventoryAnyUnparsedIsFailed() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryAnyUnparsedIsFailed] = "false"
	return e
}

// WithAnsibleInventoryCache sets the value for the configuraion ANSIBLE_INVENTORY_CACHE (Toggle to turn on inventory caching. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory configuration. This message will be removed in 2.16.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryCache(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryCache] = value
	return e
}

// WithAnsibleInventoryCachePlugin sets the value for the configuraion ANSIBLE_INVENTORY_CACHE_PLUGIN (The plugin for caching inventory. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryCachePlugin(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryCachePlugin] = value
	return e
}

// WithAnsibleInventoryCacheConnection sets the value for the configuraion ANSIBLE_INVENTORY_CACHE_CONNECTION (The inventory cache connection. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryCacheConnection(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryCacheConnection] = value
	return e
}

// WithAnsibleInventoryCachePluginPrefix sets the value for the configuraion ANSIBLE_INVENTORY_CACHE_PLUGIN_PREFIX (The table prefix for the cache plugin. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryCachePluginPrefix(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryCachePluginPrefix] = value
	return e
}

// WithAnsibleInventoryCacheTimeout sets the value for the configuraion ANSIBLE_INVENTORY_CACHE_TIMEOUT (Expiration timeout for the inventory cache plugin data. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryCacheTimeout(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryCacheTimeout] = value
	return e
}

// WithAnsibleInventoryEnabled sets the value for the configuraion ANSIBLE_INVENTORY_ENABLED (List of enabled inventory plugins, it also determines the order in which they are used.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryEnabled(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryEnabled] = value
	return e
}

// WithAnsibleInventoryExport sets the value for the configuraion ANSIBLE_INVENTORY_EXPORT (Controls if ansible-inventory will accurately reflect Ansible’s view into inventory or its optimized for exporting.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryExport(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryExport] = value
	return e
}

// WithAnsibleInventoryIgnore sets the value for the configuraion ANSIBLE_INVENTORY_IGNORE (List of extensions to ignore when using a directory as an inventory source)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryIgnore(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryIgnore] = value
	return e
}

// WithAnsibleInventoryIgnoreRegex sets the value for the configuraion ANSIBLE_INVENTORY_IGNORE_REGEX (List of patterns to ignore when using a directory as an inventory source)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryIgnoreRegex(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryIgnoreRegex] = value
	return e
}

// WithAnsibleInventoryUnparsedFailed sets the value for the configuraion ANSIBLE_INVENTORY_UNPARSED_FAILED (If ‘true’ it is a fatal error if every single potential inventory source fails to parse, otherwise this situation will only attract a warning.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryUnparsedFailed(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryUnparsedFailed] = value
	return e
}

// WithAnsibleInventoryUnparsedWarning sets the option ANSIBLE_INVENTORY_UNPARSED_WARNING to true (By default Ansible will issue a warning when no inventory was loaded and notes that it will use an implicit localhost-only inventory. These warnings can be silenced by adjusting this setting to False.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleInventoryUnparsedWarning() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryUnparsedWarning] = "true"
	return e
}

// WithoutAnsibleInventoryUnparsedWarning sets the option ANSIBLE_INVENTORY_UNPARSED_WARNING to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleInventoryUnparsedWarning() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleInventoryUnparsedWarning] = "false"
	return e
}

// WithAnsibleJinja2NativeWarning sets the option ANSIBLE_JINJA2_NATIVE_WARNING to true (Toggle to control showing warnings related to running a Jinja version older than required for jinja2_native [:Deprecated in: 2.17 :Deprecated detail: This option is no longer used in the Ansible Core code base.])
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleJinja2NativeWarning() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleJinja2NativeWarning] = "true"
	return e
}

// WithoutAnsibleJinja2NativeWarning sets the option ANSIBLE_JINJA2_NATIVE_WARNING to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleJinja2NativeWarning() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleJinja2NativeWarning] = "false"
	return e
}

// WithAnsibleLocalhostWarning sets the option ANSIBLE_LOCALHOST_WARNING to true (By default Ansible will issue a warning when there are no hosts in the inventory. These warnings can be silenced by adjusting this setting to False.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleLocalhostWarning() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleLocalhostWarning] = "true"
	return e
}

// WithoutAnsibleLocalhostWarning sets the option ANSIBLE_LOCALHOST_WARNING to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleLocalhostWarning() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleLocalhostWarning] = "false"
	return e
}

// WithAnsibleMaxDiffSize sets the value for the configuraion ANSIBLE_MAX_DIFF_SIZE (Maximum size of files to be considered for diff display)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleMaxDiffSize(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleMaxDiffSize] = value
	return e
}

// WithAnsibleModuleIgnoreExts sets the value for the configuraion ANSIBLE_MODULE_IGNORE_EXTS (List of extensions to ignore when looking for modules to load This is for rejecting script and binary module fallback extensions)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleModuleIgnoreExts(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleModuleIgnoreExts] = value
	return e
}

// WithAnsibleModuleStrictUtf8Response sets the value for the configuraion ANSIBLE_MODULE_STRICT_UTF8_RESPONSE (Enables whether module responses are evaluated for containing non UTF-8 data Disabling this may result in unexpected behavior Only ansible-core should evaluate this configuration)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleModuleStrictUtf8Response(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleModuleStrictUtf8Response] = value
	return e
}

// WithAnsibleNetconfSshConfig sets the value for the configuraion ANSIBLE_NETCONF_SSH_CONFIG (This variable is used to enable bastion/jump host with netconf connection. If set to True the bastion/jump host ssh settings should be present in ~/.ssh/config file, alternatively it can be set to custom ssh configuration file path to read the bastion/jump host settings.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleNetconfSshConfig(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleNetconfSshConfig] = value
	return e
}

// WithAnsibleNetworkGroupModules sets the value for the configuraion ANSIBLE_NETWORK_GROUP_MODULES ()
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleNetworkGroupModules(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleNetworkGroupModules] = value
	return e
}

// WithAnsibleOldPluginCacheClear sets the option ANSIBLE_OLD_PLUGIN_CACHE_CLEAR to true (Previously Ansible would only clear some of the plugin loading caches when loading new roles, this led to some behaviours in which a plugin loaded in previous plays would be unexpectedly ‘sticky’. This setting allows to return to that behaviour.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleOldPluginCacheClear() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleOldPluginCacheClear] = "true"
	return e
}

// WithoutAnsibleOldPluginCacheClear sets the option ANSIBLE_OLD_PLUGIN_CACHE_CLEAR to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleOldPluginCacheClear() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleOldPluginCacheClear] = "false"
	return e
}

// WithPager sets the value for the configuraion PAGER ()
func (e *AnsibleWithConfigurationSettingsExecute) WithPager(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[Pager] = value
	return e
}

// WithAnsibleParamikoHostKeyAutoAdd sets the option ANSIBLE_PARAMIKO_HOST_KEY_AUTO_ADD to true ()
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleParamikoHostKeyAutoAdd() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleParamikoHostKeyAutoAdd] = "true"
	return e
}

// WithoutAnsibleParamikoHostKeyAutoAdd sets the option ANSIBLE_PARAMIKO_HOST_KEY_AUTO_ADD to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleParamikoHostKeyAutoAdd() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleParamikoHostKeyAutoAdd] = "false"
	return e
}

// WithAnsibleParamikoLookForKeys sets the option ANSIBLE_PARAMIKO_LOOK_FOR_KEYS to true ()
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleParamikoLookForKeys() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleParamikoLookForKeys] = "true"
	return e
}

// WithoutAnsibleParamikoLookForKeys sets the option ANSIBLE_PARAMIKO_LOOK_FOR_KEYS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleParamikoLookForKeys() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleParamikoLookForKeys] = "false"
	return e
}

// WithAnsiblePersistentCommandTimeout sets the value for the configuraion ANSIBLE_PERSISTENT_COMMAND_TIMEOUT (This controls the amount of time to wait for response from remote device before timing out persistent connection.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePersistentCommandTimeout(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePersistentCommandTimeout] = value
	return e
}

// WithAnsiblePersistentConnectRetryTimeout sets the value for the configuraion ANSIBLE_PERSISTENT_CONNECT_RETRY_TIMEOUT (This controls the retry timeout for persistent connection to connect to the local domain socket.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePersistentConnectRetryTimeout(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePersistentConnectRetryTimeout] = fmt.Sprint(value)
	return e
}

// WithAnsiblePersistentConnectTimeout sets the value for the configuraion ANSIBLE_PERSISTENT_CONNECT_TIMEOUT (This controls how long the persistent connection will remain idle before it is destroyed.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePersistentConnectTimeout(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePersistentConnectTimeout] = fmt.Sprint(value)
	return e
}

// WithAnsiblePersistentControlPathDir sets the value for the configuraion ANSIBLE_PERSISTENT_CONTROL_PATH_DIR (Path to socket to be used by the connection persistence system.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePersistentControlPathDir(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePersistentControlPathDir] = value
	return e
}

// WithAnsiblePlaybookDir sets the value for the configuraion ANSIBLE_PLAYBOOK_DIR (A number of non-playbook CLIs have a --playbook-dir argument; this sets the default value for it.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePlaybookDir(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePlaybookDir] = value
	return e
}

// WithAnsiblePlaybookVarsRoot sets the value for the configuraion ANSIBLE_PLAYBOOK_VARS_ROOT (This sets which playbook dirs will be used as a root to process vars plugins, which includes finding host_vars/group_vars)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePlaybookVarsRoot(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePlaybookVarsRoot] = value
	return e
}

// WithAnsiblePythonModuleRlimitNofile sets the value for the configuraion ANSIBLE_PYTHON_MODULE_RLIMIT_NOFILE (Attempts to set RLIMIT_NOFILE soft limit to the specified value when executing Python modules (can speed up subprocess usage on Python 2.x. See https://bugs.python.org/issue11284). The value will be limited by the existing hard limit. Default value of 0 does not attempt to adjust existing system-defined limits.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePythonModuleRlimitNofile(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePythonModuleRlimitNofile] = value
	return e
}

// WithAnsibleRetryFilesEnabled sets the value for the configuraion ANSIBLE_RETRY_FILES_ENABLED (This controls whether a failed Ansible playbook should create a .retry file.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleRetryFilesEnabled(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleRetryFilesEnabled] = value
	return e
}

// WithAnsibleRetryFilesSavePath sets the value for the configuraion ANSIBLE_RETRY_FILES_SAVE_PATH (This sets the path in which Ansible will save .retry files when a playbook fails and retry files are enabled. This file will be overwritten after each run with the list of failed hosts from all plays.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleRetryFilesSavePath(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleRetryFilesSavePath] = value
	return e
}

// WithAnsibleRunVarsPlugins sets the value for the configuraion ANSIBLE_RUN_VARS_PLUGINS (This setting can be used to optimize vars_plugin usage depending on user’s inventory size and play selection.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleRunVarsPlugins(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleRunVarsPlugins] = value
	return e
}

// WithAnsibleShowCustomStats sets the value for the configuraion ANSIBLE_SHOW_CUSTOM_STATS (This adds the custom stats set via the set_stats plugin to the default output)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleShowCustomStats(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleShowCustomStats] = value
	return e
}

// WithAnsibleStringConversionAction sets the value for the configuraion ANSIBLE_STRING_CONVERSION_ACTION (Action to take when a module parameter value is converted to a string (this does not affect variables). For string parameters, values such as ‘1.00’, “[‘a’, ‘b’,]”, and ‘yes’, ‘y’, etc. will be converted by the YAML parser unless fully quoted. Valid options are ‘error’, ‘warn’, and ‘ignore’. Since 2.8, this option defaults to ‘warn’ but will change to ‘error’ in 2.12.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleStringConversionAction(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleStringConversionAction] = value
	return e
}

// WithAnsibleStringTypeFilters sets the value for the configuraion ANSIBLE_STRING_TYPE_FILTERS (This list of filters avoids ‘type conversion’ when templating variables Useful when you want to avoid conversion into lists or dictionaries for JSON strings, for example.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleStringTypeFilters(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleStringTypeFilters] = value
	return e
}

// WithAnsibleSystemWarnings sets the option ANSIBLE_SYSTEM_WARNINGS to true (Allows disabling of warnings related to potential issues on the system running ansible itself (not on the managed hosts) These may include warnings about 3rd party packages or other conditions that should be resolved if possible.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleSystemWarnings() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleSystemWarnings] = "true"
	return e
}

// WithoutAnsibleSystemWarnings sets the option ANSIBLE_SYSTEM_WARNINGS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleSystemWarnings() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleSystemWarnings] = "false"
	return e
}

// WithAnsibleRunTags sets the value for the configuraion ANSIBLE_RUN_TAGS (default list of tags to run in your plays, Skip Tags has precedence.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleRunTags(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleRunTags] = value
	return e
}

// WithAnsibleSkipTags sets the value for the configuraion ANSIBLE_SKIP_TAGS (default list of tags to skip in your plays, has precedence over Run Tags)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleSkipTags(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleSkipTags] = value
	return e
}

// WithAnsibleTaskDebuggerIgnoreErrors sets the option ANSIBLE_TASK_DEBUGGER_IGNORE_ERRORS to true (This option defines whether the task debugger will be invoked on a failed task when ignore_errors=True is specified. True specifies that the debugger will honor ignore_errors, False will not honor ignore_errors.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleTaskDebuggerIgnoreErrors() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleTaskDebuggerIgnoreErrors] = "true"
	return e
}

// WithoutAnsibleTaskDebuggerIgnoreErrors sets the option ANSIBLE_TASK_DEBUGGER_IGNORE_ERRORS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleTaskDebuggerIgnoreErrors() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleTaskDebuggerIgnoreErrors] = "false"
	return e
}

// WithAnsibleTaskTimeout sets the value for the configuraion ANSIBLE_TASK_TIMEOUT (Set the maximum time (in seconds) that a task can run for. If set to 0 (the default) there is no timeout.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleTaskTimeout(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleTaskTimeout] = fmt.Sprint(value)
	return e
}

// WithAnsibleTransformInvalidGroupChars sets the value for the configuraion ANSIBLE_TRANSFORM_INVALID_GROUP_CHARS (Make ansible transform invalid characters in group names supplied by inventory sources.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleTransformInvalidGroupChars(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleTransformInvalidGroupChars] = value
	return e
}

// WithAnsibleUsePersistentConnections sets the option ANSIBLE_USE_PERSISTENT_CONNECTIONS to true (Toggles the use of persistence for connections.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleUsePersistentConnections() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleUsePersistentConnections] = "true"
	return e
}

// WithoutAnsibleUsePersistentConnections sets the option ANSIBLE_USE_PERSISTENT_CONNECTIONS to false
func (e *AnsibleWithConfigurationSettingsExecute) WithoutAnsibleUsePersistentConnections() *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleUsePersistentConnections] = "false"
	return e
}

// WithAnsibleValidateActionGroupMetadata sets the value for the configuraion ANSIBLE_VALIDATE_ACTION_GROUP_METADATA (A toggle to disable validating a collection’s ‘metadata’ entry for a module_defaults action group. Metadata containing unexpected fields or value types will produce a warning when this is True.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleValidateActionGroupMetadata(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleValidateActionGroupMetadata] = value
	return e
}

// WithAnsibleVarsEnabled sets the value for the configuraion ANSIBLE_VARS_ENABLED (Accept list for variable plugins that require it.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleVarsEnabled(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleVarsEnabled] = value
	return e
}

// WithAnsiblePrecedence sets the value for the configuraion ANSIBLE_PRECEDENCE (Allows to change the group variable precedence merge order.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsiblePrecedence(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsiblePrecedence] = value
	return e
}

// WithAnsibleVaultEncryptSalt sets the value for the configuraion ANSIBLE_VAULT_ENCRYPT_SALT (The salt to use for the vault encryption. If it is not provided, a random salt will be used.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleVaultEncryptSalt(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleVaultEncryptSalt] = value
	return e
}

// WithAnsibleVerboseToStderr sets the value for the configuraion ANSIBLE_VERBOSE_TO_STDERR (Force ‘verbose’ option to use stderr instead of stdout)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleVerboseToStderr(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleVerboseToStderr] = value
	return e
}

// WithAnsibleWinAsyncStartupTimeout sets the value for the configuraion ANSIBLE_WIN_ASYNC_STARTUP_TIMEOUT (For asynchronous tasks in Ansible (covered in Asynchronous Actions and Polling), this is how long, in seconds, to wait for the task spawned by Ansible to connect back to the named pipe used on Windows systems. The default is 5 seconds. This can be too low on slower systems, or systems under heavy load. This is not the total time an async command can run for, but is a separate timeout to wait for an async command to start. The task will only start to be timed against its async_timeout once it has connected to the pipe, so the overall maximum duration the task can take will be extended by the amount specified here.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleWinAsyncStartupTimeout(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleWinAsyncStartupTimeout] = fmt.Sprint(value)
	return e
}

// WithAnsibleWorkerShutdownPollCount sets the value for the configuraion ANSIBLE_WORKER_SHUTDOWN_POLL_COUNT (The maximum number of times to check Task Queue Manager worker processes to verify they have exited cleanly. After this limit is reached any worker processes still running will be terminated. This is for internal use only.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleWorkerShutdownPollCount(value int) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleWorkerShutdownPollCount] = fmt.Sprint(value)
	return e
}

// WithAnsibleWorkerShutdownPollDelay sets the value for the configuraion ANSIBLE_WORKER_SHUTDOWN_POLL_DELAY (The number of seconds to sleep between polling loops when checking Task Queue Manager worker processes to verify they have exited cleanly. This is for internal use only.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleWorkerShutdownPollDelay(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleWorkerShutdownPollDelay] = value
	return e
}

// WithAnsibleYamlFilenameExt sets the value for the configuraion ANSIBLE_YAML_FILENAME_EXT (Check all of these extensions when looking for ‘variable’ files which should be YAML or JSON or vaulted versions of these. This affects vars_files, include_vars, inventory and vars plugins among others.)
func (e *AnsibleWithConfigurationSettingsExecute) WithAnsibleYamlFilenameExt(value string) *AnsibleWithConfigurationSettingsExecute {
	e.configurationSettings[AnsibleYamlFilenameExt] = value
	return e
}
