/*
 * ░█░█░█▀█░▀▀█░█▀█░█▀▄░▀█▀░█▀█░
 * ░░█░░█░█░▄▀░░█░█░█▀▄░░█░░█░█░
 * ░░▀░░▀▀▀░▀▀▀░▀▀▀░▀░▀░▀▀▀░▀░▀░
 * Copyright (c) 2025 kompl3xpr. All rights reserved.
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE file.
 *
 * File:               cmd/root.go
 * Last modified by:   kompl3xpr <kompl3xpr@proton.me>
 * Last modified time: 2025-10-01 09:13:46 UTC
 */

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cli-template",
	Short: "",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cli-template.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
