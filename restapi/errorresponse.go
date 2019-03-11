package restapi

const (
	unknownError           = 1000
	notAuthorized          = 2000
	invalidCredentials     = 2001
	invalidInputFormat     = 3000
	accessDenied           = 4000
	missingField           = 5000
	dbError                = 6000
	volunteerAlreadyExists = 7000
	teamAlreadyExists      = 7001
	teamAlreadyJoined      = 7002
	networkError           = 8000
)

type ErrorJson struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

// todo handel unknown code ...
var volunteerErrorResponses = map[uint32]string{
	unknownError:           "unknown error",
	notAuthorized:          "not authorized",
	invalidCredentials:     "invalid email/password",
	accessDenied:           "access denied",
	invalidInputFormat:     "invalid input format",
	missingField:           "missing field",
	dbError:                "database error",
	networkError:           "db connection error",
	volunteerAlreadyExists: "volunteer with this email already exists",
	teamAlreadyExists:      "team with this name already exists",
	teamAlreadyJoined:      "team already joined"}

func getErrorJson(errorCode uint32) *ErrorJson {
	return &ErrorJson{Code: errorCode, Message: volunteerErrorResponses[errorCode]}
}
