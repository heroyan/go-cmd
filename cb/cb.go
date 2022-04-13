package cb

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fage",
	Short: "fage's cobra demo",
	Long: "the demo show how to use cobra package",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cobra demo program, with args: %v\n", args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
	},
}

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Print images information",
	Long: "Print all images information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("image one is ubuntu 16.04")
		fmt.Println("image two is ubuntu 18.04")
		fmt.Println("image args are : " + strings.Join(args, " "))
	},
}

var echoTimes int
var cmdTimes = &cobra.Command{
	Use:   "times [string to echo]",
	Short: "Echo anything to the screen more times",
	Long: `echo things multiple times back to the user by providing a count and a string.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < echoTimes; i++ {
			fmt.Println("Echo: " + strings.Join(args, " "))
		}
	},
}

var imageListCmd = &cobra.Command{
	Use: "list",
	Short: "list images",
	Long: "list all the images",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yes, this a image list")
		fmt.Println("image args are: " + strings.Join(args, ""))
	},
}

var imageListClearCmd = &cobra.Command{
	Use: "clear",
	Short: "clear the list",
	Long: "clear the image list all",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clear the image list all")
		fmt.Println("image args are: " + strings.Join(args, ""))
	},
}

func init()  {
	rootCmd.AddCommand(imageCmd)
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	imageCmd.AddCommand(cmdTimes)
	imageCmd.AddCommand(imageListCmd)
	imageListCmd.AddCommand(imageListClearCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
