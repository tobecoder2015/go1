package main
/**
 参考 https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/03.3.md
 */
import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"text/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:",r.Method)
	if r.Method=="GET"{
		t,_:=template.ParseFiles("page/upload.gtpl")
		crutime:=time.Now().Unix()
		h:=md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token:=fmt.Sprintf("%x",h.Sum(nil))
		log.Println(t.Execute(w,token))
	}else {
		r.ParseForm()  //解析参数，默认是不会解析的

		r.ParseMultipartForm(2)
		file, handler, err := r.FormFile("uploadfile")
		if err !=nil{
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprint(w,"%v",handler.Header)
		f,err:=os.OpenFile("./tmp/"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err!=nil{
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f,file)
	}
}


func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:",r.Method)
	if r.Method=="GET"{
		t,_:=template.ParseFiles("page/login.gtpl")
		crutime:=time.Now().Unix()
		h:=md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token:=fmt.Sprintf("%x",h.Sum(nil))
		log.Println(t.Execute(w,token))
	}else {
		r.ParseForm()  //解析参数，默认是不会解析的

		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])

		fmt.Fprintf(w, "Hello"+r.Form["username"][0]+" token"+string(r.Form.Get("token")))

	}
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login) //设置访问的路由
	http.HandleFunc("/upload", upload) //设置访问的路由

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}