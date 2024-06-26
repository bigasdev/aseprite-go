package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	targetDir   string
	glyphDir    string
	craftDir    string
	filePathDir string
	templateDir string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "rune",
		Short: "Rune is a CLI application to generate .aseprite files",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello World")
		},
	}

	var power = &cobra.Command{
		Use:   "power",
		Short: "power your rune!",
		Run: func(cmd *cobra.Command, args []string) {
			err := generateAsepriteFile(args[0] + ".aseprite")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = openFile(args[0] + ".aseprite")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var glyph = &cobra.Command{
		Use:   "glyph",
		Short: "glyph your rune!",
		Run: func(cmd *cobra.Command, args []string) {
			err := generateMdFile(glyphDir, args[0]+".md")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = openFile(args[0] + ".md")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var craft = &cobra.Command{
		Use:   "craft",
		Short: "craft your rune!",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if args[1] == "video" {
				err := generateCraftFile(craftDir, args[0]+".md", Video())
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

			} else if args[1] == "article" {
				err := generateCraftFile(craftDir, args[0]+".md", Article())
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			} else if args[1] == "prototype" {
				err := generateCraftFile(craftDir, args[0]+".md", Prototype())
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			} else if args[1] == "project" {
				err := generateCraftFile(craftDir, args[0]+".md", Project())
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			} else {
				fmt.Println("Invalid type")
				os.Exit(1)

			}

			err := openFile(args[0] + ".md")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	var extract = &cobra.Command{
		Use:   "extract",
		Short: "extract your rune!",
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] == "video" {
				err := openFolder(craftDir)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			} else if args[0] == "article" {
				err := openFolder(craftDir)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			} else if args[0] == "prototype" {
				err := openFolder(craftDir)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			} else if args[0] == "project" {
				err := openFolder(craftDir)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			} else {
				fmt.Println("Invalid type")
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(power)
	rootCmd.AddCommand(glyph)
	rootCmd.AddCommand(craft)
	rootCmd.AddCommand(extract)
	startViper()

	Video()

	//loading the config paramters
	targetDir = viper.GetString("user.power_path")
	glyphDir = viper.GetString("user.glyph_path")
	craftDir = viper.GetString("user.craft_path")
	templateDir = viper.GetString("user.template_path")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generateAsepriteFile(filename string) error {
	srcFile, err := os.Open(templateDir)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	path := filepath.Join(targetDir, filename)

	filePathDir = targetDir

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

func openFile(filename string) error {
	var cmd *exec.Cmd

	path := filepath.Join(filePathDir, filename)

	cmd = exec.Command("cmd", "/c", "start", path)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func openFolder(folderpath string) error {
	var cmd *exec.Cmd

	cmd = exec.Command("cmd", "/c", "explorer", folderpath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func generateMdFile(filename string, pathDir string) error {
	path := filepath.Join(pathDir, filename)

	filePathDir = pathDir

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("# Hello World")
	if err != nil {
		return err
	}

	return nil
}

func generateCraftFile(pathDir string, filename string, template string) error {
	path := filepath.Join(pathDir, filename)

	filePathDir = pathDir

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(template)
	if err != nil {
		return err
	}

	return nil
}

func startViper() {
	viper.AddConfigPath("C:\\rune\\configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	}
}
