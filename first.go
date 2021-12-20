package main

import (
	"fmt"
	"reflect"
)

type A int32
type B int32

func main()  {
	a:=A(3)
	fmt.Println(reflect.TypeOf(a),a)
	fmt.Println(reflect.TypeOf(B(a)),B(a))

}