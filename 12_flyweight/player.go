package flyweight

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
