package main

import (
	"bufio"
	"fmt"
	"os"

	"password-manager/internal/manager"
)

func main() {
	mgr := manager.NewManager("passwords.json")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. List passwords")
		fmt.Println("2. Save a password")
		fmt.Println("3. Retrieve a password")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			names := mgr.ListNames()
			fmt.Println("Stored passwords:")
			for _, name := range names {
				fmt.Println(name)
			}
		case "2":
			fmt.Print("Enter name: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Print("Enter password: ")
			scanner.Scan()
			password := scanner.Text()
			mgr.SavePassword(name, password)
		case "3":
			fmt.Print("Enter name: ")
			scanner.Scan()
			name := scanner.Text()
			password, found := mgr.GetPassword(name)
			if found {
				fmt.Println("Password:", password)
			} else {
				fmt.Println("Password not found")
			}
		case "4":
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
