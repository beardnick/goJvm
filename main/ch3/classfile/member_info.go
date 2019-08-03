package classfile

type MemberInfo struct {
	cp ConstantPool
	//  accessFlag, nameIndex, descriptionIndex 都是常量池索引
	accessFlags     uint16          // 访问标志, public之类的
	nameIndex       uint16          // 字段或方法名(全限定名，带包的那种，如java.lang.Object)
	descriptorIndex uint16          // 字段或者方法的描述符(描述字段或方法的数据类型，参数，参数列表，返回值等)
	attributes      []AttributeInfo // 属性表
}

func (self *MemberInfo) readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	length := reader.readUint16()
	members := make([]*MemberInfo, length)
	for i := range members {
		members[i] = self.readMember(reader, cp)
	}
	return members
}
