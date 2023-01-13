package main

import (
	"fmt"
	"unsafe"
)


func test_sizeof(){
	var y bool
	var b byte
	var c rune
	var p *int
	var f1 float32
	var f2 float64
	var s string
	fmt.Println("bool size:",unsafe.Sizeof(y))
	fmt.Println("byte size:",unsafe.Sizeof(b))
	fmt.Println("rune size:",unsafe.Sizeof(c))
	fmt.Println("pointer size:",unsafe.Sizeof(p))
	fmt.Println("float32 size:",unsafe.Sizeof(f1))
	fmt.Println("float64 size:",unsafe.Sizeof(f2))

	// 对不同长度的字符串,unsafe.Sizeof() 函数的返回值都为 16,
	// string 类型对应一个结构体，该结构体有两个域，第一个域指向该字符串的指针，第二个域为字符串的长度，
	// 每个域占 8 个字节，但是并不包含指针指向的字符串的内容，t
	fmt.Println("string size:",unsafe.Sizeof(s))  

}



// unsafe.pointer
// 任意类型的指针值都可以转换为unsafe.Pointer (A pointer value of any type can be converted to a Pointer.)
// unsafe.Pointer可以转换为任意类型的指针值   (A Pointer can be converted to a pointer value of any type.)
// uintptr可以转换为unsafe.Pointer         (A uintptr can be converted to a Pointer.)
// unsafe.Pointer可以转换为uintptr         (A Pointer can be converted to a uintptr.)
func test_pointer(){
	

}



func main(){
	test_sizeof()


}