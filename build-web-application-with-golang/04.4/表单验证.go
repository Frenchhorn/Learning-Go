package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.html")
		t.Execute(w, token)
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// 验证token的合法性
			fmt.Println("Token存在")
		} else {
			// 不存在token报错
			fmt.Println("Token不存在")
			template.HTMLEscape(w, []byte("ERROR"))
		}
		fmt.Println("username length:", len(r.Form["username"][0])) //输出到服务器端
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
