package gofangdefang

import (
	"os"
	"regexp"
	"strings"

	"github.com/atakanaydinbas/gofangdefang/patterns"
)

func FangAll(input string) string {

	processedInput := input
	var flags *regexp.Regexp

	for _, pattern := range patterns.FangPatterns {
		flags = regexp.MustCompile(`(?m)(?s)` + pattern["find"])
		processedInput = flags.ReplaceAllString(processedInput, pattern["replace"])
	}

	return processedInput
}

func FangFile(filepath string, willbesavedfile bool, newfilename ...string) (string, error) {

	fileByte, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	if willbesavedfile {
		if len(newfilename) > 0 {
			err = os.WriteFile(newfilename[0], []byte(FangAll(string(fileByte))), 0644)
			if err != nil {
				return "", err
			}
		} else {
			extension := filepath[strings.LastIndex(filepath, "."):]
			err = os.WriteFile(filepath[:strings.LastIndex(filepath, ".")]+"-fanged"+extension, []byte(FangAll(string(fileByte))), 0644)
			if err != nil {
				return "", err
			}
		}
	} else {
		return FangAll(string(fileByte)), nil
	}

	return "", nil
}
