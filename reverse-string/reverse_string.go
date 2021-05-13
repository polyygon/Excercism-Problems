package reverse


func Reverse(s string) string {
	//convert to a rune to get code points 
	
	//convert input string into rune 
	cp := []rune(s)
	/*	i = index counting up from beginging of rune 
	 	j = index counting down from end of rune
	  	example: input "party on wayne", 14 characters
	 	cp[0] <-> cp[13] = "earty on waynp"
		cp[1] <-> cp[12] = "enrty on wayap"
		cp[2] <-> cp[11] = "enyty on warap"
	*/ 

	for i, j := 0, len(cp)-1; i < j; i, j = i+1, j-1 {
        	//swap rune values via i, j 
        	cp[i], cp[j] = cp[j], cp[i]

        }
	return string(cp)
} 