package utils

import (
	"errors"
	"eyesStars/app/common"
	"strconv"
)

/**
 * @author eyesYeager
 * @date 2023/1/30 22:20
 */

func String2UintSlice(slice []string) (err error, result []uint) {
	result = make([]uint, len(slice))
	for i, v := range slice {
		uintV, err := strconv.Atoi(v)
		if err != nil {
			return common.CustomError{}.SetErrorMsg("Wrong parameter type, cannot be converted to uint type"), result
		}
		result[i] = uint(uintV)
	}
	return err, result
}

func Uint64ToInt64Slice(slice []uint64) (err error, result []int64) {
	defer func() {
		cErr := recover()
		if cErr != nil {
			err = errors.New("uint64 cannot be converted to int64")
		}
	}()
	result = make([]int64, len(slice))
	for _, v := range slice {
		result = append(result, int64(v))
	}
	return err, result
}
