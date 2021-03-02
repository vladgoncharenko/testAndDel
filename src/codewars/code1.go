package codewars

import (
	"strconv"
)

//Write a function that accepts an array
//of 10 integers (between 0 and 9), that
//returns a string of those numbers in the form of a phone number.

//CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})
//returns "(123) 456-7890"

func CreatePhoneNumber(numbers [10]uint) string {
	str := sliceToStr(numbers)
	return "(" + str[:3] + ")" + " " + str[3:6] + "-" + str[6:]

}

func sliceToStr(numbers [10]uint) string {
	var str string
	for _, v := range numbers {
		str += strconv.Itoa(int(v))
	}
	return str
}
