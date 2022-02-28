package utils

import "regexp"

// PatternRegex 匹配正则表达式
func PatternRegex(input string, regex string) bool {
	compile, _ := regexp.Compile(regex)
	return compile.MatchString(input)
}