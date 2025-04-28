package utils

import (
	"fmt"
	"os"
	"strings"
	"urge/cmd"
	"urge/log"
	"urge/model"
)

func InputParser() ([]model.Image, error) {
	if cmd.Flags.Input != "" && cmd.Flags.File != "" {
		log.WarnLogger.Printf("either provide input: %s or file: %s", cmd.Flags.Input, cmd.Flags.File)
		os.Exit(1)
	}
	if cmd.Flags.Input == "" && cmd.Flags.File == "" {
		log.WarnLogger.Printf("provide input or file")
		os.Exit(1)
	}
	var input []string
	if cmd.Flags.Input != "" {
		input = strings.Split(strings.Trim(cmd.Flags.Input, " "), ",")
		if in := MultipleInputParser(input); len(in) > 0 {
			return in, nil
		} else {
			return nil, fmt.Errorf("problem ocurred to intialize images")
		}
	}

	if cmd.Flags.File != "" {

	}
	return nil, fmt.Errorf("error in parsing input")

}

func MultipleInputParser(input []string) []model.Image {
	var images []model.Image
	for _, url := range input {
		image := model.Image{
			Image:  nil,
			Url:    url,
			Resize: cmd.Flags.Resize,
			Filter: cmd.Flags.Filter,
			Output: "",
		}
		images = append(images, image)
	}
	return images
}
