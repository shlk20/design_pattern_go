package singleton

import "sync"

type Singleton interface {
	showMessage() string
}

type singleton struct{}

func (s *singleton) showMessage() string {
	return "Hello world!"
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
