package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int{
	out:=make(chan int)
	go func(){
		i := 0
		for{
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int,c chan int){
	for n:=range c {
		fmt.Printf("worker id %d channel %d\n",id,n)
	}

}
//<-chan int  像这样的只能接收值
//
//chan<- int  像这样的只能发送值
func createWorker(id int) chan<- int{
	c:=make(chan int)
	go worker(id,c)
	return c
}
func main() {
	var c1,c2=generator(),generator()
	var w=createWorker(0)
	n:=0
	 	hasvalue:=false
	for{
		var ac chan<- int
	 	if hasvalue{
			fmt.Println("hasvalue ",w)
	 		ac=w
		}
		select {
		    //n收数据
			case n= <-c1:
				hasvalue=true
				//n收数据
			case n= <-c2:
				hasvalue=true
				//ac收数据
			case ac <-n:
				hasvalue=false

		}
	}
}
