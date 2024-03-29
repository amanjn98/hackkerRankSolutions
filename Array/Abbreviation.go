package Array

import (
	"strings"
)

func GetAbbrevName(fullname string) string {
	nameArray := strings.Split(fullname, " ")
	newName := ""
	for i := 0; i < len(nameArray)-1; i++ {
		newName += string(nameArray[0][0]) + "."
	}
	return newName
}

// file: name_test.go
