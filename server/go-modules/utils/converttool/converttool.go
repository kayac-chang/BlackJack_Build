package converttool

import (
	"fmt"
	"strconv"
	"strings"
)

func GetStringToInt(source string, defaultValue int) int {
	res, err := strconv.Atoi(source)
	if err != nil {
		return defaultValue
	}
	return res
}

func GetStringToBool(source string, defaultValue bool) bool {
	res, err := strconv.ParseBool(source)
	if err != nil {
		return defaultValue
	}
	return res
}

func GetStringToFloat64(source string, defaultValue float64) (float64, error) {
	res, err := strconv.ParseFloat(source, 64)
	if err != nil {
		return defaultValue, err
	}
	return res, err
}

func IntArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
