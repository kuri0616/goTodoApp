package apperrors

type ErrCode string

const (
	Unkonwn          ErrCode = "U000"
	InsertDataFailed ErrCode = "S001"
	Nodata           ErrCode = "S002"
	GetDataFailed    ErrCode = "S003"
	DeleteDataFailed ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"

	RequireAuthorizationHeader ErrCode = "A001"
	CanNotMakeValidator        ErrCode = "A002"
	IllegalToken               ErrCode = "A003"
	NotMatchUser               ErrCode = "A004"
)

func (e ErrCode) Wrap(err error, message string) error {
	return &TodoAppError{
		ErrCode: e,
		Message: message,
		Err:     err,
	}
}
