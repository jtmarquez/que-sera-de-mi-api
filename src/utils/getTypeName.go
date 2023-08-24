package getTipeName

import "reflect"

func GetTypeName(v interface{}) string {
	return reflect.TypeOf(v).String()
}
