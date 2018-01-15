package vm

type VM struct {
	Stack     *Stack
	Registers map[string]string
}

func NewVM() *VM {
	return &VM{
		Stack:     NewStack(),
		Registers: make(map[string]string),
	}
}
