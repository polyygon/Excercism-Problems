
package acronym

import (
	"strings"
    "unicode"
 )
// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	//quick function to remove any characters that are not letters or numbers
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != rune('\'')
	
	}
	//need to filter out ' character'
	newWord := strings.FieldsFunc(s, f)

	tla := ""
	for _, i := range newWord {
		tla += strings.ToUpper(string(i[0]))

	}
	
	return tla
}
