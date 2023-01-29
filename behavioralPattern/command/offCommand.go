package command

type OffCommand struct {
	Device Device
}

func (offCommand *OffCommand) execute() {
	offCommand.Device.off()
}
