package config

import (
	"fmt"
	"os"
	"strings"
	"urge/cmd"
	"urge/log"
)

func InputParser() ([]string, error) {
	if cmd.Flags.Input != "" && cmd.Flags.File != "" {
		log.WarnLogger.Printf("either provide input: %s or file: %s", cmd.Flags.Input, cmd.Flags.File)
		os.Exit(1);
	}
	if cmd.Flags.Input == "" && cmd.Flags.File == "" {
		log.WarnLogger.Printf("provide input or file");
		os.Exit(1);
	}
	if cmd.Flags.Input != "" {
		input := strings.Split(strings.Trim(cmd.Flags.Input, " "), ",")
		return input, nil
	}

	if cmd.Flags.File != "" {

	}
	return nil, fmt.Errorf("error in parsing input")

}