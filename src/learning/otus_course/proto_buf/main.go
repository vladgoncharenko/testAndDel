package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

//https://developers.google.com/protocol-buffers/docs/gotutorial

//go:generate ls -la
//go:generate protoc --go_out=. Person.proto

func main() {
	p := new(Person)
	p.Name = "User_name"
	p.Mobile = append(p.Mobile, "094-333-23-22")
	data, err := proto.Marshal(p)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
	fmt.Println(p.String())

	p1 := Person{}
	err = proto.Unmarshal(data, &p1)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", p1)
}
