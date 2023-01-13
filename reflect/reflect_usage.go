package main


import (
	"fmt"
	"reflect"
)



// 反射可以将接口类型变量 转换为"反射类型对象"；

// 反射可以将 "反射类型对象"转换为 接口类型变量；

// 如果要修改 "反射类型对象" 其类型必须是 可写的；

func test_reflect(){
	var i interface{} = 1
	fmt.Printf("%T %v\n",i,i)

	t:= reflect.TypeOf(i)
	v:= reflect.ValueOf(i)

	fmt.Printf("%T \n",t)
	fmt.Printf("%T \n",v)

	ii:=v.Interface()

	fmt.Printf("%T \n",ii)


	fmt.Println(v.CanSet())
}


func main(){

	test_reflect()
}

