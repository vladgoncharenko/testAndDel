package main

import (
	"errors"
	"fmt"
)

func main() {

	err := makeError()
	if err != nil {
		fmt.Println(err)
	}
}

func makeError() error {
	err := errors.New("some error")
	return fmt.Errorf("make error %w", err)
}
