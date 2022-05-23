package database

import "main/models"

func AutoMigrateDB() error {
    db, connErr := GetDatabaseConnection()

    if connErr != nil {
        return connErr
    }

    // migrate user models
    err := db.AutoMigrate(&models.User{})

    return err
}
