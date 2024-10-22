package main

import (
	"fmt"
	"miggolang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!", user.Username)
	fmt.Printf("Type in your commands\n")
	repl.Start(os.Stdin, os.Stdout)
}