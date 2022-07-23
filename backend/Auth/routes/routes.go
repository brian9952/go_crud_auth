package routes

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"main/libs"
	"main/models"
	"net/http"
	"strings"
    "time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthData struct {
    DataStr string `json:"data"`
}

type UserData struct {
    Username string `json:"username"`
    HashPassword string `json:"password"`
    Role string `json:"user_role"`
}

type LoginResponse struct {
    StatusType string `json:"status_type"`
    Username string `json:"username"`
    Role string `json:"role"`
    Token string `json:"token"`
}

func createLoginResponse(username string, role string, token string) (*LoginResponse) {
    var lr *LoginResponse = new(LoginResponse)

    lr.StatusType = "1"
    lr.Username = username
    lr.Role = role
    lr.Token = token

    return lr
}

func processBase64Data(dataStr string) (*models.User, error) {
    var newUser *models.User = new(models.User)

    // decode to str
    decodedStr, err := base64.StdEncoding.DecodeString(dataStr)
    if err != nil {
        return nil, fmt.Errorf("Failed to decode string")
    }

    // parse into struct
    dataArr := strings.Split(string(decodedStr), ":")
    newUser.Username = dataArr[0]
    newUser.HashPassword = dataArr[1]
    newUser.Role = dataArr[2]

    return newUser, nil
}

func decodeLoginString(dataStr string) ([]string, error) {
    // decode to str
    decodedStr, err := base64.StdEncoding.DecodeString(dataStr)
    if err != nil {
        return nil, err
    }

    // split string
    dataArr := strings.Split(string(decodedStr), ":")
    return dataArr, nil
}

func GetHashPassword(pass string) (string, error){
    bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
    return string(bytes), err
}

func generateToken(username string, role string) (string, error) {
    key := []byte(libs.Auth_key)

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["authorized"] = true
    claims["username"] = username
    claims["role"] = role
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString(key)
    if err != nil {
        log.Default().Println("Something went wrong")
        log.Default().Println(err.Error())
        return "", err
    }

    return tokenString, nil
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
    // change header
    w.Header().Set("Content-Type", "application/json")

    log.Default().Println(r.Header.Get("API-Token"))

    // check if authorized
    if r.Header.Get("Authorized") != "1" {
        err := libs.CreateErrorMessage("Error: You are not authorized")
        json.NewEncoder(w).Encode(err)
        return
    } 

    // db conn
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := libs.CreateErrorMessage("Error: Unable to connect to the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    // init pointer 
    var user *models.User
    var data *AuthData

    // get base64 data
    jsonErr := json.NewDecoder(r.Body).Decode(&data)
    if jsonErr != nil {
        err := libs.CreateErrorMessage("Error: Decoding the data")
        json.NewEncoder(w).Encode(err)
        return
    }

    // decode base64
    dataArr, decodeErr := decodeLoginString(data.DataStr)
    if decodeErr != nil {
        err := libs.CreateErrorMessage("Error: Decoding the data")
        json.NewEncoder(w).Encode(err)
        return
    }

    // query db
    db.Where("username = ?",dataArr[0]).First(&user)
    if user.Username == "" {
        err := libs.CreateErrorMessage("Error: Username is incorrect")
        json.NewEncoder(w).Encode(err)
        return
    }

    // check password
    pwdErr := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(dataArr[1]))
    if pwdErr != nil {
        err := libs.CreateErrorMessage("Error: Password is incorrect")
        json.NewEncoder(w).Encode(err)
        return
    }

    // generate jwt
    tokenStr, tokenErr := generateToken(user.Username, user.Role)
    if tokenErr != nil {
        err := libs.CreateErrorMessage("Error: Something went wrong, contact administrator if you see this message")
        json.NewEncoder(w).Encode(err)
        return
    }

    // send response
    log.Default().Println("Login Success")
    response := createLoginResponse(user.Username, user.Role, tokenStr)
    json.NewEncoder(w).Encode(response)
    return
}

func RegisterUser(w http.ResponseWriter, r *http.Request){ 
    if r.Header.Get("Authorized") != "1" {
        err := libs.CreateErrorMessage("Error: You are not authorized")
        json.NewEncoder(w).Encode(err)
        return
    }

    // db conn
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := libs.CreateErrorMessage("Error: Unable to connect to the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    // init pointer
    var data *AuthData

    // get base64 data
    jsonErr := json.NewDecoder(r.Body).Decode(&data)
    if jsonErr != nil {
        err := libs.CreateErrorMessage("Error: Decoding the data")
        json.NewEncoder(w).Encode(err)
        return
    }

    // process base64 data
    newUser, processErr := processBase64Data(data.DataStr)
    if processErr != nil {
        log.Default().Println("Error: Unable to process base64 data")
        err := libs.CreateErrorMessage("Error: Invalid input")
        json.NewEncoder(w).Encode(err)
        return
    }

    // generate hash password
    var pwdErr error
    newUser.HashPassword, pwdErr = GetHashPassword(newUser.HashPassword)
    if pwdErr != nil {
        log.Default().Println("Error in password hashing")
        err := libs.CreateErrorMessage("Error: Invalid Input")
        json.NewEncoder(w).Encode(err)
        return
    }

    // insert user
    db.Create(&newUser)
    w.Header().Set("Content-Type", "application/json")
    log.Default().Println("User creation complete")
    succ := libs.CreateSuccessMessage("New user has been created!")
    json.NewEncoder(w).Encode(succ)
    return
}

