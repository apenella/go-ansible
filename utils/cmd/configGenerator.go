package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/iancoleman/strcase"
)

type config struct {
	description string
	env         string
	name        string
	vartype     string
}

func LoadConfigs(url string) []*config {
	configs := []*config{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the section with the ID "common-options"
	commonOptionsSection := doc.Find("section#common-options")

	// Select elements within the "common-options" section
	commonOptionsSubSections := commonOptionsSection.Find("section")

	commonOptionsSubSections.Each(func(i int, selection *goquery.Selection) {
		c := &config{name: strings.TrimSpace(selection.Find("h3").Text())}

		selection.Find("dl.field-list.simple").Each(func(i int, fieldListSimpleItem *goquery.Selection) {
			if i == 0 {

				fieldListSimpleItem.Find("dt").Each(func(i int, s *goquery.Selection) {
					if s.Text() == "Description:" {
						c.description = strings.TrimSpace(s.Next().Text())
					}

					if s.Text() == "Type:" {
						c.vartype = strings.TrimSpace(s.Next().Text())
					}

					// Environment.Variable
					if s.Text() == "Variable:" {
						envvarraw := strings.TrimSpace(s.Next().Text())
						envvar := strings.Split(envvarraw, "\n")
						if len(envvar) > 1 {
							c.description = fmt.Sprintf("%s %s", c.description, envvar[1:])
						}
						c.env = strings.TrimSpace(envvar[0])
					}
				})

			}
		})
		configs = append(configs, c)
	})

	return configs
}

func generateConst(config *config) string {

	str := ""
	if config.env == "" {
		str = fmt.Sprintf("\t// Parameter '%s' can not be configured by environment variable.\n// %s\n", config.name, config.description)
	} else {

		varname := strcase.ToCamel(config.env)
		str = fmt.Sprintf("\t// %s (%s) %s\n", varname, config.vartype, config.description)
		str = fmt.Sprintf("%s\t%s = \"%s\"\n", str, varname, config.env)
	}
	return str
}

func generateConsts(configs []*config) string {
	str := "const (\n"
	for _, config := range configs {
		str = fmt.Sprintf("%s%s\n", str, generateConst(config))
	}
	str = fmt.Sprintf("%s)\n", str)
	return str
}

func generateConfigMethod(config *config) string {
	str := ""
	if config.env != "" {

		optionName := strcase.ToCamel(config.env)

		switch config.vartype {
		case "boolean":
			str = fmt.Sprintf("// With%s sets the option %s to true (%s)\n", optionName, config.env, config.description)
			str = fmt.Sprintf("%sfunc (e *ExecutorWithAnsibleConfigurationSettings) With%s() *ExecutorWithAnsibleConfigurationSettings {\n	e.configurationSettings[%s] = \"true\"\n	return e\n}\n", str, optionName, optionName)

			str = fmt.Sprintf("%s// Without%s sets the option %s to false\n", str, optionName, config.env)
			str = fmt.Sprintf("%sfunc (e *ExecutorWithAnsibleConfigurationSettings) Without%s() *ExecutorWithAnsibleConfigurationSettings {\n	e.configurationSettings[%s] = \"false\"\n	return e\n}\n", str, optionName, optionName)
		case "integer":
			str = fmt.Sprintf("// With%s sets the value for the configuraion %s (%s)\n", optionName, config.env, config.description)
			str = fmt.Sprintf("%sfunc (e *ExecutorWithAnsibleConfigurationSettings) With%s(value int) *ExecutorWithAnsibleConfigurationSettings {\n	e.configurationSettings[%s] = fmt.Sprint(value)\n	return e\n}\n", str, optionName, optionName)

		default:
			str = fmt.Sprintf("// With%s sets the value for the configuraion %s (%s)\n", optionName, config.env, config.description)
			str = fmt.Sprintf("%sfunc (e *ExecutorWithAnsibleConfigurationSettings) With%s(value string) *ExecutorWithAnsibleConfigurationSettings {\n	e.configurationSettings[%s] = value\n	return e\n}\n", str, optionName, optionName)
		}
	}
	return str
}

func generateConfigMethods(configs []*config) string {
	str := ""
	for _, config := range configs {
		str = fmt.Sprintf("%s%s\n", str, generateConfigMethod(config))
	}
	return str
}

func generateTest(config *config) string {
	str := ""

	if config.env != "" {

		optionName := strcase.ToCamel(config.env)

		switch config.vartype {
		case "boolean":
			str = fmt.Sprintf("// TestWith%s tests the method that sets %s to true\n", optionName, config.env)
			str = fmt.Sprintf("%sfunc TestWith%s(t *testing.T) {\nexec := NewExecutorWithAnsibleConfigurationSettings(nil).With%s()\nsetting := exec.configurationSettings[%s]\nexpected := \"true\"\nassert.Equal(t, setting, expected)\n}\n", str, optionName, optionName, optionName)

			str = fmt.Sprintf("%s\n// TestWithout%s tests the method that sets %s to false\n", str, optionName, config.env)
			str = fmt.Sprintf("%sfunc TestWithout%s(t *testing.T) {\nexec := NewExecutorWithAnsibleConfigurationSettings(nil).Without%s()\nsetting := exec.configurationSettings[%s]\nexpected := \"false\"\nassert.Equal(t, setting, expected)\n}\n", str, optionName, optionName, optionName)

		case "integer":
			str = fmt.Sprintf("// TestWith%s tests the method that sets the value for %s\n", optionName, config.env)
			str = fmt.Sprintf("%sfunc TestWith%s(t *testing.T) {\nvalue := 10\nexec := NewExecutorWithAnsibleConfigurationSettings(nil).With%s(value)\nsetting := exec.configurationSettings[%s]\nassert.Equal(t, setting, fmt.Sprint(value))\n}\n", str, optionName, optionName, optionName)

		default:
			str = fmt.Sprintf("// TestWith%s tests the method that sets the value for %s\n", optionName, config.env)
			str = fmt.Sprintf("%sfunc TestWith%s(t *testing.T) {\nvalue := \"testvalue\"\nexec := NewExecutorWithAnsibleConfigurationSettings(nil).With%s(value)\nsetting := exec.configurationSettings[%s]\nassert.Equal(t, setting, value)\n}\n", str, optionName, optionName, optionName)
		}
	}

	return str
}

func generateTests(configs []*config) string {
	str := ""
	for _, config := range configs {
		str = fmt.Sprintf("%s%s\n", str, generateTest(config))
	}
	return str
}

type generateFunc func([]*config) string

func main() {

	args := os.Args[1:]
	invalidOptionsMsgErr := fmt.Sprintf("Invalid option.\n\n%s <options>\n\n OPTIONS:\n - const: Generate constants\n - method: generated methods\n - test: Generate tests\n\n", path.Base(os.Args[0]))

	if len(args) != 1 {
		log.Fatalf(invalidOptionsMsgErr)
	}

	var f generateFunc

	// Request the HTML page.
	url := "https://docs.ansible.com/ansible/latest/reference_appendices/config.html"

	switch args[0] {
	case "const":
		f = generateConsts
	case "method":
		f = generateConfigMethods
	case "test":
		f = generateTests
	default:
		log.Fatalf(invalidOptionsMsgErr)
	}

	configs := LoadConfigs(url)
	fmt.Println(f(configs))

}
