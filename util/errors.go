package util

import "errors"

var (
	ErrEmptyToken   = errors.New("token值为空")
	ErrExpiredToken = errors.New("令牌失效")
	// TagRetErrStru 用户中心返回错误
	TagRetErrStru error = errors.New("get retErrStru form user center")
	// ErrNotFound 未找到
	ErrNotFound error = errors.New("not found")
	// ErrFound 错误查找
	ErrFound error = errors.New("found")
	// ErrNotImpl 未启用
	ErrNotImpl error = errors.New("not implementation")
	// ErrUnSurported 不支持
	ErrUnSurported error = errors.New("unsurported")
	// ErrInvalidFormat 无效格式
	ErrInvalidFormat error = errors.New("invalid format")
	// ErrRespMsg 返回结果出错
	ErrRespMsg error = errors.New("respMsg.ResultStru.Code != 200")
	// Err401 401错误
	Err401 error = errors.New("401 from provider")
	// Err404 404错误
	Err404 error = errors.New("404 from provider")
	// Err500 500错误
	Err500 error = errors.New("500 from provider")
	// ErrPost post方法中的错误
	ErrPost error = errors.New("err from Post()")
	// ErrAccessQuick 登录过快
	ErrAccessQuick error = errors.New("interface access quickly")
	// ErrWeChatLink 未连上微信公众号
	ErrWeChatLink error = errors.New("wechat public account not link")
	// ErrParamNull 参数为空
	ErrParamNull error = errors.New("param null")
	// ErrParamWrong 参数错误
	ErrParamWrong error = errors.New("param wrong")
	// ErrServResult server result错误
	ErrServResult error = errors.New("server result error")
	// ErrNotLogin 用户未登录
	ErrNotLogin error = errors.New("user not login")
	// ErrOpenIDNull 微信openid为空
	ErrOpenIDNull error = errors.New("wechat openid is null")
	// ErrDBAccess 数据库连接错误
	ErrDBAccess error = errors.New("database access error")
	// ErrNoBinding openid未绑定
	ErrNoBinding error = errors.New("openid no binding")
	// ErrGetUCAccToken 获取用户中心访问令牌失败
	ErrGetUCAccToken error = errors.New("failed to get AccessToken of userCenter")
	// ErrInvalidWebOAuthToken 无效的web oauth令牌
	ErrInvalidWebOAuthToken error = errors.New("invalid web oauth token")
	// ErrGetWebOauthAccessToken 获取web oauth2.0访问令牌失败
	ErrGetWebOauthAccessToken error = errors.New("failed to get accessToken of web oauth2.0")
	// ErrRefreshToken 刷新访问令牌失败
	ErrRefreshToken error = errors.New("failed to refresh access_token")
	// ErrUIDNull 管理员uid不存在
	ErrUIDNull error = errors.New("manager uid is null")
	// ErrRefuseAuthorization 微信用户无权限
	ErrRefuseAuthorization error = errors.New("weixin user refused to authorize")
	// ErrGetCode 获取web oauth2.0 Code失败
	ErrGetCode error = errors.New("failed to get code of web oauth2.0")
	// ErrTrieNoInit Trie未初始化
	ErrTrieNoInit error = errors.New("Trie not init")
	// ErrFileWriteType 无效的文档写入类型
	ErrFileWriteType error = errors.New("Invaild File write type")
	// ErrPhoneInvalid 无效的电话
	ErrPhoneInvalid error = errors.New("phone si invalid")
	// ErrGetUserInfo 获取用户信息失败
	ErrGetUserInfo error = errors.New("failed to get userInfo")
	// ErrOperationDB 操作数据库错误
	ErrOperationDB error = errors.New("operation db failed")
	// ErrAddIntegral 添加用户积分情况失败
	ErrAddIntegral error = errors.New("failed to add integral")
	// ErrDuplicateData 添加数据重复
	ErrDuplicateData error = errors.New("duplicate data")
	// ErrFetch 获取媒资服务器数据失败
	ErrFetchVodData error = errors.New("fetch vod data failed")
)
