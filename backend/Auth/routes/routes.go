package routes

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"main/libs"
	"net/http"
)

type LoginData struct {
    dataStr string `json:"data"`
}

func decodeStr(data string) (string, bool){
    decodedStr, err := base64.StdEncoding.DecodeString(data)
    if err != nil {
        return "", false
    }
    return string(decodedStr), true
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
    // check if authorized
    if r.Header.Get("Authorized") != "0" {
        return
    } 

    // init pointers
    var data *LoginData

    // get base64 data
    jsonErr := json.NewDecoder(r.Body).Decode(&data)
    if jsonErr != nil {
        err := libs.CreateErrorMessage("Error: Decoding the data")
        json.NewEncoder(w).Encode(err)
        return
    }
    
    // decode the data
    dataStr, isErr := decodeStr(data.dataStr)
    if isErr {

    }

}

func RegisterUser(w http.ResponseWriter, r *http.Request){ 
    log.Default().Println("THIS IS A REGISTER API")
}

