package models

type Product struct {
    ProductId int `gorm:"primary_key:auto_increment;not_null" json:"product_id"`
    ProductName string `json:"product_name" binding:"required"`
    ProductDescription string `json:"product_description" binding:"required"`
    ProductValue int `json:"product_value" binding:"required"`
}
