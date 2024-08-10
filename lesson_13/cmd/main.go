package main

import (
	"fmt"
	"os"

	"password-manager/internal/manager"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	mgr := manager.NewManager("passwords.json")
	command := os.Args[1]

	switch command {
	case "list":
		names := mgr.ListNames()
		fmt.Println("Stored passwords:")
		for _, name := range names {
			fmt.Println(name)
		}
	case "put":
		if len(os.Args) != 4 {
			fmt.Println("Usage: ./password-manager put <name> <password>")
			return
		}
		name := os.Args[2]
		password := os.Args[3]
		mgr.SavePassword(name, password)
		fmt.Printf("Password for '%s' saved.\n", name)
	case "get":
		if len(os.Args) != 3 {
			fmt.Println("Usage: ./password-manager get <name>")
			return
		}
		name := os.Args[2]
		password, found := mgr.GetPassword(name)
		if found {
			fmt.Printf("Password for '%s': %s\n", name, password)
		} else {
			fmt.Printf("Password for '%s' not found.\n", name)
		}
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  ./password-manager list              - List all stored passwords")
	fmt.Println("  ./password-manager put <name> <password> - Save a password")
	fmt.Println("  ./password-manager get <name>        - Retrieve a password")
}
