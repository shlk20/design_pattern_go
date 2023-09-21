package chainofresponsibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "p"}
	reception.execute(patient)
}
