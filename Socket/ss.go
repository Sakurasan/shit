package main

import (
	"fmt"
	// "io/ioutil"
	"net"
	"os"
	// "time"
)

var (
	buffer = make([]byte, 2048)
	str    string
)

func main() {
	laddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080") //定义一个本机IP和端口。

	fmt.Println("Initiating server... (Ctrl-C to stop)")
	Listener, err := net.ListenTCP("tcp", laddr)
	defer Listener.Close()
	if err != nil {
		fmt.Println("监听出错\n", err)
		os.Exit(-1)
	}

	for {

		conn, err := Listener.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println("链接失败")
			os.Exit(0)
		}
		fmt.Println("等待连接...")

		go func(conn net.Conn) {
			fmt.Println("新连接: ", conn.RemoteAddr())
			for {
				n, err := conn.Read(buffer)
				if err != nil {
					fmt.Printf("Client %v quit.\n", conn.RemoteAddr())
					conn.Close()
					return
				}
				str = string(buffer[0:n])
				fmt.Printf("Client:%s  Send:%s\n", conn.RemoteAddr(), str)
				conn.Write([]byte("收到:(" + str + ")"))
			}
		}(conn)

	}

}
