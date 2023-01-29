package command

import "fmt"

type TV struct {
}

func (tv *TV) on() {
	fmt.Println("Tivi is on")
}

func (tv *TV) off() {
	fmt.Println("Tivi is off")
}
