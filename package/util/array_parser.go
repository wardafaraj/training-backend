package util

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseStringToInt32Array(data string) []int32 {
	// Split the string into individual integer tokens
	numbersStr := strings.Split(data, ",")
	// Slice to store parsed integers
	var numbers []int32
	// Iterate over each token and parse it as an integer
	for _, numStr := range numbersStr {
		num, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing integer %s: %v\n", numStr, err)
			return []int32{}
		}
		numbers = append(numbers, int32(num))
	}
	return numbers
}
func ParseStringToStringArray(data string) []string {
	// Split the string into individual integer tokens
	stringsStr := strings.Split(data, ",")
	// Slice to store parsed integers
	var strings []string
	return append(strings, stringsStr...)
}

func Int32ToStringSlice(intArray []int32) []string {
	stringArray := make([]string, len(intArray))
	for i, num := range intArray {
		stringArray[i] = strconv.Itoa(int(num)) // Convert int32 to int, then to string
	}
	return stringArray
}
