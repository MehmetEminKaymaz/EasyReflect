package EasyReflect

import "reflect"

func newStructLike(Astruct interface{}) (newStruct interface{}){ //this method return settable and adressable struct
	ptr:=reflect.New(reflect.TypeOf(Astruct))
	newStruct=ptr.Elem()
	return newStruct
}

func FieldNames(Astruct interface{}) (fieldNames []string){
	value:=reflect.ValueOf(Astruct)
	for i:=0;i<value.NumField();i++{
		fieldNames=append(fieldNames,value.Type().Field(i).Name)
	}

	return
}

func EmbeddedFields(Astruct interface{}) (embedded []interface{}){//return all embedded fields in a struct if exist
	value:=reflect.ValueOf(Astruct)

	for i:=0;i<value.NumField();i++{
		if value.Type().Field(i).Anonymous==true{
			embedded=append(embedded,value.Type().Field(i))
		}
	}

	return
}

func NonEmbeddedFields(Astruct interface{}) (Nonembedded []interface{}){
	value :=reflect.ValueOf(Astruct)

	for i:=0;i<value.NumField();i++{
		if value.Type().Field(i).Anonymous==false{
			Nonembedded=append(Nonembedded,value.Type().Field(i))
		}
	}

	return
}

func FieldTypes(Astruct interface{}) (types []interface{}){
	value:=reflect.ValueOf(Astruct)

	for i:=0;i<value.NumField();i++{
		types=append(types,value.Type().Field(i).Type)
	}

	return types
}


