package main

import "fmt"

type person struct {

	fname string
	lname string
	age int
}
func (p person ) fullname() string{

	return   p.fname + p.lname
}
func main(){



	p1:= person{"Alex","Bwanakweli",45}
println(p1.fname,p1.lname,p1.age)
fmt.Println(p1.fullname())

	}
