package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"exercise/utils"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	GetFileMd5("./demo.txt")
}

func GetFileMd5(path string) (fileMd5 string, err error) {
	f, err := os.Open(path)
	utils.Check(err)
	defer func() {
		f.Close()
	}()
	md5hash := sha1.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		fmt.Println("fileMd51:", fileMd5)
		return fileMd5, err
	}
	fileMd5 = hex.EncodeToString(md5hash.Sum(nil))
	fmt.Println("fileMd52:", fileMd5)
	return fileMd5, nil
}

func ReadAll(path string) (fileMD5 string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return fileMD5, err
	}
	defer func() {
		_ = f.Close()
	}()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		return fileMD5, err
	}
	hash := sha1.New()
	hash.Write(body)
	fileMD5 = hex.EncodeToString(hash.Sum(nil))
	return fileMD5, nil
}

func ReadBuf(path string) (fileMD5 string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return fileMD5, err
	}
	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 1024)
	reader := bufio.NewReader(f)
	md5hash := sha1.New()
	for {
		n, err := reader.Read(buf)
		if err != nil { // 遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				goto stop
			}
			return fileMD5, err
		}
		md5hash.Write(buf[:n])
	}
stop:
	fileMD5 = hex.EncodeToString(md5hash.Sum(nil))
	return fileMD5, nil
}
