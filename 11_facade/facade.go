package facade

import "fmt"

type Facade interface {
	Call() string
}

type facadeImpl struct {
	one NumberOne
	two NumberTwo
}

func CallFacade() Facade {
	return &facadeImpl{
		one: CallNumberOne(),
		two: CallNumberTwo(),
	}
}

func (f *facadeImpl) Call() string {
	a := f.one.TestNumberOne()
	b := f.two.TestNumberTwo()
	return fmt.Sprintf("%s %s", a, b)
}

type NumberOne interface {
	TestNumberOne() string
}

type numberOneImpl struct{}

func (f *numberOneImpl) TestNumberOne() string {
	return "Call the number one succesfully."
}

func CallNumberOne() NumberOne {
	return &numberOneImpl{}
}

type NumberTwo interface {
	TestNumberTwo() string
}

type numberTwoImpl struct{}

func (f *numberTwoImpl) TestNumberTwo() string {
	return "Call the number two succesfully."
}

func CallNumberTwo() NumberTwo {
	return &numberTwoImpl{}
}
