package samplefactory

import "fmt"

type Shape interface {
	Draw() string
}

func DrawShape(shapeType int) Shape {
	if shapeType == 1 {
		return &Rectangle{}
	} else if shapeType == 2 {
		return &Square{}
	} else if shapeType == 3 {
		return &Circle{}
	}
	return nil
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
