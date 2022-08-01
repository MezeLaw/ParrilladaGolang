package utils

import (
	"Parrillada/models/comida"
	"reflect"
)

func RetornarNombreStructHijoComida(plato comida.IComida) string {

	childStruct := reflect.TypeOf(plato)

	return childStruct.String()
}
