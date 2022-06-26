package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	//OpenFile()
	//ReadAllDir("./demo.txt")
	//ReadFile("./demo.txt")
	//WriteFile("./demo.txt", "测试数据")
	//AppendToFile("./demo.txt", "dasdasdasdsadasdas")
	MkAllDir("./demo/demo/haha/heihei")
}

//OpenFile 打开文件
func OpenFile() {
	file, err := os.Open("./demo.txt")
	Check(err)
	file.Close()
}

func ReadFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	Check(err)
	fmt.Println(string(data))
}

//ReadAllDir 读取文件夹
func ReadAllDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

//WriteFile 这种会覆盖掉原先内容
func WriteFile(fileName, data string) {
	err := ioutil.WriteFile(fileName, []byte(data), os.ModePerm)
	Check(err)
}

//AppendToFile 追加文件内容
func AppendToFile(fileName, data string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer func() {
		_ = file.Close()
	}()
	Check(err)
	_, _ = file.Write([]byte(data))
}

//CreateFile 创建文件并返回文件指针
func CreateFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	defer file.Close()
	Check(err)
	return file
}

//MkOneDir 创建单个文件夹
func MkOneDir(dir string) {
	err := os.Mkdir(dir, os.ModePerm)
	Check(err)
	_ = os.RemoveAll(dir)
}

//MkAllDir 创建多层文件夹
func MkAllDir(dirs string) {
	// 如果不存在，才创建
	if !IsExist(dirs) {
		err := os.MkdirAll(dirs, os.ModePerm)
		Check(err)
		_ = os.RemoveAll(strings.Split(dirs, "/")[0])
	}
}
func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
