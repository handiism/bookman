package main

import (
	"bookman/config"
	"bookman/helper"
)

func main() {
	db := config.OpenDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	helper.DatabaseMigration(db)
}
