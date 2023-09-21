package observer

import "testing"

func TestObserver(t *testing.T) {
	shirtItem := NewItem("Nike Shirt")

	observerOne := &Customer{
		id: "abc@gmail.com",
	}
	observerTwo := &Customer{
		id: "efg@gmail.com",
	}
	shirtItem.register(observerOne)
	shirtItem.register(observerTwo)

	shirtItem.updateAvailability()
}
