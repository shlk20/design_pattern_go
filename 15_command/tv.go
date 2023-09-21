package command

import "fmt"

type TV struct {
	isRunning bool
}

func (tv *TV) on() {
	tv.isRunning = true
	fmt.Println("Turning tv on")
}

func (tv *TV) off() {
	tv.isRunning = false
	fmt.Println("Turning tv off")
}
