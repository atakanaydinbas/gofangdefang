package gofangdefang

import (
	"regexp"
)

func FangAll(input string) string {

	processedInput := input
	var flags *regexp.Regexp

	for _, pattern := range FangPatterns {
		flags = regexp.MustCompile(`(?m)(?s)` + pattern["find"])
		processedInput = flags.ReplaceAllString(processedInput, pattern["replace"])
	}

	return processedInput
}

var FangPatterns = []map[string]string{
	{
		"find":    `(\[:\/\/\])`,
		"replace": "://",
	},
	{
		"find":    `((\ *[\[\]\(\)\{\}]{1}\ *)[.,]([\[\]\(\)\{\}]{1}\ *))`,
		"replace": ".",
	},
	{
		"find":    `((\ *[\[\]\(\)\{\}]{1}\ *):(\ *[\[\]\(\)\{\}]{1}\ *))`,
		"replace": ":",
	},
	{
		"find":    `((\ *[\[\]\(\)\{\}]*\ *)DOT(\ *[\[\]\(\)\{\}]*\ *))`,
		"replace": ".",
	},
	{
		"find":    `((\ *[\[\]\(\)\{\}-]+\ *)(?:dot|punto|punkt)(\ *[\[\]\(\)\{\}-]*\ *))`,
		"replace": ".",
	},
	{
		"find":    `((\ *[\[\]\(\)\{\}-]*\ *)(?:dot|punto|punkt)(\ *[\[\]\(\)\{\}-]+\ *))`,
		"replace": ".",
	},
	{
		"find":    `((\ *[\[\]\(\)\{\}]+\ *)(?:(?:AT)|(?:ET)|(?:ARROBA))(\ *[\[\]\(\)\{\}]*\ *))`,
		"replace": "@",
	},
	{
		"find":    `((\ *[\[\]\(\)\{\}]*\ *)(?:(?:AT)|(?:ET)|(?:ARROBA))(\ *[\[\]\(\)\{\}]+\ *))`,
		"replace": "@",
	},
	{
		"find":    `([a-z])\ *(?:AT|ET|ARROBA)\ *([a-z])`,
		"replace": `$1@$2`,
	},
	{
		"find":    `(([\[\]\(\)\{\}]{1}\ *)www(\ *[\[\]\(\)\{\}]{1}\ *))`,
		"replace": "www",
	},
	{
		"find":    `:\/\/\/+`,
		"replace": "://",
	},
	{
		"find":    `:\/\/ *`,
		"replace": "://",
	},
	{
		"find":    `: +\/\/`,
		"replace": "://",
	},
	{
		"find":    `(?:[\[\(\{]+\ *)htt(ps?)(?:\ *[\[\]\(\)\{\}]*\ *)`,
		"replace": `htt$1`,
	},
	{
		"find":    `(?:[\[\]\(\)\{\}]*\ *)htt(ps?)(?:\ *[\[\]\(\)\{\}]+\ *)`,
		"replace": `htt$1`,
	},
	{
		"find":    `h[xA-Z]{2}(ps?[^.])`,
		"replace": `htt$1`,
	},
	{
		"find":    `htt(ps?)\/`,
		"replace": `htt$1:/`,
	},
	{
		"find":    `(https?:\/\/)\ *`,
		"replace": `$1`,
	},
	{
		"find":    `(https?)\ *:`,
		"replace": `$1:`,
	},
	{
		"find":    `\b(x{4}:\/\/)`,
		"replace": `http://`,
	},
	{
		"find":    `\b((?:x{5}|x{4}s):\/\/)`,
		"replace": `https://`,
	},
	{
		"find":    `(\[-\])`,
		"replace": "-",
	},
	{
		"find":    `(\ +@\ +)`,
		"replace": "@",
	},
	{"find": `(\^.)`, "replace": "."},
	{"find": `(<\.>)`, "replace": "."},
}
