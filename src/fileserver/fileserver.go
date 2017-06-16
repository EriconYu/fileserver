package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego/httplib"
)

func main() {
	port := ":8085"

	arg_num := len(os.Args)
	if arg_num > 1 {
		port = ":" + os.Args[1]
	}

	// 开启web服务
	fmt.Println("/******************局域网内网共享工具********************")
	fmt.Println("\t【文件服务器端口】：")
	fmt.Println("\t", port[1:])
	fmt.Println("\t【本机内网地址】：")
	getInternal()
	fmt.Println("\t【本机公网地址】：")

	//以下代码需要访问公网
	req := httplib.Get("http://getip.stackbang.com")
	externalip, err := req.String()
	if err != nil {
		fmt.Println("\t获取公网地址失败：")
		fmt.Println("\t", err)
	} else {
		strs := strings.Split(externalip, ",")
		fmt.Println("\t" + strs[2][14:])
	}

	fmt.Printf("\t在浏览器中访问http://对应网卡的ip%s即可\n", port)
	fmt.Println("********************************************************/")

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	server := http.Server{
		Addr:           port,            // 监听的地址和端口
		Handler:        nil,             // 所有请求需要调用的Handler（实际上这里说是ServeMux更确切）如果为空则设置为DefaultServeMux
		ReadTimeout:    0 * time.Second, // 读的最大Timeout时间
		WriteTimeout:   0 * time.Second, // 写的最大Timeout时间
		MaxHeaderBytes: 256,             // 请求头的最大长度
		TLSConfig:      nil,             // 配置TLS
	}

	errListen := server.ListenAndServe()

	if errListen != nil {
		fmt.Println("ListenAndServer: ", errListen)
	}
	fmt.Println("exit...")
}
