package common

/**
 * @author eyesYeager
 * @date 2023/1/30 18:04
 */

type CustomError struct {
	message string
}

func (customError CustomError) Error() string {
	return customError.message
}

func (customError CustomError) SetErrorMsg(msg string) CustomError {
	customError.message = msg
	return customError
}
