package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)
const prefix="/list/"
type testError string
func (e testError)Error()string{
	return e.Message()
}
func (e testError)Message()string  {
	return string(e)
}
func HandlerFileList(write http.ResponseWriter,r *http.Request)error{
	if strings.Index(r.URL.Path,prefix)!=0{
		return testError("路径必须带"+prefix)
	}
	path:=r.URL.Path[len("/list/"):]
	file,err:=os.Open(path)
	if err!=nil{
		//panic(err)
		return err
	}
	defer file.Close()
	all,err:=ioutil.ReadAll(file)
	if err!=nil{
		//panic(err)
		return err
	}
	write.Write(all)
	return nil
}
