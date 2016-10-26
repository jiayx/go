package stringutil

// InArray 查找一个字符串是否在一个数组里
func InArray(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}

	return false
}
