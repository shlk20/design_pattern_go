package builder

import "testing"

func TestNormalBuilder(t *testing.T) {
	normalBuilder := getBuilder("normal")
	director := newDirector(normalBuilder)
	house := director.buildHouse()

	if house.door != "normal door" {
		t.Fatal("It is not the normal house's door.")
	}
	if house.window != "normal window" {
		t.Fatal("It is not the normal huuse's window.")
	}
	if house.floor != 4 {
		t.Fatal("It is not the normal house's floor.")
	}
}

func TestIglooBuilder(t *testing.T) {
	iglooBuilder := getBuilder("igloo")
	director := newDirector(iglooBuilder)
	house := director.buildHouse()

	if house.door != "igloo door" {
		t.Fatal("It is not the igloo's door.")
	}
	if house.window != "igloo window" {
		t.Fatal("It is not the igloo's window.")
	}
	if house.floor != 1 {
		t.Fatal("It is not the igloo's floor.")
	}
}
