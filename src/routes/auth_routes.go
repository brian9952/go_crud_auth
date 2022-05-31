package routes

import (
	"encoding/json"
	"log"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
    "golang.org/x/crypto/bcrypt"
)

func GetHashPassword(pass string) (string, error){
    bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
    return string(bytes), err
}

type LoginDetails struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Error struct {
    ErrorType string `json:"errortype"`
    Message string `json:"message"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
    db, connErr := database.GetDatabaseConnection()

    if connErr != nil {
        var err Error
        err.ErrorType = "db_conn"
        err.Message = "Error: Unable to connect to the database"
        w.Header().Set("Content-type", "application/json")
        json.NewEncoder(w).Encode(err)
        return
    }

    var newuser models.User
    err := json.NewDecoder(r.Body).Decode(newuser)
    if err != nil {
        var err Error
        err.ErrorType = "json_decoding"
        err.Message = "Error decoding the data"
        w.Header().Set("Content-type", "application/json")
        json.NewEncoder(w).Encode(err)
        return
    }

    // generate hash password
    newuser.HashPassword, err = GetHashPassword(newuser.HashPassword)
    if err != nil {
        log.Default().Println("Error in password hashing")
        return
    }

    // insert user
    db.Create(&newuser)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newuser)

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
    db, connErr := database.GetDatabaseConnection()

    if connErr != nil {
        var err Error
        err.Message = "Error: Unable to connect to the database"
        w.Header().Set("Content-type", "application/json")
        json.NewEncoder(w).Encode(err)
        return
    }

    var loginDetails LoginDetails

    err := json.NewDecoder(r.Body).Decode(&loginDetails)
    if err != nil {
        var 
    }

}
