package classfile

func ReadAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	count := reader.ReadUint16()
	attributes := make([]AttributeInfo, count)
	for i := range attributes {
		attributes[i] = ReadAttribute(reader, cp)
	}
	return attributes
}

func ReadAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	nameIndex := reader.ReadUint16()
	attrName := cp.getContantString(nameIndex)
	attrLen := reader.ReadUint32()
	return NewAttributeInfo(attrName, attrLen, cp)
}

func NewAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		//return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		//return &UnparsedAttribute{attrName, attrLen, nil}
	}
	return nil
}
