/*
Copyright © 2020 Oliver Jakoubek <info@jakoubek.net>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		addTask(cmd)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().IntP("project_id", "p", -1, "ID of the project")
	addCmd.Flags().StringP("title", "n", "", "Title of the task")
	addCmd.Flags().IntP("category_id", "c", -1, "ID of the category")
	addCmd.Flags().StringP("tags", "t", "", "Tags of the task")
}

func addTask(cmd *cobra.Command) {
	projectID, _ := cmd.Flags().GetInt("project_id")
	if projectID == -1 {
		panic("No project given")
	}

	title, _ := cmd.Flags().GetString("title")
	if title == "" {
		panic("No title given")
	}

	categoryID, _ := cmd.Flags().GetInt("category_id")

	tagsString, _ := cmd.Flags().GetString("tags")
	if tagsString == "" {
		panic("No tags given")
	}
	tags := strings.Split(tagsString, ",")

	fmt.Println("Project :", projectID)
	fmt.Println("Title   :", title)
	fmt.Println("Category:", categoryID)
	fmt.Println("Tags    :")
	for _, t := range tags {
		fmt.Println(t)
	}
}
