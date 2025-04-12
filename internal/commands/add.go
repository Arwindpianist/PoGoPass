package commands

import (
	"PoGoPass/internal/db"
	"PoGoPass/internal/ui"
	"bufio"
	"encoding/base64"
	"os"
	"strings"

	"github.com/sethvargo/go-password/password"
)

// Add function to handle password entry
func Add(dbFile string, masterKey []byte) {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for the service name
	ui.PrintInfo("Enter service name (e.g. Facebook): ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Prompt for the username
	ui.PrintInfo("Enter username (e.g. your email): ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	// Prompt for the password
	var passwordVal string
	ui.PrintInfo("Enter password (leave blank to auto-generate): ")
	passwordInput, _ := reader.ReadString('\n')
	passwordInput = strings.TrimSpace(passwordInput)

	if passwordInput == "" {
		// Auto-generate password if none is provided
		gen, err := password.Generate(16, 4, 4, false, false)
		if err != nil {
			ui.PrintError("Failed to generate password: " + err.Error())
			os.Exit(1)
		}
		passwordVal = gen
		ui.PrintSuccess("Auto-generated password: " + passwordVal)
	} else {
		passwordVal = passwordInput
	}

	// Encode password
	encoded := base64.StdEncoding.EncodeToString([]byte(passwordVal))

	// Open the encrypted database
	database, err := db.OpenEncryptedDB(dbFile, masterKey)
	if err != nil {
		ui.PrintError("DB error: " + err.Error())
		os.Exit(1)
	}
	defer database.Close()

	// Insert new password entry into the database
	_, err = database.Exec("INSERT INTO passwords (service, username, password) VALUES (?, ?, ?)", name, username, encoded)
	if err != nil {
		ui.PrintError("Insert error: " + err.Error())
		os.Exit(1)
	}

	// Success message
	ui.PrintSuccess("✅ Password saved securely.")
}
