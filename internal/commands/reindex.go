package commands

import (
	"PoGoPass/internal/db"
	"PoGoPass/internal/ui"
)

func Reindex(dbFile string, masterKey []byte) {
	database, err := db.OpenEncryptedDB(dbFile, masterKey)
	if err != nil {
		ui.PrintError("❌ Failed to open database: " + err.Error())
		return
	}
	defer database.Close()

	tx, err := database.Begin()
	if err != nil {
		ui.PrintError("❌ Failed to begin transaction: " + err.Error())
		return
	}

	_, err = tx.Exec("CREATE TABLE IF NOT EXISTS passwords_backup (service TEXT, username TEXT, password TEXT)")
	if err != nil {
		tx.Rollback()
		ui.PrintError("❌ Failed to create backup table: " + err.Error())
		return
	}

	_, err = tx.Exec("INSERT INTO passwords_backup(service, username, password) SELECT service, username, password FROM passwords ORDER BY id")
	if err != nil {
		tx.Rollback()
		ui.PrintError("❌ Failed to backup data: " + err.Error())
		return
	}

	_, err = tx.Exec("DROP TABLE passwords")
	if err != nil {
		tx.Rollback()
		ui.PrintError("❌ Failed to drop original table: " + err.Error())
		return
	}

	_, err = tx.Exec("CREATE TABLE passwords (id INTEGER PRIMARY KEY AUTOINCREMENT, service TEXT, username TEXT, password TEXT)")
	if err != nil {
		tx.Rollback()
		ui.PrintError("❌ Failed to recreate passwords table: " + err.Error())
		return
	}

	_, err = tx.Exec("INSERT INTO passwords(service, username, password) SELECT service, username, password FROM passwords_backup")
	if err != nil {
		tx.Rollback()
		ui.PrintError("❌ Failed to restore data: " + err.Error())
		return
	}

	_, err = tx.Exec("DROP TABLE passwords_backup")
	if err != nil {
		tx.Rollback()
		ui.PrintError("❌ Failed to remove backup table: " + err.Error())
		return
	}

	if err = tx.Commit(); err != nil {
		ui.PrintError("❌ Failed to commit transaction: " + err.Error())
		return
	}

	ui.PrintSuccess("✅ Reindex completed successfully.")
}
