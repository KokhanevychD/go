package algo

import (
	"strings"
)

// ReverseVowels func - reverse all vowels in given string
func ReverseVowels(str string) string {
	strArr := make([]string, len(str))
	vowels := "aeiouAEIOU"
	idx2 := len(str) - 1

	for idx := 0; idx < len(str); idx++ {
		if strings.Contains(vowels, string(str[idx])) && idx < idx2 {
			swaped := false
			for ; idx2 > idx; idx2-- {
				if strings.Contains(vowels, string(str[idx2])) {
					strArr[idx] = string(str[idx2])
					strArr[idx2] = string(str[idx])
					idx2--
					swaped = true
					break
				} else {
					strArr[idx2] = string(str[idx2])
				}
			}
			if !swaped {
				strArr[idx] = string(str[idx])
			}
		} else {
			strArr[idx] = string(str[idx])
		}
		if idx == idx2 {
			break
		}
	}
	res := strings.Join(strArr, "")
	return res
}
