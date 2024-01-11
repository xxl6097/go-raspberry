package utils

import "time"

const timezone = "Asia/Shanghai"

func GetTimeFormat(format string) string {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		location = time.FixedZone("CST", 8*3600) //替换上海时区方式
	}
	date := time.Now()
	date.In(location)
	return date.Format(format)

}

func GetTimeDir() string {
	return GetTimeFormat("2006/01/02/15/04/05/")
}
