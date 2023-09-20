package bridge

import "fmt"

type Meazza struct{}

func (m *Meazza) PrepareMatch() {
	fmt.Println("Preparing match in Meazza")
}
