package struct2slice

import (
	"reflect"
)
const testVersion = 1
type Manystrings struct {
	string1 string
	string2 string
	string3 string
}

var Session2slice []string

func Struct2slice(somestruct Manystrings) []string {

	v := reflect.ValueOf(somestruct)
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).String()
	}
	return values
}