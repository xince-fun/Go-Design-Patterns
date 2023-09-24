package singleton

type singleton struct {
	count int
}

var instance *singleton = &singleton{}

func GetInstance() *singleton {
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
