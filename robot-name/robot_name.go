package robotname

import (
	"fmt"
	"math/rand"
)

//two uppercase letters with 3 digits RX543
//create two constants for digit and code gen
const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberBytes = "0123456789"

//function for generating 3 random numbers
func Numbers() string {
	numbers := make([]byte, 3)
	for i := range numbers {
        numbers[i] = numberBytes[rand.Intn(len(numberBytes))]
    }
    return string(numbers)
}

//function for generating 2 random letters 
func Letters() string {
	letters := make([]byte, 2)
    for i := range letters {
        letters[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(letters)
}

//create struct for Robot name 
type Robot struct {
	name string 
}	

//create method to generate a bane 
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	r.name = fmt.Sprintf("%s%s", Letters(), Numbers())
	return r.name, nil
}

//create method to reset name to empty string
func (r *Robot) Reset() (string, error) {
	r.name = ""
	return r.name, nil 
}

