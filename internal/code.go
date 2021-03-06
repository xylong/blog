package internal

const (
	SUCCESS       = 0
	InvalidParams = 400
	ERROR         = 10000

	RedisError = 10001

	UNAUTHORIZED     = 20001
	TokenExpired     = 20002
	TokenNotValidYet = 20003
	TokenMalformed   = 20004
	TokenInvalid     = 20005

	TagGetFail    = 40001
	TagExist      = 40002
	TagAddFail    = 40003
	TagUpdateFail = 40004
	TagDeleteFail = 40005

	CategoryGetFail    = 50001
	CategoryExist      = 50002
	CategoryAddFail    = 50003
	CategoryUpdateFail = 50004
	CategoryDeleteFail = 50005

	ArticleGetFail    = 60001
	ArticleAddFail    = 60002
	ArticleUpdateFail = 60003
	ArticleDeleteFail = 60004

	UploadInvalidFile = 70001
	UploadFormat      = 70002
	UploadSize        = 70003
	UploadSave        = 70004
)
