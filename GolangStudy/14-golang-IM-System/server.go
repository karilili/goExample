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

	//在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线User
func (this *Server) ListernMessager() {
	for {
		fmt.Println("server ListernMessager...")
		msg := <-this.Message

		//将msg发送给全部的在线User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	//...当前链接的业务
	// fmt.Println("链接建立成功")
	defer fmt.Println("Handler end")

	user := NewUser(conn, this)

	//用户上线，将用户加入到onlineMap中
	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			fmt.Println("server Handle Read Client Msg")
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			//提取用户的消息（去除'\n'）
			msg := string(buf[:n-1])

			//用户针对msg进行广播
			user.DoMessage(msg)

			//用户的任意消息，代表当前用户是活跃的
			isLive <- true
		}
	}()

	//当前handler阻塞
	for {
		select {
		case <-isLive:
			//当前用户是活跃的，应该重置定时器
			//不做任何事情，为了激活select，更新下面的定时器

			//需要把isLive放到上面
			//当这个case触发后，下面的case中的代码虽然不会执行，但是条件会执行
			//而对于time.After，只要再次执行，就会重置定时器
			fmt.Println("Case isLive")
		case <-time.After(time.Second * 10):
			//已经超时
			//将当前user强制关闭
			user.SendMsg("你被踢了")
			//销毁用户的资源
			//！关闭channe后，读取channe时会立即返回，不会阻塞
			//！因此user中的listenMessage函数会导致CPU占用率高
			//close(user.C)
			//关闭连接
			conn.Close()
			//退出当前handler
			fmt.Println("Case time return")
			return //runtime.Goexit()	或者调用这个函数
		}
		fmt.Println("server Handle for ...")
	}
}

// 启动服务器的接口
func (this *Server) Start() {
	//socket lister
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//close lister socket
	defer listener.Close()

	//启动监听Message的goroutine
	go this.ListernMessager()

	//死循环，不断监听
	for {
		//accpet
		fmt.Println("server accpet...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listen accpet err:", err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}
}
