package builder

type IBuilder interface {
	setWindow()
	setDoor()
	setFloor()
	getHouse() House
}

type House struct {
	window string
	door   string
	floor  int
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	} else if builderType == "igloo" {
		return newIglooBuilder()
	}

	return nil
}

type NormalBuilder struct {
	window string
	door   string
	floor  int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setWindow() {
	n.window = "normal window"
}

func (n *NormalBuilder) setDoor() {
	n.door = "normal door"
}

func (n *NormalBuilder) setFloor() {
	n.floor = 4
}

func (n *NormalBuilder) getHouse() House {
	return House{
		window: n.window,
		door:   n.door,
		floor:  n.floor,
	}
}

type IglooBuilder struct {
	window string
	door   string
	floor  int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) setWindow() {
	i.window = "igloo window"
}

func (n *IglooBuilder) setDoor() {
	n.door = "igloo door"
}

func (n *IglooBuilder) setFloor() {
	n.floor = 1
}

func (n *IglooBuilder) getHouse() House {
	return House{
		window: n.window,
		door:   n.door,
		floor:  n.floor,
	}
}

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setWindow()
	d.builder.setDoor()
	d.builder.setFloor()
	return d.builder.getHouse()
}
