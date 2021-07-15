package utils

import "strings"

func IsEmptyQueryParmas(param string) bool {
	return strings.ReplaceAll(param, " ", "") == ""
}
