package stdoutcallback

// func TestGetResultsFunc(t *testing.T) {
// 	tests := []struct {
// 		desc     string
// 		callback string
// 		res      StdoutCallbackResultsFunc
// 	}{
// 		{
// 			desc:     "Testing get JSON results func",
// 			callback: JSONStdoutCallback,
// 			res:      results.JSONStdoutCallbackResults,
// 		},
// 		{
// 			desc:     "Testing get default results func",
// 			callback: DefaultStdoutCallback,
// 			res:      output.DefaultStdoutCallbackResults,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.desc, func(t *testing.T) {
// 			t.Log(test.desc)

// 			f := GetResultsFunc(test.callback)
// 			assert.Equal(t, runtime.FuncForPC(reflect.ValueOf(test.res).Pointer()), runtime.FuncForPC(reflect.ValueOf(f).Pointer()))

// 		})
// 	}

// }
