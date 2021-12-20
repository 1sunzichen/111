package main

import "fmt"

func main(){
	//
	arr3:=[...]int{2,3,4,5,6}
	//fmt.Println(arr1,arr2,arr3)
	//for i:=0;i<len(arr3);i++{
	//	fmt.Println(i,arr3[i])
	//}
	for i,v:=range arr3{
		fmt.Println(i,v)
	}
	arr:=[...]int{0,1,2,3,4,5}
	s1:=append(arr[0:],12)
	fmt.Println(s1,arr[0:],arr)
}