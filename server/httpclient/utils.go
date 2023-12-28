package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

import "time"

// "2023-05-29 15:10:41"
func GetNowStr() string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

// GetHTTPResponse 处理HTTP结果，返回序列化的json
func GetHTTPResponse(resp *http.Response, url string, err error, result interface{}) error {
	body, err := GetHTTPResponseOrg(resp, url, err)
	if err == nil {
		// log.Println(string(body))
		if len(body) != 0 {
			err = json.Unmarshal(body, &result)
			if err != nil {
				log.Printf("请求接口%s解析json结果失败! ERROR: %s\n", url, err)
			}
		}
	}
	return err
}

// GetHTTPResponseOrg 处理HTTP结果，返回byte
func GetHTTPResponseOrg(resp *http.Response, url string, err error) ([]byte, error) {
	if resp == nil {
		log.Printf("请求接口%s失败! ERROR: %s\n", url, err)
		return nil, err
	}

	defer resp.Body.Close()
	if err != nil {
		log.Printf("请求接口%s失败! ERROR: %s\n", url, err)
		return nil, err
	}

	lr := io.LimitReader(resp.Body, 1024000)
	body, err := io.ReadAll(lr)

	if err != nil {
		log.Printf("请求接口%s失败! ERROR: %s\n", url, err)
	}

	// 300及以上状态码都算异常
	if resp.StatusCode >= 300 {
		errMsg := fmt.Sprintf("请求接口 %s 失败! 返回内容: %s ,返回状态码: %d\n", url, string(body), resp.StatusCode)
		log.Println(errMsg)
		err = fmt.Errorf(errMsg)
	}

	return body, err
}
