package main

import "fmt"

func main(){
	m:=map[string]string{
		"name":"ekko",
		"1name":"ekko",
		"1name2":"ekko",
		"1name3":"ekko",
		"1name4":"ekko",
	}
	fmt.Println(m)
	//遍历
	for k,v:=range m{
		fmt.Println(k,v)
	}
	//是否存在  零值
    if name,ok:=m["names"];ok{
    	fmt.Println(name)
	}else{
		fmt.Println("key does not exist")
	}
	//删除
	delete(m,"1name1")
    fmt.Println(m)
}
