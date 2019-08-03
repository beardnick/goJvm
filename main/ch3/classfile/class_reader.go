package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// 读取u2表, 第一个u2代表数量，后续的代表具体的数据
func (self *ClassReader) readUint16s() []uint16 {
	count := self.readUint16()
	table := make([]uint16, count)
	for i := range table {
		table[i] = self.readUint16()
	}
	return table
}

func (self *ClassReader) readBytes(length uint32) []byte {
	s := self.data[:length]
	self.data = self.data[length:]
	return s
}
