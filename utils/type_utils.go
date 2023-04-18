package utils

import (
	"errors"
	"fmt"
)

func Float64ToUint64(f float64) (uint64, error) {
	if f < 0 {
		return 0, errors.New(fmt.Sprintf("cannot convert negative value %f to uint64", f))
	}
	return uint64(f), nil
}

func IntToUint64(i int) (uint64, error) {
	if i < 0 {
		return 0, fmt.Errorf("intToInt64: cannot convert negative value %d to uint64", i)
	}
	return uint64(i), nil
}
