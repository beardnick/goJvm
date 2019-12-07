package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

type DeprecatedAttribute struct {
	name_index uint16
	cp         ConstantPool
}

func (self *DeprecatedAttribute) readInfo(reader *ClassReader) {

}

func (self *DeprecatedAttribute) String() string {
	return self.cp.getContantString(self.name_index)
}

type SyntheticAttribute struct {
	name_index uint16
	cp         ConstantPool
}

func (self *SyntheticAttribute) readInfo(reader *ClassReader) {
}

func (self *SyntheticAttribute) String() string {
	return self.cp.getContantString(self.name_index)
}

type SourceFileAtrribute struct {
	source_file_index uint16
	cp                ConstantPool
}

func (self *SourceFileAtrribute) readInfo(reader *ClassReader) {
	self.source_file_index = reader.ReadUint16()
}

func (self *SourceFileAtrribute) FileName() string {
	return self.cp.getContantString(self.source_file_index)
}

type CodeAttribute struct {
	name_index uint16
	max_stack  uint16
	max_locals uint16

	code            []uint8
	exception_table []ExceptionsAttribute
	attribute_info  []AttributeInfo

	cp ConstantPool
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.name_index = reader.ReadUint16()
	_ = reader.ReadUint32()
	self.max_stack = reader.ReadUint16()
	self.max_locals = reader.ReadUint16()
	code_length := reader.ReadUint32()
	self.code = make([]uint8, code_length)
	for i, _ := range self.code {
		self.code[i] = reader.ReadUint8()
	}
	self.exception_table = ReadExceptions(reader, self.cp)
	self.attribute_info = ReadAttributes(reader, self.cp)
}

// @todo: 完成剩余的属性解析 <25-11-19> //
type ExceptionsAttribute struct {
}

func ReadExceptions(reader *ClassReader, cp ConstantPool) []ExceptionsAttribute {
	return nil
}

type LineNumberTableAttribute struct {
}

type LocalVariableTableAttribute struct {
}

type ConstantValueAttribute struct {
}

type SourceFileAttribute struct {
}

type UnparsedAttribute struct {
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
}

func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
}
