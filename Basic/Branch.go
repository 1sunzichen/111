package main
import (
	"fmt"
	//"io/ioutil"
)
func grade(score int) string{
	switch{
	case score<60:
		return "差"
	default:
		return "优"
	}
}
func main(){
	var a=grade(28)
    fmt.Print(a)
}