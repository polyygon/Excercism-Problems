package raindrops

import "strconv"

func Convert(input int) string {
	//create array of values to divide by 
	rainDrops := []string{"Pling", "Plang", "Plong", strconv.Itoa(input)}
	output :=""
	
	//check to see if divisible by 3
	if input%3==0 {
		output +=rainDrops[0]
	} 
	//check to see if divisible by 5 
	if input%5==0 {
		output +=rainDrops[1]
	} 
	//check to see if divisible by 7
	if input%7==0 {
		output +=rainDrops[2]
	} 
	//check to see if output is empty string 
	if output =="" { 
		output +=rainDrops[3]
	}
	//return output string 
	return output
}