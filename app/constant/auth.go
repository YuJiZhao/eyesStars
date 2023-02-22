package constant

/**
 * @author eyesYeager
 * @date 2023/2/6 14:44
 */

type auth struct {
	AuthFlag string // token是否存在的标识
	SToken   string // 短token名称
	LToken   string // 长token名称
	CUid     string // Context携带用户id的key
	CRole    string // Context携带用户角色的key
}

var Auth = auth{
	AuthFlag: "flag",
	SToken:   "sAuth",
	LToken:   "lAuth",
	CUid:     "uid",
	CRole:    "role",
}
