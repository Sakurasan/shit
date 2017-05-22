package main

import (
	"fmt"
	"net"
	"os"
	// "time"
)

var (
	address  = ":8080"
	buffer   = make([]byte, 2048)
	str, msg string
)

func main() {

	fmt.Println("正在连接...")
	conn, err := net.Dial("tcp", address)
	// conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		fmt.Println("Serve not found")
		os.Exit(-1)
	}
	defer conn.Close()

	fmt.Println("连接成功")

	for {

		fmt.Printf("输入发送信息：")
		fmt.Scanf("%s", &str)
		if str == "exit" {
			fmt.Println("退出")
			os.Exit(1)
		}

		send, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Printf("发送失败：%s", send)
			os.Exit(0)
		}

		len, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("读取消息错误")
			os.Exit(0)
		}

		msg = string(buffer[0:len])
		fmt.Printf("Serve:%s \n", msg)

	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("错误信息：%s", err)
	}
}
