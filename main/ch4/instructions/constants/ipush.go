package constants

import (
	"goJvm/main/ch4/instructions/base"
	"goJvm/main/ch4/jvm/rtda"
)

type BIPUSH struct {
	val int8 // push byte
}

type SIPUSH struct {
	val int16 // push short
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {

}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
}
