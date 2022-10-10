package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	demo1()
}

func demo1() {
	type Person struct {
		Name   string
		Age    int64
		Weight float64
	}
	p1 := Person{
		Name:   "七米",
		Age:    18,
		Weight: 71.5,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

	//------------------------------------------------------------------

	var p2 interface{}
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	//fmt.Printf("p2:%#v\n", p2.(map[string]interface{})["Name1"].(string))
	v, ok := p2.(map[string]interface{})["Name"].(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println(ok)
	}
	TypeJudge(v)
}

func TypeJudge(item ...interface{}) {
	for index, value := range item {
		switch value.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index, value)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index, value)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", index, value)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", index, value)
		case int:
			fmt.Printf("第%v个参数是int类型，值是%v\n", index, value)
		case int8:
			fmt.Printf("第%v个参数是int8类型，值是%v\n", index, value)
		case int16:
			fmt.Printf("第%v个参数是int16类型，值是%v\n", index, value)
		case int32:
			fmt.Printf("第%v个参数是int32类型，值是%v\n", index, value)
		case int64:
			fmt.Printf("第%v个参数是int64类型，值是%v\n", index, value)
		case uint:
			fmt.Printf("第%v个参数是uint类型，值是%v\n", index, value)
		case uint8:
			fmt.Printf("第%v个参数是uint8类型，值是%v\n", index, value)
		case uint16:
			fmt.Printf("第%v个参数是uint16类型，值是%v\n", index, value)
		case uint32:
			fmt.Printf("第%v个参数是uint32类型，值是%v\n", index, value)
		case uint64:
			fmt.Printf("第%v个参数是uint64类型，值是%v\n", index, value)
		default:
			fmt.Printf("没有找到相匹配的类型\n", value)
		}

	}
}
