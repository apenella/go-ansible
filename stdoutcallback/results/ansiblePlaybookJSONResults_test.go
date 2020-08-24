package results

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdoutCallbackJSONResults(t *testing.T) {

	tests := []struct {
		desc           string
		inputResult    string
		expectedResult string
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
			expectedResult: `[127.0.0.1] (Print line)	That's a message to debug
[192.198.1.1] (Print line)	That's a message to debug

Host: 127.0.0.1
 Changed: 0 Failures: 0 Ignored: 0 Ok: 1 Rescued: 0 Skipped: 0 Unreachable: 0

Host: 192.168.1.1
 Changed: 0 Failures: 0 Ignored: 0 Ok: 1 Rescued: 0 Skipped: 0 Unreachable: 0

`,
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
			expectedResult: `[127.0.0.1] (Print line)	That's a message to debug

Host: 127.0.0.1
 Changed: 0 Failures: 0 Ignored: 0 Ok: 1 Rescued: 0 Skipped: 0 Unreachable: 0

`,
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
			expectedResult: "",
			err:            errors.New("(results::JSONStdoutCallbackResults) Host 127.0.0.1 finished with 1 failures"),
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
			expectedResult: "",
			err:            errors.New("(results::JSONStdoutCallbackResults) Host 127.0.0.1 finished with 1 unrecheable hosts"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			wbuff := bytes.Buffer{}
			writer := io.Writer(&wbuff)
			reader := bufio.NewReader(strings.NewReader(test.inputResult))
			err := JSONStdoutCallbackResults("prefix", reader, writer)
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
										Msg:     "That's a message to debug",
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
			reader := bufio.NewReader(strings.NewReader(test.inputResult))
			res, _ := JSONParse(reader)

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

func TestSkipLine(t *testing.T) {
	tests := []struct {
		desc string
		line string
		res  bool
	}{
		{
			desc: "Test matching line",
			line: "Playbook run took 1 days, 10 hours, 53 minutes, 27 seconds",
			res:  true,
		},
		{
			desc: "Test not matching line",
			line: "line: 'Playbook run took 1 days, 10 hours, 53 minutes, 27 seconds'",
			res:  false,
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		res := skipLine(test.line)
		assert.Equal(t, test.res, res, "Unexpected result")
	}
}
