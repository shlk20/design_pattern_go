package facade

import "testing"

func TestFacade(t *testing.T) {
	facade := CallFacade()
	res := facade.Call()
	if res != "Call the number one succesfully. Call the number two succesfully." {
		t.Fatal("Facade testing error, ", res)
	}
}
