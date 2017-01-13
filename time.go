package goutils

import "time"

/**
* @Author: shuxian
* @Date:   12-Jan-2017
* @Email:  shuxian@jawave.com/printfcoder@gmail.com
* @Project: jawave-center/register
* @Last modified by:   shuxian
* @Last modified time: 13-Jan-2017
 */

// GetTimeNowRFC3339 获取RFC3339的现在格式
func GetTimeNowRFC3339() string {
	return time.Now().Format(time.RFC3339)
}
