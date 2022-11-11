package requests_test

import (
    "fmt"
    "github.com/occos/requests"
    "testing"
)

func TestGet(t *testing.T) {
    //测试Get请求
    a1 := requests.Args{
        Headers: map[string]interface{}{
            "Content-Type": "application/json",
        },
        Params: map[string]interface{}{
            "q1": "v1",
            "q2": "v2",
        },
    }
    client1 := requests.Client{}
    result1 := client1.Get("https://echo.apifox.com/get", a1)
    fmt.Println(string(result1))
}

func TestPost(t *testing.T) {
    //测试Post请求
    a2 := requests.Args{
        Headers: map[string]interface{}{
            "Content-Type": "application/json",
        },
        Params: map[string]interface{}{
            "q1": "v1",
            "q2": "v2",
        },
        Body: map[string]interface{}{
            "name":    "Jason",
            "age":     18,
            "address": "Asia/Shanghai",
        },
    }

    client2 := requests.Client{}
    result2 := client2.Post("https://echo.apifox.com/post", a2)
    fmt.Println(string(result2))
}
