package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	s:="我爱爬长城"//utf-8 三个字节 - 一个汉字
	for _,b:=range[]byte(s){
		fmt.Printf("%X ",b)
	}
	fmt.Println()
	//rune 2个字节 -一个汉字
	for i,ch:=range s{//ch is rune
		fmt.Printf("(%d %X %c) ",i,ch,ch)
	}
	//几个rune 数量
	fmt.Println(utf8.RuneCountInString(s))
	bytes:=[]byte(s)
	for len(bytes)>0{
		//字节  ，ch 解析出来第一和字符  从下一个位置继续解析
		ch,size:=utf8.DecodeRune(bytes)
		bytes=bytes[size:]
		fmt.Printf("%c ",ch)
	}
}
