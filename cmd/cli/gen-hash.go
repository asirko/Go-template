package main

import (
	"fmt"
	"github.com/asirko/go-template/internal/core/util"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("A single argument is required: the password to be hashed")
	}

	hashedPassword, err := util.HashPassword(os.Args[1])
	if err != nil {
		fmt.Println("Error hashing password")
	}
	fmt.Println("hashedPassword:")
	fmt.Println(hashedPassword)

}
