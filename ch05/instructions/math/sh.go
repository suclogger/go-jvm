package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)



type ISHL struct{ base.NoOperandsInstruction }     // int左位移
type ISHR struct{ base.NoOperandsInstruction }     // int算术右位移
type IUSHR struct{ base.NoOperandsInstruction }    // int逻辑右位移
type LSHL struct{ base.NoOperandsInstruction }    // long左位移
type LSHR struct{ base.NoOperandsInstruction }    // long算术右位移
type LUSHR struct{ base.NoOperandsInstruction }    // long逻辑右位移

func (self *ISHL) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// int 只有32位,只需要取前5个比特
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

func (self *ISHR) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := v2 & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}


func (self *LSHR) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}







