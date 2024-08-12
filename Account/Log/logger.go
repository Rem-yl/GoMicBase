package logger

import (
	"log"
	"os"
)

// Initialize sets up the global logger with desired settings.
func Init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
}
