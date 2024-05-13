package apperrors

type ErrCode string

const (
	Unkonwn          ErrCode = "U000"
	InsertDataFailed ErrCode = "S001"
	Nodata           ErrCode = "S002"
	GetDataFailed    ErrCode = "S003"
	DeleteDataFailed ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"
)

func (e ErrCode) Wrap(err error, message string) error {
	return &TodoAppError{
		ErrCode: e,
		Message: message,
		Err:     err,
	}
}
