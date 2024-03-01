package defaultresult

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"testing"

	"github.com/apenella/go-ansible/v2/mocks"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestPrint(t *testing.T) {

	longMessageLine := randStringBytes(512_000)
	// errContext := "(results::Print)"

	tests := []struct {
		desc   string
		input  string
		res    string
		err    error
		result *DefaultResults
		trans  []transformer.TransformerFunc
	}{
		{
			desc:   "Testing default stdout callback",
			result: NewDefaultResults(),
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
			err: &errors.Error{},
		},
		{
			desc:   "Testing very long lines reading",
			result: NewDefaultResults(),
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
			err: &errors.Error{},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			var reader io.Reader

			wbuff := bytes.Buffer{}
			writer := io.Writer(&wbuff)

			reader = bufio.NewReader(strings.NewReader(test.input))

			err := test.result.Print(context.TODO(), reader, writer, WithTransformers(test.trans...))
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				assert.Equal(t, test.res, wbuff.String(), "Unexpected value")
			}
		})
	}
}

func TestOutput(t *testing.T) {

	buff := bytes.Buffer{}
	longMessageLine := randStringBytes(512_000) + "\n"
	errContext := "(results::output)"

	tests := []struct {
		desc              string
		reader            io.Reader
		writer            io.Writer
		err               error
		res               string
		trans             []transformer.TransformerFunc
		prepareAssertFunc func(io.Reader, io.Writer)
	}{
		{
			desc:   "Testing process an output message",
			reader: strings.NewReader("output message"),
			writer: io.Writer(&buff),
			res:    "output message\n",
		},
		{
			desc:   "Testing process a long output message",
			reader: bytes.NewReader([]byte(longMessageLine)),
			writer: io.Writer(&buff),
			res:    longMessageLine,
		},
		{
			desc:   "Testing error reading output message",
			reader: mocks.NewMockIOReader(),
			writer: io.Writer(&buff),
			res:    longMessageLine,
			prepareAssertFunc: func(reader io.Reader, writer io.Writer) {
				if reader != nil {
					reader.(*mocks.MockIOReader).On(
						"Read",
						mock.Anything,
					).Return(0, errors.New(errContext, "error while reading"))
				}
			},
			err: errors.New(errContext, "error while reading"),
		},
		{
			desc:   "Testing error writing output message",
			reader: bytes.NewReader([]byte(longMessageLine)),
			writer: mocks.NewMockIOWriter(),
			res:    longMessageLine,
			prepareAssertFunc: func(reader io.Reader, writer io.Writer) {
				if writer != nil {
					writer.(*mocks.MockIOWriter).On(
						"Write",
						mock.Anything,
					).Return(0, errors.New(errContext, "error while writing"))
				}
			},
			err: errors.New(errContext, "error while writing"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			buff.Reset()

			if test.prepareAssertFunc != nil {
				test.prepareAssertFunc(test.reader, test.writer)
			}

			err := output(context.TODO(), test.reader, test.writer, test.trans...)
			if err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				assert.Equal(t, test.res, buff.String(), "Unexpected value")
			}
		})
	}
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
