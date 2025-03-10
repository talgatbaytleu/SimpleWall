package logger

import (
	"log"
	"os"
)

func LogMessage(message string) {
	log.SetOutput(os.Stdout)
	log.Println(message)
}

func LogError(err error) {
	log.Println(err)
	log.SetOutput(os.Stderr)
}
