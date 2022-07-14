package libs

import (
    "os"
    )

var Private_Key = os.Getenv("AUTH_API_SECRET")

type Status struct {
    StatusType int `json:"status_type"`// -1 error, 0 warning, 1 success
    StatusMessage string `json:"status_message"`
}

func CreateErrorMessage(message string) *Status {
    var s *Status = new(Status)
    s.StatusType = -1
    s.StatusMessage = message
    return s
}

func CreateSuccessMessage(message string) *Status {
    var s *Status = new(Status)
    s.StatusType = 1
    s.StatusMessage = message
    return s
}