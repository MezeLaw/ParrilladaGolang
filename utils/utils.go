package utils

import (
	"Parrillada/models/comida"
	"reflect"
)

func RetornarStructHijoNombre(plato comida.IComida) string {

	childStruct := reflect.TypeOf(plato)

	return childStruct.String()
}
