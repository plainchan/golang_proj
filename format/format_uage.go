package main

import "fmt"



func main(){
	a,b,c:=1,2,3
	fmt.Println("Hello World",a)
	fmt.Printf("Hello World %d\n",b)
	fmt.Print("Hello World ",c)
}


// General
//     %v 以默认的方式打印变量的值
//     %T 打印变量的类型

// Integer
//     %+d 带符号的整型，fmt.Printf("%+d", 255)输出+255
//     %q 打印单引号
//     %o 不带零的八进制
//     %#o 带零的八进制
//     %x 小写的十六进制
//     %X 大写的十六进制
//     %#x 带0x的十六进制
//     %U 打印Unicode字符
//     %#U 打印带字符的Unicode
//     %b 打印整型的二进制


// Float
//     %f (=%.6f) 6位小数点
//     %e (=%.6e) 6位小数点（科学计数法）
//     %g 用最少的数字来表示
//     %.3g 最多3位数字来表示
//     %.3f 最多3位小数来表示


// String Width (以5做例子）
//     %5s 最小宽度为5
//     %-5s 最小宽度为5（左对齐）
//     %.5s 最大宽度为5
//     %5.7s 最小宽度为5，最大宽度为7
//     %-5.7s 最小宽度为5，最大宽度为7（左对齐）
//     %5.3s 如果宽度大于3，则截断
//     %05s 如果宽度小于5，就会在字符串前面补零

// Struct
//     %v 正常打印。比如：{sam {12345 67890}}
//     %+v 带字段名称。比如：{name:sam phone:{mobile:12345 office:67890}
//     %#v 用Go的语法打印。
//     比如main.People{name:"sam", phone:main.Phone{mobile:"12345", office:"67890"}}


// Boolean
//     %t 打印true或false

// Pointer
//     %p 带0x的指针
//     %#p 不带0x的指针

