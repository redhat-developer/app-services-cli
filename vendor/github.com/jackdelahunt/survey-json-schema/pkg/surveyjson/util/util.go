package util

import (
	"encoding/json"
)

// ToStructFromMapStringInterface marshals a generic map[string]interface{} to a struct by marshalling to json and back
// Use JSON for the marshalling instead of YAML because sub-structs will get marshalled into map[interface{}]interface{}
// when using YAML, but map[string]interface{} when using JSON and vault libraries can't handle map[interface{}]interface{}
func ToStructFromMapStringInterface(m map[string]interface{}, str interface{}) error {
	j, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(j, str)
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
