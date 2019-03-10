package restapi

const (
	unknownError       = 1000
	notAuthorized      = 2000
	invalidCredentials = 2001
	invalidInputFormat = 3000
	accessDenied       = 4000
	missingField       = 5000
	dbError            = 6000
	alreadyExists      = 7000
)

type ErrorJson struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

// todo handel unknown code ...
var volunteerErrorResponses = map[uint32]string{
	unknownError:       "unknown error",
	notAuthorized:      "not authorized",
	invalidCredentials: "invalid email/password",
	accessDenied:       "access denied",
	invalidInputFormat: "invalid input format",
	missingField:       "missing field",
	dbError:            "database error",
	alreadyExists:      "already exists"}

func getErrorJson(errorCode uint32) *ErrorJson {
	return &ErrorJson{Code: errorCode, Message: volunteerErrorResponses[errorCode]}
}
