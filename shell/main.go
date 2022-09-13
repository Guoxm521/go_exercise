package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	//fmt.Println("开始")
	//cmd := exec.Command("D:/BianCheng/Git/bin/bash", "./01_shell.sh")
	//bytes, err := cmd.Output()
	//if err != nil {
	//	fmt.Println("cmd.Output:", err)
	//	return
	//}
	//fmt.Println(string(bytes))
	demo2()
}

func demo() {
	cmd := exec.Command("D:/BianCheng/Git/bin/bash", "./01_shell.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	var index int
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
		index++
	}

	cmd.Wait()
}

func demo2() {
	cmd := exec.Command("D:/BianCheng/Git/bin/bash", "./01_shell.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
	}
	cmd.Start()
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		cmdRe := string(in.Bytes())
		fmt.Println(cmdRe)
	}

}
