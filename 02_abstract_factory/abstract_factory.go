package abstractfactory

import "fmt"

type Shape interface {
	Draw() string
}

type Rectangle struct{}

func (r *Rectangle) Draw() string {
	return fmt.Sprintf("Inside Rectangle::Draw() method.")
}

type Square struct{}

func (s *Square) Draw() string {
	return fmt.Sprintf("Inside Square::Draw() method.")
}

type Circle struct{}

func (c *Circle) Draw() string {
	return fmt.Sprint("Inside Circle::Draw() method.")
}

type Color interface {
	Fill() string
}

type Red struct{}

func (r *Red) Fill() string {
	return fmt.Sprintf("Inside Red::Fill() method.")
}

type Green struct{}

func (g *Green) Fill() string {
	return fmt.Sprintf("Inside Green::Fill() method.")
}

type Blue struct{}

func (b *Blue) Fill() string {
	return fmt.Sprint("Inside Blue::Fill() method.")
}

type AbstractFactory interface {
	CreateShape(shapeType int) Shape
	CreateColor(colorType int) Color
}

type ShapFactory struct{}

func (sf *ShapFactory) CreateShape(shapeType int) Shape {
	if shapeType == 1 {
		return &Rectangle{}
	} else if shapeType == 2 {
		return &Square{}
	} else if shapeType == 3 {
		return &Circle{}
	}
	return nil
}

func (sf *ShapFactory) CreateColor(colorType int) Color {
	if colorType == 1 {
		return &Red{}
	} else if colorType == 2 {
		return &Green{}
	} else if colorType == 3 {
		return &Blue{}
	}
	return nil
}
