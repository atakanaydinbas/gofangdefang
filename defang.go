package gofangdefang

import (
	"regexp"
	"strings"

	"github.com/atakanaydinbas/gofangdefang/patterns"
)

func DefangAll(input string) string {

	output := input
	for _, pattern := range patterns.DefangPatterns {
		re := regexp.MustCompile(pattern["find"])
		matches := re.FindAllString(input, -1)

		for _, match := range matches {
			modified := strings.Replace(match, pattern["change"], pattern["replace"], -1)
			output = strings.Replace(output, match, modified, -1)
		}

	}
	return output
}
