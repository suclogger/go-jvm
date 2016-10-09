package classpath

import (
	"os"
	"path/filepath"
	"strings"
)
func newWildcardEntry(path string) CompositeEntry {
	// remove ending *
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//* 通配符不能递归子目录下的 jar 文件
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") ||  strings.HasSuffix(path, ".JAR") {
			jarEntry  := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	// filepath.WalK 在包括根目录的每个节点执行 walkFn 方法,根据字母顺序遍历
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}

