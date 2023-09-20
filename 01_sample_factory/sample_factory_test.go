package samplefactory

import "testing"

func TestRectangle(t *testing.T) {
	shape := DrawShape(1)
	s := shape.Draw()
	if s != "Inside Rectangle::Draw() method." {
		t.Fatal("testing draw rectangle failed")
	}
}

func TestSqaure(t *testing.T) {
	shape := DrawShape(2)
	s := shape.Draw()
	if s != "Inside Square::Draw() method." {
		t.Fatal("testing draw square failed")
	}
}

func TestCircle(t *testing.T) {
	shape := DrawShape(3)
	s := shape.Draw()
	if s != "Inside Circle::Draw() method." {
		t.Fatal("testing draw circle failed")
	}
}
