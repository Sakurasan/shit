package main

import (
	"fmt"
	"net"
	"os"
)

var (
	str string
)

func main() {
	// 创建连接
	fmt.Println("Initiating server... (Ctrl-C to stop)")
	// conn, err := net.Dial("udp", ":8080")
	raddr, _ := net.ResolveUDPAddr("udp", ":8080")
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Println("连接失败!", err)
		return
	}
	defer conn.Close()
	for {
		// 发送数据
		fmt.Printf("请输入：")
		fmt.Scanf("%s", &str)

		msg := []byte(str)
		_, err := conn.Write(msg)

		if err != nil {
			fmt.Println("发送数据失败!", err)
			return
		}
		if str == "quit" {
			fmt.Println("退出 Socket_udp")
			os.Exit(1)
		}
		// 接收数据
		data := make([]byte, 2048)
		_, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败!", err)
			return
		}
		fmt.Println(remoteAddr)
		fmt.Printf("Serve return:%s\n", data)
	}
}
