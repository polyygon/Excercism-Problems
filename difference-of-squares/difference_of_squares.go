package diffsquares

func SquareOfSum(input int) int {
	//given a random integer input  
	//create slice 
	//fill slice with for loop up to input 
	//square each elemenet of slice 
	
	//create slice the size of input + 1 & initalize variables 
	s :=make([]int, input+1)
	sum := 0

	for i :=0; i <=input; i++ {
		s[i] = i
		sum = s[i] + sum
		
	}
	return sum*sum 

}

func SumOfSquares(input int) int {
	//create a slice of the size of input + 1 & initalize variables 
	s :=make([]int, input+1)
	square := 0

	for i :=0; i <=input; i++ {
		s[i] = i*i
		square = s[i] + square
		
	}
	return square 

}

func Difference(input int) int {
		return SquareOfSum(input) - SumOfSquares(input)
}