package main

import "fmt"

// 创建数组
// [capacity]data_type{element_values}

func test_construct(){

	var nums1 = [3]int{1,2,3}
	nums2:=[...]int{1,2,3,4}

	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Printf("%T\n",nums1)
	fmt.Printf("%T\n",nums2)
}



func test_visited(){
	nums:= [10]int{1,2,3,4,5,6,7,8,9,10}
	//访问数组
	for i:=0;i<len(nums);i++ {
		fmt.Printf("%d\t",nums[i])
	}
	fmt.Println()
}
func test_range_visited(){
	nums:= [10]int{1,2,3,4,5,6,7,8,9,10}

	//访问数组
	for _,v := range nums{
		fmt.Printf("%d\t",v)
	}
	fmt.Println()
}

func main(){

	test_construct()
	test_visited()
	test_range_visited()
}