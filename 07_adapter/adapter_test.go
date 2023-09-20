package adapter

import "testing"

func TestAdapter(t *testing.T) {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	pc := &PC{}
	pcAdapter := &pcAdapter{
		pc: pc,
	}

	client.InsertLightningConnectorIntoComputer(pcAdapter)
}
