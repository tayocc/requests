package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
	"strconv"
	"strings"
)

type Args struct {
	Headers map[string]interface{}
	Params  map[string]interface{}
	Body    map[string]interface{}
}

type Client struct{}

func (c *Client) Get(url string, args Args) []byte {
	return c.Request("GET", url, args)
}

func (c *Client) Post(url string, args Args) []byte {
	return c.Request("POST", url, args)
}

//func (c *Client) Put(url string, data []byte, args ...Args) {
///
//}
//
//func (c *Client) Delete(url string, args ...Args) {
//
//}
//
//func (c *Client) Head(url string, args ...Args) {
//
//}
//
//func (c *Client) Patch(url string, data []byte, args ...Args) {
//
//}

func (c *Client) Request(method string, url string, args Args) []byte {
	//解析Params
	var urlPath string
	if args.Params != nil {
		urlPath = setParams(url, args.Params)
	} else {
		urlPath = url
	}

	fmt.Println("==============================================")
	fmt.Println(urlPath)
	fmt.Println("==============================================")

	//解析Body
	var bytesData []byte
	var tmpData *bytes.Buffer
	if args.Body != nil {
		bytesData, _ = json.Marshal(args.Body)
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
	if args.Headers != nil {
		setHeaders(Req, args.Headers)
	}

	res, err := client.Do(Req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

// 接口转字符串
func interface2string(inter interface{}) string {

	var newString string
	if inter == nil {
		return newString
	}

	switch inter.(type) {
	case float64:
		ft := inter.(float64)
		newString = strconv.FormatFloat(ft, 'f', -1, 64)

	case float32:
		ft := inter.(float32)
		newString = strconv.FormatFloat(float64(ft), 'f', -1, 64)

	case int:
		it := inter.(int)
		newString = strconv.Itoa(it)

	case uint:
		it := inter.(uint)
		newString = strconv.Itoa(int(it))

	case int8:
		it := inter.(int8)
		newString = strconv.Itoa(int(it))

	case uint8:
		it := inter.(uint8)
		newString = strconv.Itoa(int(it))

	case int16:
		it := inter.(int16)
		newString = strconv.Itoa(int(it))

	case uint16:
		it := inter.(uint16)
		newString = strconv.Itoa(int(it))

	case int32:
		it := inter.(int32)
		newString = strconv.Itoa(int(it))

	case uint32:
		it := inter.(uint32)
		newString = strconv.Itoa(int(it))

	case int64:
		it := inter.(int64)
		newString = strconv.FormatInt(it, 10)

	case uint64:
		it := inter.(uint64)
		newString = strconv.FormatUint(it, 10)

	case string:
		newString = inter.(string)

	case []byte:
		newString = string(inter.([]byte))

	default:
		newValue, _ := json.Marshal(inter)
		newString = string(newValue)
	}

	return newString
}

// 设置请求头
func setHeaders(r *http.Request, headers map[string]interface{}) {
	for key, value := range headers {
		r.Header.Add(key, interface2string(value))
	}
}

// 设置查询参数
func setParams(url string, params map[string]interface{}) string {
	tmpParams := url2.Values{}
	baseUrl, _ := url2.Parse(url)
	for key, value := range params {
		tmpParams.Set(key, interface2string(value))
	}
	baseUrl.RawQuery = tmpParams.Encode()
	urlPath := baseUrl.String()
	return urlPath
}
