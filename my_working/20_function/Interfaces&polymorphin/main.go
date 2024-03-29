package main

import (
	"fmt"
)

type person struct {
	first string
	last string
}
type secretAgent struct {
	person
	ltk bool
}
//func(reciver_name Type) method_name(parameter_list)(return_types){...code...}
func (s secretAgent) speak(){
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak.")
}
func (p person) speak(){
	fmt.Println("I am", p.first, p.last, " - the person speark.")
}
type human interface {
	speak()
}

func bar(h human){

	switch h.(type){
	case person:
		fmt.Println("I was passed into bar.", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed intor bar", h.(secretAgent).first)
	}

	fmt.Println("I was passed into bar.", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			first: "vinh",
			last: "bui",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Quynh",
			last: "le",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last: "Yes",
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	// Conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
