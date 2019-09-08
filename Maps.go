package EasyReflect

import "reflect"

func NewMap(key interface{},value interface{}) (newmp interface{}){

	newmp=reflect.MakeMap(reflect.MapOf(reflect.TypeOf(key),reflect.TypeOf(value)))
	return
}

