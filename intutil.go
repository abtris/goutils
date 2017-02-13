package goutils

/**
* 整形工具集
* @Author: shuxian
* @Date:   14-Oct-2016
* @Email:  shuxian@jawave.com/printfcoder@gmail.com
* @Last modified by:   shuxian
* @Last modified time: 07-Feb-2017
 */

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

//GenerateRandomNumber 生成count个[start,end)结束的不重复的随机数  RandInt 指定范围内的随机数
func GenerateRandomNumber(start int, end int, count int) ([]int, error) {
	//范围检查
	if end < start {
		return nil, fmt.Errorf("utils:GenerateRandomNumber随机范围结束数字小于开始数字")
	}

	var temp = count
	if (end - start) < count {
		temp = end - start
	}

	var nums []int
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < temp {
		//生成随机数
		num := r.Intn(end-start) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums, nil
}

// RandInt 指定范围内的随机数
func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return r.Intn(max-min) + min
}

//StringsIsFloat 判断字符串是不是float 类型的数据
func StringsIsFloat(in ...string) bool {
	reg, _ := regexp.Compile(`^\d*\.\d*$`)
	for _, v := range in {
		if !reg.MatchString(v) {
			return false
		}
	}
	return true
}

//StringsIsInt 判断字符串是不是int 类型的数据
func StringsIsInt(in ...string) bool {
	reg, _ := regexp.Compile(`^\d+$`)
	for _, v := range in {
		if !reg.MatchString(v) {
			return false
		}
	}
	return true
}

//StringIsInt 判断字符串是不是int, 类型的数据，是则顺便转换
func StringIsInt(in string) bool {
	reg, _ := regexp.Compile(`^-?\d+$`)
	if !reg.MatchString(in) {
		return false
	}
	return true
}

//StringIsIntAndToi 判断字符串是不是int, 类型的数据，是则顺便转换
func StringIsIntAndToi(in string) (int, bool) {
	reg, _ := regexp.Compile(`^-?\d+$`)
	if !reg.MatchString(in) {
		return -1, false
	}
	ret, _ := strconv.Atoi(in)
	return ret, true
}

//StringIsFloat64AndParse 判断字符串是不是float64, 类型的数据，是则顺便转换
func StringIsFloat64AndParse(in string) (float64, bool) {
	ret, err := strconv.ParseFloat(in, 32)
	if err != nil {
		return 0, false
	}

	return ret, true
}

// StringIsNotEmptyAndToa 判断字符串是不是长度大于0，是则传出
func StringIsNotEmptyAndToa(in string) (string, bool) {
	if len(in) == 0 {
		return "", false
	}
	return in, true
}

// IntStartWith 整形数据是否以另一个整形开头
func IntStartWith(origin, start int) bool {
	reg := regexp.MustCompile("^" + strconv.Itoa(start))
	return reg.MatchString(strconv.Itoa(origin))
}

// IntIsBetween the number in is between start and end
func IntIsBetween(in, start, end int) bool {
	if in < start {
		return false
	}

	if in > end {
		return false
	}

	return true
}

// IntIsIn the number is in the array
func IntIsIn(in int, arr ...int) bool {

	for i, j := 0, len(arr); i < j; i++ {
		if in == arr[i] {
			return true
		}
	}

	return false
}
