package classfile
type ConstantClassInfo struct {
	cp ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	// class 存放的也是索引
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) String() string {
	return self.cp.getUtf8(self.nameIndex)
}