package main

import (
	"encoding/json"
	"exercise/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)

var dirMap map[string]float64

func main() {
	dirPath := "E:\\images"
	dirMap = make(map[string]float64)
	//method1(dirPath, dirMap)
	//method2(dirPath, dirMap)
	method3(dirPath, dirMap)
}

//method1 方法一
func method1(dirPath string, dirMap map[string]float64) {
	t := time.Now()
	scanDir1(dirPath, dirMap)
	var fileCount int   //文件数量
	var dirSize float64 //文件夹的大小
	for _, v := range dirMap {
		fileCount++
		dirSize += v
	}
	fmt.Println("花费的时间" + time.Since(t).String())
	fmt.Printf("文件的数量%d\n", fileCount)
	fmt.Printf("文件夹的大小%f", dirSize)
	data, _ := json.MarshalIndent(dirMap, "", " ")
	utils.WriteFile("./demo.txt", string(data))
}

// scanDir1 递归计算目录下所有文件
func scanDir1(path string, dirMap map[string]float64) {
	dirArr, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, info := range dirArr {
		if info.IsDir() {
			scanDir1(filepath.Join(path, info.Name()), dirMap)
		} else {
			dirMap[filepath.Join(path, info.Name())] = float64(info.Size()) / 1024
		}
	}
}

var wait sync.WaitGroup

//使用携程
func method2(dirPath string, dirMap map[string]float64) {
	fileSize := make(chan float64)
	t := time.Now()
	wait.Add(1)
	go scanDir2(dirPath, fileSize)
	go func() {
		defer close(fileSize)
		wait.Wait()
	}()
	var fileCount int   //文件数量
	var dirSize float64 //文件夹的大小
	for v := range fileSize {
		fileCount++
		dirSize += v
	}
	fmt.Println("花费的时间" + time.Since(t).String())
	fmt.Printf("文件的数量%d\n", fileCount)
	fmt.Printf("文件夹的大小%f", dirSize)
}

func scanDir2(path string, fileSize chan<- float64) {
	defer wait.Done()
	dirArr, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, info := range dirArr {
		if info.IsDir() {
			wait.Add(1)
			go scanDir2(filepath.Join(path, info.Name()), fileSize)
		} else {
			fileSize <- float64(info.Size()) / 1024
		}
	}
}

var syncM sync.Map

//使用sync map
func method3(dirPath string, dirMap map[string]float64) {
	dirMap = make(map[string]float64)
	t := time.Now()
	wait.Add(1)
	go scanDir3(dirPath, &syncM)
	wait.Wait()
	var fileCount int   //文件数量
	var dirSize float64 //文件夹的大小

	syncM.Range(func(key, value interface{}) bool {
		fileCount++
		v := value.(float64)
		dirSize += v
		return true
	})

	fmt.Println("花费的时间" + time.Since(t).String())
	fmt.Printf("文件的数量%d\n", fileCount)
	fmt.Printf("文件夹的大小%f", dirSize)
}

func scanDir3(path string, syncM *sync.Map) {
	defer wait.Done()
	dirAry, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, e := range dirAry {
		if e.IsDir() {
			wait.Add(1)
			go scanDir3(filepath.Join(path, e.Name()), syncM)
		} else {
			syncM.Store(filepath.Join(path, e.Name()), float64(e.Size())/1024)
		}
	}
}
