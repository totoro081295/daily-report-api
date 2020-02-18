package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/totoro081295/daily-report-api/account"
	"github.com/totoro081295/daily-report-api/category"
	"github.com/totoro081295/daily-report-api/dailycontent"
	"github.com/totoro081295/daily-report-api/db"
	"github.com/totoro081295/daily-report-api/project"
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
		&category.Category{},
		&dailycontent.DailyContent{},
		&refreshtoken.RefreshToken{},
		&project.Project{},
	)
}

// DropTable drop table
func DropTable(database *gorm.DB) {
	database.DropTableIfExists(
		&account.Account{},
		&category.Category{},
		&dailycontent.DailyContent{},
		&refreshtoken.RefreshToken{},
		&project.Project{},
	)
}
