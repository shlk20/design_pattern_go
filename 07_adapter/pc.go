package adapter

import "fmt"

type PC struct{}

func (p *PC) InsertUSBPort() {
	fmt.Println("USB connector is plugged into PC machine.")
}
