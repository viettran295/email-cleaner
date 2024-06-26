package utils

import "strings"

func ContainString(arr []string, str string) bool {
	str = trimSenderValue(str)
	for _, substr := range arr {
		if strings.Contains(str, substr){
			return true
		}
	}
	return false
}

func trimSenderValue(str string) string {
	if strings.Contains(str, "<"){
		splitStr := strings.Split(str, "<")
		emailAddr := strings.Split(splitStr[1], ">")
		return emailAddr[0]
	} else {
		return str
	}
}
