package results

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestStdoutCallbackJSONResults(t *testing.T) {

	t.Skip()

	tests := []struct {
		desc           string
		inputResult    string
		expectedResult string
		trans          []TransformerFunc
		err            error
	}{
		{
			desc: "Testing stdout callback json result",
			inputResult: `{
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
										"msg": "That's a message to debug"
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
							},
							{
								"hosts": {
									"192.198.1.1": {
										"_ansible_no_log": false,
										"_ansible_verbose_always": true,
										"action": "debug",
										"changed": false,
										"msg": "That's a message to debug"
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
						"ok": 1,
						"rescued": 0,
						"skipped": 0,
						"unreachable": 0
					},
					"192.168.1.1": {
						"changed": 0,
						"failures": 0,
						"ignored": 0,
						"ok": 1,
						"rescued": 0,
						"skipped": 0,
						"unreachable": 0
					}
				}
			}`,
			expectedResult: `{				"custom_stats": {},				"global_custom_stats": {},				"plays": [					{						"play": {							"duration": {								"end": "2020-08-07T20:51:30.942955Z",								"start": "2020-08-07T20:51:30.607525Z"							},							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",							"name": "local"						},						"tasks": [							{								"hosts": {									"127.0.0.1": {										"_ansible_no_log": false,										"_ansible_verbose_always": true,										"action": "debug",										"changed": false,										"msg": "That's a message to debug"									}								},								"task": {									"duration": {										"end": "2020-08-07T20:51:30.942955Z",										"start": "2020-08-07T20:51:30.908539Z"									},									"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",									"name": "Print line"								}							},							{								"hosts": {									"192.198.1.1": {										"_ansible_no_log": false,										"_ansible_verbose_always": true,										"action": "debug",										"changed": false,										"msg": "That's a message to debug"									}								},								"task": {									"duration": {										"end": "2020-08-07T20:51:30.942955Z",										"start": "2020-08-07T20:51:30.908539Z"									},									"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",									"name": "Print line"								}							}						]					}				],				"stats": {					"127.0.0.1": {						"changed": 0,						"failures": 0,						"ignored": 0,						"ok": 1,						"rescued": 0,						"skipped": 0,						"unreachable": 0					},					"192.168.1.1": {						"changed": 0,						"failures": 0,						"ignored": 0,						"ok": 1,						"rescued": 0,						"skipped": 0,						"unreachable": 0					}				}			}`,
			err: nil,
		},
		{
			desc: "Testing stdout callback json result skipping lines",
			inputResult: `{
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
										"msg": "That's a message to debug"
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
						"ok": 1,
						"rescued": 0,
						"skipped": 0,
						"unreachable": 0
					}
				}
			}
			Playbook run took 0 days, 0 hours, 0 minutes, 0 seconds`,
			expectedResult: `{				"custom_stats": {},				"global_custom_stats": {},				"plays": [					{						"play": {							"duration": {								"end": "2020-08-07T20:51:30.942955Z",								"start": "2020-08-07T20:51:30.607525Z"							},							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",							"name": "local"						},						"tasks": [							{								"hosts": {									"127.0.0.1": {										"_ansible_no_log": false,										"_ansible_verbose_always": true,										"action": "debug",										"changed": false,										"msg": "That's a message to debug"									}								},								"task": {									"duration": {										"end": "2020-08-07T20:51:30.942955Z",										"start": "2020-08-07T20:51:30.908539Z"									},									"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",									"name": "Print line"								}							}						]					}				],				"stats": {					"127.0.0.1": {						"changed": 0,						"failures": 0,						"ignored": 0,						"ok": 1,						"rescued": 0,						"skipped": 0,						"unreachable": 0					}				}			}`,
			err: nil,
		},
		{
			desc: "Testing stdout callback json result with failures",
			inputResult: `{
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
										"msg": "That's a message to debug"
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
						"failures": 1,
						"ignored": 0,
						"ok": 0,
						"rescued": 0,
						"skipped": 0,
						"unreachable": 0
					}
				}
			}`,
			expectedResult: `{				"custom_stats": {},				"global_custom_stats": {},				"plays": [					{						"play": {							"duration": {								"end": "2020-08-07T20:51:30.942955Z",								"start": "2020-08-07T20:51:30.607525Z"							},							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",							"name": "local"						},						"tasks": [							{								"hosts": {									"127.0.0.1": {										"_ansible_no_log": false,										"_ansible_verbose_always": true,										"action": "debug",										"changed": false,										"msg": "That's a message to debug"									}								},								"task": {									"duration": {										"end": "2020-08-07T20:51:30.942955Z",										"start": "2020-08-07T20:51:30.908539Z"									},									"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",									"name": "Print line"								}							}						]					}				],				"stats": {					"127.0.0.1": {						"changed": 0,						"failures": 1,						"ignored": 0,						"ok": 0,						"rescued": 0,						"skipped": 0,						"unreachable": 0					}				}			}`,
			err: errors.New("(results::JSONStdoutCallbackResults)", "Host 127.0.0.1 finished with 1 failures"),
		},
		{
			desc: "Testing stdout callback json result with hosts unrecheable",
			inputResult: `{
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
										"msg": "That's a message to debug"
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
			expectedResult: `{				"custom_stats": {},				"global_custom_stats": {},				"plays": [					{						"play": {							"duration": {								"end": "2020-08-07T20:51:30.942955Z",								"start": "2020-08-07T20:51:30.607525Z"							},							"id": "a0a4c5d1-62fd-b6f1-98ea-000000000006",							"name": "local"						},						"tasks": [							{								"hosts": {									"127.0.0.1": {										"_ansible_no_log": false,										"_ansible_verbose_always": true,										"action": "debug",										"changed": false,										"msg": "That's a message to debug"									}								},								"task": {									"duration": {										"end": "2020-08-07T20:51:30.942955Z",										"start": "2020-08-07T20:51:30.908539Z"									},									"id": "a0a4c5d1-62fd-b6f1-98ea-000000000008",									"name": "Print line"								}							}						]					}				],				"stats": {					"127.0.0.1": {						"changed": 0,						"failures": 0,						"ignored": 0,						"ok": 0,						"rescued": 0,						"skipped": 0,						"unreachable": 1					}				}			}`,
			err: errors.New("(results::JSONStdoutCallbackResults)", "Host 127.0.0.1 finished with 1 unrecheable hosts"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			wbuff := bytes.Buffer{}
			writer := io.Writer(&wbuff)
			reader := bufio.NewReader(strings.NewReader(test.inputResult))
			err := JSONStdoutCallbackResults(context.TODO(), reader, writer, test.trans...)
			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.expectedResult, wbuff.String(), "Unexpected value")
			}
		})
	}

}

func TestJSONParser(t *testing.T) {

	tests := []struct {
		desc        string
		inputResult string
		res         *AnsiblePlaybookJSONResults
	}{
		{
			desc: "Testing json parse",
			inputResult: `{
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
										"msg": ["That's a message to debug"]
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
						"ok": 1,
						"rescued": 0,
						"skipped": 0,
						"unreachable": 0
					}
				}
			}`,
			res: &AnsiblePlaybookJSONResults{
				CustomStats:       map[string]interface{}{},
				GlobalCustomStats: map[string]interface{}{},
				Plays: []AnsiblePlaybookJSONResultsPlay{
					{
						Play: &AnsiblePlaybookJSONResultsPlaysPlay{
							Name: "local",
							Id:   "a0a4c5d1-62fd-b6f1-98ea-000000000006",
							Duration: &AnsiblePlaybookJSONResultsPlayDuration{
								End:   "2020-08-07T20:51:30.942955Z",
								Start: "2020-08-07T20:51:30.607525Z",
							},
						},
						Tasks: []AnsiblePlaybookJSONResultsPlayTask{
							{
								Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
									Id:   "a0a4c5d1-62fd-b6f1-98ea-000000000008",
									Name: "Print line",
									Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
										End:   "2020-08-07T20:51:30.942955Z",
										Start: "2020-08-07T20:51:30.908539Z",
									},
								},
								// TODOx
								Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
									"127.0.0.1": {
										//"_ansible_no_log": false, "_ansible_verbose_always": true,
										Action:  "debug",
										Changed: false,
										Msg:     []interface{}{"That's a message to debug"},
									},
								},
							},
						},
					},
				},
				Stats: map[string]*AnsiblePlaybookJSONResultsStats{
					"127.0.0.1": {
						Changed:     0,
						Failures:    0,
						Ignored:     0,
						Ok:          1,
						Rescued:     0,
						Skipped:     0,
						Unreachable: 0,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res, _ := JSONParse([]byte(test.inputResult))
			assert.Equal(t, test.res, res, "Unexpected result")
		})
	}
}

func TestParseJSONResultsStream(t *testing.T) {

	tests := []struct {
		desc        string
		inputResult string
		res         *AnsiblePlaybookJSONResults
	}{
		{
			desc: "Testing json parse",
			inputResult: `{
				"custom_stats": {},
				"global_custom_stats": {},
				"plays": [
					{
						"play": {
							"duration": {
								"end": "2021-12-21T06:55:29.890926Z",
								"start": "2021-12-21T06:55:29.881536Z"
							},
							"id": "3982ba1a-4acb-67e8-84e1-000000000006",
							"name": "all"
						},
						"tasks": [
							{
								"hosts": {
									"127.0.0.1": {
										"_ansible_no_log": false,
										"_ansible_verbose_always": true,
										"action": "debug",
										"changed": false,
										"msg": "Your are running\n'json-stdout-ansibleplaybook'\nfirst example\n"
									}
								},
								"task": {
									"duration": {
										"end": "2021-12-21T06:55:29.890926Z",
										"start": "2021-12-21T06:55:29.886253Z"
									},
									"id": "3982ba1a-4acb-67e8-84e1-000000000008",
									"name": "json-stdout-ansibleplaybook"
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
						"ok": 1,
						"rescued": 0,
						"skipped": 0,
						"unreachable": 0
					}
				}
			}
			{
				"custom_stats": {},
				"global_custom_stats": {},
				"plays": [
					{
						"play": {
							"duration": {
								"end": "2021-12-21T06:55:29.890926Z",
								"start": "2021-12-21T06:55:29.881536Z"
							},
							"id": "3982ba1a-4acb-67e8-84e1-000000000006",
							"name": "all"
						},
						"tasks": [
							{
								"hosts": {
									"127.0.0.1": {
										"_ansible_no_log": false,
										"_ansible_verbose_always": true,
										"action": "debug",
										"changed": false,
										"msg": "Your are running\n'json-stdout-ansibleplaybook'\nfirst example\n"
									}
								},
								"task": {
									"duration": {
										"end": "2021-12-21T06:55:29.890926Z",
										"start": "2021-12-21T06:55:29.886253Z"
									},
									"id": "3982ba1a-4acb-67e8-84e1-000000000008",
									"name": "json-stdout-ansibleplaybook"
								}
							}
						]
					},
					{
						"play": {
							"duration": {
								"end": "2021-12-21T06:55:29.901245Z",
								"start": "2021-12-21T06:55:29.894953Z"
							},
							"id": "3982ba1a-4acb-67e8-84e1-00000000001a",
							"name": "all"
						},
						"tasks": [
							{
								"hosts": {
									"127.0.0.1": {
										"_ansible_no_log": false,
										"_ansible_verbose_always": true,
										"action": "debug",
										"changed": false,
										"msg": "Your are running\n'json-stdout-ansibleplaybook'\nsecond example\n"
									}
								},
								"task": {
									"duration": {
										"end": "2021-12-21T06:55:29.901245Z",
										"start": "2021-12-21T06:55:29.896772Z"
									},
									"id": "3982ba1a-4acb-67e8-84e1-00000000001c",
									"name": "json-stdout-ansibleplaybook"
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
						"ok": 2,
						"rescued": 0,
						"skipped": 0,
						"unreachable": 0
					}
				}
			}
			{
				"custom_stats": {},
				"global_custom_stats": {},
				"plays": [
					{
						"play": {
							"duration": {
								"end": "2021-12-21T06:55:29.890926Z",
								"start": "2021-12-21T06:55:29.881536Z"
							},
							"id": "3982ba1a-4acb-67e8-84e1-000000000006",
							"name": "all"
						},
						"tasks": [
							{
								"hosts": {
									"127.0.0.1": {
										"_ansible_no_log": false,
										"_ansible_verbose_always": true,
										"action": "debug",
										"changed": false,
										"msg": "Your are running\n'json-stdout-ansibleplaybook'\nfirst example\n"
									}
								},
								"task": {
									"duration": {
										"end": "2021-12-21T06:55:29.890926Z",
										"start": "2021-12-21T06:55:29.886253Z"
									},
									"id": "3982ba1a-4acb-67e8-84e1-000000000008",
									"name": "json-stdout-ansibleplaybook"
								}
							}
						]
					},
					{
						"play": {
							"duration": {
								"end": "2021-12-21T06:55:29.901245Z",
								"start": "2021-12-21T06:55:29.894953Z"
							},
							"id": "3982ba1a-4acb-67e8-84e1-00000000001a",
							"name": "all"
						},
						"tasks": [
							{
								"hosts": {
									"127.0.0.1": {
										"_ansible_no_log": false,
										"_ansible_verbose_always": true,
										"action": "debug",
										"changed": false,
										"msg": "Your are running\n'json-stdout-ansibleplaybook'\nsecond example\n"
									}
								},
								"task": {
									"duration": {
										"end": "2021-12-21T06:55:29.901245Z",
										"start": "2021-12-21T06:55:29.896772Z"
									},
									"id": "3982ba1a-4acb-67e8-84e1-00000000001c",
									"name": "json-stdout-ansibleplaybook"
								}
							}
						]
					},
					{
						"play": {
							"duration": {
								"end": "2021-12-21T06:55:29.910879Z",
								"start": "2021-12-21T06:55:29.904593Z"
							},
							"id": "3982ba1a-4acb-67e8-84e1-00000000002e",
							"name": "all"
						},
						"tasks": [
							{
								"hosts": {
									"127.0.0.1": {
										"_ansible_no_log": false,
										"_ansible_verbose_always": true,
										"action": "debug",
										"changed": false,
										"msg": "Your are running\n'json-stdout-ansibleplaybook'\nthird example\n"
									}
								},
								"task": {
									"duration": {
										"end": "2021-12-21T06:55:29.910879Z",
										"start": "2021-12-21T06:55:29.906537Z"
									},
									"id": "3982ba1a-4acb-67e8-84e1-000000000030",
									"name": "json-stdout-ansibleplaybook"
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
						"ok": 3,
						"rescued": 0,
						"skipped": 0,
						"unreachable": 0
					}
				}
			}`,
			res: &AnsiblePlaybookJSONResults{

				CustomStats:       map[string]interface{}{},
				GlobalCustomStats: map[string]interface{}{},
				Plays: []AnsiblePlaybookJSONResultsPlay{
					{
						Play: &AnsiblePlaybookJSONResultsPlaysPlay{
							Name: "all",
							Id:   "3982ba1a-4acb-67e8-84e1-000000000006",
							Duration: &AnsiblePlaybookJSONResultsPlayDuration{
								End:   "2021-12-21T06:55:29.890926Z",
								Start: "2021-12-21T06:55:29.881536Z",
							},
						},
						Tasks: []AnsiblePlaybookJSONResultsPlayTask{
							{
								Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
									Id:   "3982ba1a-4acb-67e8-84e1-000000000008",
									Name: "json-stdout-ansibleplaybook",
									Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
										End:   "2021-12-21T06:55:29.890926Z",
										Start: "2021-12-21T06:55:29.886253Z",
									},
								},
								Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
									"127.0.0.1": {
										// "_ansible_no_log": false, "_ansible_verbose_always": true,
										Action:  "debug",
										Changed: false,
										Msg:     "Your are running\n'json-stdout-ansibleplaybook'\nfirst example\n",
									},
								},
							},
						},
					},
					{
						Play: &AnsiblePlaybookJSONResultsPlaysPlay{
							Name: "all",
							Id:   "3982ba1a-4acb-67e8-84e1-00000000001a",
							Duration: &AnsiblePlaybookJSONResultsPlayDuration{
								End:   "2021-12-21T06:55:29.901245Z",
								Start: "2021-12-21T06:55:29.894953Z",
							},
						},
						Tasks: []AnsiblePlaybookJSONResultsPlayTask{
							{
								Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
									Id:   "3982ba1a-4acb-67e8-84e1-00000000001c",
									Name: "json-stdout-ansibleplaybook",
									Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
										End:   "2021-12-21T06:55:29.901245Z",
										Start: "2021-12-21T06:55:29.896772Z",
									},
								},
								Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
									"127.0.0.1": {
										// "_ansible_no_log": false, "_ansible_verbose_always": true,
										Action:  "debug",
										Changed: false,
										Msg:     "Your are running\n'json-stdout-ansibleplaybook'\nsecond example\n",
									},
								},
							},
						},
					},
					{
						Play: &AnsiblePlaybookJSONResultsPlaysPlay{
							Name: "all",
							Id:   "3982ba1a-4acb-67e8-84e1-00000000002e",
							Duration: &AnsiblePlaybookJSONResultsPlayDuration{
								End:   "2021-12-21T06:55:29.910879Z",
								Start: "2021-12-21T06:55:29.904593Z",
							},
						},
						Tasks: []AnsiblePlaybookJSONResultsPlayTask{
							{
								Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
									Id:   "3982ba1a-4acb-67e8-84e1-000000000030",
									Name: "json-stdout-ansibleplaybook",
									Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
										End:   "2021-12-21T06:55:29.910879Z",
										Start: "2021-12-21T06:55:29.906537Z",
									},
								},
								Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
									"127.0.0.1": {
										// "_ansible_no_log": false, "_ansible_verbose_always": true,
										Action:  "debug",
										Changed: false,
										Msg:     "Your are running\n'json-stdout-ansibleplaybook'\nthird example\n",
									},
								},
							},
						},
					},
				},
				Stats: map[string]*AnsiblePlaybookJSONResultsStats{
					"127.0.0.1": {
						Changed:     0,
						Failures:    0,
						Ignored:     0,
						Ok:          3,
						Rescued:     0,
						Skipped:     0,
						Unreachable: 0,
					},
				},
			},
		},
		{
			desc: "Testing json parse skipping and failing tasks",
			inputResult: `{
					"custom_stats": {},
					"global_custom_stats": {},
					"plays": [
						{
							"play": {
								"duration": {
									"end": "2022-02-08T16:51:13.677774Z",
									"start": "2022-02-08T16:51:12.808956Z"
								},
								"id": "201e881a-804c-dd08-2927-000000000006",
								"name": "all"
							},
							"tasks": [
								{
									"hosts": {
										"127.0.0.1": {
											"_ansible_no_log": false,
											"action": "command",
											"changed": true,
											"cmd": "/usr/bin/true",
											"delta": "0:00:00.002663",
											"deprecations": [
												{
													"msg": "Distribution fedora 35 on host 127.0.0.1 should use /usr/bin/python3, but is using /usr/bin/python for backward compatibility with prior Ansible releases. A future Ansible release will default to using the discovered platform python for this host. See https://docs.ansible.com/ansible/2.9/reference_appendices/interpreter_discovery.html for more information",
													"version": "2.12"
												}
											],
											"end": "2022-02-08 17:51:13.094418",
											"invocation": {
												"module_args": {
													"_raw_params": "/usr/bin/true",
													"_uses_shell": true,
													"argv": null,
													"chdir": null,
													"creates": null,
													"executable": null,
													"removes": null,
													"stdin": null,
													"stdin_add_newline": true,
													"strip_empty_ends": true,
													"warn": true
												}
											},
											"rc": 0,
											"start": "2022-02-08 17:51:13.091755",
											"stderr": "",
											"stderr_lines": [],
											"stdout": "",
											"stdout_lines": []
										}
									},
									"task": {
										"duration": {
											"end": "2022-02-08T16:51:13.112192Z",
											"start": "2022-02-08T16:51:12.818849Z"
										},
										"id": "201e881a-804c-dd08-2927-000000000008",
										"name": "ok-task"
									}
								},
								{
									"hosts": {
										"127.0.0.1": {
											"_ansible_no_log": false,
											"action": "ansible.builtin.shell",
											"changed": false,
											"skip_reason": "Conditional result was False",
											"skipped": true
										}
									},
									"task": {
										"duration": {
											"end": "2022-02-08T16:51:13.144340Z",
											"start": "2022-02-08T16:51:13.113955Z"
										},
										"id": "201e881a-804c-dd08-2927-000000000009",
										"name": "skipping-task"
									}
								},
								{
									"hosts": {
										"127.0.0.1": {
											"_ansible_no_log": false,
											"action": "command",
											"changed": true,
											"cmd": "exit -1",
											"delta": "0:00:00.003074",
											"end": "2022-02-08 17:51:13.300085",
											"failed": true,
											"invocation": {
												"module_args": {
													"_raw_params": "exit -1",
													"_uses_shell": true,
													"argv": null,
													"chdir": null,
													"creates": null,
													"executable": null,
													"removes": null,
													"stdin": null,
													"stdin_add_newline": true,
													"strip_empty_ends": true,
													"warn": true
												}
											},
											"msg": "non-zero return code",
											"rc": 255,
											"start": "2022-02-08 17:51:13.297011",
											"stderr": "",
											"stderr_lines": [],
											"stdout": "",
											"stdout_lines": []
										}
									},
									"task": {
										"duration": {
											"end": "2022-02-08T16:51:13.320178Z",
											"start": "2022-02-08T16:51:13.146031Z"
										},
										"id": "201e881a-804c-dd08-2927-00000000000a",
										"name": "failing-task"
									}
								},
								{
									"hosts": {
										"127.0.0.1": {
											"_ansible_no_log": false,
											"action": "ansible.builtin.command",
											"changed": true,
											"cmd": [
												"/usr/bin/ls",
												"/tmp/foobar.baz"
											],
											"delta": "0:00:00.002326",
											"end": "2022-02-08 17:51:13.621549",
											"failed": true,
											"failed_when_result": true,
											"invocation": {
												"module_args": {
													"_raw_params": "/usr/bin/ls /tmp/foobar.baz",
													"_uses_shell": false,
													"argv": null,
													"chdir": null,
													"creates": null,
													"executable": null,
													"removes": null,
													"stdin": null,
													"stdin_add_newline": true,
													"strip_empty_ends": true,
													"warn": true
												}
											},
											"msg": "non-zero return code",
											"rc": 2,
											"start": "2022-02-08 17:51:13.619223",
											"stderr": "/usr/bin/ls: cannot access '/tmp/foobar.baz': No such file or directory",
											"stderr_lines": [
												"/usr/bin/ls: cannot access '/tmp/foobar.baz': No such file or directory"
											],
											"stdout": "",
											"stdout_lines": []
										}
									},
									"task": {
										"duration": {
											"end": "2022-02-08T16:51:13.677774Z",
											"start": "2022-02-08T16:51:13.322404Z"
										},
										"id": "201e881a-804c-dd08-2927-00000000000b",
										"name": "failing-task-when"
									}
								}
							]
						}
					],
					"stats": {
						"127.0.0.1": {
							"changed": 2,
							"failures": 1,
							"ignored": 1,
							"ok": 2,
							"rescued": 0,
							"skipped": 1,
							"unreachable": 0
						}
					}
				}`,
			res: &AnsiblePlaybookJSONResults{

				CustomStats:       map[string]interface{}{},
				GlobalCustomStats: map[string]interface{}{},
				Plays: []AnsiblePlaybookJSONResultsPlay{
					{
						Play: &AnsiblePlaybookJSONResultsPlaysPlay{
							Name: "all",
							Id:   "201e881a-804c-dd08-2927-000000000006",
							Duration: &AnsiblePlaybookJSONResultsPlayDuration{
								End:   "2022-02-08T16:51:13.677774Z",
								Start: "2022-02-08T16:51:12.808956Z",
							},
						},
						Tasks: []AnsiblePlaybookJSONResultsPlayTask{
							{
								Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
									Id:   "201e881a-804c-dd08-2927-000000000008",
									Name: "ok-task",
									Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
										End:   "2022-02-08T16:51:13.112192Z",
										Start: "2022-02-08T16:51:12.818849Z",
									},
								},
								Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
									"127.0.0.1": {
										// "_ansible_no_log": false, "_ansible_verbose_always": true,
										Action:           "command",
										Changed:          true,
										Stdout:           "",
										StdoutLines:      []string{},
										StderrLines:      []string{},
										Cmd:              "/usr/bin/true",
										Failed:           false,
										FailedWhenResult: false,
										Skipped:          false,
										SkipReason:       "",
									},
								},
							},
							{
								Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
									Id:   "201e881a-804c-dd08-2927-000000000009",
									Name: "skipping-task",
									Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
										End:   "2022-02-08T16:51:13.144340Z",
										Start: "2022-02-08T16:51:13.113955Z",
									},
								},
								Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
									"127.0.0.1": {
										// "_ansible_no_log": false, "_ansible_verbose_always": true,
										Action:           "ansible.builtin.shell",
										Changed:          false,
										Failed:           false,
										FailedWhenResult: false,
										Skipped:          true,
										SkipReason:       "Conditional result was False",
									},
								},
							},
							{
								Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
									Id:   "201e881a-804c-dd08-2927-00000000000a",
									Name: "failing-task",
									Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
										End:   "2022-02-08T16:51:13.320178Z",
										Start: "2022-02-08T16:51:13.146031Z",
									},
								},
								Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
									"127.0.0.1": {
										// "_ansible_no_log": false, "_ansible_verbose_always": true,
										Action:           "command",
										Changed:          true,
										Msg:              "non-zero return code",
										StdoutLines:      []string{},
										Stderr:           "",
										StderrLines:      []string{},
										Cmd:              "exit -1",
										Failed:           true,
										FailedWhenResult: false,
										Skipped:          false,
										SkipReason:       "",
									},
								},
							},
							{
								Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
									Id:   "201e881a-804c-dd08-2927-00000000000b",
									Name: "failing-task-when",
									Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
										End:   "2022-02-08T16:51:13.677774Z",
										Start: "2022-02-08T16:51:13.322404Z",
									},
								},
								Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
									"127.0.0.1": {
										// "_ansible_no_log": false, "_ansible_verbose_always": true,
										Action:           "ansible.builtin.command",
										Changed:          true,
										Msg:              "non-zero return code",
										StdoutLines:      []string{},
										Stderr:           "/usr/bin/ls: cannot access '/tmp/foobar.baz': No such file or directory",
										StderrLines:      []string{"/usr/bin/ls: cannot access '/tmp/foobar.baz': No such file or directory"},
										Cmd:              []interface{}{"/usr/bin/ls", "/tmp/foobar.baz"},
										Failed:           true,
										FailedWhenResult: true,
										Skipped:          false,
										SkipReason:       "",
									},
								},
							},
						},
					},
				},
				Stats: map[string]*AnsiblePlaybookJSONResultsStats{
					"127.0.0.1": {
						Changed:     2,
						Failures:    1,
						Ignored:     1,
						Ok:          2,
						Rescued:     0,
						Skipped:     1,
						Unreachable: 0,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res, err := ParseJSONResultsStream(strings.NewReader(test.inputResult))
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			assert.Equal(t, test.res, res, "Unexpected result")
		})
	}
}

func TestAnsiblePlaybookJSONResultsString(t *testing.T) {
	tests := []struct {
		desc    string
		results *AnsiblePlaybookJSONResults
		res     string
	}{
		{
			desc:    "Testing empty result to string",
			results: &AnsiblePlaybookJSONResults{},
			res:     "",
		},
		{
			desc: "Testing json result to string",
			results: &AnsiblePlaybookJSONResults{
				Stats: map[string]*AnsiblePlaybookJSONResultsStats{
					"127.0.0.1": {
						Changed:     0,
						Failures:    0,
						Ignored:     0,
						Ok:          0,
						Rescued:     0,
						Skipped:     0,
						Unreachable: 0,
					},
				},
			},
			res: `
Host: 127.0.0.1
 Changed: 0 Failures: 0 Ignored: 0 Ok: 0 Rescued: 0 Skipped: 0 Unreachable: 0
`,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res := test.results.String()
			assert.Equal(t, test.res, res, "Unexpected result")
		})
	}
}

func TestAnsiblePlaybookJSONResultsStatsString(t *testing.T) {
	tests := []struct {
		desc  string
		stats *AnsiblePlaybookJSONResultsStats
		res   string
	}{
		{
			desc: "Testing json result stats to string",
			stats: &AnsiblePlaybookJSONResultsStats{
				Changed:     0,
				Failures:    0,
				Ignored:     0,
				Ok:          0,
				Rescued:     0,
				Skipped:     0,
				Unreachable: 0,
			},
			res: " Changed: 0 Failures: 0 Ignored: 0 Ok: 0 Rescued: 0 Skipped: 0 Unreachable: 0",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res := test.stats.String()
			assert.Equal(t, test.res, res, "Unexpected result")
		})
	}
}
