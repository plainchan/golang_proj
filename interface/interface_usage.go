package main

import "fmt"

//空接口
func test_empty_interfce(){
	var any interface{}
	fmt.Printf("type:%T,value:%v\n",any,any)

	any = 520
	fmt.Println(any)
	any = "I Love You"
	fmt.Println(any)
	any = true
	fmt.Println(any)
	any = 520.1314
	fmt.Println(any)
}



//类型断言

func test_assert(){
	var any interface{} = 1

	num,ok :=any.(string)
	if ok{
		fmt.Println("类型转换成功! num = ", num)
	}else{
		fmt.Println("类型转换失败!")
	}
}

func test_interface_assert(any interface{}){

	switch i:=any.(type){
	case int :
		fmt.Println("int ",i)
	case string :
		fmt.Println("string ",i)
	case bool :
		fmt.Println("bool ",i)
	case float32 :
		fmt.Println("float32 ",i)
	default:
		fmt.Println("unknown type")
	}
}



// type 接口类型名 interface{
// 	方法名1( 参数列表1 ) 返回值列表1
// 	方法名2( 参数列表2 ) 返回值列表2
// 	…
// }

type Animal interface{
	speak()
	eat()
}

type Cat struct{

}
type Dog struct{

}


func (c Cat)speak(){
	fmt.Println("Cats meow")
}
func (d Dog)speak(){
	fmt.Println("Dogs woof")
}
func (c Cat)eat(){
	fmt.Println("Cats eat mice")
}
func (d Dog)eat(){
	fmt.Println("Dogs eat meat")
}


func test_polymorphic(){

	var c Cat
	var d Dog
	c.speak()
	c.eat()
	d.speak()
	d.eat()
	

}


func main(){
	test_empty_interfce()
	// test_polymorphic()

	// test_assert()
	// var any interface{} = make([]int,5)
	// test_interface_assert(any)
}


