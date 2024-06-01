package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
    outputFileName string
    targetDir      string
)


func main(){
	targetDir = ".code/wiki/res/aseprite"

	var rootCmd = &cobra.Command{
		Use: "rune",
		Short: "Rune is a CLI application to generate .aseprite files",
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println("Hello World")
		},
	}

	var power = &cobra.Command{
		Use: "power",
		Short: "power your rune!",
		Run: func(cmd *cobra.Command, args []string){
			err := generateAsepriteFile(args[0] + ".aseprite")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = openAsepriteFile(args[0] + ".aseprite")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(power)

	if err := rootCmd.Execute(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func generateAsepriteFile(filename string) error {
    srcFile, err := os.Open("C:\\rune\\template.aseprite")
    if err != nil {
        return err
    }
    defer srcFile.Close()

	path := filepath.Join(targetDir, filename)

    dstFile, err := os.Create(path)
    if err != nil {
        return err
    }
    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    if err != nil {
        return err
    }

    return nil
}

func openAsepriteFile(filename string) error {
    var cmd *exec.Cmd

	path := filepath.Join(targetDir, filename)

    cmd = exec.Command("cmd", "/c", "start", path)

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}