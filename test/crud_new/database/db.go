package database

import (
    "fmt"
    "log"
    "os"
    "time"

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
    d.User = os.Getenv("POSTGRES_USER")
    d.Password = os.Getenv("POSTGRES_PASSWORD")
    d.Db = os.Getenv("POSTGRES_DB")
    d.Host = os.Getenv("POSTGRES_HOST")
    d.Port = os.Getenv("POSTGRES_PORT")
    d.SSL = os.Getenv("POSTGRES_SSL")
    d.Timezone = os.Getenv("POSTGRES_TIMEZONE")
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

// small functions

func getDSN() string {
    var d *DbParam = new(DbParam)
    d = d.InitParams()
    return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
        d.Host, d.User, d.Password, d.Db, d.Port, d.SSL, d.Timezone)
}
