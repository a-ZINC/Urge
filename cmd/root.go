package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Urge",
	Short: "Urge is a high-performance concurrent image processing CLI tool.",
	Long: color.HiBlueString(`
          _____                    _____                    _____                    _____          
         /\    \                  /\    \                  /\    \                  /\    \         
        /::\____\                /::\    \                /::\    \                /::\    \        
       /:::/    /               /::::\    \              /::::\    \              /::::\    \       
      /:::/    /               /::::::\    \            /::::::\    \            /::::::\    \      
     /:::/    /               /:::/\:::\    \          /:::/\:::\    \          /:::/\:::\    \     
    /:::/    /               /:::/__\:::\    \        /:::/  \:::\    \        /:::/__\:::\    \    
   /:::/    /               /::::\   \:::\    \      /:::/    \:::\    \      /::::\   \:::\    \   
  /:::/    /      _____    /::::::\   \:::\    \    /:::/    / \:::\    \    /::::::\   \:::\    \  
 /:::/____/      /\    \  /:::/\:::\   \:::\____\  /:::/    /   \:::\ ___\  /:::/\:::\   \:::\    \ 
|:::|    /      /::\____\/:::/  \:::\   \:::|    |/:::/____/  ___\:::|    |/:::/__\:::\   \:::\____\
|:::|____\     /:::/    /\::/   |::::\  /:::|____|\:::\    \ /\  /:::|____|\:::\   \:::\   \::/    /
 \:::\    \   /:::/    /  \/____|:::::\/:::/    /  \:::\    /::\ \::/    /  \:::\   \:::\   \/____/ 
  \:::\    \ /:::/    /         |:::::::::/    /    \:::\   \:::\ \/____/    \:::\   \:::\    \     
   \:::\    /:::/    /          |::|\::::/    /      \:::\   \:::\____\       \:::\   \:::\____\    
    \:::\__/:::/    /           |::| \::/____/        \:::\  /:::/    /        \:::\   \::/    /    
     \::::::::/    /            |::|  ~|               \:::\/:::/    /          \:::\   \/____/     
      \::::::/    /             |::|   |                \::::::/    /            \:::\    \         
       \::::/    /              \::|   |                 \::::/    /              \:::\____\        
        \::/____/                \:|   |                  \::/____/                \::/    /        
         ~~                       \|___|                                            \/____/         
`) + `

` + color.HiYellowString("Concurrent Image Processing Pipeline") + `

` + color.HiCyanString(`
Urge is a fast and scalable CLI tool designed for concurrent image processing.
It powers transformations like resizing, format conversion, and optimization
through an efficient pipeline architecture.
`) + color.HiMagentaString(`
✨ Features:
- Parallel processing using all your CPU cores
- Easy commands for batch image tasks
- Extensible pipeline for custom workflows
- Minimal setup, blazing fast results
`) + color.HiGreenString(`
Run 'urge help [command]' to explore more.
`),
}

func Execute() {
	rootCmd.Execute()
}
