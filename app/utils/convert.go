package utils

import (
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
