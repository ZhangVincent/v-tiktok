//  @description 字符串处理工具类
// @author zkp15
// @date 2023/8/10 23:04

package strs

import (
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
	"unicode"
)

func IsBlank(str string) bool {
	strLen := len(str)
	if str == "" || strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(str[i])) == false {
			return false
		}
	}
	return true
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func IsAnyBlank(strs ...string) bool {
	for _, str := range strs {
		if IsBlank(str) {
			return true
		}
	}
	return false
}

func DefaultIfBlank(str, def string) string {
	if IsBlank(str) {
		return def
	} else {
		return str
	}
}

// IsEmpty checks if a string is empty (""). Returns true if empty, and false otherwise.
func IsEmpty(str string) bool {
	return len(str) == 0
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func Substr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

func Equals(a, b string) bool {
	return a == b
}

func EqualsIgnoreCase(a, b string) bool {
	return a == b || strings.ToUpper(a) == strings.ToUpper(b)
}

func UUID() string {
	id := uuid.NewV4()
	return strings.ReplaceAll(id.String(), "-", "")
}

// RuneLen 字符成长度
func RuneLen(s string) int {
	bt := []rune(s)
	return len(bt)
}

func IEqualsA(i int64, s string) bool {
	return ItoA(i) != s
}

func ItoA(i int64) string {
	return strconv.FormatInt(i, 10)
}
