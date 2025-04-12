package commands

import (
	"PoGoPass/internal/db"
	"PoGoPass/internal/ui"
	"fmt"
)

func List(dbFile string, masterKey []byte) {
	// Open the encrypted database
	database, err := db.OpenEncryptedDB(dbFile, masterKey)
	if err != nil {
		ui.PrintError("❌ Failed to open database: " + err.Error())
		return
	}
	defer database.Close()

	// Display header
	ui.PrintInfo("📋 Listing encrypted entries:")
	ui.PrintSuccess("Note: Passwords shown here are encrypted.\nUse the `show` command with an ID to view the decrypted password.\n")

	// Query for all entries in the database
	rows, err := database.Query("SELECT id, service, username, password FROM passwords ORDER BY id")
	if err != nil {
		ui.PrintError("❌ Failed to query database: " + err.Error())
		return
	}
	defer rows.Close()

	// Check if there are any entries
	hasEntries := false

	// Iterate over the rows and print details
	for rows.Next() {
		var id int
		var service, username, passwordEnc string
		if err := rows.Scan(&id, &service, &username, &passwordEnc); err != nil {
			ui.PrintError("❌ Error reading row: " + err.Error())
			continue
		}
		hasEntries = true
		ui.PrintData(fmt.Sprintf("ID: %d\nService: %s\nUsername: %s\nEncrypted Password: %s\n", id, service, username, passwordEnc))
	}

	// If no entries exist, notify the user
	if !hasEntries {
		ui.PrintWarning("❌ No password entries found. You can add entries using the `add` command.")
	}
}
