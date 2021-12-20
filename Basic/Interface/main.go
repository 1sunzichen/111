package main

import (
	"fmt"
	"goBasic1231/Basic/Interface/infra"
	"goBasic1231/Basic/Interface/mock"
	"goBasic1231/Basic/Interface/testing"
)
//func retrieve(url string)string{
//	resp,err:=http.Get(url)
//	if err!=nil{
//		panic(err)
//	}
//	defer resp.Body.Close()
//	bytes,_:=ioutil.ReadAll(resp.Body)
//	return string(bytes)
//}


//func getRetrever() retriever{
//	//return testing.Retriever{}
//	return infra.Retriever{}
//}
// something that can "GET"   接口 1
type  retriever interface{
	Get(string) string
}

type Poster interface{
	Post(url string,form map[string]string)string
}
//接口
type Re interface {
	retriever
	Poster
}
const url="https://www.baidu.com"
func download(r retriever) string{
	return r.Get(url)
}
//方法
func session (s Re) string{

	s.Post(url,
		map[string]string{
		"contents":"ccmouse",

		})
	return s.Get(url)
}
func post(poster Poster){
	poster.Post(url,
		map[string]string{
		"name":"ccmouse",
		"course":"golang",
		})
}
func main(){
    var r retriever
    r=infra.Retriever{}
    fmt.Printf("%T %v\n",r,r)
    r=testing.Retriever{}

    fmt.Printf("%T %v\n",r,r)
    r1:=mock.Retriever{Contents:"com" }
   fmt.Println(session(&r1))
	fmt.Printf("%s\n",r.Get("https://www.baidu.com"))
}
