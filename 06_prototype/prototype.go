package prototype

import "fmt"

type INode interface {
	print(string)
	clone() INode
}

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() INode {
	return &File{
		name: f.name + "_clone",
	}
}

type Folder struct {
	name     string
	children []INode
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() INode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	children := []INode{}

	for _, i := range f.children {
		copy := i.clone()
		children = append(children, copy)
	}

	cloneFolder.children = children
	return cloneFolder
}
