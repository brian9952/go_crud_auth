package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"main/database"
	"main/libs"
	l "main/libs"
	"main/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type IdProduct struct {
    Id int `json:"product_id" binding:"required"`
}

type products struct {
    ProductId int `json:"product_id"`
    ProductName string `json:"product_name"`
    ProductValue int `json:"product_value"`
    ProductDescription string `json:"product_description"`
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
        err := l.CreateAddProductMessage(2, "Internal error, please contact administrator", -1)
        json.NewEncoder(w).Encode(err)
        return
    }


    // get json data
    jsonErr := json.NewDecoder(r.Body).Decode(&newProduct)
    if jsonErr != nil {
        err := l.CreateAddProductMessage(1, "User input error", -1)
        json.NewEncoder(w).Encode(err)
        return
    }

    // insert database
    if result := db.Create(&newProduct); result.Error != nil {
        err := l.CreateAddProductMessage(2, "Internal error, please contact administrator", -1)
        json.NewEncoder(w).Encode(err)
        return
    }

    json.NewEncoder(w).Encode(libs.CreateAddProductMessage(0, "Success inserting the data", newProduct.ProductId))
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
        err := l.CreateEditProductMessage(2, "Unable to connect to the database", -1)
        json.NewEncoder(w).Encode(err)
        return
    }

    // get json data
    jsonErr := json.NewDecoder(r.Body).Decode(&editedProduct)
    if jsonErr != nil {
        err := l.CreateEditProductMessage(1, "Error when decoding user input", -1)
        json.NewEncoder(w).Encode(err)
        return
    }
    
    log.Default().Println(editedProduct.ProductId)

    // update
    result := db.Model(&editedProduct).Where("product_id = ?", editedProduct.ProductId).Updates(
        models.Product {
            ProductName: editedProduct.ProductName,
            ProductDescription: editedProduct.ProductDescription,
            ProductValue: editedProduct.ProductValue,
        })

    if result.Error != nil {
        err := l.CreateEditProductMessage(2, "There is no result on database", -1)
        json.NewEncoder(w).Encode(err)
        return 
    }

    success := l.CreateEditProductMessage(0, "Edit product success", editedProduct.ProductId)
    json.NewEncoder(w).Encode(success)
    return
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    //var jsonData *IdProduct
    // get params
    params := mux.Vars(r)
    product_id := params["id"]

    // db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := l.CreateDeleteProductMessage(2, "Unable to connect to the database", -1)
        json.NewEncoder(w).Encode(err)
        return
    }

    // debugging
    b_str, _ := io.ReadAll(r.Body)
    log.Default().Println(string(b_str))

    // delete product
    result := db.Delete(&models.Product{}, product_id)
    if result.Error != nil {
        err := l.CreateDeleteProductMessage(1, "Error deleting the data", -1)
        json.NewEncoder(w).Encode(err)
        return
    }

    // success message
    product_id_int, _ := strconv.Atoi(product_id)
    success := l.CreateDeleteProductMessage(0, "Success deleting the data", product_id_int)
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
