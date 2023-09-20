package abstractfactory

import "testing"

func TestRedRectangleFactory(t *testing.T) {
	var factory AbstractFactory

	factory = &ShapFactory{}
	s1 := factory.CreateColor(1).Fill()
	s2 := factory.CreateShape(1).Draw()
	if s1 != "Inside Red::Fill() method." {
		t.Fatal("Creating red color failed")
	}
	if s2 != "Inside Rectangle::Draw() method." {
		t.Fatal("Creating rectangle shap failed")
	}
}

func TestBlueCircleFactory(t *testing.T) {
	var factory AbstractFactory

	factory = &ShapFactory{}
	s1 := factory.CreateColor(3).Fill()
	s2 := factory.CreateShape(3).Draw()
	if s1 != "Inside Blue::Fill() method." {
		t.Fatal("Creating blue color failed")
	}
	if s2 != "Inside Circle::Draw() method." {
		t.Fatal("Creating circle shap failed")
	}
}
