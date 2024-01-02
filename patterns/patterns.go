package patterns

var DefangPatterns = []map[string]string{
	{
		"find":    `(https?://\S+)`,
		"change":  "https:",
		"replace": "hXXps:",
	},
	{
		"find":    `^(http:\/\/)`,
		"change":  "http:",
		"replace": "hXXp:",
	},
	{
		"find":    `([a-zA-Z0-9_.+-]+)@([a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+)`,
		"change":  "@",
		"replace": "[AT]",
	},
	{
		"find":    `(\d+\.\d+\.\d+\.\d+)([^0-9])`,
		"change":  ".",
		"replace": "[.]",
	},
	{
		"find":    `\d+\.\d+\.\d+\.(\d+)`,
		"change":  ".",
		"replace": "[.]",
	},
	{
		"find":    `[^.]\.([a-zA-Z0-9])`,
		"change":  ".",
		"replace": "[.]",
	},
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
