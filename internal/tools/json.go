package tools

import "encoding/json"

func ParseJson[T interface{}](input string) (T, error) {
	var output T
	err := json.Unmarshal([]byte(input), &output)
	return output, err
}
