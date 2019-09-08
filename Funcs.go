package EasyReflect

import "reflect"

func Swapp(slice interface{},i,j int) (result interface{}){
	funk:=reflect.Swapper(slice)
	funk(i,j)
	return slice
}



/*func newFunc(inputs []interface{},outputs []interface{},isvariodic bool) (funk interface{} ){
	var in []reflect.Type
	var out []reflect.Type

	for _,v:=range inputs{
		in=append(in,reflect.TypeOf(v))

	}
	for _,v:=range outputs{
		out=append(out,reflect.TypeOf(v))
	}

	typ:=reflect.FuncOf(in,out,isvariodic)

	//...makefunc from typ!
}
*/