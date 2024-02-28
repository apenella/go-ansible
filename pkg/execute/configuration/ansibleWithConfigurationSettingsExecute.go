package configuration

import (
	"context"
	"fmt"
)

//type configurationSettings map[string]string

// ConfigurationSettingsFunc is a function to set the configuration settings
type ConfigurationSettingsFunc func(*AnsibleWithConfigurationSettingsExecute)

// AnsibleWithConfigurationSettingsExecute is a builder for Ansible Cmd
type AnsibleWithConfigurationSettingsExecute struct {
	executor              ExecutorEnvVarSetter
	configurationSettings map[string]string
}

// NewAnsibleWithConfigurationSettingsExecute return a new AnsibleWithConfigurationSettingsExecute
func NewAnsibleWithConfigurationSettingsExecute(executor ExecutorEnvVarSetter, options ...ConfigurationSettingsFunc) *AnsibleWithConfigurationSettingsExecute {
	exec := &AnsibleWithConfigurationSettingsExecute{
		executor:              executor,
		configurationSettings: make(map[string]string),
	}

	for _, option := range options {
		option(exec)
	}

	return exec
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
func WithAnsibleActionWarnings() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleActionWarnings] = "true"
	}
}

// WithoutAnsibleActionWarnings sets the option ANSIBLE_ACTION_WARNINGS to false
func WithoutAnsibleActionWarnings() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleActionWarnings] = "false"
	}
}

// WithAnsibleAgnosticBecomePrompt sets the option ANSIBLE_AGNOSTIC_BECOME_PROMPT to true (Display an agnostic become prompt instead of displaying a prompt containing the command line supplied become method)
func WithAnsibleAgnosticBecomePrompt() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleAgnosticBecomePrompt] = "true"
	}
}

// WithoutAnsibleAgnosticBecomePrompt sets the option ANSIBLE_AGNOSTIC_BECOME_PROMPT to false
func WithoutAnsibleAgnosticBecomePrompt() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleAgnosticBecomePrompt] = "false"
	}
}

// WithAnsibleConnectionPath sets the value for the configuraion ANSIBLE_CONNECTION_PATH (Specify where to look for the ansible-connection script. This location will be checked before searching $PATH. If null, ansible will start with the same directory as the ansible script.)
func WithAnsibleConnectionPath(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleConnectionPath] = value
	}
}

// WithAnsibleCowAcceptlist sets the value for the configuraion ANSIBLE_COW_ACCEPTLIST (Accept list of cowsay templates that are ‘safe’ to use, set to empty list if you want to enable all installed templates. [:Version Added: 2.11])
func WithAnsibleCowAcceptlist(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCowAcceptlist] = value
	}
}

// WithAnsibleCowPath sets the value for the configuraion ANSIBLE_COW_PATH (Specify a custom cowsay path or swap in your cowsay implementation of choice)
func WithAnsibleCowPath(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCowPath] = value
	}
}

// WithAnsibleCowSelection sets the value for the configuraion ANSIBLE_COW_SELECTION (This allows you to chose a specific cowsay stencil for the banners or use ‘random’ to cycle through them.)
func WithAnsibleCowSelection(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCowSelection] = value
	}
}

// WithAnsibleForceColor sets the option ANSIBLE_FORCE_COLOR to true (This option forces color mode even when running without a TTY or the “nocolor” setting is True.)
func WithAnsibleForceColor() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleForceColor] = "true"
	}
}

// WithoutAnsibleForceColor sets the option ANSIBLE_FORCE_COLOR to false
func WithoutAnsibleForceColor() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleForceColor] = "false"
	}
}

// WithAnsibleHome sets the value for the configuraion ANSIBLE_HOME (The default root path for Ansible config files on the controller.)
func WithAnsibleHome(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleHome] = value
	}
}

// WithNoColor sets the option NO_COLOR to true (This setting allows suppressing colorizing output, which is used to give a better indication of failure and status information.)
func WithNoColor() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[NoColor] = "true"
	}
}

// WithoutNoColor sets the option NO_COLOR to false
func WithoutNoColor() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[NoColor] = "false"
	}
}

// WithAnsibleNocows sets the option ANSIBLE_NOCOWS to true (If you have cowsay installed but want to avoid the ‘cows’ (why????), use this.)
func WithAnsibleNocows() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleNocows] = "true"
	}
}

// WithoutAnsibleNocows sets the option ANSIBLE_NOCOWS to false
func WithoutAnsibleNocows() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleNocows] = "false"
	}
}

// WithAnsiblePipelining sets the option ANSIBLE_PIPELINING to true (This is a global option, each connection plugin can override either by having more specific options or not supporting pipelining at all. Pipelining, if supported by the connection plugin, reduces the number of network operations required to execute a module on the remote server, by executing many Ansible modules without actual file transfer. It can result in a very significant performance improvement when enabled. However this conflicts with privilege escalation (become). For example, when using ‘sudo:’ operations you must first disable ‘requiretty’ in /etc/sudoers on all managed hosts, which is why it is disabled by default. This setting will be disabled if ANSIBLE_KEEP_REMOTE_FILES is enabled.)
func WithAnsiblePipelining() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePipelining] = "true"
	}
}

// WithoutAnsiblePipelining sets the option ANSIBLE_PIPELINING to false
func WithoutAnsiblePipelining() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePipelining] = "false"
	}
}

// WithAnsibleAnyErrorsFatal sets the option ANSIBLE_ANY_ERRORS_FATAL to true (Sets the default value for the any_errors_fatal keyword, if True, Task failures will be considered fatal errors.)
func WithAnsibleAnyErrorsFatal() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleAnyErrorsFatal] = "true"
	}
}

// WithoutAnsibleAnyErrorsFatal sets the option ANSIBLE_ANY_ERRORS_FATAL to false
func WithoutAnsibleAnyErrorsFatal() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleAnyErrorsFatal] = "false"
	}
}

// WithAnsibleBecomeAllowSameUser sets the option ANSIBLE_BECOME_ALLOW_SAME_USER to true (This setting controls if become is skipped when remote user and become user are the same. I.E root sudo to root. If executable, it will be run and the resulting stdout will be used as the password.)
func WithAnsibleBecomeAllowSameUser() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecomeAllowSameUser] = "true"
	}
}

// WithoutAnsibleBecomeAllowSameUser sets the option ANSIBLE_BECOME_ALLOW_SAME_USER to false
func WithoutAnsibleBecomeAllowSameUser() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecomeAllowSameUser] = "false"
	}
}

// WithAnsibleBecomePasswordFile sets the value for the configuraion ANSIBLE_BECOME_PASSWORD_FILE (The password file to use for the become plugin. –become-password-file. If executable, it will be run and the resulting stdout will be used as the password.)
func WithAnsibleBecomePasswordFile(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecomePasswordFile] = value
	}
}

// WithAnsibleBecomePlugins sets the value for the configuraion ANSIBLE_BECOME_PLUGINS (Colon separated paths in which Ansible will search for Become Plugins.)
func WithAnsibleBecomePlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecomePlugins] = value
	}
}

// WithAnsibleCachePlugin sets the value for the configuraion ANSIBLE_CACHE_PLUGIN (Chooses which cache plugin to use, the default ‘memory’ is ephemeral.)
func WithAnsibleCachePlugin(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCachePlugin] = value
	}
}

// WithAnsibleCachePluginConnection sets the value for the configuraion ANSIBLE_CACHE_PLUGIN_CONNECTION (Defines connection or path information for the cache plugin)
func WithAnsibleCachePluginConnection(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCachePluginConnection] = value
	}
}

// WithAnsibleCachePluginPrefix sets the value for the configuraion ANSIBLE_CACHE_PLUGIN_PREFIX (Prefix to use for cache plugin files/tables)
func WithAnsibleCachePluginPrefix(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCachePluginPrefix] = value
	}
}

// WithAnsibleCachePluginTimeout sets the value for the configuraion ANSIBLE_CACHE_PLUGIN_TIMEOUT (Expiration timeout for the cache plugin data)
func WithAnsibleCachePluginTimeout(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCachePluginTimeout] = fmt.Sprint(value)
	}
}

// WithAnsibleCallbacksEnabled sets the value for the configuraion ANSIBLE_CALLBACKS_ENABLED (List of enabled callbacks, not all callbacks need enabling, but many of those shipped with Ansible do as we don’t want them activated by default. [:Version Added: 2.11])
func WithAnsibleCallbacksEnabled(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCallbacksEnabled] = value
	}
}

// WithAnsibleCollectionsOnAnsibleVersionMismatch sets the value for the configuraion ANSIBLE_COLLECTIONS_ON_ANSIBLE_VERSION_MISMATCH (When a collection is loaded that does not support the running Ansible version (with the collection metadata key requires_ansible).)
func WithAnsibleCollectionsOnAnsibleVersionMismatch(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCollectionsOnAnsibleVersionMismatch] = value
	}
}

