package command_test

import (
	"fmt"
	"log"

	"github.com/gotoolkit/command"
)

func ExampleNew() {
	cmd := command.New()
	_, err := cmd.Command("echo", "test").CombinedOutput()
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("fortune is available at %s\n", cmd)
}
