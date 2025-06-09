package main

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("format: ./bcrypt <command> <args>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "hash":
		if len(os.Args) < 3 {
			fmt.Println("format: ./bcrypt hash <plaintext>")
			os.Exit(1)
		}

		fmt.Println(hash(os.Args[2]))
	case "compare":
		if len(os.Args) < 4 {
			fmt.Println("format: ./bcrypt compare <plaintext> <hash>")
			os.Exit(1)
		}

		msg, err := compare(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg)
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])
	}
}

func hash(pass string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error hashing: %v\n", pass)
	}

	return string(hashedBytes)
}

func compare(pass, hash string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return "", errors.New("password is invalid")
	}

	return "password is valid", nil
}
