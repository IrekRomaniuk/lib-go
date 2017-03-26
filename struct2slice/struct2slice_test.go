package struct2slice

import (
	"testing"
	_ "github.com/stretchr/testify/assert"
	 "math/rand"
	"reflect"
)

const targetTestVersion = 1

func TestStruct2slice(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}

	input := Manystrings{"1","2","3"}
	want := []string{"1","2","3"}
	r := rand.Intn(len(want))

	if output := Struct2slice(input);  output[r] == want[r] {
		t.Fatalf("Struct to Slice: got %v, want %v",output[r], want[r])
	}
}

//go test -run TestStrings
func TestStrings(t *testing.T) {
	egs := []interface{Strings()[]string}{
		&Manystrings{"one", "two", "three"},
		&Manystrings{},
		&Manystrings{"a", "", ""},
		&Manystrings{"", "b", ""},
		&Manystrings{"", "", "c"},
		//... any other test cases
	}
	for _, ex := range egs {
		got, want := ex.Strings(), toStrings(ex)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%v.Strings() = %v, want %v", ex, got, want)
		}
	}
}