package bridge

import "fmt"

type Milan struct {
	stadium Stadium
}

func (m *Milan) KickOff() {
	m.stadium.PrepareMatch()
	fmt.Println("AC match kicked off")
}

func (m *Milan) SetStadium(s Stadium) {
	m.stadium = s
}
