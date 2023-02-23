package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

// this file content read env function

/*
GetInt64 Get config type int64 from env, config file in order
*/
func GetInt64(camelKey string, def int64) int64 {
	v := os.Getenv(camelKey)
	if result, err := strconv.ParseInt(v, 10, 64); err == nil {
		return result
	}
	return def
}

/*
GetString Get config type string from env, config file in order
*/
func GetString(camelKey string, def string) string {
	result := os.Getenv(camelKey)
	if len(result) > 0 {
		return result
	}
	return def
}

/*
GetBool Get config type bool from env, config file in order
*/
func GetBool(camelKey string, def bool) bool {
	result := os.Getenv(camelKey)
	if b, err := strconv.ParseBool(result); len(result) > 0 && err == nil {
		return b
	}
	return def
}
