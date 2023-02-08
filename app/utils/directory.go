package utils

import "os"

/**
 * @author eyesYeager
 * @date 2023/1/28 17:14
 */

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
