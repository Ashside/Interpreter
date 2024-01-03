package main

import (
	"Interp/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Ashside's Interp is running!\n", current.Username)
	repl.Start(os.Stdin, os.Stdout)
}
