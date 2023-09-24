package singleton

import "sync"

type singleton struct {
	count int
}

var (
	once     sync.Once
	instance *singleton
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})

	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
