package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main(){
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

			err = initGitRepo()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = initialCommit()
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
    srcFile, err := os.Open("template.aseprite")
    if err != nil {
        return err
    }
    defer srcFile.Close()

    dstFile, err := os.Create(filename)
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

func initGitRepo() error {
    cmd := exec.Command("git", "init")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func initialCommit() error {
    cmd := exec.Command("git", "add", ".")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return err
    }

    cmd = exec.Command("git", "commit", "-m", "Initial commit")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    
	return cmd.Run()
}

func openAsepriteFile(filename string) error {
    var cmd *exec.Cmd

    cmd = exec.Command("cmd", "/c", "start", filename)

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}