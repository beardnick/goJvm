package classfile

import (
	"math"
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

//cp_info {
//  u1 tag;
//  u1 info[];
//}

// tags
const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}

	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}

	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
	return nil
}

type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	self.val = int32(reader.ReadUint32())
}

type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	self.val = math.Float64frombits(reader.ReadUint64())
}

type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	self.val = int64(reader.ReadUint64())
}

type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	self.val = math.Float32frombits(reader.ReadUint32())
}

// utf8格式字符串
type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := reader.ReadUint16()
	self.str = string(reader.ReadBytes(uint32(length)))
}

// 指向ConstantUtf8Info的一个索引
type ConstantStringInfo struct {
	str_index uint16
	cp        ConstantPool
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.str_index = reader.ReadUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getContantString(self.str_index)
}

type ConstantClassInfo struct {
	name_index uint16
	cp         ConstantPool
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.name_index = reader.ReadUint16()
}

func (self *ConstantClassInfo) String() string {
	return self.cp.getContantString(self.name_index)
}

type ConstantNameAndTypeInfo struct {
	name_index uint16
	desc_index uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.name_index = reader.ReadUint16()
	self.desc_index = reader.ReadUint16()
}

type ConstantMemberrefInfo struct {
	cp              ConstantPool
	class_index     uint16
	name_type_index uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.class_index = reader.ReadUint16()
	self.name_type_index = reader.ReadUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	if classInfo, ok := self.cp.GetIndex(self.class_index).(*ConstantClassInfo); ok {
		return self.cp.getContantString(classInfo.name_index)
	}
	return ""
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (name, desc string) {
	if nameAndTypeInfo, ok := self.cp.GetIndex(self.name_type_index).(*ConstantNameAndTypeInfo); ok {
		return self.cp.getContantString(nameAndTypeInfo.name_index), self.cp.getContantString(nameAndTypeInfo.desc_index)
	}
	return "", ""
}

type ConstantMethodrefInfo struct {
	member ConstantMemberrefInfo
}

func (self *ConstantMethodrefInfo) readInfo(reader *ClassReader) {
	self.member.readInfo(reader)
}

type ConstantInterfaceMethodrefInfo struct {
	member ConstantMemberrefInfo
}

func (self *ConstantInterfaceMethodrefInfo) readInfo(reader *ClassReader) {
	self.member.readInfo(reader)
}

type ConstantFieldrefInfo struct {
	member ConstantMemberrefInfo
}

func (self *ConstantFieldrefInfo) readInfo(reader *ClassReader) {
	self.member.readInfo(reader)
}

type ConstantMethodTypeInfo struct {
	member ConstantMemberrefInfo
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.member.readInfo(reader)
}

type ConstantMethodHandleInfo struct{}
type ConstantInvokeDynamicInfo struct{}

func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {}
