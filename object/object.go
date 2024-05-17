package object

import (
	"fmt"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	NULL        = "Null"
	INTEGER_OBJ = "Integer"
	BOOLEAN_OBJ = "Boolean"
)

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

type Boolean struct {
	Value bool
}

func (i *Boolean) Inspect() string {
	return fmt.Sprintf("%t", i.Value)
}
func (i *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

type Null struct{}

func (i *Null) Inspect() string {
	return fmt.Sprintf("%s", "null")
}
func (i *Null) Type() ObjectType {
	return NULL
}