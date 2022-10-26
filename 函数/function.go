package main

import "fmt"

// func 函数名(形式参数列表)(返回值列表){
//     函数体
// }
func add(x,y int)int{
	return x+y
}


// 匿名函数
// func(参数列表)(返回参数列表){
//     函数体
// }
// 匿名函数可以在声明后调用
func test_anonymous(){
	x:=1
	fmt.Println(x)
	func(x *int){
		(*x)++
	}(&x)
	fmt.Println(x)
}



//函数传入函数
func test_param_func(){
	
}


func main(){
	test_anonymous()
}