package commands

import (
	"PoGoPass/internal/crypto"
	"PoGoPass/internal/db"
	"PoGoPass/internal/ui"
	"encoding/base64"
	"fmt"
)

// Show securely displays a password entry after verifying the master password again
func Show(dbFile string, masterKey []byte, id int) {
	// Validate ID
	if id <= 0 {
		ui.PrintError("❌ Invalid ID. Please specify a valid ID.")
		return
	}

	// Prompt user to re-enter master password
	ui.PrintInfo("🔐 Re-enter your master password to view the password entry:")
	reenteredPassword := ui.PromptForPassword()

	// Verify re-entered master password
	valid := crypto.VerifyMasterKey(reenteredPassword)
	if !valid {
		ui.PrintError("❌ Master password verification failed. Cannot show entry.")
		return
	}


	// Open the encrypted database
	database, err := db.OpenEncryptedDB(dbFile, masterKey)
	if err != nil {
		ui.PrintError("❌ Failed to open database: " + err.Error())
		return
	}
	defer database.Close()

	// Query for the password entry
	row := database.QueryRow("SELECT service, username, password FROM passwords WHERE id = ?", id)

	var service, username, passwordEnc string
	err = row.Scan(&service, &username, &passwordEnc)
	if err != nil {
		ui.PrintError("❌ No entry found with the provided ID.")
		return
	}

	// Decode the password
	password, err := base64.StdEncoding.DecodeString(passwordEnc)
	if err != nil {
		ui.PrintError("❌ Failed to decode password.")
		return
	}

	// Display the decrypted password
	ui.PrintSuccess("🔐 Password Entry:")
	fmt.Println("Service :", service)
	fmt.Println("Username:", username)
	fmt.Println("Password:", string(password))
}
