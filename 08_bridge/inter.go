package bridge

import "fmt"

type Inter struct {
	stadium Stadium
}

func (i *Inter) KickOff() {
	i.stadium.PrepareMatch()
	fmt.Println("Inter Milan match kicked off")
}

func (i *Inter) SetStadium(s Stadium) {
	i.stadium = s
}
