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
	Task  *AnsiblePlaybookJSONResultsPlayTaskItem `json:"task"`
	Hosts map[string]map[string]interface{}       `json:"hosts"`
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
	str := fmt.Sprintf(" Changed: %s\n", strconv.Itoa(s.Changed))
	str = fmt.Sprintf("%s Failures: %s\n", str, strconv.Itoa(s.Failures))
	str = fmt.Sprintf("%s Ignored: %s\n", str, strconv.Itoa(s.Ignored))
	str = fmt.Sprintf("%s Ok: %s\n", str, strconv.Itoa(s.Ok))
	str = fmt.Sprintf("%s Rescued: %s\n", str, strconv.Itoa(s.Rescued))
	str = fmt.Sprintf("%s Skipped: %s\n", str, strconv.Itoa(s.Skipped))
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

// AnsibleJsonParse parse output and retrive playbook stats
// func (r *PlaybookResults) AnsibleJsonParse(e *Executor) error {
// 	stdout := e.Stdout

// 	r.RawStdout = stdout

// 	//verify json
// 	if !gjson.Valid(stdout) {
// 		return errors.New("(ansible:Run) -> invalid json returned after ansible run")
// 	}

// 	// retrive stats from ansible run
// 	//changed
// 	value := gjson.Get(stdout, "stats.*.changed")
// 	if !value.Exists() {
// 		return errors.New("(ansible:Run) -> cannot retrive changed from ansible run")
// 	} else {
// 		r.Changed = value.Int()
// 	}
// 	//changed
// 	value = gjson.Get(stdout, "stats.*.failures")
// 	if !value.Exists() {
// 		return errors.New("(ansible:Run) -> cannot retrive failures from ansible run")
// 	} else {
// 		r.Failures = value.Int()
// 	}
// 	//ignored
// 	value = gjson.Get(stdout, "stats.*.ignored")
// 	if !value.Exists() {
// 		return errors.New("(ansible:Run) -> cannot retrive ignored from ansible run")
// 	} else {
// 		r.Ignored = value.Int()
// 	}
// 	//ok
// 	value = gjson.Get(stdout, "stats.*.ok")
// 	if !value.Exists() {
// 		return errors.New("(ansible:Run) -> cannot retrive ok from ansible run")
// 	} else {
// 		r.Ok = value.Int()
// 	}
// 	//rescued
// 	value = gjson.Get(stdout, "stats.*.rescued")
// 	if !value.Exists() {
// 		return errors.New("(ansible:Run) -> cannot retrive rescued from ansible run")
// 	} else {
// 		r.Rescued = value.Int()
// 	}
// 	//skipped
// 	value = gjson.Get(stdout, "stats.*.skipped")
// 	if !value.Exists() {
// 		return errors.New("(ansible:Run) -> cannot retrive skipped from ansible run")
// 	} else {
// 		r.Skipped = value.Int()
// 	}
// 	//unreachable
// 	value = gjson.Get(stdout, "stats.*.unreachable")
// 	if !value.Exists() {
// 		return errors.New("(ansible:Run) -> cannot retrive unreachable from ansible run")
// 	} else {
// 		r.Unreachable = value.Int()
// 	}

// 	return nil
// }

// PlaybookResultsChecks return a error if a critical issue is found in playbook stats
// func (r *PlaybookResults) PlaybookResultsChecks() error {
// 	if r == nil {
// 		return errors.New("(ansible:PlaybookResultsChecks) -> passed result is nil")
// 	}
// 	if r.Unreachable > 0 {
// 		return errors.New("(ansible:Run) -> host is not reachable")
// 	}
// 	if r.Failures > 0 {
// 		return errors.New("(ansible:Run) -> one of tasks defined in playbook is failing")
// 	}
// 	return nil
// }
