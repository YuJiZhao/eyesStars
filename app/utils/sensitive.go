package utils

import (
	"errors"
	"strings"
)

/**
 * @author eyesYeager
 * @date 2023/2/18 13:42
 */

// SensitiveEmail 邮箱脱敏(例：e*********@163.com)
func SensitiveEmail(email string) (err error, result string) {
	split := strings.Split(email, "@")
	if len(split) != 2 {
		return errors.New("bad mailbox format"), result
	}
	result = string(split[0][0])
	for i := 1; i < len(split[0]); i++ {
		result += "*"
	}
	result += "@" + split[1]
	return err, result
}
