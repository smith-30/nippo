package cmd

import (
	"time"

	"github.com/smith-30/nippo/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//
		// create logger
		//
		zl, err := logger.NewDefaultLogger()
		if err != nil {
			panic(err)
		}
		zl.Info("start process", zap.String("version_info", verInfo()))

		// for logging panic
		defer func() {
			if err := recover(); err != nil {
				zl.Panic(err)
			}
		}()

		time.Sleep(1 * time.Minute)
	},
}

func init() {
	// merge command
	rootCmd.AddCommand(runCmd)
}
