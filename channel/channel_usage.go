package main

import (
	"fmt"
	"time"
	// "sync"
)

func test_channe_read_write(){

	channel:=make(chan int,10)
	channel<-1
	channel<-2
	channel<-3

	close(channel)

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)  
	//缓存读取完毕,也可以继续读取，读取的值为零
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func test_goroutine_sync(){
	done:=make(chan bool,1)
	defer close(done)

	//协程
	go func(done chan bool){
		fmt.Println("worker routine")
		time.Sleep(time.Second)
		done<- true
	}(done)

	fmt.Println("boss routine")
	
	<-done   //<-done用来从channel done中接收数据，这个表达式会一直被block,直到有数据可以接收。
	fmt.Println("worker done")
}


func test_channel_range(){
	channel:=make(chan int)
	go func(){
		for i:=0;i<10;i++{
			channel <- i 
		}

		close(channel) //close 可以用于退出range,否则会陷入range死循环
	}()

	for i:=range channel{
		fmt.Println(i)
	}
}





func main(){
	test_channe_read_write()
	// test_goroutine_sync()
	// fmt.Println("Hello World")

}