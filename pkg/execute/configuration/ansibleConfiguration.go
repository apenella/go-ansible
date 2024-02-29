package configuration

const (
	// AnsibleActionWarnings (boolean) By default Ansible will issue a warning when received from a task action (module or action plugin) These warnings can be silenced by adjusting this setting to False.
	AnsibleActionWarnings = "ANSIBLE_ACTION_WARNINGS"

	// AnsibleAgnosticBecomePrompt (boolean) Display an agnostic become prompt instead of displaying a prompt containing the command line supplied become method
	AnsibleAgnosticBecomePrompt = "ANSIBLE_AGNOSTIC_BECOME_PROMPT"

	// AnsibleConnectionPath (path) Specify where to look for the ansible-connection script. This location will be checked before searching $PATH. If null, ansible will start with the same directory as the ansible script.
	AnsibleConnectionPath = "ANSIBLE_CONNECTION_PATH"

	// AnsibleCowAcceptlist (list) Accept list of cowsay templates that are ‘safe’ to use, set to empty list if you want to enable all installed templates. [:Version Added: 2.11]
	AnsibleCowAcceptlist = "ANSIBLE_COW_ACCEPTLIST"

	// AnsibleCowPath (string) Specify a custom cowsay path or swap in your cowsay implementation of choice
	AnsibleCowPath = "ANSIBLE_COW_PATH"

	// AnsibleCowSelection () This allows you to chose a specific cowsay stencil for the banners or use ‘random’ to cycle through them.
	AnsibleCowSelection = "ANSIBLE_COW_SELECTION"

	// AnsibleForceColor (boolean) This option forces color mode even when running without a TTY or the “nocolor” setting is True.
	AnsibleForceColor = "ANSIBLE_FORCE_COLOR"

	// AnsibleHome (path) The default root path for Ansible config files on the controller.
	AnsibleHome = "ANSIBLE_HOME"

	// NoColor (boolean) This setting allows suppressing colorizing output, which is used to give a better indication of failure and status information.
	NoColor = "NO_COLOR"

	// AnsibleNocows (boolean) If you have cowsay installed but want to avoid the ‘cows’ (why????), use this.
	AnsibleNocows = "ANSIBLE_NOCOWS"

	// AnsiblePipelining (boolean) This is a global option, each connection plugin can override either by having more specific options or not supporting pipelining at all. Pipelining, if supported by the connection plugin, reduces the number of network operations required to execute a module on the remote server, by executing many Ansible modules without actual file transfer. It can result in a very significant performance improvement when enabled. However this conflicts with privilege escalation (become). For example, when using ‘sudo:’ operations you must first disable ‘requiretty’ in /etc/sudoers on all managed hosts, which is why it is disabled by default. This setting will be disabled if ANSIBLE_KEEP_REMOTE_FILES is enabled.
	AnsiblePipelining = "ANSIBLE_PIPELINING"

	// AnsibleAnyErrorsFatal (boolean) Sets the default value for the any_errors_fatal keyword, if True, Task failures will be considered fatal errors.
	AnsibleAnyErrorsFatal = "ANSIBLE_ANY_ERRORS_FATAL"

	// AnsibleBecomeAllowSameUser (boolean) This setting controls if become is skipped when remote user and become user are the same. I.E root sudo to root. If executable, it will be run and the resulting stdout will be used as the password.
	AnsibleBecomeAllowSameUser = "ANSIBLE_BECOME_ALLOW_SAME_USER"

	// AnsibleBecomePasswordFile (path) The password file to use for the become plugin. –become-password-file. If executable, it will be run and the resulting stdout will be used as the password.
	AnsibleBecomePasswordFile = "ANSIBLE_BECOME_PASSWORD_FILE"

	// AnsibleBecomePlugins (pathspec) Colon separated paths in which Ansible will search for Become Plugins.
	AnsibleBecomePlugins = "ANSIBLE_BECOME_PLUGINS"

	// AnsibleCachePlugin () Chooses which cache plugin to use, the default ‘memory’ is ephemeral.
	AnsibleCachePlugin = "ANSIBLE_CACHE_PLUGIN"

	// AnsibleCachePluginConnection () Defines connection or path information for the cache plugin
	AnsibleCachePluginConnection = "ANSIBLE_CACHE_PLUGIN_CONNECTION"

	// AnsibleCachePluginPrefix () Prefix to use for cache plugin files/tables
	AnsibleCachePluginPrefix = "ANSIBLE_CACHE_PLUGIN_PREFIX"

	// AnsibleCachePluginTimeout (integer) Expiration timeout for the cache plugin data
	AnsibleCachePluginTimeout = "ANSIBLE_CACHE_PLUGIN_TIMEOUT"

	// AnsibleCallbacksEnabled (list) List of enabled callbacks, not all callbacks need enabling, but many of those shipped with Ansible do as we don’t want them activated by default. [:Version Added: 2.11]
	AnsibleCallbacksEnabled = "ANSIBLE_CALLBACKS_ENABLED"

	// AnsibleCollectionsOnAnsibleVersionMismatch () When a collection is loaded that does not support the running Ansible version (with the collection metadata key requires_ansible).
	AnsibleCollectionsOnAnsibleVersionMismatch = "ANSIBLE_COLLECTIONS_ON_ANSIBLE_VERSION_MISMATCH"

	// AnsibleCollectionsPaths (pathspec) Colon separated paths in which Ansible will search for collections content. Collections must be in nested subdirectories, not directly in these directories. For example, if COLLECTIONS_PATHS includes '{{ ANSIBLE_HOME ~ "/collections" }}', and you want to add my.collection to that directory, it must be saved as '{{ ANSIBLE_HOME} ~ "/collections/ansible_collections/my/collection" }}'.
	AnsibleCollectionsPaths = "ANSIBLE_COLLECTIONS_PATHS"

	// AnsibleCollectionsScanSysPath (boolean) A boolean to enable or disable scanning the sys.path for installed collections
	AnsibleCollectionsScanSysPath = "ANSIBLE_COLLECTIONS_SCAN_SYS_PATH"

	// AnsibleColorChanged () Defines the color to use on ‘Changed’ task status
	AnsibleColorChanged = "ANSIBLE_COLOR_CHANGED"

	// AnsibleColorConsolePrompt () Defines the default color to use for ansible-console
	AnsibleColorConsolePrompt = "ANSIBLE_COLOR_CONSOLE_PROMPT"

	// AnsibleColorDebug () Defines the color to use when emitting debug messages
	AnsibleColorDebug = "ANSIBLE_COLOR_DEBUG"

	// AnsibleColorDeprecate () Defines the color to use when emitting deprecation messages
	AnsibleColorDeprecate = "ANSIBLE_COLOR_DEPRECATE"

	// AnsibleColorDiffAdd () Defines the color to use when showing added lines in diffs
	AnsibleColorDiffAdd = "ANSIBLE_COLOR_DIFF_ADD"

	// AnsibleColorDiffLines () Defines the color to use when showing diffs
	AnsibleColorDiffLines = "ANSIBLE_COLOR_DIFF_LINES"

	// AnsibleColorDiffRemove () Defines the color to use when showing removed lines in diffs
	AnsibleColorDiffRemove = "ANSIBLE_COLOR_DIFF_REMOVE"

	// AnsibleColorError () Defines the color to use when emitting error messages
	AnsibleColorError = "ANSIBLE_COLOR_ERROR"

	// AnsibleColorHighlight () Defines the color to use for highlighting
	AnsibleColorHighlight = "ANSIBLE_COLOR_HIGHLIGHT"

	// AnsibleColorOk () Defines the color to use when showing ‘OK’ task status
	AnsibleColorOk = "ANSIBLE_COLOR_OK"

	// AnsibleColorSkip () Defines the color to use when showing ‘Skipped’ task status
	AnsibleColorSkip = "ANSIBLE_COLOR_SKIP"

	// AnsibleColorUnreachable () Defines the color to use on ‘Unreachable’ status
	AnsibleColorUnreachable = "ANSIBLE_COLOR_UNREACHABLE"

	// AnsibleColorVerbose () Defines the color to use when emitting verbose messages. i.e those that show with ‘-v’s.
	AnsibleColorVerbose = "ANSIBLE_COLOR_VERBOSE"

	// AnsibleColorWarn () Defines the color to use when emitting warning messages
	AnsibleColorWarn = "ANSIBLE_COLOR_WARN"

	// Parameter 'CONNECTION_FACTS_MODULES' can not be configured by environment variable.
	// Which modules to run during a play’s fact gathering stage based on connection

	// AnsibleConnectionPasswordFile (path) The password file to use for the connection plugin. –connection-password-file.
	AnsibleConnectionPasswordFile = "ANSIBLE_CONNECTION_PASSWORD_FILE"

	// AnsibleCoverageRemoteOutput (str) Sets the output directory on the remote host to generate coverage reports to. Currently only used for remote coverage on PowerShell modules. This is for internal use only.
	AnsibleCoverageRemoteOutput = "_ANSIBLE_COVERAGE_REMOTE_OUTPUT"

	// AnsibleCoverageRemotePathFilter (str) A list of paths for files on the Ansible controller to run coverage for when executing on the remote host. Only files that match the path glob will have its coverage collected. Multiple path globs can be specified and are separated by :. Currently only used for remote coverage on PowerShell modules. This is for internal use only.
	AnsibleCoverageRemotePathFilter = "_ANSIBLE_COVERAGE_REMOTE_PATH_FILTER"

	// AnsibleActionPlugins (pathspec) Colon separated paths in which Ansible will search for Action Plugins.
	AnsibleActionPlugins = "ANSIBLE_ACTION_PLUGINS"

	// Parameter 'DEFAULT_ALLOW_UNSAFE_LOOKUPS' can not be configured by environment variable.
	// When enabled, this option allows lookup plugins (whether used in variables as {{lookup('foo')}} or as a loop as with_foo) to return data that is not marked ‘unsafe’. By default, such data is marked as unsafe to prevent the templating engine from evaluating any jinja2 templating language, as this could represent a security risk. This option is provided to allow for backward compatibility, however users should first consider adding allow_unsafe=True to any lookups which may be expected to contain data which may be run through the templating engine late

	// AnsibleAskPass (boolean) This controls whether an Ansible playbook should prompt for a login password. If using SSH keys for authentication, you probably do not need to change this setting.
	AnsibleAskPass = "ANSIBLE_ASK_PASS"

	// AnsibleAskVaultPass (boolean) This controls whether an Ansible playbook should prompt for a vault password.
	AnsibleAskVaultPass = "ANSIBLE_ASK_VAULT_PASS"

	// AnsibleBecome (boolean) Toggles the use of privilege escalation, allowing you to ‘become’ another user after login.
	AnsibleBecome = "ANSIBLE_BECOME"

	// AnsibleBecomeAskPass (boolean) Toggle to prompt for privilege escalation password.
	AnsibleBecomeAskPass = "ANSIBLE_BECOME_ASK_PASS"

	// AnsibleBecomeExe () executable to use for privilege escalation, otherwise Ansible will depend on PATH
	AnsibleBecomeExe = "ANSIBLE_BECOME_EXE"

	// AnsibleBecomeFlags () Flags to pass to the privilege escalation executable.
	AnsibleBecomeFlags = "ANSIBLE_BECOME_FLAGS"

	// AnsibleBecomeMethod () Privilege escalation method to use when become is enabled.
	AnsibleBecomeMethod = "ANSIBLE_BECOME_METHOD"

	// AnsibleBecomeUser () The user your login/remote user ‘becomes’ when using privilege escalation, most systems will use ‘root’ when no user is specified.
	AnsibleBecomeUser = "ANSIBLE_BECOME_USER"

	// AnsibleCachePlugins (pathspec) Colon separated paths in which Ansible will search for Cache Plugins.
	AnsibleCachePlugins = "ANSIBLE_CACHE_PLUGINS"

	// AnsibleCallbackPlugins (pathspec) Colon separated paths in which Ansible will search for Callback Plugins.
	AnsibleCallbackPlugins = "ANSIBLE_CALLBACK_PLUGINS"

	// AnsibleCliconfPlugins (pathspec) Colon separated paths in which Ansible will search for Cliconf Plugins.
	AnsibleCliconfPlugins = "ANSIBLE_CLICONF_PLUGINS"

	// AnsibleConnectionPlugins (pathspec) Colon separated paths in which Ansible will search for Connection Plugins.
	AnsibleConnectionPlugins = "ANSIBLE_CONNECTION_PLUGINS"

	// AnsibleDebug (boolean) Toggles debug output in Ansible. This is very verbose and can hinder multiprocessing.  Debug output can also include secret information despite no_log settings being enabled, which means debug mode should not be used in production.
	AnsibleDebug = "ANSIBLE_DEBUG"

	// AnsibleExecutable () This indicates the command to use to spawn a shell under for Ansible’s execution needs on a target. Users may need to change this in rare instances when shell usage is constrained, but in most cases it may be left as is.
	AnsibleExecutable = "ANSIBLE_EXECUTABLE"

	// AnsibleFactPath (string) This option allows you to globally configure a custom path for ‘local_facts’ for the implied ansible_collections.ansible.builtin.setup_module task when using fact gathering. If not set, it will fallback to the default from the ansible.builtin.setup module: /etc/ansible/facts.d. This does not affect  user defined tasks that use the ansible.builtin.setup module. The real action being created by the implicit task is currently    ansible.legacy.gather_facts module, which then calls the configured fact modules, by default this will be ansible.builtin.setup for POSIX systems but other platforms might have different defaults.
	AnsibleFactPath = "ANSIBLE_FACT_PATH"

	// AnsibleFilterPlugins (pathspec) Colon separated paths in which Ansible will search for Jinja2 Filter Plugins.
	AnsibleFilterPlugins = "ANSIBLE_FILTER_PLUGINS"

	// AnsibleForceHandlers (boolean) This option controls if notified handlers run on a host even if a failure occurs on that host. When false, the handlers will not run if a failure has occurred on a host. This can also be set per play or on the command line. See Handlers and Failure for more details.
	AnsibleForceHandlers = "ANSIBLE_FORCE_HANDLERS"

	// AnsibleForks (integer) Maximum number of forks Ansible will use to execute tasks on target hosts.
	AnsibleForks = "ANSIBLE_FORKS"

	// AnsibleGatherSubset (list) Set the gather_subset option for the ansible_collections.ansible.builtin.setup_module task in the implicit fact gathering. See the module documentation for specifics. It does not apply to user defined ansible.builtin.setup tasks.
	AnsibleGatherSubset = "ANSIBLE_GATHER_SUBSET"

	// AnsibleGatherTimeout (integer) Set the timeout in seconds for the implicit fact gathering, see the module documentation for specifics. It does not apply to user defined ansible_collections.ansible.builtin.setup_module tasks.
	AnsibleGatherTimeout = "ANSIBLE_GATHER_TIMEOUT"

	// AnsibleGathering () This setting controls the default policy of fact gathering (facts discovered about remote systems). This option can be useful for those wishing to save fact gathering time. Both ‘smart’ and ‘explicit’ will use the cache plugin.
	AnsibleGathering = "ANSIBLE_GATHERING"

	// AnsibleHashBehaviour (string) This setting controls how duplicate definitions of dictionary variables (aka hash, map, associative array) are handled in Ansible. This does not affect variables whose values are scalars (integers, strings) or arrays. WARNING, changing this setting is not recommended as this is fragile and makes your content (plays, roles, collections) non portable, leading to continual confusion and misuse. Don’t change this setting unless you think you have an absolute need for it. We recommend avoiding reusing variable names and relying on the combine filter and vars and varnames lookups to create merged versions of the individual variables. In our experience this is rarely really needed and a sign that too much complexity has been introduced into the data structures and plays. For some uses you can also look into custom vars_plugins to merge on input, even substituting the default host_group_vars that is in charge of parsing the host_vars/ and group_vars/ directories. Most users of this setting are only interested in inventory scope, but the setting itself affects all sources and makes debugging even harder. All playbooks and roles in the official examples repos assume the default for this setting. Changing the setting to merge applies across variable sources, but many sources will internally still overwrite the variables. For example include_vars will dedupe variables internally before updating Ansible, with ‘last defined’ overwriting previous definitions in same file. The Ansible project recommends you avoid ``merge`` for new projects. It is the intention of the Ansible developers to eventually deprecate and remove this setting, but it is being kept as some users do heavily rely on it. New projects should avoid ‘merge’.
	AnsibleHashBehaviour = "ANSIBLE_HASH_BEHAVIOUR"

	// AnsibleInventory (pathlist) Comma separated list of Ansible inventory sources
	AnsibleInventory = "ANSIBLE_INVENTORY"

	// AnsibleHttpapiPlugins (pathspec) Colon separated paths in which Ansible will search for HttpApi Plugins.
	AnsibleHttpapiPlugins = "ANSIBLE_HTTPAPI_PLUGINS"

	// Parameter 'DEFAULT_INTERNAL_POLL_INTERVAL' can not be configured by environment variable.
	// This sets the interval (in seconds) of Ansible internal processes polling each other. Lower values improve performance with large playbooks at the expense of extra CPU load. Higher values are more suitable for Ansible usage in automation scenarios, when UI responsiveness is not required but CPU usage might be a concern. The default corresponds to the value hardcoded in Ansible <= 2.1

	// AnsibleInventoryPlugins (pathspec) Colon separated paths in which Ansible will search for Inventory Plugins.
	AnsibleInventoryPlugins = "ANSIBLE_INVENTORY_PLUGINS"

	// AnsibleJinja2Extensions () This is a developer-specific feature that allows enabling additional Jinja2 extensions. See the Jinja2 documentation for details. If you do not know what these do, you probably don’t need to change this setting :)
	AnsibleJinja2Extensions = "ANSIBLE_JINJA2_EXTENSIONS"

	// AnsibleJinja2Native (boolean) This option preserves variable types during template operations.
	AnsibleJinja2Native = "ANSIBLE_JINJA2_NATIVE"

	// AnsibleKeepRemoteFiles (boolean) Enables/disables the cleaning up of the temporary files Ansible used to execute the tasks on the remote. If this option is enabled it will disable ANSIBLE_PIPELINING.
	AnsibleKeepRemoteFiles = "ANSIBLE_KEEP_REMOTE_FILES"

	// AnsibleLibvirtLxcNoseclabel (boolean) This setting causes libvirt to connect to lxc containers by passing –noseclabel to virsh. This is necessary when running on systems which do not have SELinux.
	AnsibleLibvirtLxcNoseclabel = "ANSIBLE_LIBVIRT_LXC_NOSECLABEL"

	// AnsibleLoadCallbackPlugins (boolean) Controls whether callback plugins are loaded when running /usr/bin/ansible. This may be used to log activity from the command line, send notifications, and so on. Callback plugins are always loaded for ansible-playbook.
	AnsibleLoadCallbackPlugins = "ANSIBLE_LOAD_CALLBACK_PLUGINS"

	// AnsibleLocalTemp (tmppath) Temporary directory for Ansible to use on the controller.
	AnsibleLocalTemp = "ANSIBLE_LOCAL_TEMP"

	// AnsibleLogFilter (list) List of logger names to filter out of the log file
	AnsibleLogFilter = "ANSIBLE_LOG_FILTER"

	// AnsibleLogPath (path) File to which Ansible will log on the controller. When empty logging is disabled.
	AnsibleLogPath = "ANSIBLE_LOG_PATH"

	// AnsibleLookupPlugins (pathspec) Colon separated paths in which Ansible will search for Lookup Plugins.
	AnsibleLookupPlugins = "ANSIBLE_LOOKUP_PLUGINS"

	// Parameter 'DEFAULT_MANAGED_STR' can not be configured by environment variable.
	// Sets the macro for the ‘ansible_managed’ variable available for ansible_collections.ansible.builtin.template_module and ansible_collections.ansible.windows.win_template_module.  This is only relevant for those two modules.

	// AnsibleModuleArgs () This sets the default arguments to pass to the ansible adhoc binary if no -a is specified.
	AnsibleModuleArgs = "ANSIBLE_MODULE_ARGS"

	// Parameter 'DEFAULT_MODULE_COMPRESSION' can not be configured by environment variable.
	// Compression scheme to use when transferring Python modules to the target.

	// Parameter 'DEFAULT_MODULE_NAME' can not be configured by environment variable.
	// Module to use with the ansible AdHoc command, if none is specified via -m.

	// AnsibleLibrary (pathspec) Colon separated paths in which Ansible will search for Modules.
	AnsibleLibrary = "ANSIBLE_LIBRARY"

	// AnsibleModuleUtils (pathspec) Colon separated paths in which Ansible will search for Module utils files, which are shared by modules.
	AnsibleModuleUtils = "ANSIBLE_MODULE_UTILS"

	// AnsibleNetconfPlugins (pathspec) Colon separated paths in which Ansible will search for Netconf Plugins.
	AnsibleNetconfPlugins = "ANSIBLE_NETCONF_PLUGINS"

	// AnsibleNoLog (boolean) Toggle Ansible’s display and logging of task details, mainly used to avoid security disclosures.
	AnsibleNoLog = "ANSIBLE_NO_LOG"

	// AnsibleNoTargetSyslog (boolean) Toggle Ansible logging to syslog on the target when it executes tasks. On Windows hosts this will disable a newer style PowerShell modules from writing to the event log.
	AnsibleNoTargetSyslog = "ANSIBLE_NO_TARGET_SYSLOG"

	// AnsibleNullRepresentation (raw) What templating should return as a ‘null’ value. When not set it will let Jinja2 decide.
	AnsibleNullRepresentation = "ANSIBLE_NULL_REPRESENTATION"

	// AnsiblePollInterval (integer) For asynchronous tasks in Ansible (covered in Asynchronous Actions and Polling), this is how often to check back on the status of those tasks when an explicit poll interval is not supplied. The default is a reasonably moderate 15 seconds which is a tradeoff between checking in frequently and providing a quick turnaround when something may have completed.
	AnsiblePollInterval = "ANSIBLE_POLL_INTERVAL"

	// AnsiblePrivateKeyFile (path) Option for connections using a certificate or key file to authenticate, rather than an agent or passwords, you can set the default value here to avoid re-specifying –private-key with every invocation.
	AnsiblePrivateKeyFile = "ANSIBLE_PRIVATE_KEY_FILE"

	// AnsiblePrivateRoleVars (boolean) By default, imported roles publish their variables to the play and other roles, this setting can avoid that. This was introduced as a way to reset role variables to default values if a role is used more than once in a playbook. Included roles only make their variables public at execution, unlike imported roles which happen at playbook compile time.
	AnsiblePrivateRoleVars = "ANSIBLE_PRIVATE_ROLE_VARS"

	// AnsibleRemotePort (integer) Port to use in remote connections, when blank it will use the connection plugin default.
	AnsibleRemotePort = "ANSIBLE_REMOTE_PORT"

	// AnsibleRemoteUser () Sets the login user for the target machines When blank it uses the connection plugin’s default, normally the user currently executing Ansible.
	AnsibleRemoteUser = "ANSIBLE_REMOTE_USER"

	// AnsibleRolesPath (pathspec) Colon separated paths in which Ansible will search for Roles.
	AnsibleRolesPath = "ANSIBLE_ROLES_PATH"

	// AnsibleSelinuxSpecialFs (list) Some filesystems do not support safe operations and/or return inconsistent errors, this setting makes Ansible ‘tolerate’ those in the list w/o causing fatal errors. Data corruption may occur and writes are not always verified when a filesystem is in the list. [:Version Added: 2.9]
	AnsibleSelinuxSpecialFs = "ANSIBLE_SELINUX_SPECIAL_FS"

	// AnsibleStdoutCallback () Set the main callback used to display Ansible output. You can only have one at a time. You can have many other callbacks, but just one can be in charge of stdout. See Callback plugins for a list of available options.
	AnsibleStdoutCallback = "ANSIBLE_STDOUT_CALLBACK"

	// AnsibleStrategy () Set the default strategy used for plays.
	AnsibleStrategy = "ANSIBLE_STRATEGY"

	// AnsibleStrategyPlugins (pathspec) Colon separated paths in which Ansible will search for Strategy Plugins.
	AnsibleStrategyPlugins = "ANSIBLE_STRATEGY_PLUGINS"

	// AnsibleSu (boolean) Toggle the use of “su” for tasks.
	AnsibleSu = "ANSIBLE_SU"

	// AnsibleSyslogFacility () Syslog facility to use when Ansible logs to the remote target
	AnsibleSyslogFacility = "ANSIBLE_SYSLOG_FACILITY"

	// AnsibleTerminalPlugins (pathspec) Colon separated paths in which Ansible will search for Terminal Plugins.
	AnsibleTerminalPlugins = "ANSIBLE_TERMINAL_PLUGINS"

	// AnsibleTestPlugins (pathspec) Colon separated paths in which Ansible will search for Jinja2 Test Plugins.
	AnsibleTestPlugins = "ANSIBLE_TEST_PLUGINS"

	// AnsibleTimeout (integer) This is the default timeout for connection plugins to use.
	AnsibleTimeout = "ANSIBLE_TIMEOUT"

	// AnsibleTransport () Can be any connection plugin available to your ansible installation. There is also a (DEPRECATED) special ‘smart’ option, that will toggle between ‘ssh’ and ‘paramiko’ depending on controller OS and ssh versions.
	AnsibleTransport = "ANSIBLE_TRANSPORT"

	// AnsibleErrorOnUndefinedVars (boolean) When True, this causes ansible templating to fail steps that reference variable names that are likely typoed. Otherwise, any ‘{{ template_expression }}’ that contains undefined variables will be rendered in a template or ansible action line exactly as written.
	AnsibleErrorOnUndefinedVars = "ANSIBLE_ERROR_ON_UNDEFINED_VARS"

	// AnsibleVarsPlugins (pathspec) Colon separated paths in which Ansible will search for Vars Plugins.
	AnsibleVarsPlugins = "ANSIBLE_VARS_PLUGINS"

	// AnsibleVaultEncryptIdentity () The vault_id to use for encrypting by default. If multiple vault_ids are provided, this specifies which to use for encryption. The –encrypt-vault-id cli option overrides the configured value.
	AnsibleVaultEncryptIdentity = "ANSIBLE_VAULT_ENCRYPT_IDENTITY"

	// AnsibleVaultIdMatch () If true, decrypting vaults with a vault id will only try the password from the matching vault-id
	AnsibleVaultIdMatch = "ANSIBLE_VAULT_ID_MATCH"

	// AnsibleVaultIdentity () The label to use for the default vault id label in cases where a vault id label is not provided
	AnsibleVaultIdentity = "ANSIBLE_VAULT_IDENTITY"

	// AnsibleVaultIdentityList (list) A list of vault-ids to use by default. Equivalent to multiple –vault-id args. Vault-ids are tried in order.
	AnsibleVaultIdentityList = "ANSIBLE_VAULT_IDENTITY_LIST"

	// AnsibleVaultPasswordFile (path) The vault password file to use. Equivalent to –vault-password-file or –vault-id If executable, it will be run and the resulting stdout will be used as the password.
	AnsibleVaultPasswordFile = "ANSIBLE_VAULT_PASSWORD_FILE"

	// AnsibleVerbosity (integer) Sets the default verbosity, equivalent to the number of -v passed in the command line.
	AnsibleVerbosity = "ANSIBLE_VERBOSITY"

	// AnsibleDeprecationWarnings (boolean) Toggle to control the showing of deprecation warnings
	AnsibleDeprecationWarnings = "ANSIBLE_DEPRECATION_WARNINGS"

	// AnsibleDevelWarning (boolean) Toggle to control showing warnings related to running devel
	AnsibleDevelWarning = "ANSIBLE_DEVEL_WARNING"

	// AnsibleDiffAlways (bool) Configuration toggle to tell modules to show differences when in ‘changed’ status, equivalent to --diff.
	AnsibleDiffAlways = "ANSIBLE_DIFF_ALWAYS"

	// AnsibleDiffContext (integer) How many lines of context to show when displaying the differences between files.
	AnsibleDiffContext = "ANSIBLE_DIFF_CONTEXT"

	// AnsibleDisplayArgsToStdout (boolean) Normally ansible-playbook will print a header for each task that is run. These headers will contain the name: field from the task if you specified one. If you didn’t then ansible-playbook uses the task’s action to help you tell which task is presently running. Sometimes you run many of the same action and so you want more information about the task to differentiate it from others of the same action. If you set this variable to True in the config then ansible-playbook will also include the task’s arguments in the header. This setting defaults to False because there is a chance that you have sensitive values in your parameters and you do not want those to be printed. If you set this to True you should be sure that you have secured your environment’s stdout (no one can shoulder surf your screen and you aren’t saving stdout to an insecure file) or made sure that all of your playbooks explicitly added the no_log: True parameter to tasks which have sensitive values See How do I keep secret data in my playbook? for more information.
	AnsibleDisplayArgsToStdout = "ANSIBLE_DISPLAY_ARGS_TO_STDOUT"

	// AnsibleDisplaySkippedHosts (boolean) Toggle to control displaying skipped task/host entries in a task in the default callback
	AnsibleDisplaySkippedHosts = "ANSIBLE_DISPLAY_SKIPPED_HOSTS"

	// AnsibleDocFragmentPlugins (pathspec) Colon separated paths in which Ansible will search for Documentation Fragments Plugins.
	AnsibleDocFragmentPlugins = "ANSIBLE_DOC_FRAGMENT_PLUGINS"

	// Parameter 'DOCSITE_ROOT_URL' can not be configured by environment variable.
	// Root docsite URL used to generate docs URLs in warning/error text; must be an absolute URL with valid scheme and trailing slash.

	// AnsibleDuplicateYamlDictKey (string) By default Ansible will issue a warning when a duplicate dict key is encountered in YAML. These warnings can be silenced by adjusting this setting to False.
	AnsibleDuplicateYamlDictKey = "ANSIBLE_DUPLICATE_YAML_DICT_KEY"

	// Editor ()
	Editor = "EDITOR"

	// AnsibleEnableTaskDebugger (boolean) Whether or not to enable the task debugger, this previously was done as a strategy plugin. Now all strategy plugins can inherit this behavior. The debugger defaults to activating when a task is failed on unreachable. Use the debugger keyword for more flexibility.
	AnsibleEnableTaskDebugger = "ANSIBLE_ENABLE_TASK_DEBUGGER"

	// AnsibleErrorOnMissingHandler (boolean) Toggle to allow missing handlers to become a warning instead of an error when notifying.
	AnsibleErrorOnMissingHandler = "ANSIBLE_ERROR_ON_MISSING_HANDLER"

	// AnsibleFactsModules (list) Which modules to run during a play’s fact gathering stage, using the default of ‘smart’ will try to figure it out based on connection type. If adding your own modules but you still want to use the default Ansible facts, you will want to include ‘setup’ or corresponding network module to the list (if you add ‘smart’, Ansible will also figure it out). This does not affect explicit calls to the ‘setup’ module, but does always affect the ‘gather_facts’ action (implicit or explicit).
	AnsibleFactsModules = "ANSIBLE_FACTS_MODULES"

	// AnsibleGalaxyCacheDir (path) The directory that stores cached responses from a Galaxy server. This is only used by the ansible-galaxy collection install and download commands. Cache files inside this dir will be ignored if they are world writable.
	AnsibleGalaxyCacheDir = "ANSIBLE_GALAXY_CACHE_DIR"

	// AnsibleGalaxyCollectionSkeleton (path) Collection skeleton directory to use as a template for the init action in ansible-galaxy collection, same as --collection-skeleton.
	AnsibleGalaxyCollectionSkeleton = "ANSIBLE_GALAXY_COLLECTION_SKELETON"

	// AnsibleGalaxyCollectionSkeletonIgnore (list) patterns of files to ignore inside a Galaxy collection skeleton directory
	AnsibleGalaxyCollectionSkeletonIgnore = "ANSIBLE_GALAXY_COLLECTION_SKELETON_IGNORE"

	// AnsibleGalaxyCollectionsPathWarning (bool) whether ansible-galaxy collection install should warn about --collections-path missing from configured COLLECTIONS_PATHS
	AnsibleGalaxyCollectionsPathWarning = "ANSIBLE_GALAXY_COLLECTIONS_PATH_WARNING"

	// AnsibleGalaxyDisableGpgVerify (bool) Disable GPG signature verification during collection installation.
	AnsibleGalaxyDisableGpgVerify = "ANSIBLE_GALAXY_DISABLE_GPG_VERIFY"

	// AnsibleGalaxyDisplayProgress (bool) Some steps in ansible-galaxy display a progress wheel which can cause issues on certain displays or when outputting the stdout to a file. This config option controls whether the display wheel is shown or not. The default is to show the display wheel if stdout has a tty.
	AnsibleGalaxyDisplayProgress = "ANSIBLE_GALAXY_DISPLAY_PROGRESS"

	// AnsibleGalaxyGpgKeyring (path) Configure the keyring used for GPG signature verification during collection installation and verification.
	AnsibleGalaxyGpgKeyring = "ANSIBLE_GALAXY_GPG_KEYRING"

	// AnsibleGalaxyIgnore (boolean) If set to yes, ansible-galaxy will not validate TLS certificates. This can be useful for testing against a server with a self-signed certificate.
	AnsibleGalaxyIgnore = "ANSIBLE_GALAXY_IGNORE"

	// AnsibleGalaxyIgnoreSignatureStatusCodes (list) A list of GPG status codes to ignore during GPG signature verification. See L(https://github.com/gpg/gnupg/blob/master/doc/DETAILS#general-status-codes) for status code descriptions. If fewer signatures successfully verify the collection than GALAXY_REQUIRED_VALID_SIGNATURE_COUNT, signature verification will fail even if all error codes are ignored.
	AnsibleGalaxyIgnoreSignatureStatusCodes = "ANSIBLE_GALAXY_IGNORE_SIGNATURE_STATUS_CODES"

	// AnsibleGalaxyRequiredValidSignatureCount (str) The number of signatures that must be successful during GPG signature verification while installing or verifying collections. This should be a positive integer or all to indicate all signatures must successfully validate the collection. Prepend + to the value to fail if no valid signatures are found for the collection.
	AnsibleGalaxyRequiredValidSignatureCount = "ANSIBLE_GALAXY_REQUIRED_VALID_SIGNATURE_COUNT"

	// AnsibleGalaxyRoleSkeleton (path) Role skeleton directory to use as a template for the init action in ansible-galaxy/ansible-galaxy role, same as --role-skeleton.
	AnsibleGalaxyRoleSkeleton = "ANSIBLE_GALAXY_ROLE_SKELETON"

	// AnsibleGalaxyRoleSkeletonIgnore (list) patterns of files to ignore inside a Galaxy role or collection skeleton directory
	AnsibleGalaxyRoleSkeletonIgnore = "ANSIBLE_GALAXY_ROLE_SKELETON_IGNORE"

	// AnsibleGalaxyServer () URL to prepend when roles don’t specify the full URI, assume they are referencing this server as the source.
	AnsibleGalaxyServer = "ANSIBLE_GALAXY_SERVER"

	// AnsibleGalaxyServerList (list) A list of Galaxy servers to use when installing a collection. The value corresponds to the config ini header [galaxy_server.{{item}}] which defines the server details. See Configuring the ansible-galaxy client for more details on how to define a Galaxy server. The order of servers in this list is used to as the order in which a collection is resolved. Setting this config option will ignore the GALAXY_SERVER config option.
	AnsibleGalaxyServerList = "ANSIBLE_GALAXY_SERVER_LIST"

	// AnsibleGalaxyServerTimeout (int) The default timeout for Galaxy API calls. Galaxy servers that don’t configure a specific timeout will fall back to this value.
	AnsibleGalaxyServerTimeout = "ANSIBLE_GALAXY_SERVER_TIMEOUT"

	// AnsibleGalaxyTokenPath (path) Local path to galaxy access token file
	AnsibleGalaxyTokenPath = "ANSIBLE_GALAXY_TOKEN_PATH"

	// AnsibleHostKeyChecking (boolean) Set this to “False” if you want to avoid host key checking by the underlying tools Ansible uses to connect to the host
	AnsibleHostKeyChecking = "ANSIBLE_HOST_KEY_CHECKING"

	// AnsibleHostPatternMismatch () This setting changes the behaviour of mismatched host patterns, it allows you to force a fatal error, a warning or just ignore it
	AnsibleHostPatternMismatch = "ANSIBLE_HOST_PATTERN_MISMATCH"

	// AnsibleInjectFactVars (boolean) Facts are available inside the ansible_facts variable, this setting also pushes them as their own vars in the main namespace. Unlike inside the ansible_facts dictionary, these will have an ansible_ prefix.
	AnsibleInjectFactVars = "ANSIBLE_INJECT_FACT_VARS"

	// AnsiblePythonInterpreter () Path to the Python interpreter to be used for module execution on remote targets, or an automatic discovery mode. Supported discovery modes are auto (the default), auto_silent, auto_legacy, and auto_legacy_silent. All discovery modes employ a lookup table to use the included system Python (on distributions known to include one), falling back to a fixed ordered list of well-known Python interpreter locations if a platform-specific default is not available. The fallback behavior will issue a warning that the interpreter should be set explicitly (since interpreters installed later may change which one is used). This warning behavior can be disabled by setting auto_silent or auto_legacy_silent. The value of auto_legacy provides all the same behavior, but for backwards-compatibility with older Ansible releases that always defaulted to /usr/bin/python, will use that interpreter if present.
	AnsiblePythonInterpreter = "ANSIBLE_PYTHON_INTERPRETER"

	// Parameter 'INTERPRETER_PYTHON_FALLBACK' can not be configured by environment variable.
	//

	// AnsibleInvalidTaskAttributeFailed (boolean) If ‘false’, invalid attributes for a task will result in warnings instead of errors
	AnsibleInvalidTaskAttributeFailed = "ANSIBLE_INVALID_TASK_ATTRIBUTE_FAILED"

	// AnsibleInventoryAnyUnparsedIsFailed (boolean) If ‘true’, it is a fatal error when any given inventory source cannot be successfully parsed by any available inventory plugin; otherwise, this situation only attracts a warning.
	AnsibleInventoryAnyUnparsedIsFailed = "ANSIBLE_INVENTORY_ANY_UNPARSED_IS_FAILED"

	// AnsibleInventoryCache (bool) Toggle to turn on inventory caching. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory configuration. This message will be removed in 2.16.
	AnsibleInventoryCache = "ANSIBLE_INVENTORY_CACHE"

	// AnsibleInventoryCachePlugin () The plugin for caching inventory. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.
	AnsibleInventoryCachePlugin = "ANSIBLE_INVENTORY_CACHE_PLUGIN"

	// AnsibleInventoryCacheConnection () The inventory cache connection. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.
	AnsibleInventoryCacheConnection = "ANSIBLE_INVENTORY_CACHE_CONNECTION"

	// AnsibleInventoryCachePluginPrefix () The table prefix for the cache plugin. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.
	AnsibleInventoryCachePluginPrefix = "ANSIBLE_INVENTORY_CACHE_PLUGIN_PREFIX"

	// AnsibleInventoryCacheTimeout () Expiration timeout for the inventory cache plugin data. This setting has been moved to the individual inventory plugins as a plugin option Inventory plugins. The existing configuration settings are still accepted with the inventory plugin adding additional options from inventory and fact cache configuration. This message will be removed in 2.16.
	AnsibleInventoryCacheTimeout = "ANSIBLE_INVENTORY_CACHE_TIMEOUT"

	// AnsibleInventoryEnabled (list) List of enabled inventory plugins, it also determines the order in which they are used.
	AnsibleInventoryEnabled = "ANSIBLE_INVENTORY_ENABLED"

	// AnsibleInventoryExport (bool) Controls if ansible-inventory will accurately reflect Ansible’s view into inventory or its optimized for exporting.
	AnsibleInventoryExport = "ANSIBLE_INVENTORY_EXPORT"

	// AnsibleInventoryIgnore (list) List of extensions to ignore when using a directory as an inventory source
	AnsibleInventoryIgnore = "ANSIBLE_INVENTORY_IGNORE"

	// AnsibleInventoryIgnoreRegex (list) List of patterns to ignore when using a directory as an inventory source
	AnsibleInventoryIgnoreRegex = "ANSIBLE_INVENTORY_IGNORE_REGEX"

	// AnsibleInventoryUnparsedFailed (bool) If ‘true’ it is a fatal error if every single potential inventory source fails to parse, otherwise this situation will only attract a warning.
	AnsibleInventoryUnparsedFailed = "ANSIBLE_INVENTORY_UNPARSED_FAILED"

	// AnsibleInventoryUnparsedWarning (boolean) By default Ansible will issue a warning when no inventory was loaded and notes that it will use an implicit localhost-only inventory. These warnings can be silenced by adjusting this setting to False.
	AnsibleInventoryUnparsedWarning = "ANSIBLE_INVENTORY_UNPARSED_WARNING"

	// AnsibleJinja2NativeWarning (boolean) Toggle to control showing warnings related to running a Jinja version older than required for jinja2_native [:Deprecated in: 2.17 :Deprecated detail: This option is no longer used in the Ansible Core code base.]
	AnsibleJinja2NativeWarning = "ANSIBLE_JINJA2_NATIVE_WARNING"

	// AnsibleLocalhostWarning (boolean) By default Ansible will issue a warning when there are no hosts in the inventory. These warnings can be silenced by adjusting this setting to False.
	AnsibleLocalhostWarning = "ANSIBLE_LOCALHOST_WARNING"

	// AnsibleMaxDiffSize (int) Maximum size of files to be considered for diff display
	AnsibleMaxDiffSize = "ANSIBLE_MAX_DIFF_SIZE"

	// AnsibleModuleIgnoreExts (list) List of extensions to ignore when looking for modules to load This is for rejecting script and binary module fallback extensions
	AnsibleModuleIgnoreExts = "ANSIBLE_MODULE_IGNORE_EXTS"

	// AnsibleModuleStrictUtf8Response (bool) Enables whether module responses are evaluated for containing non UTF-8 data Disabling this may result in unexpected behavior Only ansible-core should evaluate this configuration
	AnsibleModuleStrictUtf8Response = "ANSIBLE_MODULE_STRICT_UTF8_RESPONSE"

	// AnsibleNetconfSshConfig () This variable is used to enable bastion/jump host with netconf connection. If set to True the bastion/jump host ssh settings should be present in ~/.ssh/config file, alternatively it can be set to custom ssh configuration file path to read the bastion/jump host settings.
	AnsibleNetconfSshConfig = "ANSIBLE_NETCONF_SSH_CONFIG"

	// AnsibleNetworkGroupModules (list)
	AnsibleNetworkGroupModules = "ANSIBLE_NETWORK_GROUP_MODULES"

	// AnsibleOldPluginCacheClear (boolean) Previously Ansible would only clear some of the plugin loading caches when loading new roles, this led to some behaviours in which a plugin loaded in previous plays would be unexpectedly ‘sticky’. This setting allows to return to that behaviour.
	AnsibleOldPluginCacheClear = "ANSIBLE_OLD_PLUGIN_CACHE_CLEAR"

	// Pager ()
	Pager = "PAGER"

	// AnsibleParamikoHostKeyAutoAdd (boolean)
	AnsibleParamikoHostKeyAutoAdd = "ANSIBLE_PARAMIKO_HOST_KEY_AUTO_ADD"

	// AnsibleParamikoLookForKeys (boolean)
	AnsibleParamikoLookForKeys = "ANSIBLE_PARAMIKO_LOOK_FOR_KEYS"

	// AnsiblePersistentCommandTimeout (int) This controls the amount of time to wait for response from remote device before timing out persistent connection.
	AnsiblePersistentCommandTimeout = "ANSIBLE_PERSISTENT_COMMAND_TIMEOUT"

	// AnsiblePersistentConnectRetryTimeout (integer) This controls the retry timeout for persistent connection to connect to the local domain socket.
	AnsiblePersistentConnectRetryTimeout = "ANSIBLE_PERSISTENT_CONNECT_RETRY_TIMEOUT"

	// AnsiblePersistentConnectTimeout (integer) This controls how long the persistent connection will remain idle before it is destroyed.
	AnsiblePersistentConnectTimeout = "ANSIBLE_PERSISTENT_CONNECT_TIMEOUT"

	// AnsiblePersistentControlPathDir (path) Path to socket to be used by the connection persistence system.
	AnsiblePersistentControlPathDir = "ANSIBLE_PERSISTENT_CONTROL_PATH_DIR"

	// AnsiblePlaybookDir (path) A number of non-playbook CLIs have a --playbook-dir argument; this sets the default value for it.
	AnsiblePlaybookDir = "ANSIBLE_PLAYBOOK_DIR"

	// AnsiblePlaybookVarsRoot () This sets which playbook dirs will be used as a root to process vars plugins, which includes finding host_vars/group_vars
	AnsiblePlaybookVarsRoot = "ANSIBLE_PLAYBOOK_VARS_ROOT"

	// Parameter 'PLUGIN_FILTERS_CFG' can not be configured by environment variable.
	// A path to configuration for filtering which plugins installed on the system are allowed to be used. See Rejecting modules for details of the filter file’s format.  The default is /etc/ansible/plugin_filters.yml

	// AnsiblePythonModuleRlimitNofile () Attempts to set RLIMIT_NOFILE soft limit to the specified value when executing Python modules (can speed up subprocess usage on Python 2.x. See https://bugs.python.org/issue11284). The value will be limited by the existing hard limit. Default value of 0 does not attempt to adjust existing system-defined limits.
	AnsiblePythonModuleRlimitNofile = "ANSIBLE_PYTHON_MODULE_RLIMIT_NOFILE"

	// AnsibleRetryFilesEnabled (bool) This controls whether a failed Ansible playbook should create a .retry file.
	AnsibleRetryFilesEnabled = "ANSIBLE_RETRY_FILES_ENABLED"

	// AnsibleRetryFilesSavePath (path) This sets the path in which Ansible will save .retry files when a playbook fails and retry files are enabled. This file will be overwritten after each run with the list of failed hosts from all plays.
	AnsibleRetryFilesSavePath = "ANSIBLE_RETRY_FILES_SAVE_PATH"

	// AnsibleRunVarsPlugins (str) This setting can be used to optimize vars_plugin usage depending on user’s inventory size and play selection.
	AnsibleRunVarsPlugins = "ANSIBLE_RUN_VARS_PLUGINS"

	// AnsibleShowCustomStats (bool) This adds the custom stats set via the set_stats plugin to the default output
	AnsibleShowCustomStats = "ANSIBLE_SHOW_CUSTOM_STATS"

	// AnsibleStringConversionAction (string) Action to take when a module parameter value is converted to a string (this does not affect variables). For string parameters, values such as ‘1.00’, “[‘a’, ‘b’,]”, and ‘yes’, ‘y’, etc. will be converted by the YAML parser unless fully quoted. Valid options are ‘error’, ‘warn’, and ‘ignore’. Since 2.8, this option defaults to ‘warn’ but will change to ‘error’ in 2.12.
	AnsibleStringConversionAction = "ANSIBLE_STRING_CONVERSION_ACTION"

	// AnsibleStringTypeFilters (list) This list of filters avoids ‘type conversion’ when templating variables Useful when you want to avoid conversion into lists or dictionaries for JSON strings, for example.
	AnsibleStringTypeFilters = "ANSIBLE_STRING_TYPE_FILTERS"

	// AnsibleSystemWarnings (boolean) Allows disabling of warnings related to potential issues on the system running ansible itself (not on the managed hosts) These may include warnings about 3rd party packages or other conditions that should be resolved if possible.
	AnsibleSystemWarnings = "ANSIBLE_SYSTEM_WARNINGS"

	// AnsibleRunTags (list) default list of tags to run in your plays, Skip Tags has precedence.
	AnsibleRunTags = "ANSIBLE_RUN_TAGS"

	// AnsibleSkipTags (list) default list of tags to skip in your plays, has precedence over Run Tags
	AnsibleSkipTags = "ANSIBLE_SKIP_TAGS"

	// AnsibleTaskDebuggerIgnoreErrors (boolean) This option defines whether the task debugger will be invoked on a failed task when ignore_errors=True is specified. True specifies that the debugger will honor ignore_errors, False will not honor ignore_errors.
	AnsibleTaskDebuggerIgnoreErrors = "ANSIBLE_TASK_DEBUGGER_IGNORE_ERRORS"

	// AnsibleTaskTimeout (integer) Set the maximum time (in seconds) that a task can run for. If set to 0 (the default) there is no timeout.
	AnsibleTaskTimeout = "ANSIBLE_TASK_TIMEOUT"

	// AnsibleTransformInvalidGroupChars (string) Make ansible transform invalid characters in group names supplied by inventory sources.
	AnsibleTransformInvalidGroupChars = "ANSIBLE_TRANSFORM_INVALID_GROUP_CHARS"

	// AnsibleUsePersistentConnections (boolean) Toggles the use of persistence for connections.
	AnsibleUsePersistentConnections = "ANSIBLE_USE_PERSISTENT_CONNECTIONS"

	// AnsibleValidateActionGroupMetadata (bool) A toggle to disable validating a collection’s ‘metadata’ entry for a module_defaults action group. Metadata containing unexpected fields or value types will produce a warning when this is True.
	AnsibleValidateActionGroupMetadata = "ANSIBLE_VALIDATE_ACTION_GROUP_METADATA"

	// AnsibleVarsEnabled (list) Accept list for variable plugins that require it.
	AnsibleVarsEnabled = "ANSIBLE_VARS_ENABLED"

	// AnsiblePrecedence (list) Allows to change the group variable precedence merge order.
	AnsiblePrecedence = "ANSIBLE_PRECEDENCE"

	// AnsibleVaultEncryptSalt () The salt to use for the vault encryption. If it is not provided, a random salt will be used.
	AnsibleVaultEncryptSalt = "ANSIBLE_VAULT_ENCRYPT_SALT"

	// AnsibleVerboseToStderr (bool) Force ‘verbose’ option to use stderr instead of stdout
	AnsibleVerboseToStderr = "ANSIBLE_VERBOSE_TO_STDERR"

	// AnsibleWinAsyncStartupTimeout (integer) For asynchronous tasks in Ansible (covered in Asynchronous Actions and Polling), this is how long, in seconds, to wait for the task spawned by Ansible to connect back to the named pipe used on Windows systems. The default is 5 seconds. This can be too low on slower systems, or systems under heavy load. This is not the total time an async command can run for, but is a separate timeout to wait for an async command to start. The task will only start to be timed against its async_timeout once it has connected to the pipe, so the overall maximum duration the task can take will be extended by the amount specified here.
	AnsibleWinAsyncStartupTimeout = "ANSIBLE_WIN_ASYNC_STARTUP_TIMEOUT"

	// AnsibleWorkerShutdownPollCount (integer) The maximum number of times to check Task Queue Manager worker processes to verify they have exited cleanly. After this limit is reached any worker processes still running will be terminated. This is for internal use only.
	AnsibleWorkerShutdownPollCount = "ANSIBLE_WORKER_SHUTDOWN_POLL_COUNT"

	// AnsibleWorkerShutdownPollDelay (float) The number of seconds to sleep between polling loops when checking Task Queue Manager worker processes to verify they have exited cleanly. This is for internal use only.
	AnsibleWorkerShutdownPollDelay = "ANSIBLE_WORKER_SHUTDOWN_POLL_DELAY"

	// AnsibleYamlFilenameExt (list) Check all of these extensions when looking for ‘variable’ files which should be YAML or JSON or vaulted versions of these. This affects vars_files, include_vars, inventory and vars plugins among others.
	AnsibleYamlFilenameExt = "ANSIBLE_YAML_FILENAME_EXT"
)
