package main

import "fmt"

func test_type(){
	var p1 *bool
	var p2 *byte
	var p3 *rune
	var p4 *int
	var p5 *float32
	var p6 *float64
	var p7 *string
	
	fmt.Printf("p1 type: %T\n",p1)
	fmt.Printf("p2 type: %T\n",p2)
	fmt.Printf("p3 type: %T\n",p3)
	fmt.Printf("p4 type: %T\n",p4)
	fmt.Printf("p5 type: %T\n",p5)
	fmt.Printf("p6 type: %T\n",p6)
	fmt.Printf("p7 type: %T\n",p7)

}

func test_pointer(){
	
	var num int = 100

	var p1 *int
	p1=&num

	p2 := new(int)
	*p2 = 100

	p3:=&num

	fmt.Printf("%T\n",p1)
	fmt.Printf("%T\n",p2)
	fmt.Printf("%T\n",p3)

	fmt.Printf("%d\n",*p1)
	fmt.Printf("%d\n",*p2)
	fmt.Printf("%d\n",*p3)
}

func main(){

	// test_type()
	test_pointer()
}
