package constant

/**
 * @author eyesYeager
 * @date 2023/2/6 14:44
 */

// 用户角色名称
type role struct {
	Admin   string
	User    string
	Visitor string
}

var Roles = role{
	Admin:   "ROLE_admin",
	User:    "ROLE_user",
	Visitor: "",
}
