package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// request 统一请求接口
func Get(reqUrl string, params url.Values) ([]byte, error) {
	reqURL, err := url.ParseRequestURI(reqUrl)
	if err != nil {
		fmt.Printf("url.ParseRequestURI()函数执行错误,错误为:%v\n", err)
		return nil, err
	}
	fmt.Println("2 reqURL: ", reqURL, "++++++++ rawUrl: ", reqUrl)
	// 3. 处理参数，保存在reqURL.RawQuery中。
	// Encode方法将 请求参数params 编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序.
	reqURL.RawQuery = params.Encode()
	// 注意打印时，reqURL会自动将rawUrl+reqURL.RawQuery参数的形式打印，
	// 所以此时reqURL就是一个带上参数的完整url。这一步最好debug去看才能理解。
	fmt.Println("3 params.Encode(), reqURL: ", reqURL, "reqURL.RawQuery: ", reqURL.RawQuery)
	// 4. 发送HTTP请求
	// 说明: reqURL.String() String将URL重构为一个合法URL字符串。
	fmt.Println("4 Get url:", reqURL.String())
	resp, err := http.Get(reqURL.String())
	body, err := GetHTTPResponseOrg(resp, reqUrl, err)
	if err == nil {
		log.Printf("Webhook调用成功, 返回数据: %q\n", string(body))
	} else {
		log.Printf("Webhook调用失败，Err：%s\n", err)
	}
	return body, err
}

func Get1(reqUrl string, params url.Values) ([]byte, error) {
	req, err := http.NewRequest(
		"GET",
		reqUrl,
		bytes.NewBuffer(nil),
	)
	req.URL.RawQuery = params.Encode()
	if err != nil {
		log.Println("httpclient.NewRequest失败. Error: ", err)
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	body, err := GetHTTPResponseOrg(resp, reqUrl, err)
	if err == nil {
		log.Printf("Webhook调用成功, 返回数据: %q\n", string(body))
	} else {
		log.Printf("Webhook调用失败，Err：%s\n", err)
	}
	return body, err
}

func GetCustomParams() url.Values {
	return url.Values{}
}

func PostStruct(reqUrl string, header map[string]string, data interface{}) ([]byte, error) {
	dataByte, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	bodyReader := bytes.NewReader(dataByte)
	//bodyReader := bytes.NewReader([]byte(json))
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, reqUrl, bodyReader)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		request.Header.Set(k, v)
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	body, err := GetHTTPResponseOrg(resp, reqUrl, err)
	if err == nil {
		log.Printf("Webhook调用成功, 返回数据: %q\n", string(body))
	} else {
		log.Printf("Webhook调用失败，Err：%s\n", err)
	}
	return body, err
}

func PostJson(reqUrl string, header map[string]string, json string) ([]byte, error) {
	bodyReader := bytes.NewReader([]byte(json))
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, reqUrl, bodyReader)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		request.Header.Set(k, v)
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	body, err := GetHTTPResponseOrg(resp, reqUrl, err)
	if err == nil {
		log.Printf("Webhook调用成功, 返回数据: %q\n", string(body))
	} else {
		log.Printf("Webhook调用失败，Err：%s\n", err)
	}
	return body, err
}
