package main

import (
	"fmt"
	"time"
)

func worker(id int,c chan int){
		for n:=range c{
			fmt.Printf("Worker %d received %c\n",id,n)
		}
}
func bufferedChannel(){
	//
	c:=make(chan int,3)
	go worker(0,c)
	c<-'a'
	c<-'v'
	c<-'c'
	time.Sleep(time.Millisecond)
}

func channelClose(){
	c:=make(chan int)
	go worker(0,c)
	c<-'a'
	c<-'v'
	c<-'c'
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	//bufferedChannel()
	channelClose()
}
