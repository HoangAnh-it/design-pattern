package main

import (
	command "designPattern/behavioralPattern/command"
)

func main() {
	/*
		command
	*/
	myTV := &command.TV{}
	onCommand := &command.OnCommand{
		Device: myTV,
	}
	offCommand := &command.OffCommand{
		Device: myTV,
	}

	onButton := &command.Button{
		Command: onCommand,
	}

	offButton := &command.Button{
		Command: offCommand,
	}

	onButton.Press()
	offButton.Press()
}
