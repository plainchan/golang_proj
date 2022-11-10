package main

import "fmt"



func test_rune(){
	var u rune ='中'
	fmt.Printf("%c %d\n",u,u)
	fmt.Printf("%T\n",u)
}



func main(){

	test_rune()

}
