package calculator

type Stack struct {
	items []Token
}

func NewStack() *Stack {
	return &Stack{
		items: []Token{},
	}
}

func (s *Stack) Push(item Token) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() Token {
	if len(s.items) == 0 {
		return Token{}
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() Token {
	if len(s.items) == 0 {
		return Token{}
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = []Token{}
}
