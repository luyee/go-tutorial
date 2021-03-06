package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

// 首先查找tag含有Foo的可导出的struct字段(首字母大写)
// 其次查找字段名是Foo的导出字段
// 最后查找类似FOO或者FoO这样的除了首字母之外其他大小写不敏感的导出字段

// JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组。Go类型和JSON类型的对应关系如下：

// bool 代表 JSON booleans,
// float64 代表 JSON numbers,
// string 代表 JSON strings,
// nil 代表 JSON null.

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	var f interface{}
	err := json.Unmarshal(b, &f)
	fmt.Printf("%#v", f)
	if err != nil {
		fmt.Println(err)
	}

	m := f.(map[string]interface{}) //断言解析未知结构的数据

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
