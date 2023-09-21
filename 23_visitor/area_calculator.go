package visitor

import (
	"fmt"
	"math"
)

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
	a.area = s.side * s.side
	fmt.Printf("Result is %d\n", a.area)
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
	a.area = int(float32(s.radius) * float32(s.radius) * math.Pi)
	fmt.Printf("Result is %d\n", a.area)
}
func (a *AreaCalculator) visitForrectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
	a.area = s.l * s.b
	fmt.Printf("Result is %d\n", a.area)
}
