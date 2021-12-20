package main

import (
	"fmt"
	"math"
)

func triangle(){
	var a,b int=3,4
	fmt.Println(calc(a,b))
}
func calc(a,b int)int {
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	return c
}
func main() {
	
}
