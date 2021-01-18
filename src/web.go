package main
// 测试网络编程
import (
	"fmt"
	"net/http"
)

// 提供文件访问服务的HTTP处理器
func testFileServer()  {
	// 如果该路径中有index.html文件，会优先显示html文件，否则会看到文件目录
	http.ListenAndServe(":2003", http.FileServer(http.Dir("/Users/smzdm/smzdm_project/ods/")))
}

// 多路由服务器
func testMuxServe()  {
	// 绑定路径，去触发方法
	http.HandleFunc("/index", indexHandler)
	// 绑定端口，第一个参数为监听地址，第二个参数表示服务端处理程序，通常为nil，
	// 这意味着服务端调用http.DefaultServeMux处理
	err := http.ListenAndServe("localhost:2003", nil)
	fmt.Println(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("/index========")
	w.Write([]byte("这是默认首页"))
}

func main2()  {
	//testFileServer()
	testMuxServe()
}
