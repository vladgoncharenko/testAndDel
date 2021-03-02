package main

import (
	"errors"
	"fmt"
	"reflect"
)

type SomeStr struct {
	Field string
	Val   int
}

func main() {
	var i = 42
	var s = struct {
		string
		int
	}{"hello", 42}
	iv := reflect.ValueOf(i)  // тип reflect.Value
	sv := reflect.ValueOf(&s) // тип reflect.Value

	fmt.Println(sv.Kind())
	fmt.Println(iv.Type())
	fmt.Println(sv.Interface())
	fmt.Println(sv.String())

	var ivv interface{} = "hello"
	ss, ok := assertString(ivv)
	fmt.Println(ss, ok)
	s2, ok := assertString(42)
	fmt.Println(s2, ok)

	someStr := SomeStr{
		Field: "Hello",
		Val:   110,
	}

	m, _ := structToMap(someStr)

	for k, v := range m {
		fmt.Println(k, "  ", v)
	}
}

func assertString(iv interface{}) (string, bool) {
	rv := reflect.ValueOf(iv)
	s := ""
	ok := false
	if rv.Kind() == reflect.String {
		s = rv.String()
		ok = true
	}
	return s, ok
}

//Преобразование структуры в map
func structToMap(iv interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(iv)
	if v.Kind() != reflect.Struct {
		return nil, errors.New("not a struct")
	}
	t := v.Type()
	mp := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // reflect.StructField
		fv := v.Field(i)    // reflect.Value
		mp[field.Name] = fv.Interface()
	}

	return mp, nil
}

func mapToStruct(mp map[string]interface{}, iv interface{}) error {
	v := reflect.ValueOf(iv)
	if v.Kind() != reflect.Ptr {
		return errors.New("not a pointer to struct")
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return errors.New("not a pointer to struct")
	}
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // reflect.StructField
		fv := v.Field(i)    // reflect.Value
		if val, ok := mp[field.Name]; ok {
			mfv := reflect.ValueOf(val)
			if mfv.Kind() != fv.Kind() {
				return errors.New("incomatible type for " + field.Name)
			}
			fv.Set(mfv)
		}
	}
	return nil
}
