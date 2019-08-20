// myQQSer project main.go
//多人聊天室服务器端
package main

import (
	"fmt"
	"net"
)

//TCPConn代表一个TCP网络连接，实现了Conn接口。
var ConnMap map[string]*net.TCPConn

//检测错误状态类型
func checkErr(tcpConn *net.TCPConn, err error) int {
	if err != nil {
		fmt.Printf("客户端%s退出。。。\n", tcpConn.RemoteAddr().String())
		return -1
	}
	return 1
}

func say(tcpConn *net.TCPConn, nickname string) {
	for {
		data := make([]byte, 256)
		total, err := tcpConn.Read(data)
		if err != nil {
			flag := checkErr(tcpConn, err)
			if flag == 0 || flag == -1 {
				break
			}
		} else {
			fmt.Printf("%s说:%s\n", nickname, string(data[:total]))
		}

		for _, conn := range ConnMap {
			if conn.RemoteAddr().String() == tcpConn.RemoteAddr().String() {
				continue
			}
			data = append([]byte(nickname+"说:"), data[:total]...)
			conn.Write(data)
		}
	}
}

func main() {
	fmt.Println("服务器开机。。。")
	//---------------------------------------//
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	fmt.Println("开启监听连接请求。。。")
	tcpListen, _ := net.ListenTCP("tcp", tcpAddr)
	ConnMap = make(map[string]*net.TCPConn)
	for {
		fmt.Println("等待新客户端连接。。。")
		tcpConn, _ := tcpListen.AcceptTCP()
		defer tcpConn.Close()
		//添加用户连接信息到连接字典中
		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn

		//获取客户端用户名
		b := make([]byte, 256)
		numOfb, _ := tcpConn.Read(b)
		nickname := string(b[:numOfb])
		fmt.Println("已连接客户端信息:"+tcpConn.RemoteAddr().String(),
			"\t昵称："+ nickname)

		go say(tcpConn, nickname)
	}
}
