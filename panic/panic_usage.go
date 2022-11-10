package main

import "fmt"


//不能跨协程
func test_panic(){

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	
	panic("crash")
}



func main(){
	test_panic()
}