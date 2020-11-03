package util

import "encoding/json"

func PrettyJSON(i interface{}) string {
	if i == nil {
		return ""
	}
	b, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}
