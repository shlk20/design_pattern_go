package adapter

import "fmt"

type pcAdapter struct {
	pc *PC
}

func (pa *pcAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	pa.pc.InsertUSBPort()
}
