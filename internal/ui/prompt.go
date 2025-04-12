package ui

import (
	"fmt"
)

// Function to prompt for a secure password input and mask it with "****"
func PromptForPassword() string {
    var password string
    for {
        fmt.Print("🔑 Enter your master password: ")
        fmt.Scanln(&password) // Use Scanln for simpler input handling
        if password != "" {
            break
        }
        fmt.Println("❌ Password cannot be empty. Please try again.")
    }
    return password
}


func PrintHelp() {
	fmt.Println("PoGoPass Help:")
	fmt.Println("Usage:")
	fmt.Println("  add     - Add a new password entry")
	fmt.Println("           Usage: add -name <service_name> -username <username> -password <password>")
	fmt.Println("           Example: add -name Tiktok -username arwin -password yourpassword")
	fmt.Println("           If no password is provided, a random password will be generated.")
	fmt.Println()
	fmt.Println("  list    - List all saved passwords")
	fmt.Println("           Usage: list")
	fmt.Println()
	fmt.Println("  delete  - Delete a password entry by ID")
	fmt.Println("           Usage: delete -id <entry_id>")
	fmt.Println("           Example: delete -id 1")
	fmt.Println()
	fmt.Println("  show    - Show the decrypted password entry by ID")
	fmt.Println("           Usage: show -id <entry_id>")
	fmt.Println("           Example: show -id 1")
	fmt.Println()
	fmt.Println("  reindex - Rebuild ID order of the database")
	fmt.Println("           Usage: reindex")
	fmt.Println()
	fmt.Println("  help    - Show this help message")
	fmt.Println("           Usage: help")
	fmt.Println()
	fmt.Println("  quit    - Exit the application")
	fmt.Println("           Usage: quit")
}
