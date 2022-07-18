package routes

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"main/database"
	l "main/libs"
	"main/models"
	"net/http"
	"strconv"
	"strings"
)

type rawData struct {
    Data string `json:"data"`
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
    var products[] models.Product

    w.Header().Set("Content-Type", "application/json")

    // db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := l.CreateErrorMessage("Unable to connect to the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    // db query
    result := db.Find(&products)
    if result.Error != nil {
        err := l.CreateErrorMessage("Error occured while getting query")
        json.NewEncoder(w).Encode(err)
        return 
    }

    // check if empty
    if result.RowsAffected == 0 {
        err := l.CreateErrorMessage("Product is empty")
        json.NewEncoder(w).Encode(err)
        return 
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(products)
}

func processProductRaw(dataStr string) (*models.Product, error) {
    var newProduct *models.Product
    decodedStr, err := base64.StdEncoding.DecodeString(dataStr)
    if err != nil {
        return nil, err
    }

    // split string
    dataArr := strings.Split(string(decodedStr), ":")

    // input product
    newProduct = new(models.Product)
    newProduct.ProductName = dataArr[0]
    newProduct.ProductDescription = dataArr[1]
    newProduct.ProductValue, err = strconv.Atoi(dataArr[2])
    if err != nil {
        return nil, fmt.Errorf("Cannot convert string to int")
    }

    return newProduct, nil
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var rd *rawData

    w.Header().Set("Content-Type", "application/json")

    // db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := l.CreateErrorMessage("Unable to connect to the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    // get json data
    jsonErr := json.NewDecoder(r.Body).Decode(&rd)
    if jsonErr != nil {
        err := l.CreateErrorMessage("Error getting the data")
        json.NewEncoder(w).Encode(err)
        return
    }

    // decrypt base64 data
    newProduct, processErr := processProductRaw(rd.Data)
    if processErr != nil {
        err := l.CreateErrorMessage("Error processing incoming data")
        json.NewEncoder(w).Encode(err)
        return
    }

    // insert product
    db.Create(&newProduct)

    json.NewEncoder(w).Encode(l.CreateSuccessMessage("Success inserting the data"))
}
