package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//http 包
	client := http.Client{}
	//func (c *Client) Get(url string) (resp *Response, err error)
	resp, err := client.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("client.get err:", err)
		return
	}
	fmt.Println("****************************************")
	body := resp.Body
	fmt.Println("body 111:", body)
	readBodystr, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("read body err:", err)
	}
	fmt.Println("body string:", string(readBodystr))

	fmt.Println("****************************************")
	//func (h Header) Get(key string) string 获取  响应头
	ct := resp.Header.Get("Content-type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")

	fmt.Println("Content-type:", ct)
	fmt.Println("Date:", date)
	fmt.Println("server:", server)
	fmt.Println("****************************************")
	url := resp.Request.URL // 获取 状态行  信息
	code := resp.StatusCode
	status := resp.Status

	fmt.Println("url:", url) // url: http://www.baidu.com

	fmt.Println("code:", code) //code: 200

	fmt.Println("status:", status) // status: 200 OK

}
