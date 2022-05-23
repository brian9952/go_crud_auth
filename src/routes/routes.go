package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"main/database"
	"main/models"
)

type messageSuccess struct {
    Message string `json:"message"`
    Id int `json:"userId"`
}

type idStruct struct {
    Id int `json:"id"`
}

func CreateUser(w http.ResponseWriter, r *http.Request){
    var user models.User
    var err error

    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Default().Fatal("create user error")
    }

    json.Unmarshal(reqBody, &user)

    // get db connection
    db, connErr := database.GetDatabaseConnection()

    if connErr != nil {
        log.Default().Panic("Error occured while connecting to the database")
        w.WriteHeader(http.StatusServiceUnavailable)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode("Service is unavailable")
        return
    }

    result := db.Create(&user)

    if result.Error != nil && result.RowsAffected != 1 {
        log.Default().Panic("Error occured while creating new user")
        json.NewEncoder(w).Encode("Error occured while creating new user")
        return
    }

    message := messageSuccess{
        "Successfully creating a new user",
        user.Id,
    }
    json.NewEncoder(w).Encode(message)

}
func ShowUser(w http.ResponseWriter, r *http.Request){
    var user models.User
    var id idStruct
    var err error

    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Default().Fatal("Create user error")
    }

    json.Unmarshal(reqBody, &id)
    log.Default().Panic(id)

    // get db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        log.Default().Panic("Error occured while connecting to the database")
        w.WriteHeader(http.StatusServiceUnavailable)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode("Service is unavailable")
    }

    result := db.First(&user, id)
    log.Default().Panic(user)

    if result.Error != nil {
        log.Default().Panic("Error occured while connection to the database")
        json.NewEncoder(w).Encode("Error occured while showing user")
    }

    json.NewEncoder(w).Encode(user)

}
