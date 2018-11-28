package main

import "github.com/fatih/color"

const (
	ver = "0.0.1"
)

func main() {

	color.Green("Subscription/Identifiaciont server %s", ver)
	// Calling bootup from boot.go file
	// to load configuration and check files
	if !Bootup() {
		color.Red("Failed to bootup servers")
		return
	}

	color.Green("Server booted Succesfully")
}
