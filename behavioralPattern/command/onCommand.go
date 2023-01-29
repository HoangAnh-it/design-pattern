package command

type OnCommand struct {
	Device Device
}

func (command *OnCommand) execute() {
	command.Device.on()
}
