package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/totoro081295/daily-report-api/account"
	"github.com/totoro081295/daily-report-api/db"
	"github.com/totoro081295/daily-report-api/refreshtoken"
)

// Execute execute migration
func Execute() {
	database := db.ConnectDB()
	Migrate(database)
	database.Close()
}

// Migrate migration
func Migrate(database *gorm.DB) {
	database.AutoMigrate(
		&account.Account{},
		&refreshtoken.RefreshToken{},
	)
}

// DropTable drop table
func DropTable(database *gorm.DB) {
	database.DropTableIfExists(
		&account.Account{},
		&refreshtoken.RefreshToken{},
	)
}
