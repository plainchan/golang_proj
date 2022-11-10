package main

import "fmt"

func test(){
	nums:=new(int)
	*nums = 250
	fmt.Println(nums,*nums)
}

func main(){
	test()
}
