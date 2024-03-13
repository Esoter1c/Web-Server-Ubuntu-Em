package functionsString

import "strings"

func GetDir(str string, levelLess int) (strResult string) {

	arrayStr := strings.Split(str, "\\")

	for it := 0; it < len(arrayStr)-levelLess; it++ {
		strResult += arrayStr[it] + "\\"
	}

	return
}
