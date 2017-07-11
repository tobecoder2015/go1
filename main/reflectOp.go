package main

import (
	f "fmt"
	"reflect"
)


func main()  {
	var x float64  =3.4
	v:=reflect.ValueOf(x)
	f.Println("type:",v.Type())
	f.Println("kind is float64:",v.Kind()==reflect.Float64)
	f.Println("value:",v.Float())

	v2:=reflect.ValueOf(&x)
	v3:=v2.Elem()
	v3.SetFloat(5.7)
	f.Println("value:",v3.Float())
	f.Println("x value:",x)

}
