package main

import (
	"fmt"
	"requests/utils"
)

func main() {

	//测试Get请求
	a1 := utils.Args{
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
		},
		Params: map[string]interface{}{
			"id": 12131,
		},
	}
	client1 := utils.Client{}
	result1 := client1.Get("http://cmdb02.zmops.cc/api/v1/applications/lite/", a1)
	fmt.Println(string(result1))

	//测试Post请求
	a2 := utils.Args{
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
		},
		Body: map[string]interface{}{
			"name":    "zhangsan",
			"age":     18,
			"address": "shanghai",
		},
	}

	client2 := utils.Client{}
	result2 := client2.Post("http://127.0.0.1:4523/m1/1911173-0-default/v1/student", a2)
	fmt.Println(string(result2))
}
