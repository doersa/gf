package gfile

import (
    "os"
    "path/filepath"
)

var File = gFile {
    Separator : string(filepath.Separator),
}

type gFile struct {
    Separator string
}

// 判断所给路径文件/文件夹是否存在
func (f gFile) Exists(path string) bool {
    _, err := os.Stat(path)
    return err == nil || os.IsExist(err)
}

// 判断所给路径是否为文件夹
func (f gFile) IsDir(path string) bool {
    s, err := os.Stat(path)
    if (err == nil) {
        return s.IsDir()
    }
    return false
}