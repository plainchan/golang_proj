package main

import "fmt"

type Person struct{
	name  string
	age   int
	email string
}

func newPerson(name string,age int,email string) *Person{
	return &Person{name,age,email}
}

// 类方法
// 由于 Go 语言不支持 class 这样的代码块，要为 Go 类添加成员方法，
// 需要在 func 和方法名之间添加方法所属的类型声明（有的地方将其称之为接收者声明）
func (person *Person)getName() string{
	return person.name
}
func (person *Person)setName(name string){
	person.name = name
}


func main(){

	p:=newPerson("plain",18,"plain@plain.com")
	fmt.Println(p.getName())
	p.setName("plaingo")
	fmt.Println(p.getName())
}
