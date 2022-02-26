package util

import "sync"

type stack struct {
	stack []interface{}
	m     *sync.Mutex
}

func Stack() *stack {
	return &stack{
		make([]interface{}, 0),
		&sync.Mutex{},
	}
}

func (s *stack) Push(v interface{}) {
	s.m.Lock()
	s.stack = append(s.stack, v)
	s.m.Unlock()
}

func (s *stack) Pop() interface{} {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.stack) == 0 {
		return nil
	}

	next := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return next
}

func (s *stack) IsEmpty() bool {
	return len(s.stack) == 0
}
