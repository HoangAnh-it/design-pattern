package prototype

import "fmt"

type File struct {
	Name string
}

func (file *File) Clone() INode {
	return &File{Name: file.Name + "_clone"}
}

func (file *File) PrintTree(space string) {
	fmt.Println(space + file.Name)
}
