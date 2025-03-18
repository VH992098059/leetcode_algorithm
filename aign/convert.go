package aign

import "strings"

/*Z字形转换*/
func convert(s string, numRows int) string {
	flag, i := -1, 0
	if numRows < 2 {
		return s
	}
	str := make([]string, numRows)
	for _, v := range s {
		str[i] += string(v)
		if i == 0 {
			flag = -1
		}
		if i == numRows-1 {
			flag = 1
		}
		if flag == -1 {
			i++
		} else {
			i--
		}
	}
	return strings.Join(str, "")
}
