package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloServer)
	http.HandleFunc("/", IndexServer)
	http.HandleFunc("/upload", uploadServer)
	http.HandleFunc("/download", downloadServer)
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

//上传文件
func uploadServer(w http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	w.Write([]byte(fmt.Sprintf("<h1>文件上传 %s</h1>", name)))
}

//下载文件
func downloadServer(w http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	w.Write([]byte(fmt.Sprintf("<h1>文件下载 %s</h1>", name)))
}

//FileServerDemo 文件服务
func FileServerDemo() {
	err := http.ListenAndServe(":90", http.FileServer(http.Dir("E:\\images")))
	if err != nil {
		log.Printf("启动服务失败,err%s \n", err.Error())
	}
}
