package main

import (
	"ask/cmd"
	"fmt"
	"os"
)

func main() {
	defer func() {
		err := recover()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}()

	cmd.Execute()
}
