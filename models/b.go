package models

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
	Token        string `json:"token" form:"token"`
}
