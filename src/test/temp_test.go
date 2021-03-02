package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestTemp(t *testing.T) {
	var arr = [10]uint{1, 4, 5, 7, 1, 2, 5, 6, 3, 1}
	//
	str := strings.ReplaceAll(fmt.Sprint(arr), " ", "")
	//
	fmt.Println(strings.Trim(str, "[1453]"))

	//fmt.Println(codewars.CreatePhoneNumber(arr))
}

func sliceToStr(numbers [10]uint) string {
	var str string
	for _, v := range numbers {
		str += strconv.Itoa(int(v))
	}
	return str
}
