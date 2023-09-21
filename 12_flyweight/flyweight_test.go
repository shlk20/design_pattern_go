package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	game := newGame()

	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	fmt.Printf("The number of terrorists in the game: %d\n", len(game.terrorists))

	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	fmt.Printf("The number of counter-terrorists in the game: %d\n", len(game.counterTerrorists))

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
