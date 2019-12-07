package constants

import (
	"goJvm/main/ch4/instructions/base"
	"goJvm/main/ch4/jvm/rtda"
)

type ACONST_NULL struct{ base.NoOperandsInstruction }
type DCONST_0 struct{ base.NoOperandsInstruction }
type DCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_0 struct{ base.NoOperandsInstruction }
type FCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_2 struct{ base.NoOperandsInstruction }
type ICONST_M1 struct{ base.NoOperandsInstruction }
type ICONST_0 struct{ base.NoOperandsInstruction }
type ICONST_1 struct{ base.NoOperandsInstruction }
type ICONST_2 struct{ base.NoOperandsInstruction }
type ICONST_3 struct{ base.NoOperandsInstruction }
type ICONST_4 struct{ base.NoOperandsInstruction }
type ICONST_5 struct{ base.NoOperandsInstruction }
type LCONST_0 struct{ base.NoOperandsInstruction }
type LCONST_1 struct{ base.NoOperandsInstruction }

// 将null指针推进栈中
func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

func (self *DCONST_1) Execute(frame *rtda.Frame) {}

func (self *FCONST_0) Execute(frame *rtda.Frame) {}

func (self *FCONST_1) Execute(frame *rtda.Frame) {}

func (self *FCONST_2) Execute(frame *rtda.Frame) {}

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

func (self *ICONST_0) Execute(frame *rtda.Frame) {}

func (self *ICONST_1) Execute(frame *rtda.Frame) {}

func (self *ICONST_2) Execute(frame *rtda.Frame) {}

func (self *ICONST_3) Execute(frame *rtda.Frame) {}

func (self *ICONST_4) Execute(frame *rtda.Frame) {}

func (self *ICONST_5) Execute(frame *rtda.Frame) {}

func (self *LCONST_0) Execute(frame *rtda.Frame) {}

func (self *LCONST_1) Execute(frame *rtda.Frame) {}
