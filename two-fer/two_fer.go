// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// twofer takes a string as input and prints a statement. If no name is entered a default value is printed. 
package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	s :=""
	if name == "" {
		s = fmt.Sprintf("One for you, one for me.") 
	} else {
		s = fmt.Sprintf("One for %s, one for me.", name) 
	}
	
	return s
}
