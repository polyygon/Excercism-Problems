package luhn

import (
	"strings"
	"unicode"
)

// Valid takes an input and checks if it is valid per the Luhn formula
// Returns true if valid, else false


func Valid(str string) bool {
	//remove spaces from input string 
	s :=strings.Join(strings.Fields(str), "")
	sum :=0 
	//check if zero or not a digit 
	if s =="0" || isDigit(s) == false {
		return false 
	}
	
	//reverse iterate through string & convert to number 
	for i := len(s)-1; i>=0; i-- {
		c := int(s[i] -'0')
		//filter by 2nd digit from right 
		if (len(s)-i)%2==0 {
			c=c*2
				if c >9 {
					c = c-9
				}
		}
		sum+= c
	}
	if sum%10 ==0 {
		return true 
	} else {
		return false 
	}
		
}

//check to see if input is number 
func isDigit(str string) bool {
	for _, c := range str {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}