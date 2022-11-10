package main

import (
	"fmt"
	"sync"
)

// 
func add(x,y int,doneFunc func())int{

	defer doneFunc()
	fmt.Printf("add(%d+%d)=%d\n",x,y,x+y)
	return x+y
}


func main(){
	wg:=sync.WaitGroup{}
	wg.Add(10)
	for i:=0;i<10;i++ {
		go add(i,i,wg.Done)	
	}
	wg.Wait()
}