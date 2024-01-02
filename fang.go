package gofangdefang

import (
	"regexp"

	"github.com/atakanaydinbas/gofangdefang/patterns"
)

func FangAll(input string) string {

	processedInput := input
	var flags *regexp.Regexp

	for _, pattern := range patterns.DefangPatterns {
		flags = regexp.MustCompile(`(?m)(?s)` + pattern["find"])
		processedInput = flags.ReplaceAllString(processedInput, pattern["replace"])
	}

	return processedInput
}
