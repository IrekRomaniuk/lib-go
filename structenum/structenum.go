package structenum

import "reflect"

type Vertex1 struct {
	X string
	Y string
	SubVertex
}

type Vertex2 struct {
	X string
	Y string
	SubVertex SubVertex
}

type SubVertex struct {
	Z string
}
//
func get_field1(v *Vertex1, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
func get_field2(v *Vertex2, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func get_subvertex_field(v *Vertex2, field1 string, field2 string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field1).FieldByName(field2)
	return f.String()
}

