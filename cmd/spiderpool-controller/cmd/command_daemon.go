// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"runtime/debug"

	"github.com/spf13/cobra"
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: binNameController + " daemon",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if e := recover(); nil != e {
				logger.Sugar().Errorf("Panic details: %v", e)
				debug.PrintStack()
			}
		}()

		DaemonMain()
	},
}

func init() {
	controllerContext.BindControllerDaemonFlags(daemonCmd.PersistentFlags())
	if err := ParseConfiguration(); nil != err {
		logger.Sugar().Fatalf("Failed to register ENV for spiderpool-controller: %v", err)
	}
	controllerContext.Verify()

	rootCmd.AddCommand(daemonCmd)
}
