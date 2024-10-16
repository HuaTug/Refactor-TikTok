package utils

import (
	"fmt"
	"strconv"
	"time"
)

const defaultTimeFormat = `2006-01-02 15:04:34`

var timeLocation, _ = time.LoadLocation("Asia/Shanghai")

func Transfer(value interface{}) int64 {
	switch v := value.(type) {
	case int64:
		fmt.Println("Value is int64:", v)
		return v // 直接返回 int64 值
	case float64:
		// 如果是 float64，可以将其转换为 int64
		fmt.Println("Value is float64, converting to int64:", int64(v))
		return int64(v) // 返回转换后的 int64 值
	case string:
		// 如果是字符串，可以尝试解析为 int64
		if intValue, err := strconv.ParseInt(v, 10, 64); err == nil {
			fmt.Println("Value is string, converted to int64:", intValue)
			return intValue // 返回解析后的 int64 值
		} else {
			fmt.Println("Failed to convert string to int64:", err)
		}
	default:
		fmt.Println("Unsupported type")
	}
	return -1 // 如果没有匹配的类型，返回 0 或其他默认值
}

func ConvertStringToInt64(v string) (int64, error) {
	if res, err := strconv.ParseInt(v, 10, 64); err != nil {
		return -1, err
	} else {
		return res, nil
	}
}

func ConvertTimestampToString(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(defaultTimeFormat)
}

func ConvertStringToTimestampDefault(date string) int64 {
	t, _ := time.ParseInLocation(defaultTimeFormat, date, timeLocation)
	return t.Unix()
}
func StringToUnixTime(timeStr string) (int64, error) {
	// 将字符串类型的时间转换为 time.Time 类型
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return 0, err
	}
	// 将 time.Time 类型的时间转换为 UNIX 时间戳
	unixTime := t.Unix()
	return unixTime, nil
}
