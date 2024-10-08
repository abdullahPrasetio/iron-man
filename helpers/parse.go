package helpers

import "time"

func IsInt(value interface{}) bool {
	_, ok := value.(int)
	return ok
}

func IsFloat64(value interface{}) bool {
	_, ok := value.(float64)
	return ok
}

func IsString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func IsIntAndGetValue(data interface{}) (result int) {
	value, ok := data.(int)
	if ok {
		result = value
	} else {
		result = 0
	}
	return
}

func IsStringAndGetValue(data interface{}) (result string) {
	value, ok := data.(string)
	if ok {
		result = value
	}
	return
}

func IsFloat64AndGetValue(data interface{}) (result float64) {
	value, ok := data.(float64)
	if ok {
		result = value
	} else {
		result = 0
	}
	return
}

func IsMapStringString(data interface{}) (result map[string]string) {
	value, ok := data.(map[string]string)
	if ok {
		result = value
	}
	return
}

func IsMapStringStringGetValue(data interface{}) (result map[string]string) {
	value, ok := data.(map[string]string)
	if ok {
		result = value
	}
	return
}

func IsArrayStringGetValue(data interface{}) (result []string) {
	value, ok := data.([]string)
	if ok {
		result = value
	}
	return
}

func IsTimeGetValue(data interface{}) (result time.Time) {
	value, ok := data.(time.Time)
	if ok {
		result = value
	}
	return
}

func IsMapStringAnyGetValue(data interface{}) (result map[string]any) {
	value, ok := data.(map[string]interface{})
	if ok {
		result = value
	}
	return
}
