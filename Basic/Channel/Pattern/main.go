package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string) chan string {
	c := make(chan string)
	//go routine
	go func() {
		i := 0
		for { //循环给数据
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("msg%s %d", name, i)
			i++
		}
	}()
	return c
}
func msgGen2(name string, done chan struct{}) chan string {
	c := make(chan string)
	//go routine
	go func() {
		i := 0
		for { //循环给数据
			select {
			case <-time.After(3 * time.Second):
				c <- fmt.Sprintf("msg%s %d", name, i)
			case <-done:
				fmt.Println("clean up")
				time.Sleep(2 * time.Second)
				fmt.Println("clean updone")
				done <- struct{}{}
				return
			}
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("msg%s %d", name, i)
			i++
		}
	}()
	return c
}
func fanIn(c1s ...chan string) chan string {
	c := make(chan string)

	for _, a := range c1s {
		//a1:=a
		go func(a1 chan string) {
			for {
				c <- <-a1
			}
		}(a)
	}

	return c
}
func fanInSelect(c1, c2 chan string) chan string {
	c := make(chan string)
	// 开一个协程
	go func() {
		for {

			select {
			case n := <-c1:
				c <- n
			case n := <-c2:
				c <- n
			}
		}
	}()

	return c
}

//非阻塞等待
func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}

}

//超时
func timeOutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}
}

func main() {
	done := make(chan struct{})
	//m1 := msgGen("小郑")
	m2 := msgGen2("小孙", done)
	//m := fanIn(m1, m2)
	//for {
	//	fmt.Println(<-m)
	//}
	for i := 0; i < 6; i++ {
		if a, ok := timeOutWait(m2, time.Second); ok {
			fmt.Println(a)
		} else {
			fmt.Println("超时")
		}
	}
	done <- struct{}{}
	<-done
}
