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

type RefreshTokenData struct {
    RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
    StatusType int `json:"status_type"`
    StatusMessage string `json:"status_message"`
    Username string `json:"username"`
    Role string `json:"role"`
    AccessToken string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}

type RegisterResponse struct {
    StatusType int `json:"status_type"`
    StatusMessage string `json:"status_message"`
}

func createLoginResponse(
    status_type int, 
    status_message string, 
    username string, 
    role string, 
    access_token string,
    refresh_token string) (*LoginResponse) {
    var lr *LoginResponse = new(LoginResponse)

    lr.StatusType = status_type
    lr.StatusMessage = status_message
    lr.Username = username
    lr.Role = role
    lr.AccessToken = access_token
    lr.RefreshToken = refresh_token

    return lr
}

func createRegisterResponse(status_type int, status_message string) (*RegisterResponse){
    // 0 success, 1 frontend error, 2 internal error
    var rr *RegisterResponse = new(RegisterResponse)

    rr.StatusType = status_type
    rr.StatusMessage = status_message
    return rr
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
    newUser.Role = "user"

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

func generateAccessToken(user_id int, username string, role string) (string, error) {
    key := []byte(libs.Auth_key)

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = user_id
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

func generateRefreshToken(user_id int) (string, error){
    key := []byte(libs.Auth_key)

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = user_id
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
    log.Default().Println(w.Header().Get("Access-Control-Allow-Origin"))

    // check if authorized
    if r.Header.Get("Authorized") != "1" {
        err := libs.CreateErrorMessage("Error: You are not authorized")
        json.NewEncoder(w).Encode(err)
        return
    } 

    // db conn
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := createLoginResponse(3, "Server internal error",  "", "", "", "")
        json.NewEncoder(w).Encode(err)
        return
    }

    // init pointer 
    var user *models.User
    var data *AuthData

    // get base64 data
    jsonErr := json.NewDecoder(r.Body).Decode(&data)
    if jsonErr != nil {
        err := createLoginResponse(3, "Server internal error", "", "", "", "")
        json.NewEncoder(w).Encode(err)
        return
    }

    // decode base64
    dataArr, decodeErr := decodeLoginString(data.DataStr)
    if decodeErr != nil {
        err := createLoginResponse(4, "Server internal error", "", "", "", "")
        json.NewEncoder(w).Encode(err)
        return
    }

    // query db
    if result := db.Where("username = ?", dataArr[0]).First(&user); result.Error != nil {
        err := createLoginResponse(1, "Username is incorrect", "", "", "", "")
        json.NewEncoder(w).Encode(err)
        return
    }

    // check password
    pwdErr := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(dataArr[1]))
    if pwdErr != nil {
        err := createLoginResponse(2, "Password is incorrect", "", "", "", "")
        json.NewEncoder(w).Encode(err)
        return
    }

    // generate access and refresh token
    accessTokenStr, accessTokenErr := generateAccessToken(user.UserId, user.Username, user.Role)
    if accessTokenErr != nil {
        err := createLoginResponse(3, "Server internal error", "", "", "", "")
        json.NewEncoder(w).Encode(err)
        return
    }

    refreshTokenStr, refreshTokenErr := generateRefreshToken(user.UserId)
    if refreshTokenErr != nil {
        err := createLoginResponse(3, "Server internal error", "", "", "", "")
        json.NewEncoder(w).Encode(err)
        return 
    }

    // send response
    log.Default().Println("Login Success")
    response := createLoginResponse(0, "Login success", user.Username, user.Role, accessTokenStr, refreshTokenStr)
    json.NewEncoder(w).Encode(response)
    return
}

