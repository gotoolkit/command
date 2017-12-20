package main

import (
	"log"
	"os"

	"github.com/gotoolkit/command"
)

func main() {
	cmd := command.New().Command("echo", "123")
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bytes))

	cmd = command.New().Command("cat")
	file, _ := os.Open("doc.go")

	bytes, err = cmd.OutputWithStdin(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bytes))

	cmd = command.New().Command("echo", "test")
	f, _ := os.Create("test.md")

	err = cmd.WaitStdout(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bytes))
}
