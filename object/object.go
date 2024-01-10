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
	INTEGER_OBJ = "INTEGER"
)

// Integer 每当在源代码中遇到整数字面值时，需要先转换为ast.IntegerLiteral。在对该节点求值时，再将其转换为Object.Integer
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