func RegisterUser(w http.ResponseWriter, r *http.Request){ 
    // change haeder
    w.Header().Set("Content-Type", "application/json")
    log.Default().Println(w.Header().Get("Access-Control-Allow-Origin"))

    // check if authorized
    if r.Header.Get("Authorized") != "1" {
        err := libs.CreateErrorMessage("Error: You are not authorized")
        json.NewEncoder(w).Encode(err)
        return
    }

    // db conn
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := createRegisterResponse(2, "Unable to connect to the database")
        json.NewEncoder(w).Encode(err)
        return
    }

    // init pointer
    var data *AuthData

    // get base64 data
    jsonErr := json.NewDecoder(r.Body).Decode(&data)
    if jsonErr != nil {
        err := createRegisterResponse(2, "Cannot decode the data")
        json.NewEncoder(w).Encode(err)
        return
    }

    // process base64 data
    newUser, processErr := processBase64Data(data.DataStr)
    if processErr != nil {
        err := createRegisterResponse(1, "Invalid user input")
        json.NewEncoder(w).Encode(err)
        return
    }

    // generate hash password
    var pwdErr error
    newUser.HashPassword, pwdErr = GetHashPassword(newUser.HashPassword)
    if pwdErr != nil {
        err := createRegisterResponse(1, "Invalid user input")
        log.Default().Println("Error in password hashing")
        //err := libs.CreateErrorMessage("Error: Invalid Input")
        json.NewEncoder(w).Encode(err)
        return
    }

    // insert user
    db.Create(&newUser)
    log.Default().Println("User creation complete")
    succ := createRegisterResponse(0, "Register success")
    json.NewEncoder(w).Encode(succ)
    return
}

func RefreshTokenUser(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("Authorized") != "1" {
        err := libs.CreateErrorMessage("Error: You are not authorized")
        json.NewEncoder(w).Encode(err)
        return
    }

    if r.Header.Get("Authorization") == "" { // bearer is empty
        resp := createLoginResponse(2, "Token not found", "", "", "", "")
        json.NewEncoder(w).Encode(resp)
        return
    }
    
    // db conn
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err := createLoginResponse(3, "Server internal error", "", "", "", "")
        json.NewEncoder(w).Encode(err)
        return
    }

    // init pointer
    var user *models.User

    // get token data
    //jsonErr := json.NewDecoder(r.Body).Decode(&data)
    //if jsonErr != nil {
    //    err := createLoginResponse(3, "Server internal error", "", "", "", "")
    //    json.NewEncoder(w).Encode(err)
    //    return
    //}
    bearer := r.Header.Get("Authorization")
    refreshToken := strings.Split(bearer, ";")[1]

    // get secret key and authorize
    var secret_key = []byte(libs.Auth_key)

    token, jwtErr := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Error in parsing token")
        }
        return secret_key, nil
    })

    // error handling
    if jwtErr != nil {
        err := libs.CreateErrorMessage("Error: User token is invalid")
        json.NewEncoder(w).Encode(err)
        return
    }

    // check ok
    claims, ok := token.Claims.(jwt.MapClaims)
    if ok && token.Valid {
        // query db
        if result := db.Where("user_id = ?", claims["user_id"]).First(&user); result.Error != nil {
            err := createLoginResponse(3, "Server internal error", "", "", "", "")
            json.NewEncoder(w).Encode(err)
            return
        }
        
        // create new access token
        accessToken, accessTokenErr := generateAccessToken(user.UserId, user.Username, user.Role)
        if accessTokenErr != nil {
            err := createLoginResponse(3, "Server internal error", "", "", "", "")
            json.NewEncoder(w).Encode(err)
            return
        }
        
        // create new refresh token
        refreshToken, refreshTokenErr := generateRefreshToken(user.UserId)
        if refreshTokenErr != nil {
            err := createLoginResponse(3, "Server internal error", "", "", "", "")
            json.NewEncoder(w).Encode(err)
            return
        }

        // send response with new tokens
        response := createLoginResponse(0, "Refresh Success", user.Username, user.Role, accessToken, refreshToken)
        json.NewEncoder(w).Encode(response)

    } else {
        // token invalid
        err := createLoginResponse(2, "Error: Token Invalid", "", "", "", "")
        json.NewEncoder(w).Encode(err)

    }
    return
}
