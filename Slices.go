package EasyReflect

import "reflect"

func NewSliceLike(slice interface{}) (newSlice interface{}){
   newSlice=reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(slice).Elem()),0,1)
   return newSlice
}

func NewSliceFor(item interface{},len,cap int)(newSlice interface{}){
   newSlice=reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(item)),len,cap)
   return newSlice
}


