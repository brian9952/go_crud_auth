package routes

import (
	"encoding/json"
	"log"
	"main/database"
	l "main/libs"
	"main/models"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    var err *l.Error

    w.Header().Set("Content-type", "application/json")

    // db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err = l.CreateError("db_conn", "Unable to connect to the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    // json decoding
    jsonErr := json.NewDecoder(r.Body).Decode(&product)
    if jsonErr != nil {
        err = l.CreateError("json_decoding", "Error decoding the data")
        log.Default().Println(jsonErr)
        json.NewEncoder(w).Encode(err)
        return
    }

    // insert new product
    db.Create(&product)
    json.NewEncoder(w).Encode(product)
}
