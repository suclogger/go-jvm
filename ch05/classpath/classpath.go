package classpath

import (
	"os"
	"path/filepath"
)
type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}
func Parse(jreOption,cpOption string) *Classpath{
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	// 如果未提供 cpOption,默认使用当前目录
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath =  newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	// 优先使用 jre 参数指定的 jre 路径
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 其次使用当前目录下的 jre
	if exists("./jre") {
		return "./jre"
	}
	// 最后使用系统变量 JAVA_HOME 定义的 jre 路径
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return  filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	// Stat returns a FileInfo describing the named file
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) ReadClass(className string)([]byte, Entry, error) {
	// 依次尝试从启动类路径,扩展类路径,用户指定路径加载主类
	// 传入的类名是不带 class 的
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)


}
func (self *Classpath) String() string {
	return self.userClasspath.String()
}

