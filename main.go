package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    sysUrl "net/url"
    "strings"
)

type Data struct {
    Headers []map[string]interface{}
    Params  []map[string]interface{}
    Body    []map[string]interface{}
}

type Requests struct {
    Url    string
    Method string
    Data
}

func (r *Requests) Get() {
    result := baseRequest(r.Url, r.Method, Data{})
    fmt.Printf(string(result))
}

func (r *Requests) Post() {
    result := baseRequest(r.Url, r.Method, Data{})
    fmt.Printf(string(result))
}

func (r *Requests) Put() {
    result := baseRequest(r.Url, r.Method, Data{})
    fmt.Printf(string(result))
}

func (r *Requests) Delete() {
    result := baseRequest(r.Url, r.Method, Data{})
    fmt.Printf(string(result))
}

//接口转字符串
func interface2string(str interface{}) string {
    return fmt.Sprintf("%s", str)
}

//基础请求函数
func baseRequest(url string, method string, data Data) []byte {

    //解析Params
    var urlPath string
    if data.Params != nil {
        tmpParams := sysUrl.Values{}
        baseUrl, _ := sysUrl.Parse(url)
        for _, dict := range data.Params {
            for key, value := range dict {
                tmpParams.Set(key, interface2string(value))
            }
        }
        baseUrl.RawQuery = tmpParams.Encode()
        urlPath = baseUrl.String()
    } else {
        urlPath = url
    }

    //解析Body
    var bytesData []byte
    var tmpData *bytes.Buffer
    if data.Body != nil {
        bytesData, _ = json.Marshal(data.Body)
        tmpData = bytes.NewBuffer(bytesData)
    }

    var Req *http.Request
    var Err error

    client := &http.Client{}
    if tmpData != nil {
        Req, Err = http.NewRequest(strings.ToUpper(method), urlPath, tmpData)
    } else {
        Req, Err = http.NewRequest(strings.ToUpper(method), urlPath, nil)
    }
    if Err != nil {
        fmt.Println(Err)
    }

    //解析Headers
    if data.Headers != nil {
        for _, dict := range data.Headers {
            for key, value := range dict {
                Req.Header.Add(key, interface2string(value))
            }
        }
    }

    res, err := client.Do(Req)
    if err != nil {
        fmt.Println(err)
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
    }
    //fmt.Println(string(body))
    return body
}

func main() {
    param := make(map[string]interface{})
    param["id"] = 12131

    header := make(map[string]interface{})
    header["Content-Type"] = "application/json"

    body := make(map[string]interface{})
    body["name"] = "张三"
    body["age"] = 18
    body["address"] = "上海"

    //aaa := Requests{}
    //aaa.Url = "http://cmdb02.zmops.cc/api/v1/applications/lite/"
    //aaa.Method = "GET"
    //aaa.Params = append(aaa.Params, param)
    //aaa.Headers = append(aaa.Headers, header)
    //aaa.Get()

    bbb := Requests{
        Url:    "http://127.0.0.1:4523/m1/1911173-0-default/v1/student",
        Method: "post",
    }
    bbb.Headers = append(bbb.Headers, header)
    bbb.Body = append(bbb.Body, body)
    bbb.Post()
}
