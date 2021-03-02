package codewars

import (
	"fmt"
	"strings"
)

//Count the number of Duplicates
//Write a function that will return the count of distinct
//case-insensitive alphabetic characters and numeric digits
//that occur more than once in the input string.
//The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.
//
//Example
//"abcde" -> 0 # no characters repeats more than once
//"aabbcde" -> 2 # 'a' and 'b'
//"aabBcde" -> 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
//"indivisibility" -> 1 # 'i' occurs six times
//"Indivisibilities" -> 2 # 'i' occurs seven times and 's' occurs twice
//"aA11" -> 2 # 'a' and '1'
//"ABBA" -> 2 # 'A' and 'B' each occur twice

func DuplicateCount(s1 string) int {
	str := strings.ToLower(strings.TrimSpace(s1))
	fmt.Println(s1)
	m := make(map[uint8]bool)
	fmt.Println(str)
	// a bcdea
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] && !m[str[i]] {
				m[str[i]] = true
			}
		}
	}

	return len(m)
	//return count //Instead of returning '1', you have to return integer 'i' as answer of solution.
}
