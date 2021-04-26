package main

import (
	"fmt"
	"github.com/ahabhgk/monkey-interpreter-in-go/repl"
	"os"
	"os/user"
)

func main() {
	cur, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", cur.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
