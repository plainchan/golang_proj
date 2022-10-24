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

func getName(person *Person) string{
	return person.name
}


func main(){

	p:=newPerson("plain",18,"plain@plain.com")
	
	fmt.Println(p.name)
	fmt.Println(getName(p))
}
