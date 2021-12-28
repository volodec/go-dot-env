package go_dot_env

import (
	"fmt"
	"strconv"
)

func get(key string, defValue interface{}) interface{} {
	value := SearchKey(key)
	if value == "" {
		return defValue
	}
	return value
}

func String(key string, defVal string) string {
	value := get(key, defVal)

	if value == "" {
		return defVal
	}

	return fmt.Sprint(value)
}

func Int(key string, defVal int) int {
	value := String(key, "")
	if value == "" {
		return defVal
	}

	result, err := strconv.Atoi(value)

	if err != nil {
		return defVal
	}

	return result
}

func Bool(key string, defVal bool) bool {
	value := String(key, "")

	if value == "" {
		return defVal
	}

	if value == "true" {
		return true
	}

	return false
}
