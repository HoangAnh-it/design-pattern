package prototype

import "fmt"

type Folder struct {
	Name     string
	Children []INode
}

func (folder *Folder) Clone() INode {
	cloneFolder := &Folder{}
	cloneFolder.Name = folder.Name + "_clone"
	for _, node := range folder.Children {
		cloneFolder.Children = append(cloneFolder.Children, node.Clone())
	}
	return cloneFolder
}

func (folder *Folder) PrintTree(space string) {
	fmt.Println(space + folder.Name)
	for _, children := range folder.Children {
		children.PrintTree(space + space)
	}
}
