package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	var ip string
	var port1, port2 int

	flag.StringVar(&ip, "ip", "10.251.26.114", "IP address to scan")
	flag.IntVar(&port1, "port1", 21, "the port to start scan")
	flag.IntVar(&port2, "port2", 1024, "the port to end scan")
	flag.Parse()
	fmt.Printf("NFlag: %v\n", flag.NFlag()) // 返回已设置的命令行标志个数
	fmt.Printf("NArg: %v\n", flag.NArg())   // 返回处理完标志后剩余的参数个数
	fmt.Printf("Args: %v\n", flag.Args())   // 返回处理完标志后剩余的参数列表
	fmt.Printf("Arg(1): %v\n", flag.Arg(1)) // 返回处理完标志后剩余的参数列表中第 i 项

	timeout := 1 * time.Second // 超时时间
	for i := port1; i <= port2; i++ {
		fmt.Printf("hello, world")
		var address = fmt.Sprintf("%s:%d", ip, i)
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			fmt.Println(address, "是关闭的")
			continue
		}
		conn.Close()
		fmt.Println(address, "打开") //增加输出到日志；
	}
}
