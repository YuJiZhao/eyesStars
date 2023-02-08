package constant

/**
 * @author eyesYeager
 * @date 2023/1/30 17:55
 */

type starStatus struct {
	Normal uint8
	Frozen uint8
	Delete uint8
}

var StarStatus = starStatus{
	0,
	1,
	2,
}
