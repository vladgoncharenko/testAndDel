package main

import (
	"errors"
	"fmt"
)

var myError = errors.New("my error")
var errPermission = errors.New("permissions")

func main() {
	newErr := errors.New("my error")

	fmt.Println(newErr)

	fmt.Println(errors.Is(myError, myError))

	//Упаковка ошибки с помощью %w делает её доступной для errors.Is и errors.As:
	err := fmt.Errorf("access denied: %w", errPermission)
	fmt.Println(err)
	fmt.Println(errors.Is(err, errPermission))

	err = fmt.Errorf("add some info %v", err)

	fmt.Println(err)

}