// WithAnsibleCollectionsPaths sets the value for the configuraion ANSIBLE_COLLECTIONS_PATHS (Colon separated paths in which Ansible will search for collections content. Collections must be in nested subdirectories, not directly in these directories. For example, if COLLECTIONS_PATHS includes '{{ ANSIBLE_HOME ~ "/collections" }}', and you want to add my.collection to that directory, it must be saved as '{{ ANSIBLE_HOME} ~ "/collections/ansible_collections/my/collection" }}'.)
func WithAnsibleCollectionsPaths(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCollectionsPaths] = value
	}
}

// WithAnsibleCollectionsScanSysPath sets the option ANSIBLE_COLLECTIONS_SCAN_SYS_PATH to true (A boolean to enable or disable scanning the sys.path for installed collections)
func WithAnsibleCollectionsScanSysPath() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCollectionsScanSysPath] = "true"
	}
}

// WithoutAnsibleCollectionsScanSysPath sets the option ANSIBLE_COLLECTIONS_SCAN_SYS_PATH to false
func WithoutAnsibleCollectionsScanSysPath() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCollectionsScanSysPath] = "false"
	}
}

// WithAnsibleColorChanged sets the value for the configuraion ANSIBLE_COLOR_CHANGED (Defines the color to use on ‘Changed’ task status)
func WithAnsibleColorChanged(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorChanged] = value
	}
}

// WithAnsibleColorConsolePrompt sets the value for the configuraion ANSIBLE_COLOR_CONSOLE_PROMPT (Defines the default color to use for ansible-console)
func WithAnsibleColorConsolePrompt(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorConsolePrompt] = value
	}
}

// WithAnsibleColorDebug sets the value for the configuraion ANSIBLE_COLOR_DEBUG (Defines the color to use when emitting debug messages)
func WithAnsibleColorDebug(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorDebug] = value
	}
}

// WithAnsibleColorDeprecate sets the value for the configuraion ANSIBLE_COLOR_DEPRECATE (Defines the color to use when emitting deprecation messages)
func WithAnsibleColorDeprecate(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorDeprecate] = value
	}
}

// WithAnsibleColorDiffAdd sets the value for the configuraion ANSIBLE_COLOR_DIFF_ADD (Defines the color to use when showing added lines in diffs)
func WithAnsibleColorDiffAdd(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorDiffAdd] = value
	}
}

// WithAnsibleColorDiffLines sets the value for the configuraion ANSIBLE_COLOR_DIFF_LINES (Defines the color to use when showing diffs)
func WithAnsibleColorDiffLines(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorDiffLines] = value
	}
}

// WithAnsibleColorDiffRemove sets the value for the configuraion ANSIBLE_COLOR_DIFF_REMOVE (Defines the color to use when showing removed lines in diffs)
func WithAnsibleColorDiffRemove(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorDiffRemove] = value
	}
}

// WithAnsibleColorError sets the value for the configuraion ANSIBLE_COLOR_ERROR (Defines the color to use when emitting error messages)
func WithAnsibleColorError(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorError] = value
	}
}

// WithAnsibleColorHighlight sets the value for the configuraion ANSIBLE_COLOR_HIGHLIGHT (Defines the color to use for highlighting)
func WithAnsibleColorHighlight(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorHighlight] = value
	}
}

// WithAnsibleColorOk sets the value for the configuraion ANSIBLE_COLOR_OK (Defines the color to use when showing ‘OK’ task status)
func WithAnsibleColorOk(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorOk] = value
	}
}

// WithAnsibleColorSkip sets the value for the configuraion ANSIBLE_COLOR_SKIP (Defines the color to use when showing ‘Skipped’ task status)
func WithAnsibleColorSkip(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorSkip] = value
	}
}

// WithAnsibleColorUnreachable sets the value for the configuraion ANSIBLE_COLOR_UNREACHABLE (Defines the color to use on ‘Unreachable’ status)
func WithAnsibleColorUnreachable(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorUnreachable] = value
	}
}

// WithAnsibleColorVerbose sets the value for the configuraion ANSIBLE_COLOR_VERBOSE (Defines the color to use when emitting verbose messages. i.e those that show with ‘-v’s.)
func WithAnsibleColorVerbose(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorVerbose] = value
	}
}

// WithAnsibleColorWarn sets the value for the configuraion ANSIBLE_COLOR_WARN (Defines the color to use when emitting warning messages)
func WithAnsibleColorWarn(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleColorWarn] = value
	}
}

// WithAnsibleConnectionPasswordFile sets the value for the configuraion ANSIBLE_CONNECTION_PASSWORD_FILE (The password file to use for the connection plugin. –connection-password-file.)
func WithAnsibleConnectionPasswordFile(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleConnectionPasswordFile] = value
	}
}

// WithAnsibleCoverageRemoteOutput sets the value for the configuraion _ANSIBLE_COVERAGE_REMOTE_OUTPUT (Sets the output directory on the remote host to generate coverage reports to. Currently only used for remote coverage on PowerShell modules. This is for internal use only.)
func WithAnsibleCoverageRemoteOutput(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCoverageRemoteOutput] = value
	}
}

// WithAnsibleCoverageRemotePathFilter sets the value for the configuraion _ANSIBLE_COVERAGE_REMOTE_PATH_FILTER (A list of paths for files on the Ansible controller to run coverage for when executing on the remote host. Only files that match the path glob will have its coverage collected. Multiple path globs can be specified and are separated by :. Currently only used for remote coverage on PowerShell modules. This is for internal use only.)
func WithAnsibleCoverageRemotePathFilter(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCoverageRemotePathFilter] = value
	}
}

// WithAnsibleActionPlugins sets the value for the configuraion ANSIBLE_ACTION_PLUGINS (Colon separated paths in which Ansible will search for Action Plugins.)
func WithAnsibleActionPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleActionPlugins] = value
	}
}

// WithAnsibleAskPass sets the option ANSIBLE_ASK_PASS to true (This controls whether an Ansible playbook should prompt for a login password. If using SSH keys for authentication, you probably do not need to change this setting.)
func WithAnsibleAskPass() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleAskPass] = "true"
	}
}

// WithoutAnsibleAskPass sets the option ANSIBLE_ASK_PASS to false
func WithoutAnsibleAskPass() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleAskPass] = "false"
	}
}

// WithAnsibleAskVaultPass sets the option ANSIBLE_ASK_VAULT_PASS to true (This controls whether an Ansible playbook should prompt for a vault password.)
func WithAnsibleAskVaultPass() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleAskVaultPass] = "true"
	}
}

// WithoutAnsibleAskVaultPass sets the option ANSIBLE_ASK_VAULT_PASS to false
func WithoutAnsibleAskVaultPass() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleAskVaultPass] = "false"
	}
}

// WithAnsibleBecome sets the option ANSIBLE_BECOME to true (Toggles the use of privilege escalation, allowing you to ‘become’ another user after login.)
func WithAnsibleBecome() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecome] = "true"
	}
}

// WithoutAnsibleBecome sets the option ANSIBLE_BECOME to false
func WithoutAnsibleBecome() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecome] = "false"
	}
}

// WithAnsibleBecomeAskPass sets the option ANSIBLE_BECOME_ASK_PASS to true (Toggle to prompt for privilege escalation password.)
func WithAnsibleBecomeAskPass() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecomeAskPass] = "true"
	}
}

// WithoutAnsibleBecomeAskPass sets the option ANSIBLE_BECOME_ASK_PASS to false
func WithoutAnsibleBecomeAskPass() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecomeAskPass] = "false"
	}
}

// WithAnsibleBecomeExe sets the value for the configuraion ANSIBLE_BECOME_EXE (executable to use for privilege escalation, otherwise Ansible will depend on PATH)
func WithAnsibleBecomeExe(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecomeExe] = value
	}
}

// WithAnsibleBecomeFlags sets the value for the configuraion ANSIBLE_BECOME_FLAGS (Flags to pass to the privilege escalation executable.)
func WithAnsibleBecomeFlags(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecomeFlags] = value
	}
}

// WithAnsibleBecomeMethod sets the value for the configuraion ANSIBLE_BECOME_METHOD (Privilege escalation method to use when become is enabled.)
func WithAnsibleBecomeMethod(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecomeMethod] = value
	}
}

// WithAnsibleBecomeUser sets the value for the configuraion ANSIBLE_BECOME_USER (The user your login/remote user ‘becomes’ when using privilege escalation, most systems will use ‘root’ when no user is specified.)
func WithAnsibleBecomeUser(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleBecomeUser] = value
	}
}

// WithAnsibleCachePlugins sets the value for the configuraion ANSIBLE_CACHE_PLUGINS (Colon separated paths in which Ansible will search for Cache Plugins.)
func WithAnsibleCachePlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCachePlugins] = value
	}
}

// WithAnsibleCallbackPlugins sets the value for the configuraion ANSIBLE_CALLBACK_PLUGINS (Colon separated paths in which Ansible will search for Callback Plugins.)
func WithAnsibleCallbackPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCallbackPlugins] = value
	}
}

// WithAnsibleCliconfPlugins sets the value for the configuraion ANSIBLE_CLICONF_PLUGINS (Colon separated paths in which Ansible will search for Cliconf Plugins.)
func WithAnsibleCliconfPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleCliconfPlugins] = value
	}
}

// WithAnsibleConnectionPlugins sets the value for the configuraion ANSIBLE_CONNECTION_PLUGINS (Colon separated paths in which Ansible will search for Connection Plugins.)
func WithAnsibleConnectionPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleConnectionPlugins] = value
	}
}

