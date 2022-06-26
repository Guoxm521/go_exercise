package utils

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//WriteFile 这种会覆盖掉原先内容
func WriteFile(fileName, data string) {
	err := ioutil.WriteFile(fileName, []byte(data), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
