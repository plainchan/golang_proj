package main

import "fmt"

func test_construct(){
	arr1 := [] int {1,2,3}
	arr2 := make([]int,3,5)
	var arr3 []int

	fmt.Printf("arr1:%d  len:%d  capacity:%d\n",arr1,len(arr1),cap(arr1))
	fmt.Printf("arr2:%d  len:%d  capacity:%d\n",arr2,len(arr2),cap(arr2))
	fmt.Printf("arr3:%d  len:%d  capacity:%d\n",arr3,len(arr3),cap(arr3))

}

func test_append(){

	arr := [] int {1,2,3}
	fmt.Printf("arr:%d  len:%d  capacity:%d\n",arr,len(arr),cap(arr))

	arr = append(arr,[] int {4,5,6}...)
	fmt.Printf("arr:%d  len:%d  capacity:%d\n",arr,len(arr),cap(arr))

	arr = append(arr,7)
	fmt.Printf("arr:%d  len:%d  capacity:%d\n",arr,len(arr),cap(arr))
}



func main(){
	// test_construct()
	test_append()
}