// WithAnsibleDebug sets the option ANSIBLE_DEBUG to true (Toggles debug output in Ansible. This is very verbose and can hinder multiprocessing.  Debug output can also include secret information despite no_log settings being enabled, which means debug mode should not be used in production.)
func WithAnsibleDebug() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDebug] = "true"
	}
}

// WithoutAnsibleDebug sets the option ANSIBLE_DEBUG to false
func WithoutAnsibleDebug() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDebug] = "false"
	}
}

// WithAnsibleExecutable sets the value for the configuraion ANSIBLE_EXECUTABLE (This indicates the command to use to spawn a shell under for Ansible’s execution needs on a target. Users may need to change this in rare instances when shell usage is constrained, but in most cases it may be left as is.)
func WithAnsibleExecutable(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleExecutable] = value
	}
}

// WithAnsibleFactPath sets the value for the configuraion ANSIBLE_FACT_PATH (This option allows you to globally configure a custom path for ‘local_facts’ for the implied ansible_collections.ansible.builtin.setup_module task when using fact gathering. If not set, it will fallback to the default from the ansible.builtin.setup module: /etc/ansible/facts.d. This does not affect  user defined tasks that use the ansible.builtin.setup module. The real action being created by the implicit task is currently    ansible.legacy.gather_facts module, which then calls the configured fact modules, by default this will be ansible.builtin.setup for POSIX systems but other platforms might have different defaults.)
func WithAnsibleFactPath(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleFactPath] = value
	}
}

// WithAnsibleFilterPlugins sets the value for the configuraion ANSIBLE_FILTER_PLUGINS (Colon separated paths in which Ansible will search for Jinja2 Filter Plugins.)
func WithAnsibleFilterPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleFilterPlugins] = value
	}
}

// WithAnsibleForceHandlers sets the option ANSIBLE_FORCE_HANDLERS to true (This option controls if notified handlers run on a host even if a failure occurs on that host. When false, the handlers will not run if a failure has occurred on a host. This can also be set per play or on the command line. See Handlers and Failure for more details.)
func WithAnsibleForceHandlers() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleForceHandlers] = "true"
	}
}

// WithoutAnsibleForceHandlers sets the option ANSIBLE_FORCE_HANDLERS to false
func WithoutAnsibleForceHandlers() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleForceHandlers] = "false"
	}
}

// WithAnsibleForks sets the value for the configuraion ANSIBLE_FORKS (Maximum number of forks Ansible will use to execute tasks on target hosts.)
func WithAnsibleForks(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleForks] = fmt.Sprint(value)
	}
}

// WithAnsibleGatherSubset sets the value for the configuraion ANSIBLE_GATHER_SUBSET (Set the gather_subset option for the ansible_collections.ansible.builtin.setup_module task in the implicit fact gathering. See the module documentation for specifics. It does not apply to user defined ansible.builtin.setup tasks.)
func WithAnsibleGatherSubset(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGatherSubset] = value
	}
}

// WithAnsibleGatherTimeout sets the value for the configuraion ANSIBLE_GATHER_TIMEOUT (Set the timeout in seconds for the implicit fact gathering, see the module documentation for specifics. It does not apply to user defined ansible_collections.ansible.builtin.setup_module tasks.)
func WithAnsibleGatherTimeout(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGatherTimeout] = fmt.Sprint(value)
	}
}

// WithAnsibleGathering sets the value for the configuraion ANSIBLE_GATHERING (This setting controls the default policy of fact gathering (facts discovered about remote systems). This option can be useful for those wishing to save fact gathering time. Both ‘smart’ and ‘explicit’ will use the cache plugin.)
func WithAnsibleGathering(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGathering] = value
	}
}

// WithAnsibleHashBehaviour sets the value for the configuraion ANSIBLE_HASH_BEHAVIOUR (This setting controls how duplicate definitions of dictionary variables (aka hash, map, associative array) are handled in Ansible. This does not affect variables whose values are scalars (integers, strings) or arrays. WARNING, changing this setting is not recommended as this is fragile and makes your content (plays, roles, collections) non portable, leading to continual confusion and misuse. Don’t change this setting unless you think you have an absolute need for it. We recommend avoiding reusing variable names and relying on the combine filter and vars and varnames lookups to create merged versions of the individual variables. In our experience this is rarely really needed and a sign that too much complexity has been introduced into the data structures and plays. For some uses you can also look into custom vars_plugins to merge on input, even substituting the default host_group_vars that is in charge of parsing the host_vars/ and group_vars/ directories. Most users of this setting are only interested in inventory scope, but the setting itself affects all sources and makes debugging even harder. All playbooks and roles in the official examples repos assume the default for this setting. Changing the setting to merge applies across variable sources, but many sources will internally still overwrite the variables. For example include_vars will dedupe variables internally before updating Ansible, with ‘last defined’ overwriting previous definitions in same file. The Ansible project recommends you avoid “merge“ for new projects. It is the intention of the Ansible developers to eventually deprecate and remove this setting, but it is being kept as some users do heavily rely on it. New projects should avoid ‘merge’.)
func WithAnsibleHashBehaviour(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleHashBehaviour] = value
	}
}

// WithAnsibleInventory sets the value for the configuraion ANSIBLE_INVENTORY (Comma separated list of Ansible inventory sources)
func WithAnsibleInventory(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventory] = value
	}
}

// WithAnsibleHttpapiPlugins sets the value for the configuraion ANSIBLE_HTTPAPI_PLUGINS (Colon separated paths in which Ansible will search for HttpApi Plugins.)
func WithAnsibleHttpapiPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleHttpapiPlugins] = value
	}
}

// WithAnsibleInventoryPlugins sets the value for the configuraion ANSIBLE_INVENTORY_PLUGINS (Colon separated paths in which Ansible will search for Inventory Plugins.)
func WithAnsibleInventoryPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryPlugins] = value
	}
}

// WithAnsibleJinja2Extensions sets the value for the configuraion ANSIBLE_JINJA2_EXTENSIONS (This is a developer-specific feature that allows enabling additional Jinja2 extensions. See the Jinja2 documentation for details. If you do not know what these do, you probably don’t need to change this setting :))
func WithAnsibleJinja2Extensions(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleJinja2Extensions] = value
	}
}

// WithAnsibleJinja2Native sets the option ANSIBLE_JINJA2_NATIVE to true (This option preserves variable types during template operations.)
func WithAnsibleJinja2Native() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleJinja2Native] = "true"
	}
}

// WithoutAnsibleJinja2Native sets the option ANSIBLE_JINJA2_NATIVE to false
func WithoutAnsibleJinja2Native() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleJinja2Native] = "false"
	}
}

// WithAnsibleKeepRemoteFiles sets the option ANSIBLE_KEEP_REMOTE_FILES to true (Enables/disables the cleaning up of the temporary files Ansible used to execute the tasks on the remote. If this option is enabled it will disable ANSIBLE_PIPELINING.)
func WithAnsibleKeepRemoteFiles() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleKeepRemoteFiles] = "true"
	}
}

// WithoutAnsibleKeepRemoteFiles sets the option ANSIBLE_KEEP_REMOTE_FILES to false
func WithoutAnsibleKeepRemoteFiles() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleKeepRemoteFiles] = "false"
	}
}

// WithAnsibleLibvirtLxcNoseclabel sets the option ANSIBLE_LIBVIRT_LXC_NOSECLABEL to true (This setting causes libvirt to connect to lxc containers by passing –noseclabel to virsh. This is necessary when running on systems which do not have SELinux.)
func WithAnsibleLibvirtLxcNoseclabel() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleLibvirtLxcNoseclabel] = "true"
	}
}

// WithoutAnsibleLibvirtLxcNoseclabel sets the option ANSIBLE_LIBVIRT_LXC_NOSECLABEL to false
func WithoutAnsibleLibvirtLxcNoseclabel() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleLibvirtLxcNoseclabel] = "false"
	}
}

// WithAnsibleLoadCallbackPlugins sets the option ANSIBLE_LOAD_CALLBACK_PLUGINS to true (Controls whether callback plugins are loaded when running /usr/bin/ansible. This may be used to log activity from the command line, send notifications, and so on. Callback plugins are always loaded for ansible-playbook.)
func WithAnsibleLoadCallbackPlugins() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleLoadCallbackPlugins] = "true"
	}
}

// WithoutAnsibleLoadCallbackPlugins sets the option ANSIBLE_LOAD_CALLBACK_PLUGINS to false
func WithoutAnsibleLoadCallbackPlugins() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleLoadCallbackPlugins] = "false"
	}
}

// WithAnsibleLocalTemp sets the value for the configuraion ANSIBLE_LOCAL_TEMP (Temporary directory for Ansible to use on the controller.)
func WithAnsibleLocalTemp(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleLocalTemp] = value
	}
}

// WithAnsibleLogFilter sets the value for the configuraion ANSIBLE_LOG_FILTER (List of logger names to filter out of the log file)
func WithAnsibleLogFilter(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleLogFilter] = value
	}
}

