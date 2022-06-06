package results

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultStdoutCallbackResults(t *testing.T) {
	longMessageLine := randStringBytes(512_000)
	tests := []struct {
		desc         string
		input        string
		res          string
		err          error
		trans        []TransformerFunc
		closedReader bool
		closedWriter bool
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
			res: `
PLAY [local] *********************************************************************************************************************************************************************************

TASK [Print test message] ********************************************************************************************************************************************************************
ok: [127.0.0.1] => 
	msg: That's a message to test			

PLAY RECAP ***********************************************************************************************************************************************************************************
127.0.0.1                  : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   			

Playbook run took 0 days, 0 hours, 0 minutes, 0 seconds
`,
			err: nil,
		},
		{
			desc: "Testing very long lines reading",
			input: fmt.Sprintf(`
PLAY [local] *********************************************************************************************************************************************************************************

TASK [Print test message] ********************************************************************************************************************************************************************
ok: [127.0.0.1] => 
	msg: That's a message to test: %s

PLAY RECAP ***********************************************************************************************************************************************************************************
127.0.0.1                  : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   			

Playbook run took 0 days, 0 hours, 0 minutes, 0 seconds
`, longMessageLine),
			res: fmt.Sprintf(`
PLAY [local] *********************************************************************************************************************************************************************************

TASK [Print test message] ********************************************************************************************************************************************************************
ok: [127.0.0.1] => 
	msg: That's a message to test: %s

PLAY RECAP ***********************************************************************************************************************************************************************************
127.0.0.1                  : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   			

Playbook run took 0 days, 0 hours, 0 minutes, 0 seconds
`, longMessageLine),
			err: nil,
		},
		{
			desc: "Testing error on writing output",
			input: fmt.Sprintf(`
PLAY [local] *********************************************************************************************************************************************************************************

TASK [Print test message] ********************************************************************************************************************************************************************
ok: [127.0.0.1] => 
	msg: That's a message to test: %s

PLAY RECAP ***********************************************************************************************************************************************************************************
127.0.0.1                  : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   			

Playbook run took 0 days, 0 hours, 0 minutes, 0 seconds
`, longMessageLine),
			res:          "",
			err:          wrapError(errClosedWriter),
			closedWriter: true,
		},
		{
			desc: "Testing error while reading",
			input: fmt.Sprintf(`
PLAY [local] *********************************************************************************************************************************************************************************

TASK [Print test message] ********************************************************************************************************************************************************************
ok: [127.0.0.1] => 
	msg: That's a message to test: %s

PLAY RECAP ***********************************************************************************************************************************************************************************
127.0.0.1                  : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   			

Playbook run took 0 days, 0 hours, 0 minutes, 0 seconds
`, longMessageLine),
			res:          "",
			err:          wrapError(errClosedReader),
			closedReader: true,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			wbuff := bytes.Buffer{}
			writer := io.Writer(&wbuff)
			if test.closedWriter {
				writer = &closedWriter{}
			}

			var reader io.Reader
			reader = bufio.NewReader(strings.NewReader(test.input))
			if test.closedReader {
				reader = &closedReader{}
			}
			err := DefaultStdoutCallbackResults(context.TODO(), reader, writer, test.trans...)
			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, wbuff.String(), "Unexpected value")
			}
		})
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

var errClosedReader = errors.New("closed reader")

type closedReader struct{}

func (c *closedReader) Read(_ []byte) (n int, err error) {
	return 0, errClosedReader
}

var errClosedWriter = errors.New("closed writer")

type closedWriter struct{}

func (c closedWriter) Write(_ []byte) (n int, err error) {
	return 0, errClosedWriter
}
