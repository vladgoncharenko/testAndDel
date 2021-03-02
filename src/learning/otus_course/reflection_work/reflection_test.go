package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Str struct {
	Val string
}

func (s Str) String() string {
	return s.Val
}

func TestReflection(t *testing.T) {
	str := Str{Val: "Hello"}

	fmt.Println(isStringer(33))
	fmt.Println(isStringer("33"))
	fmt.Println(isStringer(str))

	fmt.Println("________________________")
	var r fmt.Stringer
	fmt.Println(isStringerReflection(str, r))
}

func isStringer(data interface{}) bool {
	_, ok := data.(fmt.Stringer)
	return ok
}

func isStringerReflection(data interface{}, inter interface{}) bool {
	ddata := reflect.TypeOf(data).Elem()

	fmt.Println(ddata)

	//dInter := reflect.TypeOf(inter)
	//fmt.Println(dInter.Implements(ddata))
	return false
}
