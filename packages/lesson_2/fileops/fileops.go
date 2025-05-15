package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const valueFile = "balance.txt"

func WriteFloatToFile(value float64) {

	valueTxt := fmt.Sprint(value)
	os.WriteFile(valueFile, []byte(valueTxt), 0644)

}

func GetFloatFromFile() (float64, error) {

	data, err := os.ReadFile(valueFile)

	if err != nil {

		return 1000, errors.New("failed to find balance file")
	}
	valueTxt := string(data)
	value, err := strconv.ParseFloat(valueTxt, 64)

	if err != nil {

		return 1000, errors.New("failed to parse balance file")
	}

	return value, nil
}
