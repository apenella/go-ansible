package helper

import "encoding/json"

// MarshalToString JSON编码为字符串
func MarshalToString(v any) string {
	bs, err := json.Marshal(v)

	s := string(bs)
	if err != nil {
		return ""
	}
	return s
}

// MarshalIndentToString JSON编码为字符串
func MarshalIndentToString(v any, prefix, indent string) string {
	s, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return ""
	}
	return string(s)
}
