package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)
type testErrors string
func (e testErrors)Error()string{
	return e.Message()
}
func (e testErrors)Message()string  {
	return string(e)
}
func errPanic(wr http.ResponseWriter,req *http.Request)error{
	panic(123)

}
func errErroring(wr http.ResponseWriter,req *http.Request)error{
	//panic(123)
	return testErrors("usererror")
}
//httptest.NewRecorder()
var testss=[]struct{
h appHandler
code int
message string
}{
{errPanic,501,"Inter server error1"},
//{errErroring,500,"usererror"},
}
func TestErrWrapper(t *testing.T){

	//httptest.NewRecorder()
  tests:=[]struct{
  	h appHandler
  	code int
  	message string
  }{
	{errPanic,501,"Inter server error1"},
	//{errErroring,500,"usererror"},
  }
  for _,tt:=range tests{
  	f:=errHandler(tt.h)
  	res:=httptest.NewRecorder()
  	req:=httptest.NewRequest(
  		http.MethodGet,"https://www.baidu.com",nil)
  	f(res,req)
  	b,_:=ioutil.ReadAll(res.Body)
  	body :=strings.Trim(string(b),"\n")
    fmt.Println(body,tt.message)
	if res.Code!=tt.code||
  		body!=tt.message{
		t.Errorf("期望 %d %s 实际 %d %s",tt.code,tt.message,res.Code,body)
	}
  }
}

func verify(res *http.Response,code int,message string,t *testing.T){
	b,_:=ioutil.ReadAll(res.Body)
	body :=strings.Trim(string(b),"\n")
	fmt.Println(body,message)
	if res.StatusCode!=code||
		body!=message{
		t.Errorf("期望 %d %s 实际 %d %s",code,message,res.StatusCode,body)
	}
}
//测试整个服务器
func TestErrWrapperInServer(t *testing.T){
		for _,tt:=range testss{
			f:=errHandler(tt.h)
			server:=httptest.NewServer(http.HandlerFunc(f))
			res,_:=http.Get(server.URL)
			verify(res,tt.code,tt.message,t)
		}
}