package main

import (
	"fmt"
	ent "internship/basic/entity"
)

func main() {
	var names = []string{"a", "b", "c", "d", "e"}
	person := &ent.Person{Name: "Alex", Gender: "Male"}
	fmt.Println(person.IsMale())
	//fmt.Println(ReturnTypes())

	for i, name := range names {
		fmt.Println(name, "at position", i)
	}
}

func ReturnTypes() (age int, name string) {
	return 22, "Alex"
}
