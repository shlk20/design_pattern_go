package bridge

import "testing"

func TestMilanMatch(t *testing.T) {
	sanSiro := &SanSiro{}
	meazza := &Meazza{}

	milan := &Milan{}
	milan.SetStadium(sanSiro)
	milan.KickOff()
	milan.SetStadium(meazza)
	milan.KickOff()

	inter := &Inter{}
	inter.SetStadium(meazza)
	inter.KickOff()
	inter.SetStadium(sanSiro)
	inter.KickOff()
}
