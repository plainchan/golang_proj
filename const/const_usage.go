package main

import "fmt"

const Pi float64 = 3.1415926

const (
	One int = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
)




func main(){
	fmt.Println(Pi)
	fmt.Println(One)
	fmt.Println(Five)
	fmt.Println(Ten)
}
