package internal

var Message = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	RedisError: "Redis异常",

	UNAUTHORIZED:     "未认证",
	TokenExpired:     "令牌过期",
	TokenNotValidYet: "令牌尚未激活",
	TokenMalformed:   "非法令牌",
	TokenInvalid:     "无效令牌",

	TagGetFail:    "标签获取失败",
	TagExist:      "标签已存在",
	TagAddFail:    "标签添加失败",
	TagUpdateFail: "标签修改失败",
	TagDeleteFail: "标签删除失败",

	CategoryGetFail:    "分类获取失败",
	CategoryExist:      "分类已存在",
	CategoryAddFail:    "分类创建失败",
	CategoryUpdateFail: "分类修改失败",
	CategoryDeleteFail: "分类删除失败",

	ArticleGetFail:    "文章获取失败",
	ArticleAddFail:    "文章创建失败",
	ArticleUpdateFail: "文章修改失败",
	ArticleDeleteFail: "文章删除失败",
}

// GetMsg 获取错误信息
func GetMsg(code int) string {
	msg, ok := Message[code]
	if ok {
		return msg
	}
	return Message[ERROR]
}
