package main

import "fmt"

//%v ­ представление по­умолчанию для типа
//%#v ­ вывести как Go код (удобно для структур) %T ­ вывести тип переменной
//%% ­ вывести символ %
//Для целых:
//%b base 2
//%d base 10
//%o base 8
//%x base 16, with lower-case letters for a-f
//Для строк:
//%s the uninterpreted bytes of the string or slice
//%q a double-quoted string safely escaped with Go syntax
//%x base 16, lower-case, two characters per byte

func main() {
	arr := []int{1, 2, 3}
	mp := map[string]string{"qwe": "asd"}
	ch := make(chan int, 3)
	var st struct {
		s string
		i int
	}

	fmt.Printf("%p\n%v\n%#v\n\n", arr, arr, arr)
	fmt.Printf("%p\n%v\n%#v\n\n", mp, mp, mp)
	fmt.Printf("%p\n%v\n%#v\n\n", ch, ch, ch)
	fmt.Printf("%v\n%#v\n\n", st, st)
}
