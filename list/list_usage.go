package  main

import (
	"fmt"
	"container/list"
)


func test_list(){
	l:=list.New()

	l.PushBack("HelloWorld")
	l.PushBack(1234)

	fmt.Println(l.Size())
}



func main(){

	test_list()
}