package main

import (
	"fmt"
	"os"
	"strings"

	"PoGoPass/internal/commands"
	"PoGoPass/internal/crypto"
	"PoGoPass/internal/ui"
)

func welcomeScreen() {
	ui.PrintData(`
 _______             ______             _______                              
/       \           /      \           /       \                             
$$$$$$$  | ______  /$$$$$$  |  ______  $$$$$$$  | ______    _______  _______ 
$$ |__$$ |/      \ $$ | _$$/  /      \ $$ |__$$ |/      \  /       |/       |
$$    $$//$$$$$$  |$$ |/    |/$$$$$$  |$$    $$/ $$$$$$  |/$$$$$$$//$$$$$$$/ 
$$$$$$$/ $$ |  $$ |$$ |$$$$ |$$ |  $$ |$$$$$$$/  /    $$ |$$      \$$      \ 
$$ |     $$ \__$$ |$$ \__$$ |$$ \__$$ |$$ |     /$$$$$$$ | $$$$$$  |$$$$$$  |
$$ |     $$    $$/ $$    $$/ $$    $$/ $$ |     $$    $$ |/     $$//     $$/ 
$$/       $$$$$$/   $$$$$$/   $$$$$$/  $$/       $$$$$$$/ $$$$$$$/ $$$$$$$/  
`)
	ui.PrintSuccess("Credits: arwindpianist")
	fmt.Println("💡 Get started by entering your master password.")
}

func main() {
	welcomeScreen()

	const dbFile = "pogo.db"

	// Check if hash file exists
	if _, err := os.Stat("masterkey.hash"); os.IsNotExist(err) {
		// First-time setup
		ui.PrintInfo("First-time setup: Create your master password.")
		password := ui.PromptForPassword()
		if password == "" {
			ui.PrintError("Password cannot be empty. Setup aborted.")
			return
		}

		hash, err := crypto.DeriveMasterKey(password)
		if err != nil {
			ui.PrintError("Error deriving master key: " + err.Error())
			return
		}

		err = crypto.SaveMasterKeyHash(hash)
		if err != nil {
			ui.PrintError("Failed to save master password hash: " + err.Error())
			return
		}

		ui.PrintSuccess("Master password saved securely. Please restart PoGoPass.")
		return
	}

	// Normal use: verify master password
	password := ui.PromptForPassword()
	if password == "" {
		ui.PrintError("Password cannot be empty. Please try again.")
		return
	}

	if !crypto.VerifyMasterKey(password) {
		ui.PrintError("Invalid master password. Access denied.")
		return
	}

	masterKey, err := crypto.DeriveMasterKey(password)
	if err != nil {
		ui.PrintError("Could not derive master key: " + err.Error())
		return
	}

	ui.PrintSuccess("Master key verified successfully")

	// Main command loop
	for {
		ui.PrintInfo("Enter command (add, list, delete, show, reindex, help, quit): ")
		var command string
		_, err := fmt.Scanln(&command)
		if err != nil {
			ui.PrintError("Error reading command: " + err.Error())
			continue
		}

		command = strings.TrimSpace(command)

		switch command {
		case "help":
			ui.PrintHelp()

		case "add":
			commands.Add(dbFile, masterKey)

		case "list":
			commands.List(dbFile, masterKey)

		case "delete":
			ui.PrintInfo("Enter the ID of the password entry to delete: ")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil || id <= 0 {
				ui.PrintError("Invalid ID. Use the `list` command to see available IDs.")
				break
			}
			commands.Delete(dbFile, masterKey, id)

		case "show":
			ui.PrintInfo("Enter ID to show (use 'list' if you're not sure): ")
			var id int
			_, err := fmt.Scanln(&id)
			if err != nil || id <= 0 {
				ui.PrintError("Invalid input. Please enter a valid numeric ID.")
				continue
			}
			commands.Show(dbFile, masterKey, id)

		case "reindex":
			commands.Reindex(dbFile, masterKey)

		case "quit":
			ui.PrintSuccess("Exiting PoGoPass. Goodbye!")
			ui.PrintData("Credits: arwindpianist")
			return

		default:
			ui.PrintError("Invalid command. Type 'help' for available commands.")
		}
	}
}
