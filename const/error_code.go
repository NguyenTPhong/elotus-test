package _const

const (
	Unauthorized      = "unauthorized"
	StatusBadRequest  = "invalid request"
	InternalServerErr = "internal server error"

	UserNotFound      = "wrong username"
	UserAlreadyExist  = "username already taken"
	UserWrongPassword = "wrong password"

	FailedToOpenFile = "failed to open file"
	InvalidFileType  = "invalid file type"
	FileTooLarge     = "file too large"
)

var ErrorCode = map[string]int{
	Unauthorized:      401,
	StatusBadRequest:  400,
	InternalServerErr: 503,

	UserNotFound:      1000,
	UserAlreadyExist:  1001,
	UserWrongPassword: 1002,

	FailedToOpenFile: 2000,
	InvalidFileType:  2001,
	FileTooLarge:     2002,
}
