package structenum

import (
	"testing"
	"fmt"
)

func TestStructenum(t *testing.T) {

//t.Fatalf
	fmt.Println("Vertex1 with anonymous field SubVertex")
	v1 := Vertex1{"A", "B",SubVertex{"C"}}
	fmt.Println(get_field1(&v1, "X"))
	fmt.Println(get_field1(&v1, "Y"))
	fmt.Println(get_field1(&v1, "Z"),v1.SubVertex.Z)

	fmt.Println("Vertex2 with named field SubVertex")
	v2 := Vertex2{"A", "B",SubVertex{"C"}}
	fmt.Println(get_field2(&v2, "X"))
	fmt.Println(get_field2(&v2, "Y"))
	//fmt.Println(get_field2(&v2, "Z"),v2.SubVertex.Z)
	fmt.Println(get_subvertex_field(&v2,"SubVertex","Z"),v2.SubVertex.Z)

}

