package log

import (
	"log"
	"os"
)

var (
	InfoLogger = log.New(os.Stdout, "[Info]", log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "[Error]", log.Lshortfile)
	WarnLogger = log.New(os.Stdout, "[Warn]", log.Lshortfile)
)
