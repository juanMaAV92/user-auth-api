package enviroment

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnvAsIntWithDefault(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("Error during env '%s' conversion to int", key))
	}

	return valueInt
}
