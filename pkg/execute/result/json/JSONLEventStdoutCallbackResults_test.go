package json

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockWriter mocks an io.Writer
type MockWriter struct {
	mock.Mock
}

func (w *MockWriter) Write(p []byte) (n int, err error) {
	args := w.Called(p)
	return args.Int(0), args.Error(1)
}

var events = `{"_event":"v2_playbook_on_play_start","_timestamp":"2025-04-01T05:17:36.646328Z","play":{"duration":{"start":"2025-04-01T05:17:36.646322Z"},"id":"cdeaff0f-de61-3a76-e91c-000000000002","name":"all","path":"site.yml:3"},"tasks":[]}`
var invalidEvent = `{"_event":"v2_playbook_on_play_start",`

func TestJSONLEventStdoutCallbackResults_Print(t *testing.T) {

	tests := []struct {
		desc        string
		context     context.Context
		writer      io.Writer
		reader      io.Reader
		results     *JSONLEventStdoutCallbackResults
		err         error
		arrangeFunc func(t *testing.T, w *MockWriter)
		assertFunc  func(t *testing.T, w *MockWriter)
	}{
		{
			desc:    "Testing JSONLEventStdoutCallbackResults Print method handle the data as expected",
			context: context.TODO(),
			writer:  &MockWriter{},
			reader:  strings.NewReader(events),
			results: NewJSONLEventStdoutCallbackResults(),
			arrangeFunc: func(t *testing.T, w *MockWriter) {
				w.On("Write", []byte(events)).Return(len([]byte("data transformed")), nil)
			},
			assertFunc: func(t *testing.T, w *MockWriter) {
				w.AssertExpectations(t)
			},
			err: nil,
		},
		{
			desc:    "Testing error in JSONLEventStdoutCallbackResults Print method when there is an error writing data to the writer",
			context: context.TODO(),
			writer:  &MockWriter{},
			reader:  strings.NewReader(events),
			results: NewJSONLEventStdoutCallbackResults(),
			arrangeFunc: func(t *testing.T, w *MockWriter) {
				w.On("Write", []byte(events)).Return(0, fmt.Errorf("error from writer"))
			},
			assertFunc: func(t *testing.T, w *MockWriter) {
				w.AssertExpectations(t)
			},
			err: errors.New("(result::json::JSONLEventStdoutCallbackResults::Print)", "Error writing to writer", fmt.Errorf("error from writer")),
		},
		{
			desc:    "Testing error in JSONLEventStdoutCallbackResults Print method when received data is not a JSON",
			context: context.TODO(),
			writer:  &MockWriter{},
			reader:  strings.NewReader(invalidEvent),
			results: NewJSONLEventStdoutCallbackResults(),
			err:     fmt.Errorf("Error reading the results stream\n\tError processing the execution output\n\terror decoding JSON: unexpected end of JSON input"),
		},
		{
			desc:    "Testing error in JSONLEventStdoutCallbackResults Print using a transformer",
			context: context.TODO(),
			writer:  &MockWriter{},
			reader:  strings.NewReader(events),
			results: NewJSONLEventStdoutCallbackResults(
				WithJSONLEventTransformers(func(data string) string {
					return "data transformed"
				}),
			),
			arrangeFunc: func(t *testing.T, w *MockWriter) {
				w.On("Write", []byte("data transformed")).Return(len([]byte("data transformed")), nil)
			},
			assertFunc: func(t *testing.T, w *MockWriter) {
				w.AssertExpectations(t)
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			if test.arrangeFunc != nil {
				test.arrangeFunc(t, test.writer.(*MockWriter))
			}

			err := test.results.Print(test.context, test.reader, test.writer)
			if err != nil && test.err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				assert.Nil(t, err, "unexpected error")
				assert.Nil(t, test.err, "expected error not reproduced")

				if test.assertFunc != nil {
					test.assertFunc(t, test.writer.(*MockWriter))
				}
			}
		})
	}
}
