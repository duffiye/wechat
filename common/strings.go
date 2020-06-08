package common

// IsNotBlank 判断字符串是否为空
func IsNotBlank(str string) bool {
	if str == "" || str == "null" || str == "NULL" {
		return false
	}
	return true
}

// IsBlank 判断字符串是否为空
func IsBlank(str string) bool {
	return !IsNotBlank(str)
}
