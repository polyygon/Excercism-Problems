package scrabble

import "strings"

//uses loops, maps and strings



func Score(input string) int {
//create map to hold key value pairs 
	m := map[rune]int{
		'A': 1, 
		'E': 1, 
		'I': 1, 
		'O': 1, 
		'U': 1, 
		'L': 1, 
		'N': 1, 
		'R': 1, 
		'S': 1, 
		'T': 1, 
		'D': 2, 
		'G': 2, 
		'B': 3, 
		'C': 3, 
		'M': 3, 
		'P': 3, 
		'F': 4, 
		'H': 4, 
		'V': 4, 
		'W': 4, 
		'Y': 4, 
		'K': 5, 
		'J': 8, 
		'X': 8, 
		'Q': 10,
		'Z': 10,
	} 

	//setup var to carry score 
	totalScore := 0
	//convert input to to all uppercase string
	upperCaseWord := strings.ToUpper(input) 

	//iterate over input word for each character 
	for _, v := range upperCaseWord {
		//add value of each letter to totalScore var
		totalScore = totalScore + m[v] 

	}

	return(totalScore)
}



//take input -> map to key value -> add 