// WithAnsibleLogPath sets the value for the configuraion ANSIBLE_LOG_PATH (File to which Ansible will log on the controller. When empty logging is disabled.)
func WithAnsibleLogPath(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleLogPath] = value
	}
}

// WithAnsibleLookupPlugins sets the value for the configuraion ANSIBLE_LOOKUP_PLUGINS (Colon separated paths in which Ansible will search for Lookup Plugins.)
func WithAnsibleLookupPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleLookupPlugins] = value
	}
}

// WithAnsibleModuleArgs sets the value for the configuraion ANSIBLE_MODULE_ARGS (This sets the default arguments to pass to the ansible adhoc binary if no -a is specified.)
func WithAnsibleModuleArgs(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleModuleArgs] = value
	}
}

// WithAnsibleLibrary sets the value for the configuraion ANSIBLE_LIBRARY (Colon separated paths in which Ansible will search for Modules.)
func WithAnsibleLibrary(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleLibrary] = value
	}
}

// WithAnsibleModuleUtils sets the value for the configuraion ANSIBLE_MODULE_UTILS (Colon separated paths in which Ansible will search for Module utils files, which are shared by modules.)
func WithAnsibleModuleUtils(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleModuleUtils] = value
	}
}

// WithAnsibleNetconfPlugins sets the value for the configuraion ANSIBLE_NETCONF_PLUGINS (Colon separated paths in which Ansible will search for Netconf Plugins.)
func WithAnsibleNetconfPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleNetconfPlugins] = value
	}
}

// WithAnsibleNoLog sets the option ANSIBLE_NO_LOG to true (Toggle Ansible’s display and logging of task details, mainly used to avoid security disclosures.)
func WithAnsibleNoLog() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleNoLog] = "true"
	}
}

// WithoutAnsibleNoLog sets the option ANSIBLE_NO_LOG to false
func WithoutAnsibleNoLog() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleNoLog] = "false"
	}
}

// WithAnsibleNoTargetSyslog sets the option ANSIBLE_NO_TARGET_SYSLOG to true (Toggle Ansible logging to syslog on the target when it executes tasks. On Windows hosts this will disable a newer style PowerShell modules from writing to the event log.)
func WithAnsibleNoTargetSyslog() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleNoTargetSyslog] = "true"
	}
}

// WithoutAnsibleNoTargetSyslog sets the option ANSIBLE_NO_TARGET_SYSLOG to false
func WithoutAnsibleNoTargetSyslog() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleNoTargetSyslog] = "false"
	}
}

// WithAnsibleNullRepresentation sets the value for the configuraion ANSIBLE_NULL_REPRESENTATION (What templating should return as a ‘null’ value. When not set it will let Jinja2 decide.)
func WithAnsibleNullRepresentation(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleNullRepresentation] = value
	}
}

// WithAnsiblePollInterval sets the value for the configuraion ANSIBLE_POLL_INTERVAL (For asynchronous tasks in Ansible (covered in Asynchronous Actions and Polling), this is how often to check back on the status of those tasks when an explicit poll interval is not supplied. The default is a reasonably moderate 15 seconds which is a tradeoff between checking in frequently and providing a quick turnaround when something may have completed.)
func WithAnsiblePollInterval(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePollInterval] = fmt.Sprint(value)
	}
}

// WithAnsiblePrivateKeyFile sets the value for the configuraion ANSIBLE_PRIVATE_KEY_FILE (Option for connections using a certificate or key file to authenticate, rather than an agent or passwords, you can set the default value here to avoid re-specifying –private-key with every invocation.)
func WithAnsiblePrivateKeyFile(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePrivateKeyFile] = value
	}
}

// WithAnsiblePrivateRoleVars sets the option ANSIBLE_PRIVATE_ROLE_VARS to true (By default, imported roles publish their variables to the play and other roles, this setting can avoid that. This was introduced as a way to reset role variables to default values if a role is used more than once in a playbook. Included roles only make their variables public at execution, unlike imported roles which happen at playbook compile time.)
func WithAnsiblePrivateRoleVars() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePrivateRoleVars] = "true"
	}
}

// WithoutAnsiblePrivateRoleVars sets the option ANSIBLE_PRIVATE_ROLE_VARS to false
func WithoutAnsiblePrivateRoleVars() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePrivateRoleVars] = "false"
	}
}

// WithAnsibleRemotePort sets the value for the configuraion ANSIBLE_REMOTE_PORT (Port to use in remote connections, when blank it will use the connection plugin default.)
func WithAnsibleRemotePort(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleRemotePort] = fmt.Sprint(value)
	}
}

// WithAnsibleRemoteUser sets the value for the configuraion ANSIBLE_REMOTE_USER (Sets the login user for the target machines When blank it uses the connection plugin’s default, normally the user currently executing Ansible.)
func WithAnsibleRemoteUser(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleRemoteUser] = value
	}
}

// WithAnsibleRolesPath sets the value for the configuraion ANSIBLE_ROLES_PATH (Colon separated paths in which Ansible will search for Roles.)
func WithAnsibleRolesPath(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleRolesPath] = value
	}
}

// WithAnsibleSelinuxSpecialFs sets the value for the configuraion ANSIBLE_SELINUX_SPECIAL_FS (Some filesystems do not support safe operations and/or return inconsistent errors, this setting makes Ansible ‘tolerate’ those in the list w/o causing fatal errors. Data corruption may occur and writes are not always verified when a filesystem is in the list. [:Version Added: 2.9])
func WithAnsibleSelinuxSpecialFs(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleSelinuxSpecialFs] = value
	}
}

// WithAnsibleStdoutCallback sets the value for the configuraion ANSIBLE_STDOUT_CALLBACK (Set the main callback used to display Ansible output. You can only have one at a time. You can have many other callbacks, but just one can be in charge of stdout. See Callback plugins for a list of available options.)
func WithAnsibleStdoutCallback(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleStdoutCallback] = value
	}
}

// WithAnsibleStrategy sets the value for the configuraion ANSIBLE_STRATEGY (Set the default strategy used for plays.)
func WithAnsibleStrategy(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleStrategy] = value
	}
}

// WithAnsibleStrategyPlugins sets the value for the configuraion ANSIBLE_STRATEGY_PLUGINS (Colon separated paths in which Ansible will search for Strategy Plugins.)
func WithAnsibleStrategyPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleStrategyPlugins] = value
	}
}

// WithAnsibleSu sets the option ANSIBLE_SU to true (Toggle the use of “su” for tasks.)
func WithAnsibleSu() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleSu] = "true"
	}
}

// WithoutAnsibleSu sets the option ANSIBLE_SU to false
func WithoutAnsibleSu() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleSu] = "false"
	}
}

// WithAnsibleSyslogFacility sets the value for the configuraion ANSIBLE_SYSLOG_FACILITY (Syslog facility to use when Ansible logs to the remote target)
func WithAnsibleSyslogFacility(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleSyslogFacility] = value
	}
}

// WithAnsibleTerminalPlugins sets the value for the configuraion ANSIBLE_TERMINAL_PLUGINS (Colon separated paths in which Ansible will search for Terminal Plugins.)
func WithAnsibleTerminalPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleTerminalPlugins] = value
	}
}

// WithAnsibleTestPlugins sets the value for the configuraion ANSIBLE_TEST_PLUGINS (Colon separated paths in which Ansible will search for Jinja2 Test Plugins.)
func WithAnsibleTestPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleTestPlugins] = value
	}
}

// WithAnsibleTimeout sets the value for the configuraion ANSIBLE_TIMEOUT (This is the default timeout for connection plugins to use.)
func WithAnsibleTimeout(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleTimeout] = fmt.Sprint(value)
	}
}

// WithAnsibleTransport sets the value for the configuraion ANSIBLE_TRANSPORT (Can be any connection plugin available to your ansible installation. There is also a (DEPRECATED) special ‘smart’ option, that will toggle between ‘ssh’ and ‘paramiko’ depending on controller OS and ssh versions.)
func WithAnsibleTransport(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleTransport] = value
	}
}

// WithAnsibleErrorOnUndefinedVars sets the option ANSIBLE_ERROR_ON_UNDEFINED_VARS to true (When True, this causes ansible templating to fail steps that reference variable names that are likely typoed. Otherwise, any ‘{{ template_expression }}’ that contains undefined variables will be rendered in a template or ansible action line exactly as written.)
func WithAnsibleErrorOnUndefinedVars() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleErrorOnUndefinedVars] = "true"
	}
}

// WithoutAnsibleErrorOnUndefinedVars sets the option ANSIBLE_ERROR_ON_UNDEFINED_VARS to false
func WithoutAnsibleErrorOnUndefinedVars() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleErrorOnUndefinedVars] = "false"
	}
}

// WithAnsibleVarsPlugins sets the value for the configuraion ANSIBLE_VARS_PLUGINS (Colon separated paths in which Ansible will search for Vars Plugins.)
func WithAnsibleVarsPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleVarsPlugins] = value
	}
}

// WithAnsibleVaultEncryptIdentity sets the value for the configuraion ANSIBLE_VAULT_ENCRYPT_IDENTITY (The vault_id to use for encrypting by default. If multiple vault_ids are provided, this specifies which to use for encryption. The –encrypt-vault-id cli option overrides the configured value.)
func WithAnsibleVaultEncryptIdentity(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleVaultEncryptIdentity] = value
	}
}

