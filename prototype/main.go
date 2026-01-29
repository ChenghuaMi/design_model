package main

import "fmt"

//原型模式

type Prototype interface {
	Clone() Prototype
}
type Stu struct {
	Name string
	Age  int
}

func (s *Stu) Clone() Prototype {
	return &Stu{
		Name: s.Name,
		Age:  s.Age,
	}
}
func main() {
	s1 := Stu{
		Name: "zs",
		Age:  0,
	}
	s2 := s1.Clone().(*Stu)
	s2.Name = "ssss"
	s2.Age = 100
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}
