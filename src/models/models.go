package models

// import modules
import (
    "gorm.io/gorm"
    )

// create user id struct
type UserId struct {
    Id string `uri:"id" binding:"required"`
}

// create user struct
type User struct {
    Id int `gorm:"primary_key;auto_increment;not_null" json:"id"`
    FirstName string `json:"firstname" binding:"required"`
    LastName string `json:"lastname" binding:"required"`
    CreatedAt int64 `gorm:"autoCreateTime:milli" json:"created_at"`
    UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

//func (u *User) FillDefaults(){
//    if u.Id == "" {
//        u.Id = uuid.NewString()
//    }
//}
