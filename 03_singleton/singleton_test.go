package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()

	if s1 != s2 {
		t.Fatal("Not the same instace!!!")
	}

	res := s1.showMessage()
	if res != "Hello world!" {
		t.Fatal("Call singleton failed")
	}
}

func TestChannel(t *testing.T) {
	// channel := make(chan struct{})

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)
	instances := [10]Singleton{}

	for i := 0; i < 10; i++ {
		go func(idx int) {
			// <-channel
			instances[idx] = GetInstance()
			waitGroup.Done()
		}(i)
	}
	// close(channel)
	waitGroup.Wait()

	for i := 1; i < 10; i++ {
		if instances[i-1] != instances[i] {
			t.Fatal("Not the same instance")
		}
	}
}
