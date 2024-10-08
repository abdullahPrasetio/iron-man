package helpers

import "strings"

func SnakeToCamelWithSeparator(s string, separator string) string {
	parts := strings.Split(s, separator)

	for i := 1; i < len(parts); i++ {
		parts[i] = strings.ToUpper(string(parts[i][0])) + strings.ToLower(parts[i][1:])
	}

	return strings.Join(parts, separator)
}

func IncludeString(text string, data []string) bool {
	found := false
	for _, validCode := range data {
		if text == validCode {
			found = true
			break
		}
	}

	return found
}

func GetString(m map[string]interface{}, key string) string {
	if val, ok := m[key].(string); ok {
		return val
	}
	return ""
}
