package funcs

import (
	"reflect"
	"sportsman/models"
)

type Structures interface {
	ValueOf() reflect.Value
	StructFields() reflect.Type
	NumOfFields() int
}

var _ Structures = &models.Computers{}
var _ Structures = &models.Person{}
var _ Structures = &models.Phones{}
