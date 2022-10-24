package main

import "fmt"

// 创建数组
// [capacity]data_type{element_values}


func main(){

	nums:= [10]int{1,2,3,4,5,6,7,8,9,10}
	//访问数组
	for i:=0;i<len(nums);i++ {
		fmt.Println(nums[i])
	}
}