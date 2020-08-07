package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func GetExt(name string) string {
	return path.Ext(name)
}

// GenerateImageName 生成文件名
func GenerateImageName(name string) string {
	ext := path.Ext(name)
	now := time.Now()
	dir := "./assets/upload/" + now.Format("200601/02/")
	_ = IsNotExistMkDir(dir)
	return fmt.Sprintf("%s%s%s",
		dir,
		strconv.FormatInt(now.Unix(), 10),
		ext)
}

func IsNotExistMkDir(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
	}
	return
}

func GetImagePath() string {
	return "upload/images/"
}

func GetImageFullPath() string {
	return "assets/" + GetImagePath()
}

func CheckImageExt(extName string) bool {
	exts := [...]string{"jpg", "jpeg", "png", "gif"}
	for _, item := range exts {
		if strings.ToUpper(item) == strings.ToUpper(extName) {
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return false
	}

	return len(content) <= 5000
}
