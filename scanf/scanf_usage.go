package main

import (
	"fmt"
)

func test_scanf(){
	var n int
	fmt.Scanf("%d",&n)
	fmt.Println(n)
	fmt.Scanf("%d",&n)
	fmt.Println(n)	
}

func test_scanln(){
	var n int
	fmt.Scanln(&n)
	fmt.Println(n)
	fmt.Scanln(&n)
	fmt.Println(n)
}




func main(){
	test_scanf()
	test_scanln()
}