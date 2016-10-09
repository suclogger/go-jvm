package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
	"fmt"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	// 逻辑与 entry_dir 一致
	absPath,err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}

}

func (self *ZipEntry) readClass(className string)([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _,f := range r.File{
		fmt.Println(f.Name)
		if f.Name == className {
			rc, err := f.Open()
			if err != nil{
				return nil, nil, err
			}
			defer rc.Close()
			data,err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data,self,nil
		}
	}
	// 通过 errors.New 来创建一个 error type
	// 字符串的拼接依然可以使用 +
	return nil,nil,errors.New("class not found:" + className)
}

func (self *ZipEntry) String() string {
	// 逻辑与 entry_dir 一致
	return self.absPath
}

