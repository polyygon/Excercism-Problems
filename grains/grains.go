package grains

import "errors"

/*
Write code that shows:
- how many grains were on a given square, and
- the total number of grains on the chessboard
*/

func Square(x int) (uint64, error) {
	//check to see if input in range 
	if x < 1 ||  x > 64 {
		return 0, errors.New("input out of bounds")
	}
	//bitshift by input amount 2*exp(x-1)
	return 1 << (x -1), nil 

	 
}

func Total() uint64 {
	return 1 <<64 -1 
}