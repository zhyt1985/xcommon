package xjson

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Unmarshal ...
func Unmarshal(bytes []byte, v interface{}) error {
	// var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	return nil
}

// Marshal ...
func Marshal(v interface{}) ([]byte, error) {
	// var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(&v)
}
