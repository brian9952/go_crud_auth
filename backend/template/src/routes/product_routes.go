package routes

import (
	"encoding/json"
	"log"
	"main/database"
	l "main/libs"
	"main/models"
	"net/http"
)

type IdProduct struct {
    Id int `json:"product_id"`
}

type EditProduct struct {
    Productid int `json:"product_id"`
    ProductName string `json:"product_name"`
    ProductDesc string `json:"product_description"`
    ProductValue int `json:"product_value"`
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
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

func ShowProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    var idProduct IdProduct
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
    jsonErr := json.NewDecoder(r.Body).Decode(&idProduct)
    if jsonErr != nil {
        err = l.CreateError("json_decoding", "Error decoding the data")
        log.Default().Println(jsonErr)
        json.NewEncoder(w).Encode(err)
        return
    }

    // search for result
    result := db.First(&product, idProduct.Id)
    if result.Error != nil {
        err = l.CreateError("querying", "Error occured while querying the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    json.NewEncoder(w).Encode(product)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    var editedProduct EditProduct
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
    jsonErr := json.NewDecoder(r.Body).Decode(&editedProduct)
    if jsonErr != nil {
        err = l.CreateError("json_decoding", "Error decoding the data")
        log.Default().Println(jsonErr)
        json.NewEncoder(w).Encode(err)
        return
    }

    // update
    result := db.Model(&product).Where("product_id = ?", editedProduct.Productid).Updates(
        models.Product {
            ProductName: editedProduct.ProductName,
            ProductDescription: editedProduct.ProductDesc,
            ProductValue: editedProduct.ProductValue,
        })

    if result.Error != nil {
        log.Default().Println("Error occured while updating query")
        return
    }

    var succ *l.Success
    succ = l.CreateSuccess("update_success", "Success updateing the data ")
    json.NewEncoder(w).Encode(succ)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    var Id IdProduct
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
    jsonErr := json.NewDecoder(r.Body).Decode(&Id)
    if jsonErr != nil {
        err = l.CreateError("json_decoding", "Error decoding the data")
        log.Default().Println(jsonErr)
        json.NewEncoder(w).Encode(err)
        return
    }

    // delete
    result := db.Delete(&product, Id.Id)
    if result.Error != nil {
        log.Default().Println("Error occured while deleting the query")
        return
    }

    var succ *l.Success
    succ = l.CreateSuccess("delete_success", "Success deleting the data")
    json.NewEncoder(w).Encode(succ)
}

func ShowAllProduct(w http.ResponseWriter, r *http.Request) {
    var products[] models.Product
    var err *l.Error

    w.Header().Set("Content-type", "application/json")
    w.Header().Set("Authorization", "False")

    log.Default().Println(r.Header.Get("IsAuthorized"))

    // get db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err = l.CreateError("db_conn", "Unable to connect to the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    result := db.Find(&products)
    if result.Error != nil {
        log.Default().Println("Error occured while getting query")
        return
    }

    if result.RowsAffected == 0 {
        err = l.CreateError("empty", "Product is empty")
        json.NewEncoder(w).Encode(err)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(products)
        
}
