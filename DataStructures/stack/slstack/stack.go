package slstack

type DataType interface{}

type Stack []DataType

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(d DataType) {
	*s = append(*s, d)
}

func (s *Stack) Pop() DataType {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1   
		element := (*s)[index] 
		*s = (*s)[:index]      
		return element
	}
}

func (s *Stack) Peek() DataType {
	if s.IsEmpty() {
		return ""
	}
	return (*s)[len(*s)-1]
}
