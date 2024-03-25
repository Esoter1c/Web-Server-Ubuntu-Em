package FunctionsString

import (
	"strings"
)

func GetDir(str string, levelLess int) (strResult string) {

	var arrayStr []string

	var charSeparator string

	if strings.ContainsAny(str, "/") {
		arrayStr = strings.Split(str, "/")
		charSeparator = "/"
	} else {
		arrayStr = strings.Split(str, "\\")
		charSeparator = "\\"
	}

	for it := 0; it < len(arrayStr)-levelLess; it++ {
		strResult += arrayStr[it] + charSeparator
	}

	return
}
