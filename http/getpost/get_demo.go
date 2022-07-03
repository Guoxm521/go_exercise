package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	getDemo()
}

func getDemo() {
	response, err := http.Get("http://www.baidu.com/")
	if err != nil {
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}
