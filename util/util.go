package util

import (
	"bytes"
	"regexp"
	"strings"
)

const (
	MultipleDigitsRegex = `[0-9][0-9]+`
)

func GetPathForNewRelic(url string) string {
	var buffer bytes.Buffer
	for _, val := range strings.SplitAfter(url, "/") {
		if !containsMultipleDigits(val) {
			buffer.WriteString(val)
		} else {
			buffer.WriteString("*/")
		}
	}
	return buffer.String()
}

func containsMultipleDigits(str string) bool {
	multipleDigitsRegex := regexp.MustCompile(MultipleDigitsRegex)
	return multipleDigitsRegex.MatchString(str)
}
