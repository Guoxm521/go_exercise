package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "docker部署镜像你好呀",
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
		for in.Scan() {
			cmdRe := string(in.Bytes())
			c.JSON(200, gin.H{
				"Message": cmdRe,
			})
		}

	})
	r.Run()
}
