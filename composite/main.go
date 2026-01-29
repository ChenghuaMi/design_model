package main

import "fmt"

//组合模式

type Node interface {
	Display(indent string)
}

type File struct {
	Name string
}

func (f *File) Display(indent string) {
	fmt.Println(indent + f.Name)
}

type Dir struct {
	Name     string
	Children []Node
}

func (d *Dir) Display(indent string) {
	fmt.Println(indent + d.Name)
	for _, f := range d.Children {
		f.Display(indent + "  ")
	}
}
func main() {
	d := &Dir{
		Name: "root",
		Children: []Node{
			&Dir{
				Name: "child1 dir",
				Children: []Node{
					&Dir{
						Name: "child2 dir",
						Children: []Node{
							&File{Name: "xxx.go"},
						},
					},
				},
			},
			&File{Name: "abc.go"},
		},
	}
	d.Display("")
}
