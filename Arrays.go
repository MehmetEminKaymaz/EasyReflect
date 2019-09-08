package EasyReflect

import "reflect"

func NewArrayLike(arry interface{}) (newarr interface{}){

	ptr:=reflect.ArrayOf(reflect.ValueOf(arry).Len(),reflect.TypeOf(arry).Elem())
	newarr=reflect.New(ptr).Elem()

	return
}

func GetItemType(arry interface{}) (typ interface{}){
	return reflect.TypeOf(arry).Elem()
}
func GetArrayType(arry interface{}) (typ interface{}){
	return  reflect.TypeOf(arry)
}

