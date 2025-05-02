package main

import (
	"sync"
	"urge/cmd"
	"urge/log"
	"urge/model"
	"urge/utils"
)

func main() {
	cmd.Execute()
	numberOfWorkers := 3

	images, err := utils.InputParser()
	if err != nil {
		log.ErrorLogger.Printf("error: %s", err)
	}
	fetchChannel := make(chan model.Image)
	resizeChannel := make(chan model.Image)
	filterChannel := make(chan model.Image)
	rotateChannel := make(chan model.Image)
	saveChannel := make(chan model.Image)

	var mainWg sync.WaitGroup
	var processWg sync.WaitGroup
	var resizeWg sync.WaitGroup
	var filterWg sync.WaitGroup
	var rotateWg sync.WaitGroup

	go utils.ProduceImages(fetchChannel, images)
	for range numberOfWorkers {
		mainWg.Add(1)
		processWg.Add(1)
		go func() {
			defer mainWg.Done()
			defer processWg.Done()
			utils.ConsumeImages(fetchChannel, rotateChannel, resizeChannel, filterChannel, saveChannel)
		}()
	}
	for range numberOfWorkers {
		mainWg.Add(1)
		rotateWg.Add(1)
		go func() {
			defer mainWg.Done()
			defer rotateWg.Done()
			utils.ConsumeRotate(rotateChannel, resizeChannel, filterChannel, saveChannel)
		}()
	}
	for range numberOfWorkers {
		mainWg.Add(1)
		resizeWg.Add(1)
		go func() {
			defer mainWg.Done()
			defer resizeWg.Done()
			utils.ConsumeResize(resizeChannel, filterChannel, saveChannel)
		}()
	}
	for range numberOfWorkers {
		mainWg.Add(1)
		filterWg.Add(1)
		go func() {
			defer mainWg.Done()
			defer filterWg.Done()
			utils.ConsumeFilter(filterChannel, saveChannel)
		}()
	}
	for range numberOfWorkers {
		mainWg.Add(1)
		go func() {
			defer mainWg.Done()
			utils.ConsumeSave(saveChannel)
		}()
	}
	go func() {
		processWg.Wait()
		close(rotateChannel)
	}()
	go func() { 
		rotateWg.Wait()
		close(resizeChannel)
	}()
	go func() {
		resizeWg.Wait()
		close(filterChannel)
	}()

	go func() {
		filterWg.Wait()
		close(saveChannel)
	}()

	mainWg.Wait()
}
