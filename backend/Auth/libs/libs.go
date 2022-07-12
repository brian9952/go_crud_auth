package libs

import "os"

var Private_Key = os.Getenv("AUTH_API_SECRET")

type Status struct {
    statusType int // -1 error, 0 warning, 1 success
    statusMessage string
}

func CreateErrorMessage(message string) *Status {
    var s *Status = new(Status)
    s.statusType = -1
    s.statusMessage = message
    return s
}