// WithAnsibleVaultIdMatch sets the value for the configuraion ANSIBLE_VAULT_ID_MATCH (If true, decrypting vaults with a vault id will only try the password from the matching vault-id)
func WithAnsibleVaultIdMatch(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleVaultIdMatch] = value
	}
}

// WithAnsibleVaultIdentity sets the value for the configuraion ANSIBLE_VAULT_IDENTITY (The label to use for the default vault id label in cases where a vault id label is not provided)
func WithAnsibleVaultIdentity(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleVaultIdentity] = value
	}
}

// WithAnsibleVaultIdentityList sets the value for the configuraion ANSIBLE_VAULT_IDENTITY_LIST (A list of vault-ids to use by default. Equivalent to multiple –vault-id args. Vault-ids are tried in order.)
func WithAnsibleVaultIdentityList(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleVaultIdentityList] = value
	}
}

// WithAnsibleVaultPasswordFile sets the value for the configuraion ANSIBLE_VAULT_PASSWORD_FILE (The vault password file to use. Equivalent to –vault-password-file or –vault-id If executable, it will be run and the resulting stdout will be used as the password.)
func WithAnsibleVaultPasswordFile(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleVaultPasswordFile] = value
	}
}

// WithAnsibleVerbosity sets the value for the configuraion ANSIBLE_VERBOSITY (Sets the default verbosity, equivalent to the number of -v passed in the command line.)
func WithAnsibleVerbosity(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleVerbosity] = fmt.Sprint(value)
	}
}

// WithAnsibleDeprecationWarnings sets the option ANSIBLE_DEPRECATION_WARNINGS to true (Toggle to control the showing of deprecation warnings)
func WithAnsibleDeprecationWarnings() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDeprecationWarnings] = "true"
	}
}

// WithoutAnsibleDeprecationWarnings sets the option ANSIBLE_DEPRECATION_WARNINGS to false
func WithoutAnsibleDeprecationWarnings() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDeprecationWarnings] = "false"
	}
}

// WithAnsibleDevelWarning sets the option ANSIBLE_DEVEL_WARNING to true (Toggle to control showing warnings related to running devel)
func WithAnsibleDevelWarning() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDevelWarning] = "true"
	}
}

// WithoutAnsibleDevelWarning sets the option ANSIBLE_DEVEL_WARNING to false
func WithoutAnsibleDevelWarning() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDevelWarning] = "false"
	}
}

// WithAnsibleDiffAlways sets the value for the configuraion ANSIBLE_DIFF_ALWAYS (Configuration toggle to tell modules to show differences when in ‘changed’ status, equivalent to --diff.)
func WithAnsibleDiffAlways(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDiffAlways] = value
	}
}

// WithAnsibleDiffContext sets the value for the configuraion ANSIBLE_DIFF_CONTEXT (How many lines of context to show when displaying the differences between files.)
func WithAnsibleDiffContext(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDiffContext] = fmt.Sprint(value)
	}
}

// WithAnsibleDisplayArgsToStdout sets the option ANSIBLE_DISPLAY_ARGS_TO_STDOUT to true (Normally ansible-playbook will print a header for each task that is run. These headers will contain the name: field from the task if you specified one. If you didn’t then ansible-playbook uses the task’s action to help you tell which task is presently running. Sometimes you run many of the same action and so you want more information about the task to differentiate it from others of the same action. If you set this variable to True in the config then ansible-playbook will also include the task’s arguments in the header. This setting defaults to False because there is a chance that you have sensitive values in your parameters and you do not want those to be printed. If you set this to True you should be sure that you have secured your environment’s stdout (no one can shoulder surf your screen and you aren’t saving stdout to an insecure file) or made sure that all of your playbooks explicitly added the no_log: True parameter to tasks which have sensitive values See How do I keep secret data in my playbook? for more information.)
func WithAnsibleDisplayArgsToStdout() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDisplayArgsToStdout] = "true"
	}
}

// WithoutAnsibleDisplayArgsToStdout sets the option ANSIBLE_DISPLAY_ARGS_TO_STDOUT to false
func WithoutAnsibleDisplayArgsToStdout() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDisplayArgsToStdout] = "false"
	}
}

// WithAnsibleDisplaySkippedHosts sets the option ANSIBLE_DISPLAY_SKIPPED_HOSTS to true (Toggle to control displaying skipped task/host entries in a task in the default callback)
func WithAnsibleDisplaySkippedHosts() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDisplaySkippedHosts] = "true"
	}
}

// WithoutAnsibleDisplaySkippedHosts sets the option ANSIBLE_DISPLAY_SKIPPED_HOSTS to false
func WithoutAnsibleDisplaySkippedHosts() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDisplaySkippedHosts] = "false"
	}
}

// WithAnsibleDocFragmentPlugins sets the value for the configuraion ANSIBLE_DOC_FRAGMENT_PLUGINS (Colon separated paths in which Ansible will search for Documentation Fragments Plugins.)
func WithAnsibleDocFragmentPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDocFragmentPlugins] = value
	}
}

// WithAnsibleDuplicateYamlDictKey sets the value for the configuraion ANSIBLE_DUPLICATE_YAML_DICT_KEY (By default Ansible will issue a warning when a duplicate dict key is encountered in YAML. These warnings can be silenced by adjusting this setting to False.)
func WithAnsibleDuplicateYamlDictKey(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleDuplicateYamlDictKey] = value
	}
}

// WithEditor sets the value for the configuraion EDITOR ()
func WithEditor(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[Editor] = value
	}
}

// WithAnsibleEnableTaskDebugger sets the option ANSIBLE_ENABLE_TASK_DEBUGGER to true (Whether or not to enable the task debugger, this previously was done as a strategy plugin. Now all strategy plugins can inherit this behavior. The debugger defaults to activating when a task is failed on unreachable. Use the debugger keyword for more flexibility.)
func WithAnsibleEnableTaskDebugger() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleEnableTaskDebugger] = "true"
	}
}

// WithoutAnsibleEnableTaskDebugger sets the option ANSIBLE_ENABLE_TASK_DEBUGGER to false
func WithoutAnsibleEnableTaskDebugger() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleEnableTaskDebugger] = "false"
	}
}

// WithAnsibleErrorOnMissingHandler sets the option ANSIBLE_ERROR_ON_MISSING_HANDLER to true (Toggle to allow missing handlers to become a warning instead of an error when notifying.)
func WithAnsibleErrorOnMissingHandler() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleErrorOnMissingHandler] = "true"
	}
}

// WithoutAnsibleErrorOnMissingHandler sets the option ANSIBLE_ERROR_ON_MISSING_HANDLER to false
func WithoutAnsibleErrorOnMissingHandler() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleErrorOnMissingHandler] = "false"
	}
}

// WithAnsibleFactsModules sets the value for the configuraion ANSIBLE_FACTS_MODULES (Which modules to run during a play’s fact gathering stage, using the default of ‘smart’ will try to figure it out based on connection type. If adding your own modules but you still want to use the default Ansible facts, you will want to include ‘setup’ or corresponding network module to the list (if you add ‘smart’, Ansible will also figure it out). This does not affect explicit calls to the ‘setup’ module, but does always affect the ‘gather_facts’ action (implicit or explicit).)
func WithAnsibleFactsModules(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleFactsModules] = value
	}
}

// WithAnsibleGalaxyCacheDir sets the value for the configuraion ANSIBLE_GALAXY_CACHE_DIR (The directory that stores cached responses from a Galaxy server. This is only used by the ansible-galaxy collection install and download commands. Cache files inside this dir will be ignored if they are world writable.)
func WithAnsibleGalaxyCacheDir(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyCacheDir] = value
	}
}

// WithAnsibleGalaxyCollectionSkeleton sets the value for the configuraion ANSIBLE_GALAXY_COLLECTION_SKELETON (Collection skeleton directory to use as a template for the init action in ansible-galaxy collection, same as --collection-skeleton.)
func WithAnsibleGalaxyCollectionSkeleton(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyCollectionSkeleton] = value
	}
}

// WithAnsibleGalaxyCollectionSkeletonIgnore sets the value for the configuraion ANSIBLE_GALAXY_COLLECTION_SKELETON_IGNORE (patterns of files to ignore inside a Galaxy collection skeleton directory)
func WithAnsibleGalaxyCollectionSkeletonIgnore(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyCollectionSkeletonIgnore] = value
	}
}

// WithAnsibleGalaxyCollectionsPathWarning sets the value for the configuraion ANSIBLE_GALAXY_COLLECTIONS_PATH_WARNING (whether ansible-galaxy collection install should warn about --collections-path missing from configured COLLECTIONS_PATHS)
func WithAnsibleGalaxyCollectionsPathWarning(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyCollectionsPathWarning] = value
	}
}

// WithAnsibleGalaxyDisableGpgVerify sets the value for the configuraion ANSIBLE_GALAXY_DISABLE_GPG_VERIFY (Disable GPG signature verification during collection installation.)
func WithAnsibleGalaxyDisableGpgVerify(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyDisableGpgVerify] = value
	}
}

