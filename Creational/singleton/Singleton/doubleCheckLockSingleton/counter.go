package singleton

import "sync"

type singleton struct {
	count int
}

var (
	instance *singleton
	mu       sync.Mutex
)

func GetInstance() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
