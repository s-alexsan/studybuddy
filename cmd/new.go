/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/s-alexsan/studybuddy/data"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new studybuddy note",
	Long:  `Creates a new studybuddy note`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	noteCmd.AddCommand(newCmd)
}

func promptGetImput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ .}}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"animal", "food", "person", "object"}
	index := -1

	var result string
	var err error

	if index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index != -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func createNewNote() {
	wordPromptContent := promptContent{
		"Please provide a word",
		"what word would you like to make a note of?",
	}

	word := promptGetImput(wordPromptContent)

	definitionPromptContent := promptContent{
		"Please provide a definition",
		fmt.Sprintf("What is the definition of %s?", word),
	}

	definition := promptGetImput(definitionPromptContent)

	categoryPromptContent := promptContent{
		"Please provide a category",
		fmt.Sprintf("What category does %s belogin to?", word),
	}

	category := promptGetSelect(categoryPromptContent)

	data.InsertNote(word, definition, category)

}
