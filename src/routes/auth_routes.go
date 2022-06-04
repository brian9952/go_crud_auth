package routes

import (
	"encoding/json"
	"log"
	"main/database"
	"main/models"
    l "main/libs"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secretkey = os.Getenv("JWT_KEY")

func GetHashPassword(pass string) (string, error){
    bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
    return string(bytes), err
}

type LoginDetails struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Token struct {
    Username string `json:"username"`
    Role string `json:"role"`
    Token_string string `json:"token"`
}

func generateToken(username string, role string) (string, error) {
    var key = []byte(secretkey)
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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var newUser models.User
    var err *l.Error

    // db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err = l.CreateError("db_conn", "Unable to connect to the database")
        w.Header().Set("Content-type", "application/json")
        json.NewEncoder(w).Encode(err)
        return
    }

    // json decoding
    jsonErr := json.NewDecoder(r.Body).Decode(&newUser)
    if jsonErr != nil {
        err = l.CreateError("json_decoding", "Error decoding the data")
        log.Default().Println(jsonErr)
        w.Header().Set("Content-type", "application/json")
        json.NewEncoder(w).Encode(err)
        return
    }

    // generate hash password
    var pwdErr error
    newUser.HashPassword, pwdErr = GetHashPassword(newUser.HashPassword)
    if pwdErr != nil {
        log.Default().Println("Error in password hashing")
        return
    }

    // insert user
    db.Create(&newUser)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newUser)

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
    var loginDetails *LoginDetails = new(LoginDetails)
    var user models.User
    var err *l.Error
    var token Token

    // db connection
    db, connErr := database.GetDatabaseConnection()
    if connErr != nil {
        err = l.CreateError("db_conn", "Unable to connect to the database")
        w.Header().Set("Content-type", "application/json")
        json.NewEncoder(w).Encode(err)
        return
    }

    // json decoding
    jsonErr := json.NewDecoder(r.Body).Decode(&loginDetails)
    if jsonErr != nil {
        err = l.CreateError("json_decoding", "Error decoding the data ")
        log.Default().Println(jsonErr)
        w.Header().Set("Content-type", "application/json")
        json.NewEncoder(w).Encode(err)
        return
    }

    // check username
    db.Where("username = ?",loginDetails.Username).First(&user)
    if user.Username == "" {
        err = l.CreateError("authen", "Username is incorrect")
        w.Header().Set("Content-type", "application/json")
        json.NewEncoder(w).Encode(err)
        return
    }

    // check password
    pwErr := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(loginDetails.Password))
    if pwErr != nil {
        err = l.CreateError("authen", "User password is incorrect")
        w.Header().Set("Content-type", "application/json")
        json.NewEncoder(w).Encode(err)
        return
    }

    // generate jwt
    tokenStr, tokenErr := generateToken(user.Username, user.Role)
    if tokenErr != nil {
        err = l.CreateError("token_creation", "Failed to generate token")
        w.Header().Set("Content-type", "application/json")
        json.NewEncoder(w).Encode(err)
        return
    }

    // send response
    token.Username = user.Username
    token.Role = user.Role
    token.Token_string  = tokenStr
    w.Header().Set("Content-type", "application/json")
    json.NewEncoder(w).Encode(token)
}
