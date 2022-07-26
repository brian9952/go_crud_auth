package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"main/database"
	l "main/libs"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
)

type idProduct struct {
    Id int `json:"product_id"`
}

func ShowAllProducts(w http.ResponseWriter, r *http.Request) {
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

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var newProduct *models.Product

    w.Header().Set("Content-Type", "application/json")

    // db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := l.CreateErrorMessage("Unable to connect to the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    // get json data
    jsonErr := json.NewDecoder(r.Body).Decode(&newProduct)
    if jsonErr != nil {
        err := l.CreateErrorMessage("Error getting the data")
        json.NewEncoder(w).Encode(err)
        return
    }

    // insert product
    db.Create(&newProduct)

    json.NewEncoder(w).Encode(l.CreateSuccessMessage("Success inserting the data"))
}

func printRawData(inp *io.ReadCloser) string {
    temp, _ := ioutil.ReadAll(*inp)
    tempstr := string(bytes.Replace(temp, []byte("\r"), []byte("\r\n"), -1))
    log.Default().Println(tempstr)
    return tempstr
}

type test struct {
    Test string `json:"test_message"`
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
    var editedProduct *models.Product

    // db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := l.CreateErrorMessage("Unable to connect to the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    // get json data
    jsonErr := json.NewDecoder(r.Body).Decode(&editedProduct)
    if jsonErr != nil {
        err := l.CreateErrorMessage("Error getting the data")
        json.NewEncoder(w).Encode(err)
        return
    }

    // update
    result := db.Model(&editedProduct).Where("product_id = ?", editedProduct.ProductId).Updates(
        models.Product {
            ProductName: editedProduct.ProductName,
            ProductDescription: editedProduct.ProductDescription,
            ProductValue: editedProduct.ProductValue,
        })

    if result.Error != nil {
        err := l.CreateErrorMessage("Error occured while updateing query")
        json.NewEncoder(w).Encode(err)
        return 
    }

    success := l.CreateSuccessMessage("Successfully updating the data")
    json.NewEncoder(w).Encode(success)
    return
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    var Id *idProduct

    // db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := l.CreateErrorMessage("Unable to connect to the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    // get json data
    jsonErr := json.NewDecoder(r.Body).Decode(&Id)
    if jsonErr != nil {
        err := l.CreateErrorMessage("Error getting the data")
        json.NewEncoder(w).Encode(err)
        return
    }

    // delete product
    result := db.Delete(&models.Product{}, Id.Id)
    if result.Error != nil {
        err := l.CreateErrorMessage("Error deleting the data")
        json.NewEncoder(w).Encode(err)
        return
    }

    // success message
    success := l.CreateSuccessMessage("Success deleting the data")
    json.NewEncoder(w).Encode(success)
    return
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
    var product *models.Product

    // get params
    params := mux.Vars(r)
    productId := params["id"]

    // db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := l.CreateErrorMessage("Unable to connect to the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    // get result
    result := db.First(&product, productId)
    if result.Error != nil {
        err := l.CreateErrorMessage("Error getting the data")
        json.NewEncoder(w).Encode(err)
        return
    }

    // success message
    json.NewEncoder(w).Encode(product)
    return
}
