package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	vendingMachine := NewVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		t.Fatal(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println("------------")

	err = vendingMachine.addItem(2)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println("------------")

	// err = vendingMachine.requestItem()
	// if err != nil {
	// 	t.Fatal(err.Error())
	// }

	err = vendingMachine.insertMoney(10)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		t.Fatal(err.Error())
	}
}
