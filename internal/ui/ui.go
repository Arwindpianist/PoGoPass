package ui

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	success = color.New(color.FgGreen).SprintFunc()
	errorMsg = color.New(color.FgRed).SprintFunc()
	info = color.New(color.FgCyan).SprintFunc()
	warn = color.New(color.FgYellow).SprintFunc()
	data = color.New(color.FgMagenta).SprintFunc()
)

// PrintSuccess prints a green success message with a ✅ icon.
func PrintSuccess(msg string) {
	fmt.Println(success("✅ " + msg))
}

// PrintError prints a red error message with a ❌ icon.
func PrintError(msg string) {
	fmt.Println(errorMsg("❌ " + msg))
}

// PrintInfo prints an informational message in cyan with an ℹ️ icon.
func PrintInfo(msg string) {
	fmt.Println(info("ℹ️ " + msg))
}

// PrintWarning prints a yellow warning message with a ⚠️ icon.
func PrintWarning(msg string) {
	fmt.Println(warn("⚠️ " + msg))
}

// PrintTitle prints a blue title header.
func PrintTitle(title string) {
	titleStyle := color.New(color.FgHiBlue, color.Bold).SprintFunc()
	fmt.Println(titleStyle("📋 " + title))
}

func PrintData(msg string) {
	fmt.Println(data(msg))
}