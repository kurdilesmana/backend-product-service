package constants

// Error Type
const (
	Err                 = "error"
	ErrGeneral          = "error general"
	ErrDataNotFound     = "Data not found"
	ErrDataAlreadyExist = "error data already exist"
	ErrFormatting       = "error formatting"
	ErrDatabase         = "error database"
	ErrTimeout          = "error timeout"
	ErrAuth             = "error authorization"
	ErrPanic            = "error panic"
)

// Error type ErrDatabase
const (
	ErrWhenBeginTx                      = "error when begin transaction"
	ErrWhenExecuteQuery                 = "error when execute query"
	ErrWhenScanResults                  = "error when scan results"
	ErrWhenSelectDB                     = "error when select data from db"
	ErrWhenUpdateDB                     = "error when update data to db"
	ErrWhenDeleteDB                     = "error when delete data from db"
	ErrWhenCommitDB                     = "error when commit db"
	ErrRollBack                         = "error when rollback data from db"
	ErrDuplicateData                    = "error when insert data - dup data"
	ErrAlreadyExistData                 = "error when insert data - already exist"
	ErrWhenCheckingBeforeDeleteDataTkpu = "error when delete data tkpu - already exist"
)

const (
	ErrConnectionToDB     = "failed to connect to db: %w"
	ErrLoadENV            = "failed to load .env file: %w"
	ErrConvertStringToInt = "error when convert string to int: %w"
)

const (
	ErrBadRequestDescription                    = "400 Bad Request"
	ErrUnauthorizedDescription                  = "401 Unauthorized (RFC 7235)"
	ErrPaymentRequiredDescription               = "402 Payment Required"
	ErrForbiddenDescription                     = "403 Forbidden"
	ErrNotFoundDescription                      = "404 Not Found"
	ErrMethodNotAllowedDescription              = "405 Method Not Allowed"
	ErrNotAcceptableDescription                 = "406 Not Acceptable"
	ErrProxyAuthenticationRequiredDescription   = "407 Proxy Authentication Required (RFC 7235)"
	ErrRequestTimeoutDescription                = "408 Request Timeout"
	ErrConflictDescription                      = "409 Conflict"
	ErrGoneDescription                          = "410 Gone"
	ErrLengthRequiredDescription                = "411 Length Required"
	ErrPreconditionFailedDescription            = "412 Precondition Failed (RFC 7232)"
	ErrPayloadTooLargeDescription               = "413 Payload Too Large (RFC 7231)"
	ErrURITooLongDescription                    = "414 URI Too Long (RFC 7231)"
	ErrUnsupportedMediaTypeDescription          = "415 Unsupported Media Type (RFC 7231)"
	ErrRangeNotSatisfiableDescription           = "416 Range Not Satisfiable (RFC 7233)"
	ErrExpectationFailedDescription             = "417 Expectation Failed"
	ErrImATeapotDescription                     = "418 I'm a teapot (RFC 2324, RFC 7168)"
	ErrMisdirectedRequestDescription            = "421 Misdirected Request (RFC 7540)"
	ErrUnprocessableEntityDescription           = "422 Unprocessable Entity (WebDAV; RFC 4918)"
	ErrLockedDescription                        = "423 Locked (WebDAV; RFC 4918)"
	ErrFailedDependencyDescription              = "424 Failed Dependency (WebDAV; RFC 4918)"
	ErrTooEarlyDescription                      = "425 Too Early (RFC 8470)"
	ErrUpgradeRequiredDescription               = "426 Upgrade Required"
	ErrPreconditionRequiredDescription          = "428 Precondition Required (RFC 6585)"
	ErrTooManyRequestsDescription               = "429 Too Many Requests (RFC 6585)"
	ErrRequestHeaderFieldsTooLargeDescription   = "431 Request Header Fields Too Large (RFC 6585)"
	ErrUnavailableForLegalReasonsDescription    = "451 Unavailable For Legal Reasons (RFC 7725)"
	ErrInternalServerErrorDescription           = "500 Internal Server Error"
	ErrNotImplementedDescription                = "501 Not Implemented"
	ErrBadGatewayDescription                    = "502 Bad Gateway"
	ErrServiceUnavailableDescription            = "503 Service Unavailable"
	ErrGatewayTimeoutDescription                = "504 Gateway Timeout"
	ErrHTTPVersionNotSupportedDescription       = "505 HTTP Version Not Supported"
	ErrVariantAlsoNegotiatesDescription         = "506 Variant Also Negotiates (RFC 2295)"
	ErrInsufficientStorageDescription           = "507 Insufficient Storage (WebDAV; RFC 4918)"
	ErrLoopDetectedDescription                  = "508 Loop Detected (WebDAV; RFC 5842)"
	ErrNotExtendedDescription                   = "510 Not Extended (RFC 2774)"
	ErrNetworkAuthenticationRequiredDescription = "511 Network Authentication Required (RFC 6585)"
)
