package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Request
//
//	url 请求地址
//	header 头部
//	requestData 请求数据
//	method 请求方法
func Request(url string, header map[string]string, requestData []byte, method string) []byte {
	//rwLock.Lock()

	//payload := strings.NewReader(requestData)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(requestData))
	//req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	for key, value := range header {
		req.Header.Set(key, value)
	}
	//过滤https证书
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//关闭连接池，不然会打满语句柄
		DisableKeepAlives: true,
	}

	//设置请求超时时间为20秒
	client := &http.Client{
		Transport: tr,
		Timeout:   20 * time.Second,
	}
	res, err := client.Do(req)
	if res != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Printf("关闭body出错: %s\n", err.Error())
			}
		}(res.Body)
		body, _ := io.ReadAll(res.Body)
		return body
	}
	if err != nil {
		fmt.Printf("请求错误: %s\n", err.Error())
		return nil
	}
	return nil

}
