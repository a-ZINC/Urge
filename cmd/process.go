package cmd

import "github.com/spf13/cobra"

type flags struct {
	Input string
	Output string
	Resize string
	Filter string
	Rotate int
	File string
	TakeFlagFromFile bool
}

var processCmd = &cobra.Command{
	Use: "process",
	Short: "Batch process images with concurrency and speed.",
}

var Flags = flags{}

func init() {
	processCmd.PersistentFlags().StringVarP(&Flags.Input, "input", "i", "", "Image or Images path")
	processCmd.PersistentFlags().StringVarP(&Flags.Output, "output", "o", "", "Path location where user want images")
	processCmd.PersistentFlags().StringVarP(&Flags.Resize, "resize", "r", "", "resize image according to given dimension")
	processCmd.PersistentFlags().StringVarP(&Flags.Filter, "filter", "f", "", "filter image according to given filter")
	processCmd.PersistentFlags().StringVarP(&Flags.File, "file", "x", "", "file url where all image url is stored")
	processCmd.PersistentFlags().IntVarP(&Flags.Rotate, "rotate", "t", 0, "rotate image according to given degree")
	processCmd.PersistentFlags().BoolVarP(&Flags.TakeFlagFromFile, "main", "y", false, "cosnider that File has all each image url and its corresponding flags so neglect user inputed flags")
	processCmd.MarkFlagRequired("input")
	rootCmd.AddCommand(processCmd);
}