package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	//链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}
	client.conn = conn

	//返回对像
	return client
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1. 公聊模式")
	fmt.Println("2. 私聊模式")
	fmt.Println("3. 修改用户名")
	fmt.Println("0. 退出")

	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>请输入合法范围内的数字<<<<<<")
		return false
	}
}

func (client *Client) Run() {
	for client.flag != 0 { //如果等于0，则跳出循环。不等于0，则一直循环
		for client.menu() != true { //如果不为true，就一直循环，执行menu函数，否则退出循环
		}
		switch client.flag {
		case 1:
			//公聊模式
			fmt.Println("公聊模式选择...")
			break
		case 2:
			//私聊模式
			fmt.Println("私聊模式...")
			break
		case 3:
			//修改用户名
			fmt.Println("修改用户名...")
			break
		case 0:
			//退出
			fmt.Println("退出...")
			break
		}
	}
}

//解析命令行参数 需要使用到flag包

var serverIp string
var serverPort int

// ./Client -ip 127.0.0.1 -port 8888

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
	//第一个形参是命令的提示词，第二个是默认值，第三个是提示信息
}

func main() {
	//命令行解析
	flag.Parse() //解析这个进程环境变量中是否有提示词

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>链接服务器失败...")
	}
	fmt.Println(">>>>>>>>链接服务器成功...")

	//启动客户端业务
	client.Run()
}
