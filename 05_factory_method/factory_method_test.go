package factorymethod

import "testing"

func TestMilan(t *testing.T) {
	_, err := getClub("Milan")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDortmund(t *testing.T) {
	_, err := getClub("Dortmund")
	if err != nil {
		t.Fatal(err.Error())
	}
}
