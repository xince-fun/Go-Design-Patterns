package singleton

import "sync"

// Singleton 是单例模式接口，导出的
// 通过该接口可以避免 GetInstance 返回一个包私有类型的指针
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var (
	once     sync.Once
	instance *singleton
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
