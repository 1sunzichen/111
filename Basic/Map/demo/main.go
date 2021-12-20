package main

import "fmt"
//不重复的最长长度的字符串
//abcaaaa  abc
//aaaa    a21
func main(){
	var a="一二三三二一四000"

	var l=make(map[rune]int)
	start:=0
	maxL:=0
	for k,v:=range []rune(a){
		ll,ok:=l[v]
		//找到重复的数 记录位置右移一位
		//fmt.Println(l[v],k,v,start,maxL)
		if ok&&ll>=start{
			start=l[v]+1
		}
		//长度 与 最大长度判断  大于 交换树值
		if k-start+1>maxL{
			maxL=k-start+1
			fmt.Println(start)
		}
		//放到map   值:索引
		l[v]=k
		//[97] 0 [98] 1 [99] 2 [97] 3
		fmt.Println(l[v],k,v,start,git,l)
	}

}

