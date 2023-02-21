package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{Name: userAddr, Addr: userAddr, C: make(chan string), conn: conn, server: server}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// Online 用户的上线业务
func (user *User) Online() {
	// 用户上线，将用户加入到OnlineMap中
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	// 广播当前用户上线消息
	user.server.BroadCast(user, "login")
}

// Offline 用户的下线业务
func (user *User) Offline() {
	// 用户下线，将用户从OnlineMap中删除
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	// 广播当前用户下线消息
	user.server.BroadCast(user, "logout")
}

// SendMsg 给当前用户对应的客户端发送消息
func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg + "\n"))
}

// DoMessage 用户处理消息
func (user *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户有哪些
		user.server.mapLock.Lock()
		for _, u := range user.server.OnlineMap {
			onlineMsg := "[" + u.Addr + "]" + u.Name + ":" + "online..."
			user.SendMsg(onlineMsg)
		}
		user.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式：rename|newName
		if len(strings.Split(msg, "|")) > 2 {
			user.SendMsg("new name can not contain \"|\"")
			return
		}
		newName := strings.Replace(msg, "rename|", "", 1)
		// 判断name是否存在
		_, ok := user.server.OnlineMap[newName]
		if ok {
			user.SendMsg("current name is used")
			return
		}
		user.server.mapLock.Lock()
		delete(user.server.OnlineMap, user.Name)
		user.server.OnlineMap[newName] = user
		user.server.mapLock.Unlock()

		user.Name = newName
		user.SendMsg("rename success: " + user.Name)

	} else if len(msg) > 4 && msg[:3] == "to|" && len(strings.Split(msg, "|")) > 2 {
		// 消息格式： to|name|消息内容
		// 1 获取对方的用户名
		split := strings.Split(msg, "|")
		remoteName := split[1]
		if remoteName == "" {
			user.SendMsg("invalid message! Format: \"to|name|message\"")
			return
		}
		// 2 根据用户名获取对方User对象
		remoteUser, ok := user.server.OnlineMap[remoteName]
		if !ok {
			user.SendMsg(remoteName + " is not exist")
			return
		}
		content := strings.Replace(msg, split[0]+"|"+split[1]+"|", "", 1)
		if content == "" {
			user.SendMsg(" message is not exist, please retry.")
			return
		}
		remoteUser.SendMsg(user.Name + " to you :" + content)
	} else {
		user.server.BroadCast(user, msg)
	}
}

// ListenMessage 监听当前user channel的方法，一旦有消息，就直接发送给客户端
func (user *User) ListenMessage() {
	for {
		msg := <-user.C
		user.conn.Write([]byte(msg + "\n"))
	}

}
