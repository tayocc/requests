package requests_test

import (
    "fmt"
    "github.com/occos/requests"
    "testing"
)

func TestClient_Get(t *testing.T) {
    args := requests.Args{
        Headers: map[string]interface{}{
            "Content-Type": "application/json",
        },
        Params: map[string]interface{}{
            "q1": "v1",
            "q2": "v2",
        },
    }
    client := requests.Client{}
    result := client.Get("https://echo.apifox.com/get", args)
    fmt.Println(string(result))
}

func TestClient_Post(t *testing.T) {
    args := requests.Args{
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

    client := requests.Client{}
    result := client.Post("https://echo.apifox.com/post", args)
    fmt.Println(string(result))
}

func TestClient_Put(t *testing.T) {
    args := requests.Args{
        Headers: map[string]interface{}{
            "Content-Type": "application/json",
        },
        Params: map[string]interface{}{
            "q1": "v1",
        },
        Body: map[string]interface{}{
            "test": "value",
        },
    }

    client := requests.Client{}
    result := client.Put("https://echo.apifox.com/put", args)
    fmt.Println(string(result))
}

func TestClient_Delete(t *testing.T) {
    args := requests.Args{
        Headers: map[string]interface{}{
            "Content-Type": "application/json",
        },
        Params: map[string]interface{}{
            "q1": "v1",
        },
        Body: map[string]interface{}{
            "b1": "v1",
            "b2": "v2",
        },
    }

    client := requests.Client{}
    result := client.Delete("https://echo.apifox.com/delete", args)
    fmt.Println(string(result))
}

func TestClient_Patch(t *testing.T) {
    args := requests.Args{
        Headers: map[string]interface{}{
            "Content-Type": "application/json",
        },
        Params: map[string]interface{}{
            "q1": "v1",
        },
    }

    client := requests.Client{}
    result := client.Patch("https://echo.apifox.com/patch", args)
    fmt.Println(string(result))
}

func TestClient_Head(t *testing.T) {
    args := requests.Args{}

    client := requests.Client{}
    result := client.Head("https://echo.apifox.com/get", args)
    fmt.Println(result)
}
