package results

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	errors "github.com/apenella/go-common-utils/error"
)

// AnsiblePlaybookJSONResults
type AnsiblePlaybookJSONResults struct {
	Playbook          string                                      `json:"-"`
	CustomStats       interface{}                                 `json:"custom_stats"`
	GlobalCustomStats interface{}                                 `json:"global_custom_stats"`
	Plays             []AnsiblePlaybookJSONResultsPlay            `json:"plays"`
	Stats             map[string]*AnsiblePlaybookJSONResultsStats `json:"stats"`
}

func (r *AnsiblePlaybookJSONResults) String() string {

	str := ""

	for _, play := range r.Plays {
		for _, task := range play.Tasks {
			name := task.Task.Name
			for host, result := range task.Hosts {
				str = fmt.Sprintf("%s[%s] (%s)	%s\n", str, host, name, result.Msg)
			}
		}
	}

	for host, stats := range r.Stats {
		str = fmt.Sprintf("%s\nHost: %s\n%s\n", str, host, stats.String())
	}

	return str
}

// CheckStats return error when is found a failure or unreachable host
func (r *AnsiblePlaybookJSONResults) CheckStats() error {
	errorMsg := ""
	for host, stats := range r.Stats {
		if stats.Failures > 0 {
			errorMsg = fmt.Sprintf("Host %s finished with %d failures", host, stats.Failures)
		}

		if stats.Unreachable > 0 {
			errorMsg = fmt.Sprintf("Host %s finished with %d unrecheable hosts", host, stats.Unreachable)
		}

		if len(errorMsg) > 0 {
			return errors.New("(results::JSONStdoutCallbackResults)", errorMsg)
		}
	}

	return nil
}

// AnsiblePlaybookJSONResultsPlay
type AnsiblePlaybookJSONResultsPlay struct {
	Play  *AnsiblePlaybookJSONResultsPlaysPlay `json:"play"`
	Tasks []AnsiblePlaybookJSONResultsPlayTask `json:"tasks"`
}

// AnsiblePlaybookJSONResultsPlaysPlay
type AnsiblePlaybookJSONResultsPlaysPlay struct {
	Name     string                                  `json:"name"`
	Id       string                                  `json:"id"`
	Duration *AnsiblePlaybookJSONResultsPlayDuration `json:"duration"`
}

/* AnsiblePlaybookJSONResultsPlayTask
'task': {
	'name': task.get_name(),
	'id': to_text(task._uuid),
	'duration': {
		'start': current_time()
	}
},
'hosts': {}
*/
type AnsiblePlaybookJSONResultsPlayTask struct {
	Task  *AnsiblePlaybookJSONResultsPlayTaskItem                 `json:"task"`
	Hosts map[string]*AnsiblePlaybookJSONResultsPlayTaskHostsItem `json:"hosts"`
}

type AnsiblePlaybookJSONResultsPlayTaskHostsItem struct {
	Action           string                 `json:"action"`
	Changed          bool                   `json:"changed"`
	Msg              interface{}            `json:"msg"`
	AnsibleFacts     map[string]interface{} `json:"ansible_facts"`
	Stdout           string                 `json:"stdout"`
	StdoutLines      []string               `json:"stdout_lines"`
	Stderr           string                 `json:"stderr"`
	StderrLines      []string               `json:"stderr_lines"`
	Cmd              interface{}            `json:"cmd"`
	Failed           bool                   `json:"failed"`
	FailedWhenResult bool                   `json:"failed_when_result"`
	Skipped          bool                   `json:"skipped"`
	SkipReason       string                 `json:"skip_reason"`
	Unreachable      string                 `json:"unreachable"`
}

type AnsiblePlaybookJSONResultsPlayTaskItem struct {
	Name     string                                          `json:"name"`
	Id       string                                          `json:"id"`
	Duration *AnsiblePlaybookJSONResultsPlayTaskItemDuration `json:"duration"`
}

type AnsiblePlaybookJSONResultsPlayTaskItemDuration struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// AnsiblePlaybookJSONResultsPlayDuration
type AnsiblePlaybookJSONResultsPlayDuration struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// AnsiblePlaybookJSONResultsStats
type AnsiblePlaybookJSONResultsStats struct {
	Changed     int `json:"changed"`
	Failures    int `json:"failures"`
	Ignored     int `json:"ignored"`
	Ok          int `json:"ok"`
	Rescued     int `json:"rescued"`
	Skipped     int `json:"skipped"`
	Unreachable int `json:"unreachable"`
}

func (s *AnsiblePlaybookJSONResultsStats) String() string {
	str := fmt.Sprintf(" Changed: %s", strconv.Itoa(s.Changed))
	str = fmt.Sprintf("%s Failures: %s", str, strconv.Itoa(s.Failures))
	str = fmt.Sprintf("%s Ignored: %s", str, strconv.Itoa(s.Ignored))
	str = fmt.Sprintf("%s Ok: %s", str, strconv.Itoa(s.Ok))
	str = fmt.Sprintf("%s Rescued: %s", str, strconv.Itoa(s.Rescued))
	str = fmt.Sprintf("%s Skipped: %s", str, strconv.Itoa(s.Skipped))
	str = fmt.Sprintf("%s Unreachable: %s", str, strconv.Itoa(s.Unreachable))

	return str
}

// JSONStdoutCallbackResults method manges the ansible' JSON stdout callback and print the result stats
func JSONStdoutCallbackResults(ctx context.Context, r io.Reader, w io.Writer, transformers ...TransformerFunc) error {

	skipPatterns := []string{
		// This pattern skips timer's callback whitelist output
		"^[\\s\\t]*Playbook run took [0-9]+ days, [0-9]+ hours, [0-9]+ minutes, [0-9]+ seconds$",
	}

	tranformers := []TransformerFunc{
		IgnoreMessage(skipPatterns),
	}

	err := output(ctx, r, w, tranformers...)
	if err != nil {
		return errors.New("(results::JSONStdoutCallbackResults)", "Error processing execution output", err)
	}

	return nil
}

// JSONParse return an AnsiblePlaybookJSONResults from
func JSONParse(data []byte) (*AnsiblePlaybookJSONResults, error) {

	result := &AnsiblePlaybookJSONResults{}

	err := json.Unmarshal(data, result)
	if err != nil {
		return nil, errors.New("(results::JSONParser)", "Unmarshall error", err)
	}

	return result, nil
}

// ParseJSONResultsStream parse the ansible' JSON stdout callback and return an AnsiblePlaybookJSONResults
func ParseJSONResultsStream(stream io.Reader) (*AnsiblePlaybookJSONResults, error) {
	decoder := json.NewDecoder(stream)
	results := &AnsiblePlaybookJSONResults{}
	for {
		err := decoder.Decode(results)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("(results::JSONParser)", "Error decoding results", err)
		}
	}

	return results, nil
}
