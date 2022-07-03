package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	PostJsonData()
}

type User struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

// 提交Form表单数据
func postFormDemo() {
	params := url.Values{}
	params.Add("user_name", "hangman")
	params.Add("pass_word", "123456")
	response, err := http.PostForm("http://www.baidu.com", params)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}

//提交JSON数据
func PostJsonData() {
	user := &User{
		UserName: "zhansgan",
		PassWord: "123456",
	}
	bs, _ := json.Marshal(user)
	data := bytes.NewReader(bs)
	resp, err := http.Post("http://www.baidu.com", "application/json", data)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
