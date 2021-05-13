package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(input string) bool {
	//take input string 
	//convert to all lowercase 
	str := strings.ToLower(input)
	//create map: type key: rune, value: bool 
	m := make(map[rune]bool)
	//iterate over input string 
	for _, v := range str {
		//check to see if value of input string is a letter
		
		if m[v] {
			return false
		}
		//check to see if code point is letter
		if unicode.IsLetter(v) {
			m[v] = true
		}
	}
	return true //is an isogram
}
