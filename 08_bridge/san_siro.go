package bridge

import "fmt"

type SanSiro struct{}

func (s *SanSiro) PrepareMatch() {
	fmt.Println("Preparing match in San Siro")
}
