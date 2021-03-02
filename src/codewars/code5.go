package codewars

import "strings"

//Return the number (count) of vowels in the given string.
//
//We will consider a, e, i, o, u as vowels for this Kata (but not y).
//
//The input string will only consist of lower case letters and/or spaces.

func GetCount(str string) (count int) {
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, v := range vowels {
		count += strings.Count(strings.ToLower(str), v)
	}
	return count
}
