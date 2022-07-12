package response

type ErrorCode int32

const (
	ErrorCodeSuccess ErrorCode = iota

	ErrorCodeFailed
	ErrorCodeInvalidParams
	ErrorCodeInvalidAuth
	ErrorCodePermissionDenied
	ErrorCodeResourceNotExist
	ErrorCodeInternalServerError
	ErrorCodeDatabaseError
	ErrorCodeInternalAPIError
	ErrorCodeExternalAPIError
	ErrorCodeLicenseExpire
)
