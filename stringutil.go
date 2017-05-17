/**
* @Author: shuxian
* @Date:   18-Oct-2016
* @Email:  shuxian@jawave.com/printfcoder@gmail.com
* @Last modified by:   shuxian
* @Last modified time: 07-Feb-2017
 */

package goutils

import (
	crand "crypto/rand"
	"fmt"
	"io"
	mrand "math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	// RandKindNum  all number
	RandKindNum = 0

	// RandKindLower lower
	RandKindLower = 1

	// RandKindUpper Upper
	RandKindUpper = 2

	// RandKindAll number with string
	RandKindAll = 3
)

// AllStringIsEmpty 判断传入的字符串是否全是空的
func AllStringIsEmpty(in ...string) bool {
	for _, str := range in {
		if str == "" {
			return false
		}
	}
	return true
}

// StringsHasOneEmpty 判断传入的字符串至少有一条是空的
func StringsHasOneEmpty(in ...string) bool {
	for _, str := range in {
		if len(str) == 0 {
			return true
		}
	}
	return false
}

// IntArrayToStringArray 将Int型数组转成字符串型数组
func IntArrayToStringArray(in []int) []string {
	ret := make([]string, 0, len(in))
	for _, v := range in {
		ret = append(ret, strconv.Itoa(v))
	}
	return ret
}

// IDArrayToSQLInString 将ID数组转成SQL in 字符串
func IDArrayToSQLInString(in []int) string {
	if len(in) > 0 {
		ret := strings.Join(IntArrayToStringArray(in), "','")
		return "'" + ret + "'"
	}
	return ""
}

// StringRandWithSize Rand ize string
func StringRandWithSize(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	mrand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll {
			// random ikind
			ikind = mrand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + mrand.Intn(scope))
	}
	return result
}

// BytesToString 判断传入的字符串至少有一条是空的
func BytesToString(in ...byte) string {
	var ret string
	for _, str := range in {
		ret += string(str)
	}
	return ret
}

// ReplaceSQLSpecialPunctuation 替换sql特殊字串
func ReplaceSQLSpecialPunctuation(in string) string {
	return strings.NewReplacer("'", "\\'", ",", "\\,").Replace(in)
}

// IsEndWith 以指定char 结尾
func IsEndWith(in, char string) bool {
	reg, _ := regexp.Compile(char + `$`)
	return reg.MatchString(in)
}

// IsStartWith 以指定char 开头
func IsStartWith(in, char string) bool {
	reg, _ := regexp.Compile(`^` + char)
	return reg.MatchString(in)
}

// MakeTimeSerialNum 以时间为基准生成以某标识开头，指定长度随机数字结尾的编号，适用于单号、流水号等
func MakeTimeSerialNum(prefix, time string, randStart, randEnd int) string {
	reg, _ := regexp.Compile(`-|:|\s`)
	return prefix + reg.ReplaceAllString(time, "") + strconv.Itoa(RandInt(randStart, randEnd))
}

// HasSubString 是否有字串
func HasSubString(in, sub string) bool {
	reg, _ := regexp.Compile(sub)
	return reg.MatchString(in)
}

//StringEqualto 是否相等，cas 大小敏感
func StringEqualto(in1, in2 string, cas bool) bool {
	var reg *regexp.Regexp
	if cas {
		reg, _ = regexp.Compile("^" + in1 + "$")

	} else {
		reg, _ = regexp.Compile("(?i)" + "^" + in1 + "$")
	}
	return reg.MatchString(in2)
}

// StringIsIn 判断字符串是否在输入串组中
func StringIsIn(in string, arg ...string) bool {
	for _, v := range arg {
		if in == v {
			return true
		}
	}
	return false
}

// StringIsHalfWidthNoSpace 判断是不是半角字符--不能有空格
func StringIsHalfWidthNoSpace(in string) bool {
	reg, _ := regexp.Compile(`^[\d\w]+$`)
	return reg.MatchString(in)
}

// StringIsHalfWidth  判断是不是半角字符
func StringIsHalfWidth(in string) bool {
	reg, _ := regexp.Compile(`^[\d\s\w]+$`)
	return reg.MatchString(in)
}

// StringSubStrBetweenHunger 找到字符区间的字符串-饥饿匹配，不处理正则自有特殊字符
func StringSubStrBetweenHunger(in, start, end string) string {
	reg, _ := regexp.Compile(`.*` + start + "(.*)" + end + ".*")
	return reg.ReplaceAllString(in, "$1")
}

// StringSubStrBetweenIndexes substr between start and end
func StringSubStrBetweenIndexes(in string, start, end int) string {
	str2 := []rune(in)
	return string(str2[len(str2)-start : len(str2)-start+end])
}

// StringSubStrAfterIndex substr after the index
func StringSubStrAfterIndex(in string, index int) string {

	str2 := []rune(in)
	if index > 0 && len(str2) > index {
		return string(str2[index:])
	} else if index < 0 && len(str2) > -index {
		return string(str2[len(str2)+index:])
	} else if index < 0 && -index >= len(str2) {
		return in
	}
	return ""
}

// StringTrimNewLine clean new line,rpls is the replacer
func StringTrimNewLine(in, rpls string) string {
	reg, _ := regexp.Compile(`\n`)
	return reg.ReplaceAllString(in, rpls)
}

// StringChinesePhoneNumOrEmail phone all email
// if return 2 means 'in' is phone, 1 means it is email
func StringChinesePhoneNumOrEmail(in string) int {
	reg, _ := regexp.Compile(`^1[34578]\d{9}$`)
	if reg.MatchString(in) {
		return 2
	}

	reg = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if reg.MatchString(in) {
		return 1
	}

	return 0
}

// NewUUID generates a random UUID according to RFC 4122
func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(crand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// StringsAreAllIn return all are in 'in'
func StringsAreAllIn(all []string, in ...string) (index int, ret bool) {
	for i, v := range all {
		if !StringIsIn(v, in...) {
			return i, false
		}
	}
	return 0, true
}

// StringsReplaceAllNonNumeric replace all non-numeric chars
// keep the remaining chars and convert it to int
// if all chars are non-numeric, then return 0
func StringsReplaceAllNonNumeric(in string) int64 {

	reg, _ := regexp.Compile(`\D`)
	tempStr := reg.ReplaceAllString(in, "")

	ret, err := strconv.ParseInt(tempStr, 10, 64)
	if err != nil {
		return 0
	}

	return ret
}
