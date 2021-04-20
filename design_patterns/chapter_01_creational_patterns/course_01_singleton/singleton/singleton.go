package singleton

// Singleton interface
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

// GetInstance returns a new instance of Singleton
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// AddOne increments counter by 1
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
