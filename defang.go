package gofangdefang

import (
	"os"
	"regexp"
	"strings"

	"github.com/atakanaydinbas/gofangdefang/patterns"
)

func DefangText(input string) string {

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

func DefangFile(filepath string, willbesavedfile bool, newfilename ...string) (string, error) {

	fileByte, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	if willbesavedfile {
		if len(newfilename) > 0 {
			err = os.WriteFile(newfilename[0], []byte(DefangText(string(fileByte))), 0644)
			if err != nil {
				return "", err
			}
		} else {
			extension := filepath[strings.LastIndex(filepath, "."):]
			err = os.WriteFile(filepath[:strings.LastIndex(filepath, ".")]+"-defanged"+extension, []byte(DefangText(string(fileByte))), 0644)
			if err != nil {
				return "", err
			}
		}
	} else {
		return DefangText(string(fileByte)), nil
	}

	return "", nil
}
