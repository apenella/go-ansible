package json

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"strings"
	"testing"

	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestPrint(t *testing.T) {

	longMessageLine := randStringBytes(512_000)
	_ = longMessageLine
	wbuff := bytes.Buffer{}

	tests := []struct {
		desc           string
		inputResult    string
		expectedResult string
		result         *JSONStdoutCallbackResults
		reader         io.Reader
		writer         io.Writer
		trans          []transformer.TransformerFunc
		err            error
	}{
		// {
		// 	desc:   "Testing error when reader is not provided in the Print method",
		// 	result: NewJSONStdoutCallbackResults(),
		// },
		// {
		// 	desc:   "Testing stdout callback json result",
		// 	result: NewJSONStdoutCallbackResults(),
		// 	writer: io.Writer(&wbuff),
		// 	reader: bufio.NewReader(strings.NewReader(`{
		// 		"custom_stats": {},
		// 		"global_custom_stats": {},
		// 		"plays": [
		// 			{
		// 				"play": {
		// 					"duration": {
		// 						"end": "2020-08-07T20:51:30.942955Z",
		// 						"start": "2020-08-07T20:51:30.607525Z"
		// 					},
		// 					"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",
		// 					"name": "local"
		// 				},
		// 				"tasks": [
		// 					{
		// 						"hosts": {
		// 							"127.0.0.1": {
		// 								"_ansible_no_log": false,
		// 								"_ansible_verbose_always": true,
		// 								"action": "debug",
		// 								"changed": false,
		// 								"msg": "That's a message to debug"
		// 							}
		// 						},
		// 						"task": {
		// 							"duration": {
		// 								"end": "2020-08-07T20:51:30.942955Z",
		// 								"start": "2020-08-07T20:51:30.908539Z"
		// 							},
		// 							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",
		// 							"name": "Print line"
		// 						}
		// 					},
		// 					{
		// 						"hosts": {
		// 							"192.198.1.1": {
		// 								"_ansible_no_log": false,
		// 								"_ansible_verbose_always": true,
		// 								"action": "debug",
		// 								"changed": false,
		// 								"msg": "That's a message to debug"
		// 							}
		// 						},
		// 						"task": {
		// 							"duration": {
		// 								"end": "2020-08-07T20:51:30.942955Z",
		// 								"start": "2020-08-07T20:51:30.908539Z"
		// 							},
		// 							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",
		// 							"name": "Print line"
		// 						}
		// 					}
		// 				]
		// 			}
		// 		],
		// 		"stats": {
		// 			"127.0.0.1": {
		// 				"changed": 0,
		// 				"failures": 0,
		// 				"ignored": 0,
		// 				"ok": 1,
		// 				"rescued": 0,
		// 				"skipped": 0,
		// 				"unreachable": 0
		// 			},
		// 			"192.168.1.1": {
		// 				"changed": 0,
		// 				"failures": 0,
		// 				"ignored": 0,
		// 				"ok": 1,
		// 				"rescued": 0,
		// 				"skipped": 0,
		// 				"unreachable": 0
		// 			}
		// 		}
		// 	}`)),
		// 	expectedResult: `{				"custom_stats": {},				"global_custom_stats": {},				"plays": [					{						"play": {							"duration": {								"end": "2020-08-07T20:51:30.942955Z",								"start": "2020-08-07T20:51:30.607525Z"							},							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",							"name": "local"						},						"tasks": [							{								"hosts": {									"127.0.0.1": {										"_ansible_no_log": false,										"_ansible_verbose_always": true,										"action": "debug",										"changed": false,										"msg": "That's a message to debug"									}								},								"task": {									"duration": {										"end": "2020-08-07T20:51:30.942955Z",										"start": "2020-08-07T20:51:30.908539Z"									},									"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",									"name": "Print line"								}							},							{								"hosts": {									"192.198.1.1": {										"_ansible_no_log": false,										"_ansible_verbose_always": true,										"action": "debug",										"changed": false,										"msg": "That's a message to debug"									}								},								"task": {									"duration": {										"end": "2020-08-07T20:51:30.942955Z",										"start": "2020-08-07T20:51:30.908539Z"									},									"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",									"name": "Print line"								}							}						]					}				],				"stats": {					"127.0.0.1": {						"changed": 0,						"failures": 0,						"ignored": 0,						"ok": 1,						"rescued": 0,						"skipped": 0,						"unreachable": 0					},					"192.168.1.1": {						"changed": 0,						"failures": 0,						"ignored": 0,						"ok": 1,						"rescued": 0,						"skipped": 0,						"unreachable": 0					}				}			}`,
		// 	err:            nil,
		// },
		// {
		// 	desc:   "Testing very long file in json output",
		// 	result: NewJSONStdoutCallbackResults(),
		// 	writer: io.Writer(&wbuff),
		// 	reader: bufio.NewReader(strings.NewReader(fmt.Sprintf(`{
		// 		"custom_stats": {},
		// 		"global_custom_stats": {},
		// 		"plays": [
		// 			{
		// 				"play": {
		// 					"duration": {
		// 						"end": "2020-08-07T20:51:30.942955Z",
		// 						"start": "2020-08-07T20:51:30.607525Z"
		// 					},
		// 					"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",
		// 					"name": "%s"
		// 				},
		// 				"tasks": []
		// 			}
		// 		],
		// 		"stats": {
		// 			"127.0.0.1": {
		// 				"changed": 0,
		// 				"failures": 0,
		// 				"ignored": 0,
		// 				"ok": 1,
		// 				"rescued": 0,
		// 				"skipped": 0,
		// 				"unreachable": 0
		// 			}
		// 		}
		// 	}`, longMessageLine))),
		// 	expectedResult: fmt.Sprintf(`{"custom_stats": {},"global_custom_stats": {},"plays": [{"play": {"duration": {"end": "2020-08-07T20:51:30.942955Z","start": "2020-08-07T20:51:30.607525Z"},"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006","name": "%s"},"tasks": []}],"stats": {"127.0.0.1": {"changed": 0,"failures": 0,"ignored": 0,"ok": 1,"rescued": 0,"skipped": 0,"unreachable": 0}}}`, longMessageLine),
		// 	err:            nil,
		// },
		// {
		// 	desc:   "Testing stdout callback json result skipping lines",
		// 	result: NewJSONStdoutCallbackResults(),
		// 	writer: io.Writer(&wbuff),
		// 	reader: bufio.NewReader(strings.NewReader(`{
		// 		"custom_stats": {},
		// 		"global_custom_stats": {},
		// 		"plays": [
		// 			{
		// 				"play": {
		// 					"duration": {
		// 						"end": "2020-08-07T20:51:30.942955Z",
		// 						"start": "2020-08-07T20:51:30.607525Z"
		// 					},
		// 					"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",
		// 					"name": "local"
		// 				},
		// 				"tasks": [
		// 					{
		// 						"hosts": {
		// 							"127.0.0.1": {
		// 								"_ansible_no_log": false,
		// 								"_ansible_verbose_always": true,
		// 								"action": "debug",
		// 								"changed": false,
		// 								"msg": "That's a message to debug"
		// 							}
		// 						},
		// 						"task": {
		// 							"duration": {
		// 								"end": "2020-08-07T20:51:30.942955Z",
		// 								"start": "2020-08-07T20:51:30.908539Z"
		// 							},
		// 							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",
		// 							"name": "Print line"
		// 						}
		// 					}
		// 				]
		// 			}
		// 		],
		// 		"stats": {
		// 			"127.0.0.1": {
		// 				"changed": 0,
		// 				"failures": 0,
		// 				"ignored": 0,
		// 				"ok": 1,
		// 				"rescued": 0,
		// 				"skipped": 0,
		// 				"unreachable": 0
		// 			}
		// 		}
		// 	}
		// 	Playbook run took 0 days, 0 hours, 0 minutes, 0 seconds`)),
		// 	expectedResult: `{				"custom_stats": {},				"global_custom_stats": {},				"plays": [					{						"play": {							"duration": {								"end": "2020-08-07T20:51:30.942955Z",								"start": "2020-08-07T20:51:30.607525Z"							},							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",							"name": "local"						},						"tasks": [							{								"hosts": {									"127.0.0.1": {										"_ansible_no_log": false,										"_ansible_verbose_always": true,										"action": "debug",										"changed": false,										"msg": "That's a message to debug"									}								},								"task": {									"duration": {										"end": "2020-08-07T20:51:30.942955Z",										"start": "2020-08-07T20:51:30.908539Z"									},									"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",									"name": "Print line"								}							}						]					}				],				"stats": {					"127.0.0.1": {						"changed": 0,						"failures": 0,						"ignored": 0,						"ok": 1,						"rescued": 0,						"skipped": 0,						"unreachable": 0					}				}			}`,
		// 	err:            nil,
		// },
		// {
		// 	desc:   "Testing stdout callback json result with failures",
		// 	result: NewJSONStdoutCallbackResults(),
		// 	writer: io.Writer(&wbuff),
		// 	reader: bufio.NewReader(strings.NewReader(`{
		// 		"custom_stats": {},
		// 		"global_custom_stats": {},
		// 		"plays": [
		// 			{
		// 				"play": {
		// 					"duration": {
		// 						"end": "2020-08-07T20:51:30.942955Z",
		// 						"start": "2020-08-07T20:51:30.607525Z"
		// 					},
		// 					"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",
		// 					"name": "local"
		// 				},
		// 				"tasks": [
		// 					{
		// 						"hosts": {
		// 							"127.0.0.1": {
		// 								"_ansible_no_log": false,
		// 								"_ansible_verbose_always": true,
		// 								"action": "debug",
		// 								"changed": false,
		// 								"msg": "That's a message to debug"
		// 							}
		// 						},
		// 						"task": {
		// 							"duration": {
		// 								"end": "2020-08-07T20:51:30.942955Z",
		// 								"start": "2020-08-07T20:51:30.908539Z"
		// 							},
		// 							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",
		// 							"name": "Print line"
		// 						}
		// 					}
		// 				]
		// 			}
		// 		],
		// 		"stats": {
		// 			"127.0.0.1": {
		// 				"changed": 0,
		// 				"failures": 1,
		// 				"ignored": 0,
		// 				"ok": 0,
		// 				"rescued": 0,
		// 				"skipped": 0,
		// 				"unreachable": 0
		// 			}
		// 		}
		// 	}`)),
		// 	expectedResult: `{				"custom_stats": {},				"global_custom_stats": {},				"plays": [					{						"play": {							"duration": {								"end": "2020-08-07T20:51:30.942955Z",								"start": "2020-08-07T20:51:30.607525Z"							},							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",							"name": "local"						},						"tasks": [							{								"hosts": {									"127.0.0.1": {										"_ansible_no_log": false,										"_ansible_verbose_always": true,										"action": "debug",										"changed": false,										"msg": "That's a message to debug"									}								},								"task": {									"duration": {										"end": "2020-08-07T20:51:30.942955Z",										"start": "2020-08-07T20:51:30.908539Z"									},									"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",									"name": "Print line"								}							}						]					}				],				"stats": {					"127.0.0.1": {						"changed": 0,						"failures": 1,						"ignored": 0,						"ok": 0,						"rescued": 0,						"skipped": 0,						"unreachable": 0					}				}			}`,
		// 	err:            errors.New("(results::JSONStdoutCallbackResults)", "Host 127.0.0.1 finished with 1 failures"),
		// },
		{
			desc:   "Testing stdout callback json result with hosts unrecheable",
			result: NewJSONStdoutCallbackResults(),
			writer: io.Writer(&wbuff),
			reader: bufio.NewReader(strings.NewReader(`{
				"custom_stats": {},
				"global_custom_stats": {},
				"plays": [
					{
						"play": {
							"duration": {
								"end": "2020-08-07T20:51:30.942955Z",
								"start": "2020-08-07T20:51:30.607525Z"
							},
							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",
							"name": "local"
						},
						"tasks": [
							{
								"hosts": {
									"127.0.0.1": {
										"_ansible_no_log": false,
										"_ansible_verbose_always": true,
										"action": "debug",
										"changed": false,
										"msg": "That's a message to debug",
										"unreachable": true
									}
								},
								"task": {
									"duration": {
										"end": "2020-08-07T20:51:30.942955Z",
										"start": "2020-08-07T20:51:30.908539Z"
									},
									"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",
									"name": "Print line"
								}
							}
						]
					}
				],
				"stats": {
					"127.0.0.1": {
						"changed": 0,
						"failures": 0,
						"ignored": 0,
						"ok": 0,
						"rescued": 0,
						"skipped": 0,
						"unreachable": 1
					}
				}
			}`)),
			expectedResult: `{
				"custom_stats": {},
				"global_custom_stats": {},
				"plays": [
					{
						"play": {
							"duration": {
								"end": "2020-08-07T20:51:30.942955Z",
								"start": "2020-08-07T20:51:30.607525Z"
							},
							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",
							"name": "local"
						},
						"tasks": [
							{
								"hosts": {
									"127.0.0.1": {
										"_ansible_no_log": false,
										"_ansible_verbose_always": true,
										"action": "debug",
										"changed": false,
										"msg": "That's a message to debug",
										"unreachable": true
									}
								},
								"task": {
									"duration": {
										"end": "2020-08-07T20:51:30.942955Z",
										"start": "2020-08-07T20:51:30.908539Z"
									},
									"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",
									"name": "Print line"
								}
							}
						]
					}
				],
				"stats": {
					"127.0.0.1": {
						"changed": 0,
						"failures": 0,
						"ignored": 0,
						"ok": 0,
						"rescued": 0,
						"skipped": 0,
						"unreachable": 1
					}
				}
			}`,
			err: errors.New("(results::JSONStdoutCallbackResults)", "Host 127.0.0.1 finished with 1 unrecheable hosts"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			t.Log(test.desc)

			wbuff.Reset()
			err := test.result.Print(context.TODO(), test.reader, test.writer, WithTransformers(test.trans...))
			if err != nil {
				assert.Equal(t, test.err, err)
			} else {
				var expectedResult, actualResult interface{}
				err := json.Unmarshal([]byte(test.expectedResult), &expectedResult)
				if err != nil {
					assert.Fail(t, "Failed to unmarshal json", err)
					return
				}

				err = json.Unmarshal(wbuff.Bytes(), &actualResult)
				if err != nil {
					assert.Fail(t, "Failed to unmarshal json", err)
					return
				}

				assert.Equal(t, expectedResult, actualResult, "Unexpected value")
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
