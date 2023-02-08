package result

/**
 * @author eyesYeager
 * @date 2023/2/6 13:58
 */

type customResult struct {
	code int
	msg  string
}

type customResults struct {
	DefaultSuccess   customResult
	DefaultFail      customResult
	ValidateError    customResult
	AuthError        customResult
	ServerError      customResult
	AuthInsufficient customResult
}

var Results = customResults{
	DefaultSuccess:   customResult{200, "success"},
	DefaultFail:      customResult{400, "fail"},
	ServerError:      customResult{10000, "服务器似乎出问题了"},
	ValidateError:    customResult{20000, "请求参数错误"},
	AuthError:        customResult{20001, "权限错误"},
	AuthInsufficient: customResult{20001, "权限不足"},
}
