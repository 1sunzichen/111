package main

import (
	"fmt"
)
//自定义错误
type MyError struct {

	Age string
}
func (e MyError) Error() string {
	 //return e.Age
	 return fmt.Sprintf("%s",e.Age)
}
func oops() error {
	return MyError{
		"the file system has gone away",
	}
}
func Example() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
	// Output: 1989-03-15 22:30:00 +0000 UTC: the file system has gone away
}
func main() {
	Example()
}
