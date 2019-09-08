package EasyReflect

import "reflect"

func newMap(key interface{},value interface{}) (newmp interface{}){

	newmp=reflect.MakeMap(reflect.MapOf(reflect.TypeOf(key),reflect.TypeOf(value)))
	return
}

