package main

import "fmt"

type a interface {
	name()
	name2()
}

type s1 struct {
}

func (s s1) name() {
	fmt.Println("name s1")
}

func (s *s1) name2() {
	fmt.Println("name2 s1")
}

type s2 struct {
}

func (s s2) name() {
	fmt.Println("name s2")
}

func (s s2) name2() {
	fmt.Println("name2 s1")
}

func main() {
	var obj a = &s1{}
	obj = s2{}

	obj.name()
}
