package command

type Button struct {
	Command Command
}

func (button *Button) Press() {
	button.Command.execute()
}
