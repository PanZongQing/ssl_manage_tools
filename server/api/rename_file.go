package api

import (
	"fmt"
	"os"
)

func RenameFile(filename1, filename2 string) {
	oldpath := "./uploaddir/" + filename1
	newpath := "./uploaddir/" + filename2
	err := os.Rename(oldpath, newpath)
	if err != nil {
		panic(err)
	}
	fmt.Println("修改文件名成功")
}
