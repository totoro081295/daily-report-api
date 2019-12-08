package main

import (
	"sync"

	"github.com/totoro081295/daily-report-api/db"
	"github.com/totoro081295/daily-report-api/db/migrations"
	"github.com/totoro081295/daily-report-api/db/seeds/sql"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	migrations.DropTable(database)
	migrations.Migrate(database)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		database.DB().Exec(sql.InsertAccounts)
		wg.Done()
	}()
	wg.Wait()
}
