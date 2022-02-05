package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "user",
		Value:    "admin",
		HttpOnly: true,
		MaxAge:   60,
	}
	cookie2 := http.Cookie{
		Name:     "user2",
		Value:    "admin2",
		HttpOnly: true,
	}
	// 将Cookie发送给浏览器
	// w.Header().Set("Set-Cookie", cookie.String())
	// 添加第二个Cookie
	// w.Header().Add("Set-Cookie", cookie2.String())

	// 直接调用http的SetCookie函数设置Cookie
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}

// getCookies 获取Cookie
func getCookies(w http.ResponseWriter, r *http.Request) {
	// 获取请求头中所有的Cookie
	cookies := r.Header["Cookie"]
	// 获取具体的cookie
	cookie, _ := r.Cookie("user")
	fmt.Println("Cookie：", cookies, cookie)
}

func main() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookies", getCookies)
	http.ListenAndServe(":8080", nil)
}
