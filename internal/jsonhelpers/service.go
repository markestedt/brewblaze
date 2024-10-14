package jsonhelpers

import "encoding/json"

func Parse[T interface{}](input string) (T, error) {
	var output T
	err := json.Unmarshal([]byte(input), &output)
	return output, err
}
