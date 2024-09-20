package main

import (
	"fmt"
	"os"
	"os/user"
	"interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is my intepreter for the Monkey programming language by Thorsten Ball.\n", user.Username)
	fmt.Printf("Try typing in commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
