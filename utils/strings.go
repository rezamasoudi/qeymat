package utils

import "strings"

func ReplacePersianNumbersToEnglishNumbers(content string, oldNew ...string) string {
	oldNew = append(oldNew, "۰", "0", "۱", "1", "۲", "2", "۳", "3", "۴", "4", "۵", "5", "۶", "6", "۷", "7", "۸", "8", "۹", "9")
	replacer := strings.NewReplacer(oldNew...)
	return replacer.Replace(content)
}
