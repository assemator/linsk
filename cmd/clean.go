package cmd

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/AlexSSD7/linsk/utils"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove the Linsk data directory.",
	Run: func(cmd *cobra.Command, args []string) {
		store := createStore()

		rmPath := store.DataDirPath()
		fmt.Fprintf(os.Stderr, "Will permanently remove '"+rmPath+"'. Proceed? (y/n) > ")

		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadBytes('\n')
		if err != nil {
			slog.Error("Failed to read answer", "error", err.Error())
			os.Exit(1)
		}

		if utils.ClearUnprintableChars(strings.ToLower(string(answer)), false) != "y" {
			fmt.Fprintf(os.Stderr, "Aborted.\n")
			os.Exit(2)
		}

		err = os.RemoveAll(rmPath)
		if err != nil {
			slog.Error("Failed to remove all", "error", err.Error(), "path", rmPath)
			os.Exit(1)
		}

		// TODO: Clean network tap allocations, if any.

		slog.Info("Deleted data directory", "path", rmPath)
	},
}