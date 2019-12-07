package base

import "goJvm/main/ch4/jvm/rtda"

// 指令的格式
// <操作码><操作数>
// <opcode><operand>
type Instruction interface {
	// 从指令中获取operand
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

func (self *NoOperandsInstruction) Execute(frame *rtda.Frame) {

}
