package main

import (
	"fmt"
	"net"
	// "os"
	"time"
)

func main() {
	fmt.Println("Initiating server... (Ctrl-C to stop)")
	// 创建监听
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("监听失败!", err)
		return
	}
	defer socket.Close()
	for {
		// 读取数据
		data := make([]byte, 4096)
		_, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败!", err)
			continue
		}
		fmt.Println(remoteAddr)
		fmt.Printf("Client :%s\n\n", data)
		// 发送数据
		senddata := []byte(time.Now().String())
		_, err = socket.WriteToUDP(senddata, remoteAddr)
		if err != nil {
			return
			fmt.Println("发送数据失败!", err)
		}
	}
}