// WithAnsibleGalaxyDisplayProgress sets the value for the configuraion ANSIBLE_GALAXY_DISPLAY_PROGRESS (Some steps in ansible-galaxy display a progress wheel which can cause issues on certain displays or when outputting the stdout to a file. This config option controls whether the display wheel is shown or not. The default is to show the display wheel if stdout has a tty.)
func WithAnsibleGalaxyDisplayProgress(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyDisplayProgress] = value
	}
}

// WithAnsibleGalaxyGpgKeyring sets the value for the configuraion ANSIBLE_GALAXY_GPG_KEYRING (Configure the keyring used for GPG signature verification during collection installation and verification.)
func WithAnsibleGalaxyGpgKeyring(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyGpgKeyring] = value
	}
}

// WithAnsibleGalaxyIgnore sets the option ANSIBLE_GALAXY_IGNORE to true (If set to yes, ansible-galaxy will not validate TLS certificates. This can be useful for testing against a server with a self-signed certificate.)
func WithAnsibleGalaxyIgnore() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyIgnore] = "true"
	}
}

// WithoutAnsibleGalaxyIgnore sets the option ANSIBLE_GALAXY_IGNORE to false
func WithoutAnsibleGalaxyIgnore() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyIgnore] = "false"
	}
}

// WithAnsibleGalaxyIgnoreSignatureStatusCodes sets the value for the configuraion ANSIBLE_GALAXY_IGNORE_SIGNATURE_STATUS_CODES (A list of GPG status codes to ignore during GPG signature verification. See L(https://github.com/gpg/gnupg/blob/master/doc/DETAILS#general-status-codes) for status code descriptions. If fewer signatures successfully verify the collection than GALAXY_REQUIRED_VALID_SIGNATURE_COUNT, signature verification will fail even if all error codes are ignored.)
func WithAnsibleGalaxyIgnoreSignatureStatusCodes(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyIgnoreSignatureStatusCodes] = value
	}
}

// WithAnsibleGalaxyRequiredValidSignatureCount sets the value for the configuraion ANSIBLE_GALAXY_REQUIRED_VALID_SIGNATURE_COUNT (The number of signatures that must be successful during GPG signature verification while installing or verifying collections. This should be a positive integer or all to indicate all signatures must successfully validate the collection. Prepend + to the value to fail if no valid signatures are found for the collection.)
func WithAnsibleGalaxyRequiredValidSignatureCount(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyRequiredValidSignatureCount] = value
	}
}

// WithAnsibleGalaxyRoleSkeleton sets the value for the configuraion ANSIBLE_GALAXY_ROLE_SKELETON (Role skeleton directory to use as a template for the init action in ansible-galaxy/ansible-galaxy role, same as --role-skeleton.)
func WithAnsibleGalaxyRoleSkeleton(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyRoleSkeleton] = value
	}
}

// WithAnsibleGalaxyRoleSkeletonIgnore sets the value for the configuraion ANSIBLE_GALAXY_ROLE_SKELETON_IGNORE (patterns of files to ignore inside a Galaxy role or collection skeleton directory)
func WithAnsibleGalaxyRoleSkeletonIgnore(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyRoleSkeletonIgnore] = value
	}
}

// WithAnsibleGalaxyServer sets the value for the configuraion ANSIBLE_GALAXY_SERVER (URL to prepend when roles don’t specify the full URI, assume they are referencing this server as the source.)
func WithAnsibleGalaxyServer(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyServer] = value
	}
}

// WithAnsibleGalaxyServerList sets the value for the configuraion ANSIBLE_GALAXY_SERVER_LIST (A list of Galaxy servers to use when installing a collection. The value corresponds to the config ini header [galaxy_server.{{item}}] which defines the server details. See Configuring the ansible-galaxy client for more details on how to define a Galaxy server. The order of servers in this list is used to as the order in which a collection is resolved. Setting this config option will ignore the GALAXY_SERVER config option.)
func WithAnsibleGalaxyServerList(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyServerList] = value
	}
}

// WithAnsibleGalaxyServerTimeout sets the value for the configuraion ANSIBLE_GALAXY_SERVER_TIMEOUT (The default timeout for Galaxy API calls. Galaxy servers that don’t configure a specific timeout will fall back to this value.)
func WithAnsibleGalaxyServerTimeout(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyServerTimeout] = value
	}
}

// WithAnsibleGalaxyTokenPath sets the value for the configuraion ANSIBLE_GALAXY_TOKEN_PATH (Local path to galaxy access token file)
func WithAnsibleGalaxyTokenPath(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleGalaxyTokenPath] = value
	}
}

// WithAnsibleHostKeyChecking sets the option ANSIBLE_HOST_KEY_CHECKING to true (Set this to “False” if you want to avoid host key checking by the underlying tools Ansible uses to connect to the host)
func WithAnsibleHostKeyChecking() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleHostKeyChecking] = "true"
	}
}

// WithoutAnsibleHostKeyChecking sets the option ANSIBLE_HOST_KEY_CHECKING to false
func WithoutAnsibleHostKeyChecking() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleHostKeyChecking] = "false"
	}
}

// WithAnsibleHostPatternMismatch sets the value for the configuraion ANSIBLE_HOST_PATTERN_MISMATCH (This setting changes the behaviour of mismatched host patterns, it allows you to force a fatal error, a warning or just ignore it)
func WithAnsibleHostPatternMismatch(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleHostPatternMismatch] = value
	}
}

// WithAnsibleInjectFactVars sets the option ANSIBLE_INJECT_FACT_VARS to true (Facts are available inside the ansible_facts variable, this setting also pushes them as their own vars in the main namespace. Unlike inside the ansible_facts dictionary, these will have an ansible_ prefix.)
func WithAnsibleInjectFactVars() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInjectFactVars] = "true"
	}
}

// WithoutAnsibleInjectFactVars sets the option ANSIBLE_INJECT_FACT_VARS to false
func WithoutAnsibleInjectFactVars() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInjectFactVars] = "false"
	}
}

// WithAnsiblePythonInterpreter sets the value for the configuraion ANSIBLE_PYTHON_INTERPRETER (Path to the Python interpreter to be used for module execution on remote targets, or an automatic discovery mode. Supported discovery modes are auto (the default), auto_silent, auto_legacy, and auto_legacy_silent. All discovery modes employ a lookup table to use the included system Python (on distributions known to include one), falling back to a fixed ordered list of well-known Python interpreter locations if a platform-specific default is not available. The fallback behavior will issue a warning that the interpreter should be set explicitly (since interpreters installed later may change which one is used). This warning behavior can be disabled by setting auto_silent or auto_legacy_silent. The value of auto_legacy provides all the same behavior, but for backwards-compatibility with older Ansible releases that always defaulted to /usr/bin/python, will use that interpreter if present.)
func WithAnsiblePythonInterpreter(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePythonInterpreter] = value
	}
}

// WithAnsibleInvalidTaskAttributeFailed sets the option ANSIBLE_INVALID_TASK_ATTRIBUTE_FAILED to true (If ‘false’, invalid attributes for a task will result in warnings instead of errors)
func WithAnsibleInvalidTaskAttributeFailed() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInvalidTaskAttributeFailed] = "true"
	}
}

// WithoutAnsibleInvalidTaskAttributeFailed sets the option ANSIBLE_INVALID_TASK_ATTRIBUTE_FAILED to false
func WithoutAnsibleInvalidTaskAttributeFailed() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInvalidTaskAttributeFailed] = "false"
	}
}

// WithAnsibleInventoryAnyUnparsedIsFailed sets the option ANSIBLE_INVENTORY_ANY_UNPARSED_IS_FAILED to true (If ‘true’, it is a fatal error when any given inventory source cannot be successfully parsed by any available inventory plugin; otherwise, this situation only attracts a warning.)
func WithAnsibleInventoryAnyUnparsedIsFailed() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryAnyUnparsedIsFailed] = "true"
	}
}

// WithoutAnsibleInventoryAnyUnparsedIsFailed sets the option ANSIBLE_INVENTORY_ANY_UNPARSED_IS_FAILED to false
func WithoutAnsibleInventoryAnyUnparsedIsFailed() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryAnyUnparsedIsFailed] = "false"
	}
}

// WithAnsibleInventoryCache sets the value for the configuraion ANSIBLE_INVENTORY_CACHE (Toggle to turn on inventory caching. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory configuration. This message will be removed in 2.16.)
func WithAnsibleInventoryCache(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryCache] = value
	}
}

// WithAnsibleInventoryCachePlugin sets the value for the configuraion ANSIBLE_INVENTORY_CACHE_PLUGIN (The plugin for caching inventory. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.)
func WithAnsibleInventoryCachePlugin(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryCachePlugin] = value
	}
}

// WithAnsibleInventoryCacheConnection sets the value for the configuraion ANSIBLE_INVENTORY_CACHE_CONNECTION (The inventory cache connection. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.)
func WithAnsibleInventoryCacheConnection(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryCacheConnection] = value
	}
}

// WithAnsibleInventoryCachePluginPrefix sets the value for the configuraion ANSIBLE_INVENTORY_CACHE_PLUGIN_PREFIX (The table prefix for the cache plugin. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.)
func WithAnsibleInventoryCachePluginPrefix(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryCachePluginPrefix] = value
	}
}

