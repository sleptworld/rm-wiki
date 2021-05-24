package Config

const (
	// ValueFormat

	ErrValueFormat = "10001"
	MsgValueFormatForID = "ID should be a number."

	ErrNoValue = "10002"
	MsgNoValueForID = "Can't find the result,check 'id' please."

	ErrBodyFormat = "20001"
	MsgBodyFormat = "Request body format error."

	ErrBodyValueMissing = "20002"
	MsgBodyValueMissing = "Missing attributes."

	ErrBodyRegistered = "20003"
	MsgBodyValueRegistered = "Have registered."

	ErrWrongData = "30000"
	MsgWrongData = "Wrong format data."

	ErrUnauthorized = "40001"
	MsgUnauthorized = "You don't have enough permissions"
)
