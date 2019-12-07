package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func (self *Stack) pop() *Frame {
	return nil
}

func (self *Stack) top() *Frame {
	return nil
}
