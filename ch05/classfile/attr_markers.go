package classfile
type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {

}

// @Deprecated 和 @Synthetic  都是指示类属性, 没有数据
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
