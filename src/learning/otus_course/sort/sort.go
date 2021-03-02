package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

func main() {
	var sl = []int{4, 5, 1, 5}
	sort.Ints(sl)
	fmt.Println(sl)

	// а что если нужно сортировать свои типы ?
	s := []User{{"vasya", 19},
		{"petya", 18}}

	fmt.Println(s)

	sort.Slice(s, func(i, j int) bool {
		return s[i].Age < s[j].Age
	})

	fmt.Println(s)
}
