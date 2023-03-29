package main

// 既然与server.go是同一个包，那就没必要import了，直接编译即可
// go build -o + 两个go文件 不太好用，会碰到不编译另一个文件的问题
// go build + 两个go文件，go build . 可以用
// 其实是我搞错了，go build -o + 可执行程序名 + 两个go文件才是正确的编译指令
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}

//创建一个server对象
//启动server服务
//处理链接业务

//nc linux下自带的网络工具
