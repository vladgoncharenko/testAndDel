package main

import "fmt"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

func main() {
	var r Role
	r = RoleAdmin
	fmt.Println(r)
}
