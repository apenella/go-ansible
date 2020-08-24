package results

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

// AnsiblePlaybookJSONResults
type AnsiblePlaybookJSONResults struct {
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
	Action  string `json:"action"`
	Changed bool   `json:"changed"`
	Msg     string `json:"msg"`
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
func JSONStdoutCallbackResults(prefix string, r io.Reader, w io.Writer) error {

	var buff bytes.Buffer
	result := &AnsiblePlaybookJSONResults{}

	reader := bufio.NewReader(r)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {

		line := scanner.Text()
		if !skipLine(line) {
			fmt.Fprintf(io.Writer(&buff), "%s", line)
		}
	}

	err := json.Unmarshal(buff.Bytes(), result)
	if err != nil {
		return err
	}

	errorMsg := ""
	for host, stats := range result.Stats {

		if stats.Failures > 0 {
			errorMsg = fmt.Sprintf("Host %s finished with %d failures", host, stats.Failures)
		}

		if stats.Unreachable > 0 {
			errorMsg = fmt.Sprintf("Host %s finished with %d unrecheable hosts", host, stats.Unreachable)
		}

		if len(errorMsg) > 0 {
			return errors.New("(results::JSONStdoutCallbackResults) " + errorMsg)
		}
	}

	fmt.Fprintln(w, result.String())

	return nil
}

func skipLine(line string) bool {
	skipPatterns := []string{
		// This pattern skips timer's callback whitelist output
		"^[\\s\\t]*Playbook run took [0-9]+ days, [0-9]+ hours, [0-9]+ minutes, [0-9]+ seconds$",
	}

	for _, pattern := range skipPatterns {
		match, _ := regexp.MatchString(pattern, line)
		if match {
			return true
		}
	}

	return false
}

func JSONParse(reader *bufio.Reader) (*AnsiblePlaybookJSONResults, error) {

	var buff bytes.Buffer
	result := &AnsiblePlaybookJSONResults{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Fprintf(io.Writer(&buff), "%s", scanner.Text())
	}

	err := json.Unmarshal(buff.Bytes(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
