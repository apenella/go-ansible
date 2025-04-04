package json

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {

	tests := []struct {
		desc     string
		input    *AnsiblePlaybookJSONLEventResults
		expected string
	}{
		{
			desc: "Testing String method with a v2_playbook_on_play_start event",
			input: &AnsiblePlaybookJSONLEventResults{
				Event:     "v2_playbook_on_play_start",
				Timestamp: "2025-03-28T20:00:00.000000",
				Play: &AnsiblePlaybookJSONResultsPlaysPlay{
					Id:   "play-id",
					Name: "play-name",
					Duration: &AnsiblePlaybookJSONResultsPlayDuration{
						Start: "2025-03-28T20:00:00.000000",
						End:   "2025-03-28T20:00:00.000000",
					},
					Path: "play-path",
				},
			},
			expected: "Event: v2_playbook_on_play_start\nTimestamp: 2025-03-28T20:00:00.000000\nPlay: &{Duration:[2025-03-28T20:00:00.000000 - 2025-03-28T20:00:00.000000] Id:play-id Name:play-name Path:play-path}\n",
		},
		{
			desc: "Testing String method with a v2_runner_on_start event",
			input: &AnsiblePlaybookJSONLEventResults{
				Event:     "v2_runner_on_start",
				Timestamp: "2025-03-28T20:00:00.000000",
				Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
					Id:   "task-id",
					Name: "task-name",
					Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
						Start: "2025-03-28T20:00:00.000000",
						End:   "2025-03-28T20:00:00.000000",
					},
					Path: "task-path",
				},
			},
			expected: "Event: v2_runner_on_start\nTimestamp: 2025-03-28T20:00:00.000000\nTask: &{Duration:[2025-03-28T20:00:00.000000 - 2025-03-28T20:00:00.000000] Id:task-id Name:task-name Path:task-path}\n",
		},
		{
			desc: "Testing String method with a v2_playbook_on_task_start event",
			input: &AnsiblePlaybookJSONLEventResults{
				Event:     "v2_playbook_on_task_start",
				Timestamp: "2025-03-28T20:00:00.000000",
				Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
					Id:   "task-id",
					Name: "task-name",
					Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
						Start: "2025-03-28T20:00:00.000000",
						End:   "2025-03-28T20:00:00.000000",
					},
					Path: "task-path",
				},
			},
			expected: "Event: v2_playbook_on_task_start\nTimestamp: 2025-03-28T20:00:00.000000\nTask: &{Duration:[2025-03-28T20:00:00.000000 - 2025-03-28T20:00:00.000000] Id:task-id Name:task-name Path:task-path}\n",
		},
		{
			desc: "Testing String method with a v2_playbook_on_handler_task_start event",
			input: &AnsiblePlaybookJSONLEventResults{
				Event:     "v2_playbook_on_handler_task_start",
				Timestamp: "2025-03-28T20:00:00.000000",
				Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
					Id:   "task-id",
					Name: "task-name",
					Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
						Start: "2025-03-28T20:00:00.000000",
						End:   "2025-03-28T20:00:00.000000",
					},
					Path: "task-path",
				},
			},
			expected: "Event: v2_playbook_on_handler_task_start\nTimestamp: 2025-03-28T20:00:00.000000\nTask: &{Duration:[2025-03-28T20:00:00.000000 - 2025-03-28T20:00:00.000000] Id:task-id Name:task-name Path:task-path}\n",
		},
		{
			desc: "Testing String method with a v2_playbook_on_stats event",
			input: &AnsiblePlaybookJSONLEventResults{
				Event:     "v2_playbook_on_stats",
				Timestamp: "2025-03-28T20:00:00.000000",
				Stats: map[string]*AnsiblePlaybookJSONResultsStats{
					"host1": {
						Changed:     1,
						Failures:    2,
						Ignored:     3,
						Ok:          4,
						Rescued:     5,
						Skipped:     6,
						Unreachable: 7,
					},
				},
				CustomStats: map[string]int{
					"host1": 1,
				},
				GlobalCustomStats: map[string]string{
					"total_runtime": "1s",
				},
			},
			expected: "Event: v2_playbook_on_stats\nTimestamp: 2025-03-28T20:00:00.000000\nStats:\n  host1:\n     Changed: 1 Failures: 2 Ignored: 3 Ok: 4 Rescued: 5 Skipped: 6 Unreachable: 7\nCustomStats: map[host1:1]\nGlobalCustomStats: map[total_runtime:1s]\n",
		},
		{
			desc: "Testing String method with a v2_runner_on_ok event",
			input: &AnsiblePlaybookJSONLEventResults{
				Event:     "v2_runner_on_ok",
				Timestamp: "2025-03-28T20:00:00.000000",
				Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
					Id:   "task-id",
					Name: "task-name",
					Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
						Start: "2025-03-28T20:00:00.000000",
						End:   "2025-03-28T20:00:00.000000",
					},
					Path: "task-path",
				},
				Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
					"host1": {
						Action:           "action",
						Changed:          true,
						Msg:              "msg",
						AnsibleFacts:     map[string]interface{}{"fact1": "value1"},
						Stdout:           "stdout",
						StdoutLines:      []interface{}{"stdout-line1", "stdout-line2"},
						Stderr:           "stderr",
						StderrLines:      []interface{}{"stderr-line1", "stderr-line2"},
						Cmd:              "cmd",
						Failed:           true,
						FailedWhenResult: true,
						Skipped:          true,
						SkipReason:       "skip-reason",
						Unreachable:      true,
					},
				},
			},
			expected: "Event: v2_runner_on_ok\nTimestamp: 2025-03-28T20:00:00.000000\nTask: &{Duration:[2025-03-28T20:00:00.000000 - 2025-03-28T20:00:00.000000] Id:task-id Name:task-name Path:task-path}\nHosts:\n  host1:\n    &{Action:action Changed:true Msg:msg AnsibleFacts:map[fact1:value1] Stdout:stdout StdoutLines:[stdout-line1 stdout-line2] Stderr:stderr StderrLines:[stderr-line1 stderr-line2] Cmd:cmd Failed:true FailedWhenResult:true Skipped:true SkipReason:skip-reason Unreachable:true}\n",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			t.Log(test.desc)

			res := test.input.String()
			assert.Equal(t, test.expected, res)
		})
	}
}

