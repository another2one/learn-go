package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// FormData
//
//	url 请求地址
//	header 头部
//	params 其他请求参数
//	paramName 文件名称
//	path 本地文件路径
func FormData(url string, header map[string]string, params map[string]string, paramName, path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("打开文件错误: %s\n", err.Error())
		return nil
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	//fmt.Printf("请求参数：%+v",params)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		fmt.Printf("文件错误: %s\n", err.Error())
		return nil
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)

	}
	err = writer.Close()
	if err != nil {
		fmt.Printf("文件关闭错误: %s\n", err.Error())
		return nil
	}

	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
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
