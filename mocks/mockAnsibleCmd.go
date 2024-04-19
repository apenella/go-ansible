package mocks

type MockAnsibleCmd struct {
	cmd []string
	err error
}

func NewMockAnsibleCmd(cmd []string, err error) *MockAnsibleCmd {
	return &MockAnsibleCmd{
		cmd: cmd,
	}
}

func (c *MockAnsibleCmd) Command() ([]string, error) {
	return c.cmd, c.err
}

func (c *MockAnsibleCmd) String() string {
	return ""
}
