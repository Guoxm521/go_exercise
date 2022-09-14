package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	r.GET("/go/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "docker部署镜像你好呀 go",
		})
	})
	r.GET("/go1", func(c *gin.Context) {
		cmd := exec.Command("D:/BianCheng/Git/bin/bash", "./01_shell.sh")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error=>", err.Error())
		}
		cmd.Start()
		in := bufio.NewScanner(stdout)
		cmdRe := ""
		for in.Scan() {
			cmdRe = string(in.Bytes())
		}
		c.JSON(200, gin.H{
			"Message": cmdRe,
		})
	})
	r.GET("/go2", func(c *gin.Context) {
		w := c.Writer
		header := w.Header()
		//在响应头添加分块传输的头字段Transfer-Encoding: chunked
		header.Set("Transfer-Encoding", "chunked")
		header.Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		cmd := exec.Command("D:/BianCheng/Git/bin/bash", "./01_shell.sh")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error=>", err.Error())
		}
		cmd.Start()
		in := bufio.NewScanner(stdout)
		for in.Scan() {
			w.WriteString(in.Text() + "\n")
			w.(http.Flusher).Flush()
			fmt.Println(in.Text())
			//c.Stream(func(w io.Writer) bool {
			//	w.Write(in.Bytes())
			//	w.Write([]byte{'\n'})
			//	d.Flush()
			//	return false
			//})

		}
		w.(http.Flusher).Flush()
		cmd.Wait()

	})
	r.Run()
}
