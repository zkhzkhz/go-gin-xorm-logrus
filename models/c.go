package models

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}
