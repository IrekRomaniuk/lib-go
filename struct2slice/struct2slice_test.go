package struct2slice

import (
	"testing"
	_ "github.com/stretchr/testify/assert"
	 "math/rand"
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
