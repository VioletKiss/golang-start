package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{Ip: ip, Port: port, OnlineMap: make(map[string]*User), Message: make(chan string)}
	return server
}

// ListenMessage 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线user
func (server *Server) ListenMessage() {
	for {
		msg := <-server.Message
		server.mapLock.Lock()
		for _, i2 := range server.OnlineMap {
			i2.C <- msg
		}
		server.mapLock.Unlock()
	}
}

// BroadCast 广播消息的方法
func (server *Server) BroadCast(user *User, msg string) {

	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	server.Message <- sendMsg
}

func (server *Server) Handler(conn net.Conn) {
	// .. 处理当前链接的业务
	//fmt.Println("connect success")

	user := NewUser(conn, server)
	fmt.Println("[" + user.Addr + "]" + user.Name + ":" + "online")

	// 用户上线
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			read, err := conn.Read(buf)
			if read == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err", err)
				return
			}
			// 提取用户的消息，去除'\n'
			msg := string(buf[:read-1])

			// 用户针对msg进行消息处理
			user.DoMessage(msg)

			// 用户任意消息，代表用户活跃
			isLive <- true

		}
	}()
	// 当前handle阻塞
	for {
		select {
		case <-isLive:
		// 当前用户活跃，重置定时器
		// 不做任何事情，更新下面的定时器 time.After在重新执行时重置
		case <-time.After(time.Minute * 5):
			//已经超时
			//将当前的User强制关闭
			user.SendMsg("you are removed")

			// 销毁资源
			close(user.C)

			// 关闭链接
			conn.Close()

			// 退出handler
			return // runtime.Goexit()
		}
	}
}

func (server *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close socket listen
	defer listener.Close()

	// 启动监听Message的goroutine
	go server.ListenMessage()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		// do handler
		go server.Handler(conn)
	}

}
