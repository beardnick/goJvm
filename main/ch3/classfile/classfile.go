package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic        uint32 // 标记Java文件的魔数
	minorVersion uint16 // 副版本号
	majorVersion uint16 // 主版本号
	// constantPoolCount uint16    // 常量池大小
	constantPool ConstantPool // 常量池
	accessFlag   uint16       // 访问标志 , public 之类的
	thisClass    uint16       // this指针
	superClass   uint16       // super指针
	// interfaceCount uint16       // 实现的接口数
	interfaces []uint16 // 接口列表
	// fieldCount        uint16    // 字段数
	fileds []*MemberInfo // 字段列表
	// methodCount uint16          // 方法数
	methods []*MemberInfo // 方法列表
	// attributeCount uint16       // 属性数
	attributes []AttributeInfo // 属性列表
}

//  将文件解析成ClassFile结构体
func (self *ClassFile) Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := ClassReader{classData}
	cf = &ClassFile{}
	cf.read(&cr)
	return
}

// 下列方法只要出错就panic，所以就不返回error了

func (self *ClassFile) read(reader *ClassReader) {
	self.magic = self.ReadAndCheckMagic(reader)
	self.minorVersion, self.majorVersion = self.ReadAndCheckVersion(reader)
	self.constantPool = self.ReadConstantPool(reader)
	self.accessFlag = reader.ReadUint16()
	self.thisClass = reader.ReadUint16()
	self.superClass = reader.ReadUint16()
	self.interfaces = reader.ReadUint16s()
	//self.fileds =
}

func (self *ClassFile) ReadAndCheckMagic(reader *ClassReader) (magic uint32) {
	magic = reader.ReadUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic !")
	}
	return
}

func (self *ClassFile) ReadAndCheckVersion(reader *ClassReader) (minor, major uint16) {
	minor = reader.ReadUint16()
	major = reader.ReadUint16()
	switch major {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if minor == 0 {
			return
		}
	}
	panic(fmt.Sprintln("java.lang.UnSupportedClassVersionError!  : ", " major : ", major, " minor: ", minor))
}

func (self *ClassFile) ReadConstantPool(reader *ClassReader) ConstantPool {
	return ReadConstantPool(reader)
}

//func (self *ClassFile) readMemberInfo(reader *ClassReader) []*MemberInfo {

//}

//func (self *ClassFile) readFields(reader *ClassReader) []*MemberInfo {

//}

//func (self *ClassFile) readMethods(reader *ClassReader) []*MemberInfo {

//}

//func (self *ClassFile) readAttributes(reader *ClassReader) []*AttributeInfo {

//}

func (self *ClassFile) PrintClassInfo() {
	fmt.Printf("majorVersion:%v\n", self.majorVersion)
	fmt.Printf("thisClass:%v\n", self.constantPool.GetIndex(self.thisClass))
	fmt.Printf("superClass:%v\n", self.constantPool.GetIndex(self.superClass))
	//fmt.Printf("interfaces:%v\n", self.constantPool.GetIndex(self.interfaces))
	fmt.Printf("interfaces:%v\n", self.interfaces)
}
