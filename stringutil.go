package utils

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
字符串工具集

Author: SX
Date: 2016-4-8
*/

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

//AllStringIsEmpty 判断传入的字符串是否全是空的
func AllStringIsEmpty(in ...string) bool {
	for _, str := range in {
		if str == "" {
			return false
		}
	}
	return true
}

//StringsHasOneEmpty 判断传入的字符串至少有一条是空的
func StringsHasOneEmpty(in ...string) bool {
	for _, str := range in {
		if len(str) == 0 {
			return true
		}
	}
	return false
}

//IntArrayToStringArray 将Int型数组转成字符串型数组
func IntArrayToStringArray(in []int) []string {
	ret := make([]string, 0, len(in))
	for _, v := range in {
		ret = append(ret, strconv.Itoa(v))
	}
	return ret
}

//IDArrayToSQLInString 将ID数组转成SQL in 字符串
func IDArrayToSQLInString(in []int) string {
	if len(in) > 0 {
		ret := strings.Join(IntArrayToStringArray(in), "','")
		return "'" + ret + "'"
	}
	return ""
}

//Krand 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

//StringsHasOneEmpty 判断传入的字符串至少有一条是空的
func BytesToString(in ...byte) string {
	var ret string
	for _, str := range in {
		ret += string(str)
	}
	return ret
}

func ReplaceSQLSpecialPunctuation(in string) string {
	return strings.NewReplacer("'", "\\'", ",", "\\,").Replace(in)
}

//IsEndWith 以指定char 结尾
func IsEndWith(in, char string) bool {
	reg, _ := regexp.Compile(char + `$`)
	return reg.MatchString(in)
}

//IsStartWith 以指定char 开头
func IsStartWith(in, char string) bool {
	reg, _ := regexp.Compile(`^` + char)
	return reg.MatchString(in)
}

//MakeTimeSerialNum 以时间为基准生成以某标识开头，指定长度随机数字结尾的编号，适用于单号、流水号等
func MakeTimeSerialNum(prefix, time string, randStart, randEnd int) string {
	reg, _ := regexp.Compile(`-|:|\s`)
	return prefix + reg.ReplaceAllString(time, "") + strconv.Itoa(RandInt(randStart, randEnd))
}

func HasSubString(in, sub string) bool {
	reg, _ := regexp.Compile(sub)
	return reg.MatchString(in)
}

//StringEqualto
func StringEqualto(in1, in2 string, cas bool) bool {
	var reg *regexp.Regexp
	if cas {
		reg, _ = regexp.Compile("^" + in1 + "$")

	} else {
		reg, _ = regexp.Compile("(?i)" + "^" + in1 + "$")
	}
	return reg.MatchString(in2)
}

//StringIsIn 判断字符串是否在输入串组中
func StringIsIn(in string, arg ...string) bool {
	for _, v := range arg {
		if in == v {
			return true
		}
	}
	return false
}

//StringIsHalfWidthNoSpace 判断是不是半角字符--不能有空格
func StringIsHalfWidthNoSpace(in string) bool {
	reg, _ := regexp.Compile(`^[\d\w]+$`)
	return reg.MatchString(in)
}

//StringIsHalfWidth  判断是不是半角字符
func StringIsHalfWidth(in string) bool {
	reg, _ := regexp.Compile(`^[\d\s\w]+$`)
	return reg.MatchString(in)
}

//StringSubStrBetweenHunger 找到字符区间的字符串-饥饿匹配，不处理正则自有特殊字符
func StringSubStrBetweenHunger(in, start, end string) string {
	reg, _ := regexp.Compile(`.*` + start + "(.*)" + end + ".*")
	return reg.ReplaceAllString(in, "$1")
}

//StringTrimNewLine 清除换行,rpls为替换符
func StringTrimNewLine(in, rpls string) string {
	reg, _ := regexp.Compile(`\n`)
	return reg.ReplaceAllString(in, rpls)
}
