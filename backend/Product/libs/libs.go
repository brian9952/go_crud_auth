package libs

import (
    "os"
    )

var (
    Auth_key = os.Getenv("AUTH_SECRET")
    Gateway_api_key = os.Getenv("PRODUCT_API_SECRET")
)

type Status struct { // 0 success, 1 frontend, 2 backend
    StatusType int `json:"status_type"`
    StatusMessage string `json:"status_message"`
    ProductId int `json:"product_id"`
}

func CreateAddProductMessage(status_type int, status_message string, product_id int) *Status { // 0 success, 1 frontend error, 2 internal error
    var s *Status = new(Status)
    s.StatusType = status_type
    s.StatusMessage = status_message
    s.ProductId = product_id
    return s
}

func CreateDeleteProductMessage(status_type int, status_message string, product_id int) *Status {
    var s *Status = new(Status)
    s.StatusType = status_type
    s.StatusMessage = status_message
    s.ProductId = product_id
    return s
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
