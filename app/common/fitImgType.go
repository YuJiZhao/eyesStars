package common

/**
 * @author eyesYeager
 * @date 2023/2/1 9:33
 */

var fitImgName = make(map[string]bool)

func init() {
	fitImgName[".jpg"] = true
	fitImgName[".png"] = true
	fitImgName[".jpeg"] = true
}

func FitImgType(extendedName string) bool {
	return fitImgName[extendedName]
}
