package rtda

import (
	"math"
)

type Frame struct {
	lower        *Frame
	local        LocalVar
	operandStack *OperandStack
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		local:        NewLocalVar(maxLocals),
		operandStack: NewOperandStack(maxStack),
	}
}

// 数组元素必须可以容纳一个int
// 两个连续的必须可以容纳下一个long或者double
// 所以总的来说最小就是32位的一个数据类型
type LocalVar []Slot

type Object struct {
}

type Slot struct {
	num int32
	ref *Object
}

func NewLocalVar(maxLocals uint) LocalVar {
	if maxLocals <= 0 {
		return nil
	}
	return make([]Slot, maxLocals)
}

func (self LocalVar) SetInt(index uint, value int32) {
	self[index].num = value
}

// long转换为两个int
func (self LocalVar) SetLong(index uint, value int64) {
	self[index].num = int32(value)
	self[index+1].num = int32(value >> 32)
}

// float转为int来存储
func (self LocalVar) SetFloat(index uint, value float32) {
	bits := math.Float32bits(value)
	self[index].num = int32(bits)
}

func (self LocalVar) SetDouble(index uint, value float64) {
	bits := math.Float64bits(value)
	self.SetLong(index, int64(bits))
}

func (self LocalVar) SetRef(index uint, value *Object) {
	self[index].ref = value
}

func (self LocalVar) GetInt(index uint) int32 {
	return self[index].num
}

func (self LocalVar) GetLong(index uint) int64 {
	//low := int64(self[index].num)
	// 这样是错的，这样就转成0了
	//high := int64(self[index].num << 32)
	// 这样是错的，符号不对
	// return low + high
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (self LocalVar) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

func (self LocalVar) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self LocalVar) GetRef(index uint) *Object {
	return self[index].ref
}

type OperandStack struct {
	size  uint
	slots []Slot
}

func NewOperandStack(maxStack uint) *OperandStack {
	if maxStack <= 0 {
		return nil
	}
	return &OperandStack{
		slots: make([]Slot, maxStack),
	}
}

func (self *OperandStack) PushInt(value int32) {
	self.slots[self.size].num = value
	self.size++
}

// long转换为两个int
func (self *OperandStack) PushLong(value int64) {
	self.slots[self.size].num = int32(value)
	self.size++
	self.slots[self.size].num = int32(value >> 32)
	self.size++
}

// float转为int来存储
func (self *OperandStack) PushFloat(value float32) {
	bits := math.Float32bits(value)
	self.slots[self.size].num = int32(bits)
	self.size++
}

func (self *OperandStack) PushDouble(value float64) {
	bits := math.Float64bits(value)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PushRef(value *Object) {
	self.slots[self.size].ref = value
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

func (self *OperandStack) PopLong() int64 {
	self.size--
	high := uint32(self.slots[self.size].num)
	self.size--
	low := uint32(self.slots[self.size].num)
	return int64(high)<<32 | int64(low)
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return ref
}
