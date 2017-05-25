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

// ArrayStringMoveItems move items
func ArrayStringMoveItems(in []string, items ...string) []string {
	ret := make([]string, 0)
	for i, j := 0, len(in); i < j; i++ {
		var beIn bool
		for k, l := 0, len(items); k < l; k++ {
			if in[i] == items[k] {
				beIn = true
				break
			}
		}
		if !beIn {
			ret = append(ret, in[i])
		}
	}

	return ret
}
