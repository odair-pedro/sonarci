package logging

import (
	"log"
	"os"
)

func Setup() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}
