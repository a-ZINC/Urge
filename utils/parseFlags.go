package utils

import (
	"fmt"
	"os"
	"path/filepath"
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
		fmt.Println("input", input)
		if in := MultipleInputParser(input); len(in) > 0 {
			return in, nil
		} else {
			return nil, fmt.Errorf("problem ocurred to parse inputs")
		}
	}

	if cmd.Flags.File != "" {

	}
	return nil, fmt.Errorf("error in parsing input")

}

func MultipleInputParser(input []string) []model.Image {
	var images []model.Image
	curr_directory, err := os.Getwd()
	if err != nil {
		log.WarnLogger.Printf("error in getting current directory: %s", err)
	}
	directoryPath := filepath.Join(curr_directory, "transform_"+filepath.Base(curr_directory))
	err = os.MkdirAll(directoryPath, 0755)
	if err != nil {
		log.WarnLogger.Printf("error in creating directory: %s", err)
	}
	for _, url := range input {
		url = strings.Trim(url, " ")
		if strings.HasPrefix(url, "http") || strings.HasPrefix(url, "https") {
			images, err = linkParser(images, url, directoryPath)
			if err != nil {
				log.WarnLogger.Printf("error in parsing link: %s", err)
			}
		} else {
			images, err = fileParser(images, url, directoryPath)
			if err != nil {
				log.WarnLogger.Printf("error in parsing file: %s", err)
			}
		}
	}
	return images
}

func fileParser(images []model.Image, url string, directoryPath string) ([]model.Image, error) {
	info, err := os.Stat(url)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		entries, err := os.ReadDir(url)
		if err != nil {
			return nil, err
		}
		var imageUrls []string
		for _, entry := range entries {
			fmt.Println("entry", entry.Name())
			if entry.IsDir() {
				images = append(images, MultipleInputParser([]string{filepath.Join(url, entry.Name())})...)
			}
			image := filepath.Ext(entry.Name())
			if image == ".jpeg" || image == ".jpg" || image == ".png" || image == ".gif" {
				imageUrls = append(imageUrls, filepath.Join(url, entry.Name()))
			}
		}
		images = append(images, MultipleInputParser(imageUrls)...)
	} else {
		image := model.Image{
			Image:     nil,
			Url:       url,
			Resize:    cmd.Flags.Resize,
			Filter:    cmd.Flags.Filter,
			Rotate:    cmd.Flags.Rotate,
			OutputUrl: directoryPath,
		}
		images = append(images, image)
	}
	return images, nil
}

func linkParser(images []model.Image, url string, directoryPath string) ([]model.Image, error) {
	image := model.Image{
		Image:     nil,
		Url:       url,
		Resize:    cmd.Flags.Resize,
		Filter:    cmd.Flags.Filter,
		Rotate:    cmd.Flags.Rotate,
		OutputUrl: directoryPath,
	}
	return append(images, image), nil
}
