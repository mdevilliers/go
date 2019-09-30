package env

import (
	"os"
	"strconv"
)

// FromEnvWithDefaultStr return the found value or the default
func FromEnvWithDefaultStr(name string, defaultValue string) string {

	value, ok := os.LookupEnv(name)

	if ok {
		return value
	}

	return defaultValue
}

// FromEnvWithDefaultBool return the found value or the default
func FromEnvWithDefaultBool(name string, defaultValue bool) bool {
	value, ok := os.LookupEnv(name)

	if ok {

		b, err := strconv.ParseBool(value)

		if err != nil {
			return defaultValue
		}
		return b
	}

	return defaultValue

}
