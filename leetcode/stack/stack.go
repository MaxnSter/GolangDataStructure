package stack

// a Basic Stack, not goroutine safe
type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{data: make([]interface{}, 0)}
}

func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Empty() bool {
	return s.Size() <= 0
}

func (s *Stack) Pop() {
	if s.Empty() {
		return
	}
	s.data = s.data[:s.Size()-1]
}

func (s *Stack) Top() interface{} {
	if s.Empty() {
		panic("Top a empty stack")
	}

	return s.data[s.Size()-1]
}

func (s *Stack) Reverse() {
	if s.Empty() {
		return
	}

	b, e := 0, s.Size()-1
	for b < e {
		s.data[b], s.data[e] = s.data[e], s.data[b]
		b++
		e--
	}
}

//dump stack to a slice, not clear stack
//len(*des) must be the same as stack Size
func (s *Stack) Dump(des *[]interface{}) {
	if s.Empty() {
		return
	}

	copy(*des, s.data)
}
