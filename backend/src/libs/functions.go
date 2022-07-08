package functions

// error message
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

// success message
type Success struct {
    SuccessType string `json:"successtype"`
    Message string `json:"message"`
}

func CreateSuccess(successType string, message string) *Success {
    var succ *Success = new(Success)
    succ.SuccessType = successType
    succ.Message = message
    return succ
}
