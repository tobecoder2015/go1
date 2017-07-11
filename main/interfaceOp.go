package main

import (
	"fmt"
	"reflect"
)

type Men interface {
	sing(song string )
	sayHi()
}

type Human struct {
	name string
	age int
	phone string
}


type Student struct {
	Human
	school string
	major string
}

func (h Human)sayHi()  {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}


func (s Student)sayHi()  {
	fmt.Printf("Hi, I am %s you can call me on %s ,my school is %s\n", s.name,s.phone,s.school)
}


func (h Human)sing(song string)  {
	fmt.Printf(" %s sing song  %s\n", h.name, song)
}


func (s Student)sing(song string)  {
	fmt.Printf(" %s sing song  %s\n", s.name, song)
}

func main()  {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", "Math"}
	mike.sayHi();


	var wqs Men;
	wqs = Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", "Math"}
	wqs.sing("七里香");


	p := new(Human)
	//查看类型
	fmt.Println(reflect.TypeOf(p))
	p.age=30
	p.name="xxx"
	p.sayHi()
}


