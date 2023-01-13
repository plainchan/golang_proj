package main

import "fmt"

type Person struct{
	name  string 
	age   int    
	email string 
}


func test_struct(){
	
}

//结构体标签
type test struct{
	
}

func main(){
	p:= Person{
		name: "plain",
		age:  18,
		// email: "plain@plain.com",
	}
	
	fmt.Println(p)
}