func TestUnmarshalJSONV2PlaybookOnPlayStart(t *testing.T) {
	t.Log("Testing unmarshal JSON with v2_playbook_on_play_start event")
	t.Parallel()

	input := []byte(`{
		"_event": "v2_playbook_on_play_start",
		"_timestamp": "2025-03-24T12:34:56Z",
		"play": {
			"name": "Example Play",
			"id": "12345-uuid-play",
			"path": "/path/to/playbook.yml",
			"duration": {
				"start": "2025-03-24T12:34:56Z"
			}
		},
		"tasks": []
	}`)

	var result AnsiblePlaybookJSONLEventResults
	expected := AnsiblePlaybookJSONLEventResults{
		Event:     "v2_playbook_on_play_start",
		Timestamp: "2025-03-24T12:34:56Z",
		Play: &AnsiblePlaybookJSONResultsPlaysPlay{
			Name: "Example Play",
			Id:   "12345-uuid-play",
			Path: "/path/to/playbook.yml",
			Duration: &AnsiblePlaybookJSONResultsPlayDuration{
				Start: "2025-03-24T12:34:56Z",
			},
		},
		Tasks: []AnsiblePlaybookJSONResultsPlayTask{},
	}

	err := json.Unmarshal(input, &result)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestUnmarshalJSONV2OnRunnerStart(t *testing.T) {
	t.Log("Testing unmarshal JSON with v2_runner_on_start event")
	t.Parallel()

	input := []byte(`{
		"_event": "v2_runner_on_start",
		"_timestamp": "2025-03-24T12:35:00Z",
		"task": {
			"name": "Example Task",
			"id": "67890-uuid-task",
			"path": "/path/to/playbook.yml",
			"duration": {
				"start": "2025-03-24T12:35:00Z"
			}
		},
		"hosts": {}
	}`)
	var result AnsiblePlaybookJSONLEventResults
	expected := AnsiblePlaybookJSONLEventResults{
		Event:     "v2_runner_on_start",
		Timestamp: "2025-03-24T12:35:00Z",
		Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
			Name: "Example Task",
			Id:   "67890-uuid-task",
			Path: "/path/to/playbook.yml",
			Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
				Start: "2025-03-24T12:35:00Z",
			},
		},
		Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{},
	}

	err := json.Unmarshal(input, &result)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUnmarshalJSONV2PlaybookOnTaskStart(t *testing.T) {
	t.Log("Testing unmarshal JSON with v2_playbook_on_task_start event")
	t.Parallel()

	input := []byte(`{
	"_event": "v2_playbook_on_task_start",
	"_timestamp": "2025-03-24T12:36:00Z",
	"task": {
		"name": "Start Example Task",
		"id": "67890-uuid-task",
		"path": "/path/to/playbook.yml",
		"duration": {
			"start": "2025-03-24T12:36:00Z"
		}
	},
	"hosts": {}
}`)
	var result AnsiblePlaybookJSONLEventResults
	expected := AnsiblePlaybookJSONLEventResults{
		Event:     "v2_playbook_on_task_start",
		Timestamp: "2025-03-24T12:36:00Z",
		Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
			Name: "Start Example Task",
			Id:   "67890-uuid-task",
			Path: "/path/to/playbook.yml",
			Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
				Start: "2025-03-24T12:36:00Z",
			},
		},
		Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{},
	}

	err := json.Unmarshal(input, &result)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUnmarshalJSONV2OnHandlerTaskStart(t *testing.T) {
	t.Log("Testing unmarshal JSON with v2_playbook_on_handler_task_start event")
	t.Parallel()

	input := []byte(`{
		"_event": "v2_playbook_on_handler_task_start",
		"_timestamp": "2025-03-24T12:37:00Z",
		"task": {
			"name": "Example Handler Task",
			"id": "abcdef-handler-task",
			"path": "/path/to/handlers.yml",
			"duration": {
				"start": "2025-03-24T12:37:00Z"
			}
		},
		"hosts": {}
	}`)
	var result AnsiblePlaybookJSONLEventResults
	expected := AnsiblePlaybookJSONLEventResults{
		Event:     "v2_playbook_on_handler_task_start",
		Timestamp: "2025-03-24T12:37:00Z",
		Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
			Name: "Example Handler Task",
			Id:   "abcdef-handler-task",
			Path: "/path/to/handlers.yml",
			Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
				Start: "2025-03-24T12:37:00Z",
			},
		},
		Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{},
	}

	err := json.Unmarshal(input, &result)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUnmarshalJSONV2PlaybookOnStats(t *testing.T) {
	t.Log("Testing unmarshal JSON with v2_playbook_on_stats event")
	t.Parallel()

	input := []byte(`{
		"_event": "v2_playbook_on_stats",
		"_timestamp": "2025-03-24T12:40:00Z",
		"stats": {
			"host1": {
				"ok": 5,
				"changed": 2,
				"failed": 0,
				"skipped": 1,
				"unreachable": 0
			},
			"host2": {
				"ok": 3,
				"changed": 1,
				"failed": 1,
				"skipped": 0,
				"unreachable": 0
			}
		},
		"custom_stats": {
			"host1": {
				"custom_metric": "100"
			}
		},
		"global_custom_stats": {
			"total_runtime": "30s"
		}
	}`)

	var result AnsiblePlaybookJSONLEventResults
	expected := AnsiblePlaybookJSONLEventResults{
		Event:     "v2_playbook_on_stats",
		Timestamp: "2025-03-24T12:40:00Z",
		Stats: map[string]*AnsiblePlaybookJSONResultsStats{
			"host1": {
				Ok:          5,
				Changed:     2,
				Failures:    0,
				Skipped:     1,
				Unreachable: 0,
			},
			"host2": {
				Ok:          3,
				Changed:     1,
				Failures:    0,
				Skipped:     0,
				Unreachable: 0,
			},
		},
		CustomStats: map[string]interface{}{
			"host1": map[string]interface{}{
				"custom_metric": "100",
			},
		},
		GlobalCustomStats: map[string]interface{}{
			"total_runtime": "30s",
		},
	}

	err := json.Unmarshal(input, &result)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUnmarshalJSONV2RunnerOnOk(t *testing.T) {
	t.Log("Testing unmarshal JSON with v2_runner_on_ok event")
	t.Parallel()

	input := []byte(`{
		"_event": "v2_runner_on_ok",
		"_timestamp": "2025-03-24T12:38:00Z",
		"task": {
			"name": "Example Task",
			"id": "67890-uuid-task",
			"path": "/path/to/playbook.yml",
			"duration": {
				"start": "2025-03-24T12:35:00Z",
				"end": "2025-03-24T12:38:00Z"
			}
		},
		"hosts": {
			"host1": {
				"action": "shell",
				"stdout": "Task executed successfully"
			}
		}
	}`)
	var result AnsiblePlaybookJSONLEventResults
	expected := AnsiblePlaybookJSONLEventResults{
		Event:     "v2_runner_on_ok",
		Timestamp: "2025-03-24T12:38:00Z",
		Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
			Name: "Example Task",
			Id:   "67890-uuid-task",
			Path: "/path/to/playbook.yml",
			Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
				Start: "2025-03-24T12:35:00Z",
				End:   "2025-03-24T12:38:00Z",
			},
		},
		Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
			"host1": {
				Action: "shell",
				Stdout: "Task executed successfully",
			},
		},
	}

	err := json.Unmarshal(input, &result)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUnmarshalJSONV2RunnerOnFailed(t *testing.T) {
	t.Log("Testing unmarshal JSON with v2_runner_on_failed event")
	t.Parallel()

	input := []byte(`{
    "_event": "v2_runner_on_failed",
    "_timestamp": "2025-03-24T12:39:00Z",
    "task": {
      "name": "Failing Task",
      "id": "abcdef-failed-task",
      "path": "/path/to/playbook.yml",
      "duration": {
        "start": "2025-03-24T12:38:30Z",
        "end": "2025-03-24T12:39:00Z"
      }
    },
    "hosts": {
      "host2": {
        "action": "command",
        "stderr": "Error executing command",
        "failed": true
      }
    }
}`)
	var result AnsiblePlaybookJSONLEventResults
	expected := AnsiblePlaybookJSONLEventResults{
		Event:     "v2_runner_on_failed",
		Timestamp: "2025-03-24T12:39:00Z",
		Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
			Name: "Failing Task",
			Id:   "abcdef-failed-task",
			Path: "/path/to/playbook.yml",
			Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
				Start: "2025-03-24T12:38:30Z",
				End:   "2025-03-24T12:39:00Z",
			},
		},
		Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
			"host2": {
				Action: "command",
				Stderr: "Error executing command",
				Failed: true,
			},
		},
	}

	err := json.Unmarshal(input, &result)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUnmarshalJSONV2RunnerOnUnreachable(t *testing.T) {
	t.Log("Testing unmarshal JSON with v2_runner_on_unreachable event")
	t.Parallel()

	input := []byte(`{
		"_event": "v2_runner_on_unreachable",
		"_timestamp": "2025-03-24T12:39:30Z",
		"task": {
			"name": "Unreachable Task",
			"id": "ghijkl-unreachable-task",
			"path": "/path/to/playbook.yml",
			"duration": {
				"start": "2025-03-24T12:39:00Z",
				"end": "2025-03-24T12:39:30Z"
			}
		},
		"hosts": {
			"host3": {
				"action": "ping",
				"msg": "Host unreachable",
				"unreachable": true
			}
		}
	}`)
	var result AnsiblePlaybookJSONLEventResults
	expected := AnsiblePlaybookJSONLEventResults{
		Event:     "v2_runner_on_unreachable",
		Timestamp: "2025-03-24T12:39:30Z",
		Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
			Name: "Unreachable Task",
			Id:   "ghijkl-unreachable-task",
			Path: "/path/to/playbook.yml",
			Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
				Start: "2025-03-24T12:39:00Z",
				End:   "2025-03-24T12:39:30Z",
			},
		},
		Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
			"host3": {
				Action:      "ping",
				Msg:         "Host unreachable",
				Unreachable: true,
			},
		},
	}

	err := json.Unmarshal(input, &result)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUnmarshalJSONV2RunnerOnSkipped(t *testing.T) {
	t.Log("Testing unmarshal JSON with v2_runner_on_skipped event")
	t.Parallel()

	input := []byte(`{
		"_event": "v2_runner_on_skipped",
		"_timestamp": "2025-03-24T12:40:00Z",
		"task": {
			"name": "Skipped Task",
			"id": "mnopqr-skipped-task",
			"path": "/path/to/playbook.yml",
			"duration": {
				"start": "2025-03-24T12:39:50Z",
				"end": "2025-03-24T12:40:00Z"
			}
		},
		"hosts": {
			"host4": {
				"action": "copy",
				"msg": "Task skipped",
				"skipped": true
			}
		}
	}`)
	var result AnsiblePlaybookJSONLEventResults
	expected := AnsiblePlaybookJSONLEventResults{
		Event:     "v2_runner_on_skipped",
		Timestamp: "2025-03-24T12:40:00Z",
		Task: &AnsiblePlaybookJSONResultsPlayTaskItem{
			Name: "Skipped Task",
			Id:   "mnopqr-skipped-task",
			Path: "/path/to/playbook.yml",
			Duration: &AnsiblePlaybookJSONResultsPlayTaskItemDuration{
				Start: "2025-03-24T12:39:50Z",
				End:   "2025-03-24T12:40:00Z",
			},
		},
		Hosts: map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem{
			"host4": {
				Action:  "copy",
				Msg:     "Task skipped",
				Skipped: true,
			},
		},
	}

	err := json.Unmarshal(input, &result)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
