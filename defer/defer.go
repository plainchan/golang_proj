package main

import "fmt"


//延迟调用
//在return后调用

func test_defer() int{

	num:=1

	defer func(){
		fmt.Println("Hello World ",num)
		num+=3
	}()
	defer func(){
		fmt.Println("Hello World ",num)
		num+=2
	}()
	defer func(){
		fmt.Println("Hello World ",num)
		num+=1
	}()

	return num
}

func main(){

	num := test_defer()
	fmt.Println(num)
}
