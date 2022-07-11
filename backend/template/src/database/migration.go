package database

import "main/models"

func AutoMigrateDB() error {
    db, connErr := GetDatabaseConnection()

    if connErr != nil {
        return connErr
    }

    // migrate models
    err := db.AutoMigrate(&models.User{}, &models.Product{})
    return err
}