// WithAnsibleInventoryCacheTimeout sets the value for the configuraion ANSIBLE_INVENTORY_CACHE_TIMEOUT (Expiration timeout for the inventory cache plugin data. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.)
func WithAnsibleInventoryCacheTimeout(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryCacheTimeout] = value
	}
}

// WithAnsibleInventoryEnabled sets the value for the configuraion ANSIBLE_INVENTORY_ENABLED (List of enabled inventory plugins, it also determines the order in which they are used.)
func WithAnsibleInventoryEnabled(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryEnabled] = value
	}
}

// WithAnsibleInventoryExport sets the value for the configuraion ANSIBLE_INVENTORY_EXPORT (Controls if ansible-inventory will accurately reflect Ansible’s view into inventory or its optimized for exporting.)
func WithAnsibleInventoryExport(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryExport] = value
	}
}

// WithAnsibleInventoryIgnore sets the value for the configuraion ANSIBLE_INVENTORY_IGNORE (List of extensions to ignore when using a directory as an inventory source)
func WithAnsibleInventoryIgnore(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryIgnore] = value
	}
}

// WithAnsibleInventoryIgnoreRegex sets the value for the configuraion ANSIBLE_INVENTORY_IGNORE_REGEX (List of patterns to ignore when using a directory as an inventory source)
func WithAnsibleInventoryIgnoreRegex(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryIgnoreRegex] = value
	}
}

// WithAnsibleInventoryUnparsedFailed sets the value for the configuraion ANSIBLE_INVENTORY_UNPARSED_FAILED (If ‘true’ it is a fatal error if every single potential inventory source fails to parse, otherwise this situation will only attract a warning.)
func WithAnsibleInventoryUnparsedFailed(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryUnparsedFailed] = value
	}
}

// WithAnsibleInventoryUnparsedWarning sets the option ANSIBLE_INVENTORY_UNPARSED_WARNING to true (By default Ansible will issue a warning when no inventory was loaded and notes that it will use an implicit localhost-only inventory. These warnings can be silenced by adjusting this setting to False.)
func WithAnsibleInventoryUnparsedWarning() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryUnparsedWarning] = "true"
	}
}

// WithoutAnsibleInventoryUnparsedWarning sets the option ANSIBLE_INVENTORY_UNPARSED_WARNING to false
func WithoutAnsibleInventoryUnparsedWarning() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleInventoryUnparsedWarning] = "false"
	}
}

// WithAnsibleJinja2NativeWarning sets the option ANSIBLE_JINJA2_NATIVE_WARNING to true (Toggle to control showing warnings related to running a Jinja version older than required for jinja2_native [:Deprecated in: 2.17 :Deprecated detail: This option is no longer used in the Ansible Core code base.])
func WithAnsibleJinja2NativeWarning() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleJinja2NativeWarning] = "true"
	}
}

// WithoutAnsibleJinja2NativeWarning sets the option ANSIBLE_JINJA2_NATIVE_WARNING to false
func WithoutAnsibleJinja2NativeWarning() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleJinja2NativeWarning] = "false"
	}
}

// WithAnsibleLocalhostWarning sets the option ANSIBLE_LOCALHOST_WARNING to true (By default Ansible will issue a warning when there are no hosts in the inventory. These warnings can be silenced by adjusting this setting to False.)
func WithAnsibleLocalhostWarning() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleLocalhostWarning] = "true"
	}
}

// WithoutAnsibleLocalhostWarning sets the option ANSIBLE_LOCALHOST_WARNING to false
func WithoutAnsibleLocalhostWarning() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleLocalhostWarning] = "false"
	}
}

// WithAnsibleMaxDiffSize sets the value for the configuraion ANSIBLE_MAX_DIFF_SIZE (Maximum size of files to be considered for diff display)
func WithAnsibleMaxDiffSize(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleMaxDiffSize] = value
	}
}

// WithAnsibleModuleIgnoreExts sets the value for the configuraion ANSIBLE_MODULE_IGNORE_EXTS (List of extensions to ignore when looking for modules to load This is for rejecting script and binary module fallback extensions)
func WithAnsibleModuleIgnoreExts(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleModuleIgnoreExts] = value
	}
}

// WithAnsibleModuleStrictUtf8Response sets the value for the configuraion ANSIBLE_MODULE_STRICT_UTF8_RESPONSE (Enables whether module responses are evaluated for containing non UTF-8 data Disabling this may result in unexpected behavior Only ansible-core should evaluate this configuration)
func WithAnsibleModuleStrictUtf8Response(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleModuleStrictUtf8Response] = value
	}
}

// WithAnsibleNetconfSshConfig sets the value for the configuraion ANSIBLE_NETCONF_SSH_CONFIG (This variable is used to enable bastion/jump host with netconf connection. If set to True the bastion/jump host ssh settings should be present in ~/.ssh/config file, alternatively it can be set to custom ssh configuration file path to read the bastion/jump host settings.)
func WithAnsibleNetconfSshConfig(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleNetconfSshConfig] = value
	}
}

// WithAnsibleNetworkGroupModules sets the value for the configuraion ANSIBLE_NETWORK_GROUP_MODULES ()
func WithAnsibleNetworkGroupModules(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleNetworkGroupModules] = value
	}
}

// WithAnsibleOldPluginCacheClear sets the option ANSIBLE_OLD_PLUGIN_CACHE_CLEAR to true (Previously Ansible would only clear some of the plugin loading caches when loading new roles, this led to some behaviours in which a plugin loaded in previous plays would be unexpectedly ‘sticky’. This setting allows to return to that behaviour.)
func WithAnsibleOldPluginCacheClear() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleOldPluginCacheClear] = "true"
	}
}

// WithoutAnsibleOldPluginCacheClear sets the option ANSIBLE_OLD_PLUGIN_CACHE_CLEAR to false
func WithoutAnsibleOldPluginCacheClear() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleOldPluginCacheClear] = "false"
	}
}

// WithPager sets the value for the configuraion PAGER ()
func WithPager(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[Pager] = value
	}
}

// WithAnsibleParamikoHostKeyAutoAdd sets the option ANSIBLE_PARAMIKO_HOST_KEY_AUTO_ADD to true ()
func WithAnsibleParamikoHostKeyAutoAdd() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleParamikoHostKeyAutoAdd] = "true"
	}
}

// WithoutAnsibleParamikoHostKeyAutoAdd sets the option ANSIBLE_PARAMIKO_HOST_KEY_AUTO_ADD to false
func WithoutAnsibleParamikoHostKeyAutoAdd() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleParamikoHostKeyAutoAdd] = "false"
	}
}

// WithAnsibleParamikoLookForKeys sets the option ANSIBLE_PARAMIKO_LOOK_FOR_KEYS to true ()
func WithAnsibleParamikoLookForKeys() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleParamikoLookForKeys] = "true"
	}
}

// WithoutAnsibleParamikoLookForKeys sets the option ANSIBLE_PARAMIKO_LOOK_FOR_KEYS to false
func WithoutAnsibleParamikoLookForKeys() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleParamikoLookForKeys] = "false"
	}
}

// WithAnsiblePersistentCommandTimeout sets the value for the configuraion ANSIBLE_PERSISTENT_COMMAND_TIMEOUT (This controls the amount of time to wait for response from remote device before timing out persistent connection.)
func WithAnsiblePersistentCommandTimeout(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePersistentCommandTimeout] = value
	}
}

// WithAnsiblePersistentConnectRetryTimeout sets the value for the configuraion ANSIBLE_PERSISTENT_CONNECT_RETRY_TIMEOUT (This controls the retry timeout for persistent connection to connect to the local domain socket.)
func WithAnsiblePersistentConnectRetryTimeout(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePersistentConnectRetryTimeout] = fmt.Sprint(value)
	}
}

// WithAnsiblePersistentConnectTimeout sets the value for the configuraion ANSIBLE_PERSISTENT_CONNECT_TIMEOUT (This controls how long the persistent connection will remain idle before it is destroyed.)
func WithAnsiblePersistentConnectTimeout(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePersistentConnectTimeout] = fmt.Sprint(value)
	}
}

// WithAnsiblePersistentControlPathDir sets the value for the configuraion ANSIBLE_PERSISTENT_CONTROL_PATH_DIR (Path to socket to be used by the connection persistence system.)
func WithAnsiblePersistentControlPathDir(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePersistentControlPathDir] = value
	}
}

// WithAnsiblePlaybookDir sets the value for the configuraion ANSIBLE_PLAYBOOK_DIR (A number of non-playbook CLIs have a --playbook-dir argument; this sets the default value for it.)
func WithAnsiblePlaybookDir(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePlaybookDir] = value
	}
}

// WithAnsiblePlaybookVarsRoot sets the value for the configuraion ANSIBLE_PLAYBOOK_VARS_ROOT (This sets which playbook dirs will be used as a root to process vars plugins, which includes finding host_vars/group_vars)
func WithAnsiblePlaybookVarsRoot(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePlaybookVarsRoot] = value
	}
}

