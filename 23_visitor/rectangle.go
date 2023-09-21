package visitor

type Rectangle struct {
	l int
	b int
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForrectangle(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}
