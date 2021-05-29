package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get http get
func Get(url string) (body []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : url=%v , statusCode=%v", url, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// Post http post
func Post(url, contentType string, data []byte) ([]byte, error) {
	body := bytes.NewBuffer(data)
	response, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", url, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// PostWithJSON http post
func PostWithJSON(url string, data []byte) ([]byte, error) {
	body := bytes.NewBuffer(data)
	response, err := http.Post(url, "application/json", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v, err=%v", url, response.StatusCode, err)
	}
	return ioutil.ReadAll(response.Body)
}

// PostWithAuthorizationSerial POST 认证信息 证书编号
func PostWithAuthorizationSerial(url string, data []byte, authorization string, serial string) ([]byte, error) {
	body := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Printf("创建HTTP REQUEST错误[%v]", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", authorization)
	request.Header.Set("Wechatpay-Serial", serial)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v, err=%v", url, response.StatusCode, err)
	}
	return ioutil.ReadAll(response.Body)
}