// WithAnsiblePythonModuleRlimitNofile sets the value for the configuraion ANSIBLE_PYTHON_MODULE_RLIMIT_NOFILE (Attempts to set RLIMIT_NOFILE soft limit to the specified value when executing Python modules (can speed up subprocess usage on Python 2.x. See https://bugs.python.org/issue11284). The value will be limited by the existing hard limit. Default value of 0 does not attempt to adjust existing system-defined limits.)
func WithAnsiblePythonModuleRlimitNofile(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePythonModuleRlimitNofile] = value
	}
}

// WithAnsibleRetryFilesEnabled sets the value for the configuraion ANSIBLE_RETRY_FILES_ENABLED (This controls whether a failed Ansible playbook should create a .retry file.)
func WithAnsibleRetryFilesEnabled(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleRetryFilesEnabled] = value
	}
}

// WithAnsibleRetryFilesSavePath sets the value for the configuraion ANSIBLE_RETRY_FILES_SAVE_PATH (This sets the path in which Ansible will save .retry files when a playbook fails and retry files are enabled. This file will be overwritten after each run with the list of failed hosts from all plays.)
func WithAnsibleRetryFilesSavePath(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleRetryFilesSavePath] = value
	}
}

// WithAnsibleRunVarsPlugins sets the value for the configuraion ANSIBLE_RUN_VARS_PLUGINS (This setting can be used to optimize vars_plugin usage depending on user’s inventory size and play selection.)
func WithAnsibleRunVarsPlugins(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleRunVarsPlugins] = value
	}
}

// WithAnsibleShowCustomStats sets the value for the configuraion ANSIBLE_SHOW_CUSTOM_STATS (This adds the custom stats set via the set_stats plugin to the default output)
func WithAnsibleShowCustomStats(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleShowCustomStats] = value
	}
}

// WithAnsibleStringConversionAction sets the value for the configuraion ANSIBLE_STRING_CONVERSION_ACTION (Action to take when a module parameter value is converted to a string (this does not affect variables). For string parameters, values such as ‘1.00’, “[‘a’, ‘b’,]”, and ‘yes’, ‘y’, etc. will be converted by the YAML parser unless fully quoted. Valid options are ‘error’, ‘warn’, and ‘ignore’. Since 2.8, this option defaults to ‘warn’ but will change to ‘error’ in 2.12.)
func WithAnsibleStringConversionAction(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleStringConversionAction] = value
	}
}

// WithAnsibleStringTypeFilters sets the value for the configuraion ANSIBLE_STRING_TYPE_FILTERS (This list of filters avoids ‘type conversion’ when templating variables Useful when you want to avoid conversion into lists or dictionaries for JSON strings, for example.)
func WithAnsibleStringTypeFilters(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleStringTypeFilters] = value
	}
}

// WithAnsibleSystemWarnings sets the option ANSIBLE_SYSTEM_WARNINGS to true (Allows disabling of warnings related to potential issues on the system running ansible itself (not on the managed hosts) These may include warnings about 3rd party packages or other conditions that should be resolved if possible.)
func WithAnsibleSystemWarnings() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleSystemWarnings] = "true"
	}
}

// WithoutAnsibleSystemWarnings sets the option ANSIBLE_SYSTEM_WARNINGS to false
func WithoutAnsibleSystemWarnings() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleSystemWarnings] = "false"
	}
}

// WithAnsibleRunTags sets the value for the configuraion ANSIBLE_RUN_TAGS (default list of tags to run in your plays, Skip Tags has precedence.)
func WithAnsibleRunTags(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleRunTags] = value
	}
}

// WithAnsibleSkipTags sets the value for the configuraion ANSIBLE_SKIP_TAGS (default list of tags to skip in your plays, has precedence over Run Tags)
func WithAnsibleSkipTags(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleSkipTags] = value
	}
}

// WithAnsibleTaskDebuggerIgnoreErrors sets the option ANSIBLE_TASK_DEBUGGER_IGNORE_ERRORS to true (This option defines whether the task debugger will be invoked on a failed task when ignore_errors=True is specified. True specifies that the debugger will honor ignore_errors, False will not honor ignore_errors.)
func WithAnsibleTaskDebuggerIgnoreErrors() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleTaskDebuggerIgnoreErrors] = "true"
	}
}

// WithoutAnsibleTaskDebuggerIgnoreErrors sets the option ANSIBLE_TASK_DEBUGGER_IGNORE_ERRORS to false
func WithoutAnsibleTaskDebuggerIgnoreErrors() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleTaskDebuggerIgnoreErrors] = "false"
	}
}

// WithAnsibleTaskTimeout sets the value for the configuraion ANSIBLE_TASK_TIMEOUT (Set the maximum time (in seconds) that a task can run for. If set to 0 (the default) there is no timeout.)
func WithAnsibleTaskTimeout(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleTaskTimeout] = fmt.Sprint(value)
	}
}

// WithAnsibleTransformInvalidGroupChars sets the value for the configuraion ANSIBLE_TRANSFORM_INVALID_GROUP_CHARS (Make ansible transform invalid characters in group names supplied by inventory sources.)
func WithAnsibleTransformInvalidGroupChars(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleTransformInvalidGroupChars] = value
	}
}

// WithAnsibleUsePersistentConnections sets the option ANSIBLE_USE_PERSISTENT_CONNECTIONS to true (Toggles the use of persistence for connections.)
func WithAnsibleUsePersistentConnections() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleUsePersistentConnections] = "true"
	}
}

// WithoutAnsibleUsePersistentConnections sets the option ANSIBLE_USE_PERSISTENT_CONNECTIONS to false
func WithoutAnsibleUsePersistentConnections() ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleUsePersistentConnections] = "false"
	}
}

// WithAnsibleValidateActionGroupMetadata sets the value for the configuraion ANSIBLE_VALIDATE_ACTION_GROUP_METADATA (A toggle to disable validating a collection’s ‘metadata’ entry for a module_defaults action group. Metadata containing unexpected fields or value types will produce a warning when this is True.)
func WithAnsibleValidateActionGroupMetadata(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleValidateActionGroupMetadata] = value
	}
}

// WithAnsibleVarsEnabled sets the value for the configuraion ANSIBLE_VARS_ENABLED (Accept list for variable plugins that require it.)
func WithAnsibleVarsEnabled(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleVarsEnabled] = value
	}
}

// WithAnsiblePrecedence sets the value for the configuraion ANSIBLE_PRECEDENCE (Allows to change the group variable precedence merge order.)
func WithAnsiblePrecedence(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsiblePrecedence] = value
	}
}

// WithAnsibleVaultEncryptSalt sets the value for the configuraion ANSIBLE_VAULT_ENCRYPT_SALT (The salt to use for the vault encryption. If it is not provided, a random salt will be used.)
func WithAnsibleVaultEncryptSalt(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleVaultEncryptSalt] = value
	}
}

// WithAnsibleVerboseToStderr sets the value for the configuraion ANSIBLE_VERBOSE_TO_STDERR (Force ‘verbose’ option to use stderr instead of stdout)
func WithAnsibleVerboseToStderr(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleVerboseToStderr] = value
	}
}

// WithAnsibleWinAsyncStartupTimeout sets the value for the configuraion ANSIBLE_WIN_ASYNC_STARTUP_TIMEOUT (For asynchronous tasks in Ansible (covered in Asynchronous Actions and Polling), this is how long, in seconds, to wait for the task spawned by Ansible to connect back to the named pipe used on Windows systems. The default is 5 seconds. This can be too low on slower systems, or systems under heavy load. This is not the total time an async command can run for, but is a separate timeout to wait for an async command to start. The task will only start to be timed against its async_timeout once it has connected to the pipe, so the overall maximum duration the task can take will be extended by the amount specified here.)
func WithAnsibleWinAsyncStartupTimeout(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleWinAsyncStartupTimeout] = fmt.Sprint(value)
	}
}

// WithAnsibleWorkerShutdownPollCount sets the value for the configuraion ANSIBLE_WORKER_SHUTDOWN_POLL_COUNT (The maximum number of times to check Task Queue Manager worker processes to verify they have exited cleanly. After this limit is reached any worker processes still running will be terminated. This is for internal use only.)
func WithAnsibleWorkerShutdownPollCount(value int) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleWorkerShutdownPollCount] = fmt.Sprint(value)
	}
}

// WithAnsibleWorkerShutdownPollDelay sets the value for the configuraion ANSIBLE_WORKER_SHUTDOWN_POLL_DELAY (The number of seconds to sleep between polling loops when checking Task Queue Manager worker processes to verify they have exited cleanly. This is for internal use only.)
func WithAnsibleWorkerShutdownPollDelay(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleWorkerShutdownPollDelay] = value
	}
}

// WithAnsibleYamlFilenameExt sets the value for the configuraion ANSIBLE_YAML_FILENAME_EXT (Check all of these extensions when looking for ‘variable’ files which should be YAML or JSON or vaulted versions of these. This affects vars_files, include_vars, inventory and vars plugins among others.)
func WithAnsibleYamlFilenameExt(value string) ConfigurationSettingsFunc {
	return func(e *AnsibleWithConfigurationSettingsExecute) {
		e.configurationSettings[AnsibleYamlFilenameExt] = value
	}
}
