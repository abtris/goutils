package goutils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// GenerateRandomNumber generate the number of count random number, between ->[start,end)
func GenerateRandomNumber(start int, end int, count int) ([]int, error) {
	//范围检查
	if end < start {
		return nil, fmt.Errorf("utils:GenerateRandomNumber random number should between start and end")
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

// RandInt rand num between min and max
func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return r.Intn(max-min) + min
}

// StringsIsFloat return the string are all float
func StringsIsFloat(in ...string) bool {
	reg, _ := regexp.Compile(`^\d*\.\d*$`)
	for _, v := range in {
		if !reg.MatchString(v) {
			return false
		}
	}
	return true
}

// StringsIsInt return the string are all int
func StringsIsInt(in ...string) bool {
	reg, _ := regexp.Compile(`^\d+$`)
	for _, v := range in {
		if !reg.MatchString(v) {
			return false
		}
	}
	return true
}

// StringIsInt return the string is int
func StringIsInt(in string) bool {
	reg, _ := regexp.Compile(`^-?\d+$`)
	if !reg.MatchString(in) {
		return false
	}
	return true
}

// StringIsIntAndToi return the string are all int, if true , convert them all
func StringIsIntAndToi(in string) (int, bool) {
	reg, _ := regexp.Compile(`^-?\d+$`)
	if !reg.MatchString(in) {
		return -1, false
	}
	ret, _ := strconv.Atoi(in)
	return ret, true
}

// StringIsFloat64AndParse  return the string are all float64, if true , convert them all
func StringIsFloat64AndParse(in string) (float64, bool) {
	ret, err := strconv.ParseFloat(in, 32)
	if err != nil {
		return 0, false
	}

	return ret, true
}

// StringIsNotEmptyAndToa return the string is not empty, if true , return both it and result
func StringIsNotEmptyAndToa(in string) (string, bool) {
	if len(in) == 0 {
		return "", false
	}
	return in, true
}

// IntStartWith return the num if starts with 'start'
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
