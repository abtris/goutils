package goutils

/**
* @Author: shuxian
* @Date:   07-Feb-2017
* @Email:  shuxian@jawave.com/printfcoder@gmail.com
* @Last modified by:   shuxian
* @Last modified time: 07-Feb-2017
 */

// ArrayStringMoveDuplicates move repeated items
func ArrayStringMoveDuplicates(in []string) []string {

	ret := make([]string, 0)
	transitMap := make(map[string]bool)
	for i, j := 0, len(in); i < j; i++ {
		if ok, _ := transitMap[in[i]]; !ok {
			ret = append(ret, in[i])
			transitMap[in[i]] = true
		}
	}

	return ret
}
