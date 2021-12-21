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
