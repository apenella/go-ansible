package results

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultStdoutCallbackResults(t *testing.T) {
	tests := []struct {
		desc  string
		input string
		res   string
		err   error
	}{
		{
			desc: "Testing default stdout callback",
			input: `
PLAY [local] *********************************************************************************************************************************************************************************

TASK [Print test message] ********************************************************************************************************************************************************************
ok: [127.0.0.1] => 
	msg: That's a message to test			

PLAY RECAP ***********************************************************************************************************************************************************************************
127.0.0.1                  : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   			

Playbook run took 0 days, 0 hours, 0 minutes, 0 seconds
`,
			res: ` ── 
 ── PLAY [local] *********************************************************************************************************************************************************************************
 ── 
 ── TASK [Print test message] ********************************************************************************************************************************************************************
 ── ok: [127.0.0.1] => 
 ── 	msg: That's a message to test			
 ── 
 ── PLAY RECAP ***********************************************************************************************************************************************************************************
 ── 127.0.0.1                  : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   			
 ── 
 ── Playbook run took 0 days, 0 hours, 0 minutes, 0 seconds
`,
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			wbuff := bytes.Buffer{}
			writer := io.Writer(&wbuff)
			reader := bufio.NewReader(strings.NewReader(test.input))
			err := DefaultStdoutCallbackResults("", reader, writer)
			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, wbuff.String(), "Unexpected value")
			}
		})
	}
}
