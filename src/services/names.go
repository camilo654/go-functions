package services

import "strings"

func GetNames(names string) (int, []string) {
	namesArray := strings.Split(names, ",")
	return len(namesArray), namesArray
}
