package classfile

type ConstantPool []ConstantInfo

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {

}

func readConstantPool(reader *ClassReader) ConstantPool {
	length := reader.readUint16()
	cp := make([]ConstantInfo, length)
	for i := 1; i < length; i++ {
		// #imp cp给进去干什么，所有的数据都在ClassReader里面?
		// 虽然现在cp中没有数据，但是后期各个常量需要获取cp中的数据，所以要将cp传入当做上下文
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		// ConstantLongInfo, ConstantDoubleInfo 占两个单位大小
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}
