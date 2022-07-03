package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/hello", helloServer)
	http.HandleFunc("/", IndexServer)
	http.HandleFunc("/upload", uploadServer)
	http.HandleFunc("/download/", downloadServer)
	log.Printf("About to listen on 8088. Go to http://127.0.0.1:8088/")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Printf("启动服务失败,err%s\n", err.Error())
	}
	//FileServer()
}

func helloServer(w http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	w.Write([]byte(fmt.Sprintf("<h1>hello %s</h1>", name)))
}

func IndexServer(w http.ResponseWriter, request *http.Request) {
	bytes, err := ioutil.ReadFile("./index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("没有找到文件")
		return
	}
	io.WriteString(w, string(bytes))
}

const (
	maxUploadSize = 2 * 1024 * 2014
	uploadPath    = "demogo"
)

//上传文件
func uploadServer(w http.ResponseWriter, request *http.Request) {
	if err := request.ParseMultipartForm(maxUploadSize); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("文件大小超过限制")
		return
	}
	file, _, _err := request.FormFile("file")
	if _err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("无效的file")
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("读取文件失败")
		return
	}
	fileType := http.DetectContentType(fileBytes)
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("获取文件扩展名失败")
		return
	}
	_time := time.Now().Unix()
	_fileName := strconv.FormatInt(_time, 16)
	newPath := filepath.Join(uploadPath, _fileName+fileEndings[len(fileEndings)-1])
	_, err = os.Stat(uploadPath)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(uploadPath, 0666)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println("文件夹不存在，创建失败")
				return
			}
		}
	}
	newFile, _ := os.Create(newPath)
	newFile.Write(fileBytes)
	w.Write([]byte("上传成功"))
}

//下载文件
func downloadServer(w http.ResponseWriter, request *http.Request) {
	path := strings.Split(request.URL.Path, "/download/")[1]
	newPath := filepath.Join(uploadPath, path)
	file, err := os.Open("./" + newPath)
	fmt.Println("newPath------>", newPath)
	fmt.Println(file)
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(buff)
}

//FileServerDemo 文件服务
func FileServerDemo() {
	err := http.ListenAndServe(":90", http.FileServer(http.Dir("E:\\images")))
	if err != nil {
		log.Printf("启动服务失败,err%s \n", err.Error())
	}
}
