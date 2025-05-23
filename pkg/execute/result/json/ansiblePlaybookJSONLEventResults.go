package json

import "fmt"

// AnsiblePlaybookJSONLEventResults represents the structure of the JSON lines generated by an Ansible playbook execution using the ansible.posix.jsonl callback plugin
type AnsiblePlaybookJSONLEventResults struct {
	Event     string `json:"_event"`
	Timestamp string `json:"_timestamp"`

	CustomStats       interface{}                                             `json:"custom_stats"`
	GlobalCustomStats interface{}                                             `json:"global_custom_stats"`
	Hosts             map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem `json:"hosts,omitempty"`
	Play              *AnsiblePlaybookJSONResultsPlaysPlay                    `json:"play,omitempty"`
	Stats             map[string]*AnsiblePlaybookJSONResultsStats             `json:"stats"`
	Task              *AnsiblePlaybookJSONResultsPlayTaskItem                 `json:"task,omitempty"`
	Tasks             []AnsiblePlaybookJSONResultsPlayTask                    `json:"tasks,omitempty"`
}

// String returns a string representation of AnsiblePlaybookJSONLEventResults
func (a AnsiblePlaybookJSONLEventResults) String() string {

	str := fmt.Sprintf("Event: %s\nTimestamp: %s\n", a.Event, a.Timestamp)

	if a.Play != nil {
		str = fmt.Sprintf("%sPlay: %+v\n", str, a.Play)
	}

	if a.Task != nil {
		str = fmt.Sprintf("%sTask: %+v\n", str, a.Task)
	}

	if len(a.Hosts) > 0 {

		str = fmt.Sprintf("%sHosts:\n", str)
		for k, v := range a.Hosts {
			str = fmt.Sprintf("%s  %s:\n    %s\n", str, k, fmt.Sprintf("%+v", v))
		}
	}

	if len(a.Tasks) > 0 {
		str = fmt.Sprintf("%sTasks:\n", str)

		for _, task := range a.Tasks {
			str = fmt.Sprintf("%s  %+v\n", str, task)
		}
	}

	if len(a.Stats) > 0 {
		str = fmt.Sprintf("%sStats:\n", str)
		for k, v := range a.Stats {
			str = fmt.Sprintf("%s  %s:\n    %s\n", str, k, fmt.Sprintf("%+v", v))
		}
	}

	if a.CustomStats != nil {
		str = fmt.Sprintf("%sCustomStats: %+v\n", str, a.CustomStats)
	}

	if a.GlobalCustomStats != nil {
		str = fmt.Sprintf("%sGlobalCustomStats: %+v\n", str, a.GlobalCustomStats)
	}

	return str

}
