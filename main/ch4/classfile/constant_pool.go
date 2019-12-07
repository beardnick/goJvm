package classfile

type ConstantPool []ConstantInfo

func ReadConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.ReadUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func (self ConstantPool) GetIndexs(indexs []uint16) []ConstantInfo {
	infos := make([]ConstantInfo, len(indexs))
	for i, v := range indexs {
		infos[i] = self.GetIndex(v)
	}
	return infos
}

func (self ConstantPool) GetIndex(index uint16) ConstantInfo {
	if int(index) > len(self) {
		panic("no such constant")
	}
	return self[index]
}

func ReadConstantPool(reader *ClassReader) ConstantPool {
	length := int(reader.ReadUint16())
	cp := make([]ConstantInfo, length)
	for i := 1; i < length; i++ {
		// #imp cp给进去干什么，所有的数据都在ClassReader里面?
		// 虽然现在cp中没有数据，但是后期各个常量需要获取cp中的数据，所以要将cp传入当做上下文
		cp[i] = ReadConstantInfo(reader, cp)
		switch cp[i].(type) {
		// ConstantLongInfo, ConstantDoubleInfo 占两个单位大小
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (self ConstantPool) getContantStrings(indexs []uint16) []string {
	results := make([]string, len(indexs))
	for i, k := range indexs {
		results[i] = self.getContantString(k)
	}
	return results
}

func (self ConstantPool) getContantString(index uint16) string {
	if utf8Info, ok := self.GetIndex(index).(*ConstantUtf8Info); ok {
		return utf8Info.str
	}
	return ""
}
