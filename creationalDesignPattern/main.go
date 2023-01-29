package main

import (
	prototype "designPattern/creationalDesignPattern/prototype"
	// builder "designPattern/creationalDesignPattern/builder"
	// singleton "designPattern/creationalDesignPattern/singleton"
	// "fmt"
)

func main() {
	/* singleton */
	// for i := 0; i <= 1; i++ {
	// 	go func() {
	// 		x := singleton.GetInstance()
	// 		fmt.Printf("type=%T, value=%v, add=%p\n", x, x, x)
	// 	}()
	// }

	/* builder */
	// manufacturingDirector := builder.ManufacturingDirector{}
	// bicycleBuilder := &builder.BicycleBuilder{}
	// manufacturingDirector.SetBuilder(bicycleBuilder)
	// manufacturingDirector.StartManufacturing()
	// myBicycle := bicycleBuilder.GetVehicle()
	// fmt.Println(myBicycle)

	/*
		Prototype
	*/
	file1 := &prototype.File{Name: "file1"}
	file2 := &prototype.File{Name: "file2"}
	folder1 := &prototype.Folder{
		Name:     "folder1",
		Children: []prototype.INode{file1},
	}
	folder2 := &prototype.Folder{
		Name:     "folder2",
		Children: []prototype.INode{file1, file2, folder1},
	}

	folder1.PrintTree("	")
	folder2.PrintTree("	")
}
