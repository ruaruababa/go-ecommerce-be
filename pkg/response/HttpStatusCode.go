package response

const (
	ErrorCodeSuccess = 200
	ErrorCodeBadRequest = 400
	ErrorCodeUnauthorized = 401
	ErrorCodeForbidden = 403
	ErrorCodeNotFound = 404
	ErrorCodeInternalServerError = 500
	ErrorCodeServiceUnavailable = 503
	ErrorCodeGatewayTimeout = 504
	ErrorCodeWrongInformation= 1000
)

var msg = map[int]string{
	ErrorCodeSuccess: "Success",
	ErrorCodeBadRequest: "Bad Request",
	ErrorCodeUnauthorized: "Unauthorized",
	ErrorCodeForbidden: "Forbidden",
	ErrorCodeNotFound: "Not Found",
	ErrorCodeInternalServerError: "Internal Server Error",
	ErrorCodeServiceUnavailable: "Service Unavailable",
	ErrorCodeGatewayTimeout: "Gateway Timeout",
	ErrorCodeWrongInformation: "Wrong Infomation",
}