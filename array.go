package goutils

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
