pakage main

import "net"

type User struct{
	Name string
	Addr string
	C  chan string
	conn net.Conn
}

//创建一个用户的API
func NewUser(Conn net.conn) *User{
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name : userAddr
		Addr : userAddr
		c : make(chan string)
		conn : conn
	}

	//启动监听当前user channel消息的goroutine
	go user.ListenMessage()
	return user
}

//监听当前User channel的方法，一旦有消息，就直接发送给对端客户端
func (this *User) ListenMessage(){
	for{
		msg:=<-this.C
        this.conn.Write([]byte(msg+"\n"))
	}
}