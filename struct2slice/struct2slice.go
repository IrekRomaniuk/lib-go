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

func Struct2slice(x Manystrings) []string {

	v := reflect.ValueOf(x)
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).String()
	}
	return values
}

func (ms *Manystrings) Strings() []string {
	return []string{ms.string1, ms.string2, ms.string3}
}

func toStrings(x interface{}) []string {
	v := reflect.Indirect(reflect.ValueOf(x))
	var r []string
	for i := 0; i < v.NumField(); i++ {
		r = append(r, v.Field(i).String())
	}
	return r
}