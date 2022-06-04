package functions

type Error struct {
    ErrorType string `json:"errortype"`
    Message string `json:"message"`
}

func CreateError(errorType string, message string) *Error {
    var err *Error = new(Error)
    err.ErrorType = errorType
    err.Message = message
    return err
}
