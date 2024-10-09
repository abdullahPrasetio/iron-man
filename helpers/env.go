package helpers

import "os"

func GetEnvWithDefaultString(key, data string) (value string) {
	value = data
	if os.Getenv(key) != "" {
		value = os.Getenv(key)
	}

	return
}
