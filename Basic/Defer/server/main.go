package main

import (
	"goBasic1231/Basic/Defer/server/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,res *http.Request) error



func errHandler(handler appHandler) func(http.ResponseWriter,*http.Request){
	return func(writer http.ResponseWriter,request *http.Request){
		err:=handler(writer,request)
		defer func(){
			r:=recover()
			if r!=nil{

			log.Printf("Panic:%v",r)
			http.Error(writer,http.StatusText(http.StatusNotFound),http.StatusInternalServerError)
			}
		}()
		if err!=nil{
			code:=http.StatusOK
			switch{
			case os.IsNotExist(err):
				http.Error(writer,http.StatusText(http.StatusNotFound),http.StatusNotFound)
			case os.IsPermission(err):
				code=http.StatusForbidden
			default:
				code=http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code),code)
		}
	}
}
func main() {
	http.HandleFunc("/list/",errHandler(filelisting.HandlerFileList))
	http.ListenAndServe(":8888",nil)
}
