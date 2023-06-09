package errno

type ErrNoCode int64

const (
	SuccessCode ErrNoCode = 10000 + iota

	// Service
	ServiceErrorCode
	ParamErrorCode
	ExecuteTimeoutCode

	// User
	AuthorizationExpiredCode
	AuthorizationFailedCode
	DuplicatedEmailCode
	InvalidEmailFormatCode
	UserNotFoundCode
	WrongPasswordCode

	// Chat
	UnknownMessageTypeCode
	UpgradeFailedCode

	// Friend
	AlreadyAppliedFriendshipCode
	ApplicationsNotFoundCode
)

const (
	SuccessMsg string = "ok"
)
