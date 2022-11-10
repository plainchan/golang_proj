package main

import "fmt"

func test_map_construct(){
	roman:=make(map[int]string)
	roman = map[int]string {
		1:"I",
		2:"II",
		3:"III",
		4:"IV",
		5:"V",
	}
	roman[6]="VI"
	roman[7]="VII"
	fmt.Println(roman[7])
}


func main(){
	test_map_construct()
}
