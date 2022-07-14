package main

import (
	"fmt"
	"main/database"
)

func main() {
    fmt.Println("Initializing DB...")
    database.CreateDBConnection()
    database.AutoMigrateDB()

    fmt.Println("Migration complete!")
}
