package slstack

type AnyData interface{}

type Stack []AnyData

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(d AnyData) {
	*s = append(*s, d)
}

func (s *Stack) Pop() AnyData {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1   
		element := (*s)[index] 
		*s = (*s)[:index]      
		return element
	}
}

func (s *Stack) Peek() AnyData {
	if s.IsEmpty() {
		return ""
	}
	return (*s)[len(*s)-1]
}
