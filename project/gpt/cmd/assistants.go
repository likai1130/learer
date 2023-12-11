/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// assistantsCmd represents the assistants command
var assistantsCmd = &cobra.Command{
	Use:   "assistants",
	Short: "使用命令创建一个助手",
	Long:  `使用命令创建一个助手. 例如: gpt assistants create`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("assistants called")
	},
}

func init() {
	rootCmd.AddCommand(assistantsCmd)
}
