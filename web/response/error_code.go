package response

type ErrorCode int32

const (
	ErrorCodeSuccess ErrorCode = 0
	ErrorCodeFailed  ErrorCode = 1
)
