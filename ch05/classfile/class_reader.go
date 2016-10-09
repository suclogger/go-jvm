package classfile

import (
	"encoding/binary"
)

type  ClassReader struct {
	// 简单包装了一下[]byte
	data []byte
}
// u(n) 代表n个字节
// 每个字节可以存放两个16进制数
// u2 代表 4 个16进制数
// 8 bit无符号整型, 读取 u1 ,对应 go 中的 uint8
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data=self.data[1:]
	return val
}

// 16 bit无符号整型, 读取 u2 ,对应 go 中的 uint16
func (self *ClassReader) readUint16() uint16 {
	// class 文件中连续字节的数据以大端方式存储
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// 32 bit无符号整型, 读取 u4 ,对应 go 中的 uint32
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

// 64 bit无符号整型, 读取 u8 ,对应 go 中的 uint64, 虚拟机规范并没有定义u8
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// class文件中相同类型的数据一般按表（table）的形式存储,表由表头和表项（item）构成,表头是u2或u4整数,表头是n,后面就紧跟这n个表项数据
// 这个方法用于读取u2表
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}

