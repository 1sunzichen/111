package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen() chan string{
	c:=make(chan string)
	//go routine
	go func(){
        i:=0
        for{//循环给数据
        	time.Sleep(time.Duration(rand.Intn(2000))*time.Millisecond)
        	c <-fmt.Sprintf("msg%d",i)
        	i++
		}
	}()
	return c
}

func main() {
	m:=msgGen()
	for{
		fmt.Println(<-m)
	}
}
