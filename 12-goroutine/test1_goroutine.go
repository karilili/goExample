//多线程实行了并发处理，但如果没有多核，则无法实现并行
//通过时间片轮询的方式，进程/线程的数量越多，切换成本就越大，也就越浪费
//软件架构与程序设计，优化性能，目标之一也是减少切换成本
//多线程开发设计越来越复杂
//进程占据虚拟内存越4GB，而线程则是4MB

//多进程调度模型 存在高调度CPU消耗与高内存消耗的问题
//协程 通过协程调度器 绑定一个线程。一个线程意味着CPU切换成本低了，但依然能实现用户层面的并发执行
//但如果一个协程阻塞，则可能影响下一个协程的执行

//协程（co-routine） --->  Goroutine
//这一章所讨论的线程，是内核层面的

//G M P
//M指内核线程，P是Processors，G是协程。一般一个P绑定一个M，一个P又绑定一队儿G
//P就是协程调度器，管理一个队列的协程，一个P一次执行一个G，P的数量通过宏GOMAXPROCS进行设置
//一个程序真正的并发数量即GOMAXPROCS
//存在一个全局G队列，当每个P的G队列满了，则多出来的G就会放在全局队列中

//调度器设计策略 复用线程 利用并行 抢占 全局G队列
//复用线程：work stealing机制：空闲线程获取另一个队列中的协程
//hand off机制：如果一个G阻塞，则会创建一个新线程，让P及其G队列到这个新线程上执行
//而原来的G则留在原来的M上，且M被休眠或销毁
//利用并行： GOMAXPROCS限定P的个数，从而只会使用部分CPU内核
//抢占：每个G最多占用10ms的CPU，其他G会在时间到了后抢占该CPU
//全局G队列：M空闲时，优先从其他P队列中偷取G，如果其他队列中没有，则会从全局G队列中偷取

package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine: i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	//创建一个go程 去执行newTask() 流程  //使用go关键字即可创建协程
	go newTask()

	fmt.Println("main goroutine exit") //main推出后，子goroutine都会死掉

	// i := 0
	// for {
	// 	i++
	// 	fmt.Printf("main Goroutine: i=%d\n", i)
	// 	time.Sleep(1 * time.Second)
	// }
}
