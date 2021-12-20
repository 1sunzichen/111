package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile(filename string){
	//可读可写 写入不了
	file,err:=os.OpenFile(filename,os.O_CREATE|os.O_EXCL,0616)
	if err!=nil{
		//panic(err)
		if pathErr,ok:=err.(*os.PathError);!ok{
			panic(pathErr)
		}else{
			fmt.Printf("%s,%s,%s\n",pathErr.Op,pathErr.Path,pathErr.Err)
		}
		return
	}
	defer file.Close()
	wr:=bufio.NewWriter(file)
	defer wr.Flush()
	for i := 0; i <101 ; i++ {
		fmt.Fprintln(wr,i)
	}
}
func main() {
	writeFile("hi.txt")
}
