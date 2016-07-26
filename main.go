package main

import (
	"log"
	"os"
	"os/exec"
	"github.com/pugnascotia/cron"
)

func main() {

	// args exluding program name
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatalln("Usage: mantra cron_spec cmd [ args ... ]")
	}

	program := args[1]
	programArgs := args[2:]

	err := cron.Run(args[0], func() {
		cmd := exec.Command(program, programArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	})

	if (err != nil) {
		log.Fatalln("Failed to run, exiting...")
	}
}
