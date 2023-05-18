package errno

var (
	ServiceError   = New(ServiceErrorCode, "service internal error")
	ParamError     = New(ParamErrorCode, "parameter error")
	ExecuteTimeout = New(ExecuteTimeoutCode, "executed timeout")

	AuthorizationExpired = New(AuthorizationExpiredCode, "authorization has expired")
	AuthorizationFailed  = New(AuthorizationFailedCode, "authorization failed")
	DuplicatedEmail      = New(DuplicatedEmailCode, "duplicated email")
	InvalidEmailFormat   = New(InvalidEmailFormatCode, "invalid email format")
	UserNotFound         = New(UserNotFoundCode, "user not found")
	WrongPassword        = New(WrongPasswordCode, "wrong password")
)
