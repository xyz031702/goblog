package helpers

import (
	"time"
	"html/template"
	"strings"
)

// 格式化时间
func DateFormat(date time.Time, layout string) string {
	return date.Format(layout)
}

// 截取字符串
func Substr(source string, start, end int) string {
	rs := []rune(source)
	length := len(rs)
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	return string(rs[start:end])
}

func Unescaped (x string) interface{} {
	return template.HTML(x)
}

func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) > n {
		return string(runes[:n])
	}
	return s
}

//返回资源路径
func StaticUrl(url ...string) string {
	if len(url) > 0 {
		return "/static/" + strings.Trim(url[0], "/")
	}

	return "/static/"
}