package internal

import (
	"errors"
	"os"
	"strconv"
)

var ErrNotExist = errors.New("value does not exist or not found")

func GetString(variable string) (string, error) {
	envString, ok := os.LookupEnv(variable)
	if !ok {
		return "", ErrNotExist
	}
	return envString, nil
}

func GetInt(variable string) (int, error) {
	envString, ok := os.LookupEnv(variable)
	if !ok {
		return -1, ErrNotExist
	}
	envInt, err := strconv.Atoi(envString)
	if err != nil {
		return -1, err
	}
	return envInt, nil
}
