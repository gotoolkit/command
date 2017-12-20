package main

import (
	"log"

	"github.com/gotoolkit/command"
)

func main() {
	cmd := command.New()
	_, err := cmd.Command("echo", "123").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(string(d))
}
