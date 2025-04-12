package commands

import (
	"PoGoPass/internal/db"
	"PoGoPass/internal/ui"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Delete(dbFile string, masterKey []byte, id int) {
	database, err := db.OpenEncryptedDB(dbFile, masterKey)
	if err != nil {
		ui.PrintError("❌ Failed to open database: " + err.Error())
		return
	}
	defer database.Close()

	row := database.QueryRow("SELECT service, username FROM passwords WHERE id = ?", id)
	var service, username string
	err = row.Scan(&service, &username)
	if err != nil {
		ui.PrintError("❌ No entry found with the provided ID.")
		return
	}

	// Show entry summary before confirmation
	ui.PrintInfo(fmt.Sprintf("🚨 You are about to delete:\n  Service: %s\n  Username: %s\n", service, username))
	ui.PrintInfo("❓ Are you sure you want to delete this entry? (y/N): ")

	reader := bufio.NewReader(os.Stdin)
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))

	if confirm != "y" && confirm != "yes" {
		ui.PrintWarning("❌ Deletion cancelled.")
		return
	}

	// Proceed to delete
	_, err = database.Exec("DELETE FROM passwords WHERE id = ?", id)
	if err != nil {
		ui.PrintError("❌ Failed to delete entry: " + err.Error())
		return
	}

	ui.PrintSuccess("✅ Entry deleted successfully.")
}
