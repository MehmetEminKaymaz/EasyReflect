# EasyReflect
Usable and Simple Package to Reflection Go Programming Language

```GO
 Test1:=[]int{1,2,3,4,5}

    EasyReflect.NewSliceLike(Test1) //create a new slice from Test1
    //so this function return an integer slice
    EasyReflect.NewSliceFor((int)(0),10,20)//create a slice that can store the first parameter (int)

    EasyReflect.EmbeddedFields(Human{})//it return embedded fields if exist
    
    EasyReflect.NonEmbeddedFields(Human{})//it return non embedded fields if exist
    
    EasyReflect.FieldTypes(Human{})//it return fieldstype of struct
    
    EasyReflect.FieldNames(Human{}) //it return fieldnames (Name,Age etc...)
    
    EasyReflect.NewArrayLike([5]int{}) //same usage like newsliceLike
    
    EasyReflect.GetArrayType([5]int{}) //return type of an Array ([5]int)
    
    EasyReflect.GetItemType([5]int{}) //return type of an array item (int)
    
    EasyReflect.NewStructLike(Human{}) //return a new struct same as passed parameter as struct (creates a new human)
    
    EasyReflect.Swapp(Test1,0,2)//return slice with swap i'th and j'th element
    
    EasyReflect.NewMap("key",2)//it return a map => map[string]int
```
