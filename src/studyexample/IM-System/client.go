package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建对象
	client := &Client{ServerIp: serverIp, ServerPort: serverPort}
	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println(" net.Dial err:", err)
		return nil
	}
	client.conn = conn
	client.flag = 999
	client.Name = conn.RemoteAddr().String()
	return client
}

func (client *Client) DealResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)
	//上面这句与下面for循环实现效果一致
	//for {
	//	buf := make([]byte, 1024)
	//	client.conn.Read(buf)
	//	fmt.Println(buf)
	//}
}

func (client *Client) menu() bool {
	var check string

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&check)

	if check == "" {
		return false
	}
	value, err := strconv.Atoi(check)
	if err != nil {
		fmt.Println(">>>>>请输入合法范围内的数字<<<<<")
		return false
	}

	if value >= 0 && value <= 3 {
		client.flag = value
		return true
	} else {
		fmt.Println(">>>>>请输入合法范围内的数字<<<<<")
		return false
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for !client.menu() {

		}
		switch client.flag {
		case 1:
			// 公聊模式
			client.PublicChat()
			break
		case 2:
			// 私聊模式
			client.PrivateChat()
			break
		case 3:
			// 更新用户名
			client.UpdateName()
			break
		}

	}
}

func (client *Client) sendMsg(msg string) {
	_, err := client.conn.Write([]byte(msg + "\n"))
	if err != nil {
		fmt.Println(" conn.Writer err: ", err)
		return
	}
}

// SelectUsers 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who"
	client.sendMsg(sendMsg)

}

// PrivateChat 私聊模式
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println(">>>>>请输入聊天对象[用户名], exit退出.")
	fmt.Scanln(&remoteName)
	for remoteName != "exit" {

		fmt.Println(">>>>>正在和" + remoteName + "私聊，请输入聊天内容，exit退出")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// 发给服务器
			// 消息不为空则发送
			sendMsg := "to|" + remoteName + "|" + chatMsg
			if len(chatMsg) != 0 {
				client.sendMsg(sendMsg)
			}
			chatMsg = ""
			fmt.Println(">>>>>正在和" + remoteName + "私聊，请输入聊天内容，exit退出")
			fmt.Scanln(&chatMsg)
		}
		client.SelectUsers()
		fmt.Println(">>>>>请输入聊天对象[用户名], exit退出.")
		fmt.Scanln(&remoteName)
	}
}

// PublicChat 公聊模式
func (client *Client) PublicChat() {
	var chatMsg string

	fmt.Println(">>>>>正在公聊，请输入聊天内容，exit退出")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {
		// 发给服务器
		// 消息不为空则发送
		if len(chatMsg) != 0 {
			client.sendMsg(chatMsg)
		}
		chatMsg = ""
		fmt.Println(">>>>>正在公聊，请输入聊天内容，exit退出")
		fmt.Scanln(&chatMsg)
	}
}

// UpdateName 更新用户名
func (client *Client) UpdateName() {
	fmt.Print(">>>>>请输入用户名：")
	fmt.Scanln(&client.Name)
	client.sendMsg("rename|" + client.Name)
	time.Sleep(time.Millisecond * 5)
}

var serverIp string
var ServerPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认127.0.0.1）")
	flag.IntVar(&ServerPort, "port", 8888, "设置服务器端口（默认8888）")

}

func main() {
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, ServerPort)
	if client == nil {
		fmt.Println(">>>>> 链接服务器失败 <<<<<")
	}
	fmt.Println(">>>>> 链接服务器成功 <<<<<")

	// 单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	time.Sleep(time.Millisecond * 5)

	// 启动客户端
	client.Run()
}
