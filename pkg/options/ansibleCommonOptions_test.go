package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCommandConnectionOptions(t *testing.T) {

	t.Log("Testing generate command coonection options")

	options := &AnsibleConnectionOptions{
		AskPass:       true,
		Connection:    "local",
		PrivateKey:    "pk",
		SCPExtraArgs:  "scp-extra-args",
		SFTPExtraArgs: "sftp-extra-args",
		SSHCommonArgs: "ssh-common-args",
		SSHExtraArgs:  "ssh-extra-args",
		Timeout:       10,
		User:          "user",
	}

	opts, _ := options.GenerateCommandConnectionOptions()

	expected := []string{"--ask-pass",
		"--connection",
		"local",
		"--private-key",
		"pk",
		"--scp-extra-args",
		"scp-extra-args",
		"--sftp-extra-args",
		"sftp-extra-args",
		"--ssh-common-args",
		"ssh-common-args",
		"--ssh-extra-args",
		"ssh-extra-args",
		"--user",
		"user",
		"--timeout",
		"10",
	}

	assert.Equal(t, expected, opts)

}

func TestCommandConnectionOptionsString(t *testing.T) {

	t.Log("Testing generate command connection options string")

	options := &AnsibleConnectionOptions{
		AskPass:       true,
		Connection:    "local",
		PrivateKey:    "pk",
		SCPExtraArgs:  "scp-extra-args",
		SFTPExtraArgs: "sftp-extra-args",
		SSHCommonArgs: "ssh-common-args",
		SSHExtraArgs:  "ssh-extra-args",
		Timeout:       10,
		User:          "user",
	}

	cmd := options.String()

	expected := " --ask-pass --connection local --private-key pk --scp-extra-args scp-extra-args --sftp-extra-args sftp-extra-args --ssh-common-args ssh-common-args --ssh-extra-args ssh-extra-args --user user --timeout 10"

	assert.Equal(t, expected, cmd)
}

func TestGenerateCommandPrivilegeEscalationOptions(t *testing.T) {

	t.Log("Testing generate command privilege escalation options")

	options := &AnsiblePrivilegeEscalationOptions{
		Become:        true,
		BecomeMethod:  "become-method",
		BecomeUser:    "become-user",
		AskBecomePass: true,
	}

	opts, _ := options.GenerateCommandPrivilegeEscalationOptions()

	expected := []string{
		"--ask-become-pass",
		"--become",
		"--become-method",
		"become-method",
		"--become-user",
		"become-user",
	}

	assert.Equal(t, expected, opts)

}

func TestCommandPrivilegeEscalationOptionsString(t *testing.T) {

	t.Log("Testing generate command privilege escalation options string")

	options := &AnsiblePrivilegeEscalationOptions{
		Become:        true,
		BecomeMethod:  "become-method",
		BecomeUser:    "become-user",
		AskBecomePass: true,
	}

	cmd := options.String()

	expected := " --ask-become-pass --become --become-method become-method --become-user become-user"

	assert.Equal(t, expected, cmd)
}
