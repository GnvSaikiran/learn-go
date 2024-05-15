package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

//go:embed passwords.txt
var passwords string

func fileEmbed() {
	pwds := strings.Split(passwords, "\n")
	if len(os.Args) > 1 {
		for _, pwd := range pwds {
			if pwd == os.Args[1] {
				fmt.Println(true)
				return
			}
		}

		fmt.Println(false)
	}
}

//go:embed help
var helpInfo embed.FS

func embedDir() {
	if len(os.Args) == 1 {
		printHelpFiles()
		return
	}

	data, err := helpInfo.ReadFile("help/" + os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func printHelpFiles() {
	fmt.Println("Contents:")
	fs.WalkDir(helpInfo, "help", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			_, name, _ := strings.Cut(path, "help/")
			fmt.Println(name)
		}

		return nil
	})
}

//go:embed rights_texts
var rightsDir embed.FS

func printRights() {
	if len(os.Args) == 1 {
		fmt.Println("Provide a language")
		return
	}

	lang := os.Args[1]
	path_to_file := fmt.Sprintf("rights_texts/%s_rights.txt", lang)
	data, err := rightsDir.ReadFile(path_to_file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
