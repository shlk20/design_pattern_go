package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	pizza := &VeggieMania{}

	fmt.Printf("Price of veggeMania is %d\n", pizza.getPrice())

	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	fmt.Printf("Price of veggeMania with cheese is %d\n", pizzaWithCheese.getPrice())

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
