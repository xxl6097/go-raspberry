package httpclient

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func getMD5(cla, clt, method, secret string) string {
	msg := cla + clt + method + secret //"hipikids1704858534000GETHipikids@28f11030af9c11eeb6ea1b717e93b065@hipikids" //
	fmt.Println(msg)
	return MD5(msg)
}
func TestSum(t *testing.T) {
	cla := "hipikids"
	secret := "Hipikids@28f11030af9c11eeb6ea1b717e93b065@hipikids"

	nowtime := time.Now()
	seconds := float64(8 * 3600)
	nanoseconds := int64(seconds * 1e9)
	nano := nowtime.UnixNano() - nanoseconds
	milliseconds := nano / (1000 * 1000)

	// 将毫秒时间戳转换为字符串
	fmt.Println("milliseconds:", milliseconds)
	timeString := strconv.FormatInt(milliseconds, 10)
	// 打印结果
	fmt.Println("UTC毫秒时间戳:", timeString)
	cls := getMD5(cla, timeString, "GET", secret)
	fmt.Println("hello:", cls)
	//https://itest.clife.net/clink/api/hipikids/secret/education/hipikids/all?cla=hipikids&clt=1704858534000&cls=0e7773501b037adfb5ff6648dc9bfb54
	host := "https://itest.clife.net/clink/api/hipikids/secret/education/hipikids/all"
	params := url.Values{} // Values是一个map类型
	params.Set("cla", cla) // Set的时候会把string转成[]string
	params.Set("clt", timeString)
	params.Set("cls", cls)
	//host += "?cla="
	//host += cla
	//host += "&clt="
	//host += timeString
	//host += "&cls="
	//host += cls
	fmt.Println(host)
	_byte, _ := Get(host, params)
	if _byte != nil {
		fmt.Println(string(_byte))
	}

}
