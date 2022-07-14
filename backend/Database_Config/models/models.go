package models

import (
    "gorm.io/gorm"
    )

type User struct {
    UserId int `gorm:"primary_key;auto_increment;not_null" json:"user_id"`
    Username string `json:"username" binding:"required"`
    HashPassword string `json:"password" binding:"required"`
    Role string `json:"user_role" binding:"required"`
    CreatedAt int64 `gorm:"autoCreateTime:milli" json:"created_at"`
    UpdateAt int64 `gorm:"autoCreateTime:milli" json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Product struct {
    ProductId int `gorm:"primary_key:auto_increment;not_null" json:"product_id"`
    ProductName string `json:"product_name" binding:"required"`
    ProductDescription string `json:"product_description" binding:"required"`
    ProductValue int `json:"product_value" binding:"required"`
}

