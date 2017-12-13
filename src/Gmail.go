package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Sendmail(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	// 解析url 传递的参数,对于POST 则解析响应包的主体(req//注意 :如果没有调用ParseForm 方法,下面无法获取表单的数据
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//fmt.Fprintf(w, "Hello astaxie!") // 这个写入到w 的是输出到客户端的
	var user string = r.Form.Get("user")
	var password string = r.Form.Get("password")

	var host string = r.Form.Get("host")

	var to string = r.Form.Get("to")

	var subject string = r.Form.Get("subject")
	var body string = r.Form.Get("body")
	//解析表格
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Mainweb.html")
		t.Execute(w, nil)
	} else if r.Method == "POST" {

		err := SendToMail(user, password, host, to, subject, body, "string")
		fmt.Fprintln(w, "send mail ....")
		if err != nil {
			fmt.Fprintln(w, "send error")
			//fmt.Fprintln(w, &host)
			//fmt.rintln(w, &user)
		} else {
			fmt.Fprintln(w, "send sussessful")
		}
	}
}

func main() {
	http.HandleFunc("/", Sendmail)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListernAndServe:", err)
	}
}
