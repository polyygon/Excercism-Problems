package clock

import "fmt"

//create type for clock
type Clock struct {
	m int 
}

const minutesPerDay = 1440 

//generate a new clock given two numbers 
func New(h int, m int) Clock {
	c :=(h*60 + m) % minutesPerDay

	if c < 0 {
		c += minutesPerDay
	}
	return Clock{c}
}

//add minutes
func (c Clock) Add(m int) Clock {
	return New(0, c.m+m)
}

//subtract minutes 
func (c Clock) Subtract(m int) Clock {
	return New(0, c.m-m)
}
//print in a string 
func (c Clock) String() string {
	h :=c.m/60
	m :=c.m%60
	return fmt.Sprintf("%02d:%02d", h, m)
}