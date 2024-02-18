package transformer

import (
	"fmt"
	"regexp"
	"time"
)

const (
	// PrefixTokenSeparator is and string printed between prefix and ansible output
	PrefixTokenSeparator = "\u2500\u2500"

	// DefaultLogFormatLayout is the default format when is used the LogFormat transformer
	DefaultLogFormatLayout = "2006-01-02 15:04:05"
)

// TransformerFunc is used to enrich or update messages before to be printed out
type TransformerFunc func(string) string

// Prepend is a transformer function that includes a prefix to a message expression
func Prepend(expression string) TransformerFunc {
	return func(message string) string {
		return fmt.Sprintf("%s %s", expression, message)
	}
}

// Append is a transformer function that includes a suffix to a message expression
func Append(expression string) TransformerFunc {
	return func(message string) string {
		return fmt.Sprintf("%s %s", message, expression)
	}
}

// LogFormat is a transformer function that includes a time reference at the beginning of a message expression
func LogFormat(layout string, f func(string) string) TransformerFunc {
	return func(message string) string {

		datetime := f(layout)
		return fmt.Sprintf("%s\t%s", datetime, message)
	}
}

// IgnoreMessage is a transformer function that returns a blank string when the message match to any skipping pattern
func IgnoreMessage(skipPatterns []string) TransformerFunc {
	return func(message string) string {
		for _, pattern := range skipPatterns {
			match, _ := regexp.MatchString(pattern, message)
			if match {
				return ""
			}
		}

		return message
	}
}

// Now returns a time value according to layout
func Now(layout string) string {
	return time.Now().Format(layout)
}
