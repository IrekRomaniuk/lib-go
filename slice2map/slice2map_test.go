package slice2map

import (
	"testing"
	_ "github.com/stretchr/testify/assert"
	"reflect"
)

const targetTestVersion = 1

func TestSlice2Map(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}

	input := []string{"ae1:1", "ae2:0", "ae3:0"}
	want := map[string]int{"ae1":1,"ae2":0,"ae3":0}

	if output := Slice2Map(input); reflect.DeepEqual(input, want) {
		t.Fatalf("Slice to Map: got %v, want %v",output, want)
	}
}


