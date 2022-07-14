package database

import (
    "fmt"
    "log"
    "os"
    "time"
    "main/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    )

// global variables
var (
    DBConn *gorm.DB
    )

// db connection parameters
type DbParam struct {
    User string
    Password string
    Db string
    Host string
    Port string
    SSL string
    Timezone string
}

// main functions

func (d *DbParam) InitParams() *DbParam {
    d.User = os.Getenv("DB_USER")
    d.Password = os.Getenv("DB_PASSWORD")
    d.Db = os.Getenv("DB_NAME")
    d.Host = os.Getenv("DB_HOST")
    d.Port = os.Getenv("DB_PORT")
    d.SSL = os.Getenv("DB_SSL")
    d.Timezone = os.Getenv("DB_TIMEZONE")
    return d;
}

func CreateDBConnection() error {
    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN: getDSN(),
    }), &gorm.Config{})

    if err != nil {
        log.Default().Fatal("Fatal Error: Error occured while connecting with the database")
    } else {
        log.Default().Println("Connected to the database")
    }

    sqlDb, err := db.DB()

    sqlDb.SetConnMaxIdleTime(time.Minute * 5)

    // set max number of connection in idle connection
    sqlDb.SetMaxIdleConns(10)

    // set max number of open connections
    sqlDb.SetMaxOpenConns(100)

    // set max amount time a connection may be reused
    sqlDb.SetConnMaxLifetime(time.Hour)

    DBConn = db;
    return err;
}

func GetDatabaseConnection() (*gorm.DB, error) {
    sqlDb, err := DBConn.DB()

    if err != nil {
        return DBConn, err
    }

    if err := sqlDb.Ping(); err != nil {
        return DBConn, err
    }

    return DBConn, nil
}

func AutoMigrateDB() error {
    db, connErr := GetDatabaseConnection()

    if connErr != nil {
        return connErr
    }

    err := db.AutoMigrate(&models.User{}, &models.Product{})
    return err
}

// small functions

func getDSN() string {
    var d *DbParam = new(DbParam)
    d = d.InitParams()
    return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
        d.Host, d.User, d.Password, d.Db, d.Port, d.SSL, d.Timezone)
}
