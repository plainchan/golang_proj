package main

import "fmt"


//不初始化数据默认值
func test_non_init_num(){
	var (
		n1 bool
		n2 byte
		n3 rune
		n4 int
		n5 float32
		n6 float64
		n7 string
	)

	fmt.Println(n1,n2,n3,n4,n5,n6,n7)
	
}

// nil 代表 the zero value
// nil is a predeclared identifier representing the zero value for a
// pointer, channel, func, interface, map, or slice type.
func test_nil(){

	var p *int
	var c <-chan int
	var f func(int)int
	var i interface{}
	var m map[string]string
	var s []int
	
	
	fmt.Println("Uninitialized pointer is nil, ",p==nil)
	fmt.Println("Uninitialized channel is nil, ",c==nil)
	fmt.Println("Uninitialized func is nil, ",f==nil)
	fmt.Println("Uninitialized interface is nil, ",i==nil)
	fmt.Println("Uninitialized map is nil, ",m==nil)
	fmt.Println("Uninitialized slice is nil, ",s==nil)
	
	

}





func main(){
	test_non_init_num()
	test_nil()
}
