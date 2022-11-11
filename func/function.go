package main

import (
	"fmt"
	"reflect"
)

// func 函数名(形式参数列表)(返回值列表){
//     函数体
// }
func add(x,y int)int{
	return x+y
}
func sub(x,y int)int{
	return x-y
}
func mul(x,y int)int{
	return x*y
}


// 匿名函数
// func(参数列表)(返回参数列表){
//     函数体
// }
// 匿名函数可以在声明后调用
func test_anonymous(){
	x:=1
	fmt.Println(x)
	func(x *int){
		(*x)++
	}(&x)
	fmt.Println(x)
}

//引用传参
func test_param_pointer(x *int){
	*x*=2
}




//函数作为形参
func test_param_func(x,y int,process func(int,int)int) int{
	return process(x,y)
}


//变长参数
func test_various_paran(nums...int){
	for _,v := range nums{
		fmt.Println(v)
	}

}

// 任意类型的变长参数（泛型）
// 反射
func test_myPrintf(args ...interface{}) {
    for _, arg := range args {
        switch reflect.TypeOf(arg).Kind() {
        case reflect.Int:
            fmt.Println(arg, "is an int value.")
        case reflect.String:
            fmt.Printf("\"%s\" is a string value.\n", arg)
        case reflect.Array:
            fmt.Println(arg, "is an array type.")
		case reflect.Map:
			fmt.Println(arg, "is an map type.")
        default:
            fmt.Println(arg, "is an unknown type.")
        }
    }
}

func main(){
	// test_anonymous()
	// fmt.Println(test_param_func(4,5,mul))
	// test_various_paran()
	test_myPrintf(1,3,[...]int{1,2,3},map[string]string{"name":"plain","sex":"male"})
}