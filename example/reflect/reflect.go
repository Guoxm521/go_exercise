package main

import (
	"fmt"
	"reflect"
)

func main() {
	userInfo := userInfo{
		Account:   "123213",
		Password:  "密码",
		AccountId: "AccountId",
	}
	//GenerateToken(userInfo)
	//getType(userInfo)
	getValue(userInfo)
}

type userInfo struct {
	Account   string `form:"account"`
	Password  string `form:"password"`
	AccountId string `form:"account_id"`
}

func GenerateToken(demo interface{}) {
	fmt.Println(demo)
	t := reflect.TypeOf(demo)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
	v := reflect.ValueOf(demo)
	k := v.Kind()
	fmt.Println(v)
	fmt.Println(k)
}

func getType(demo interface{}) {
	t := reflect.TypeOf(demo)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
		fmt.Println(field.Type)
	}

}

func getValue(demo interface{}) {
	value := reflect.ValueOf(demo)
	t := reflect.TypeOf(demo)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
		fmt.Println(value.Field(i))
	}
}
