package main

import (
	"fmt"
	"errors"
)

func divide(x,y int)(v int,err error){
	if y==0 {
		err = errors.New("divide can't be zero")
		return 
	}
	v = x/y
	return 
}

func main(){

	v,err:=divide(10,0)
	fmt.Println(v,err)
}