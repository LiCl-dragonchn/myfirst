// myQQCli project main.go
//多人聊天室客户端
package main

import (
	"fmt"
	"net"
	"os"
)

var nickname string

func reader(conn *net.TCPConn) {
	buff := make([]byte, 256)
	for {
		total, err := conn.Read(buff)
		if err != nil {
			fmt.Println("服务器连接断开。。。")
			os.Exit(2)
		}
		fmt.Printf("%s\n", buff[0:total])
	}
}

func main() {
	fmt.Println("客户端开启。。。")
	//------------------------------------------------------//
	fmt.Println("请输入昵称")
	fmt.Scanln(&nickname)
	fmt.Println("你的昵称是:", nickname)

	TcpAdd, _ := net.ResolveTCPAddr("tcp", "172.16.2.83:8080")
	conn, err := net.DialTCP("tcp", nil, TcpAdd)
	if err != nil {
		fmt.Println("无法连接到服务器。。。")
		os.Exit(1)
	}
	//发送用户名到服务器
	b := []byte(nickname)
	conn.Write(b)
	defer conn.Close()
	go reader(conn)

	//输入与发送信息
	for {
		var msg string
		fmt.Scan(&msg)
		fmt.Print(nickname + "说:")
		fmt.Println(msg)
		b = []byte(msg)
		conn.Write(b)
	}
}
