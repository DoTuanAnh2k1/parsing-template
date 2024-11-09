package parsing

import "regexp"

func checkOption(template string) string {
	re := regexp.MustCompile(`\[(.*?)\]`)

	template = re.ReplaceAllStringFunc(template, func(s string) string {
		if s == "[]" {
			return ""
		}
		return s
	})

	return template
}
