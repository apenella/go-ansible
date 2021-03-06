package results

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrepend(t *testing.T) {
	desc := "Testing prepend transformer"
	t.Run(desc, func(t *testing.T) {
		t.Log(desc)

		message := "my message"
		trans := Prepend("prefix")
		message = trans(message)

		assert.Equal(t, "prefix my message", message)
	})
}

func TestAppend(t *testing.T) {
	desc := "Testing append transformer"
	t.Run(desc, func(t *testing.T) {
		t.Log(desc)

		message := "my message"
		trans := Append("suffix")
		message = trans(message)

		assert.Equal(t, "my message suffix", message)
	})

}

func TestLogFormat(t *testing.T) {
	desc := "Testing log format transformer"
	t.Run(desc, func(t *testing.T) {
		t.Log(desc)
		f := func(layout string) string {
			return time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC).Format(layout)
		}

		message := "my message"
		trans := LogFormat(DefaultLogFormatLayout, f)
		message = trans(message)

		assert.Equal(t, "2019-01-01 00:00:00	my message", message)
	})
}

func TestIgnoreMessage(t *testing.T) {

	tests := []struct {
		desc string
		line string
		res  string
	}{
		{
			desc: "Test matching line",
			line: "Playbook run took 1 days, 10 hours, 53 minutes, 27 seconds",
			res:  "",
		},
		{
			desc: "Test not matching line",
			line: "line: 'Playbook run took 1 days, 10 hours, 53 minutes, 27 seconds'",
			res:  "line: 'Playbook run took 1 days, 10 hours, 53 minutes, 27 seconds'",
		},
	}

	for _, test := range tests {

		patterns := []string{
			// This pattern skips timer's callback whitelist output
			"^[\\s\\t]*Playbook run took [0-9]+ days, [0-9]+ hours, [0-9]+ minutes, [0-9]+ seconds$",
		}

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)
			trans := IgnoreMessage(patterns)
			line := trans(test.line)

			assert.Equal(t, test.res, line)
		})
	}
}
