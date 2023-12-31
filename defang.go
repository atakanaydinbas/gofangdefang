package gofangdefang

import (
	"regexp"
	"strings"
)

func DefangAll(input string) string {

	output := input
	for _, pattern := range DefangPatterns {
		re := regexp.MustCompile(pattern["find"])
		matches := re.FindAllString(input, -1)

		for _, match := range matches {
			modified := strings.Replace(match, pattern["change"], pattern["replace"], -1)
			output = strings.Replace(output, match, modified, -1)
		}

	}
	return output
}

var DefangPatterns = []map[string]string{
	{
		"find":    `(?:\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b|\b[a-zA-Z0-9-]+\.[a-zA-Z]{2,}\b)`,
		"change":  ".",
		"replace": "[.]",
	},

	{
		"find":    `^(http:\/\/)`,
		"change":  "http:",
		"replace": "hXXp:",
	},

	{
		"find":    `^(https:\/\/)`,
		"change":  "https",
		"replace": "hXXps:",
	},

	{
		"find":    `([a-zA-Z0-9_.+-]+)@([a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+)`,
		"change":  "@",
		"replace": "[AT]",
	},
}
