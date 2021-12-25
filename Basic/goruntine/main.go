package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [11]int
	for i:=0;i<10;i++{
		go func(){
			for{
				a[i]++
			 fmt.Printf("%d\n",i)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
